package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v3"
)

type ManagementGroup struct {
	DisplayName string             `yaml:"name"`
	Id          string             `yaml:"id"`
	parentId    string             // not marshable (lowercase)
	Children    *[]ManagementGroup `yaml:"children"`
}

// Sets the ParentID in nested ManagementGroup structs
func (mg *ManagementGroup) setParents() {
	if mg.Children != nil {
		var tmp []ManagementGroup
		for _, child := range *mg.Children {
			child.parentId = mg.Id
			tmp = append(tmp, child)
			child.setParents()
		}
		*mg.Children = tmp
	}
}

// Returns a flat array of Management Group structs
func flattenMg(mg ManagementGroup) []ManagementGroup {
	mgs := []ManagementGroup{}
	mgs = append(mgs, mg)

	children := mg.Children
	if children != nil {
		for _, c := range *children {
			mgs = append(mgs, flattenMg(c)...)
		}
	}
	return mgs
}

// For "pretty printing" of ManagementGroup struct :)
func (mg *ManagementGroup) print(level int) {
	fmt.Printf("%s%s (%s)\n", strings.Repeat(" ", level), mg.DisplayName, mg.Id)
	if mg.parentId != "" {
		fmt.Printf("%sparentId: %s\n", strings.Repeat(" ", level+3), mg.parentId)
	}
	if mg.Children != nil {
		level += 5
		fmt.Printf("%sChildren: \n", strings.Repeat(" ", level-2))
		for _, c := range *mg.Children {
			c.print(level) // Recurse!
		}
	}
}

func main() {
	// Input parameters
	var printFile = flag.Bool("printFile", false, "to print file content")
	var file = flag.String("file", "mg.yml", "name of file to read")
	flag.Parse()

	// Read file content
	data, err := ioutil.ReadFile(*file)
	if err != nil {
		log.Printf("data.Get err #%v", err)
	}

	// Unmarshal from yaml to ManagementGroup
	mg := ManagementGroup{}
	err = yaml.Unmarshal(data, &mg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	// TODO: Add parent id to all nodes?
	mg.setParents()

	// Print file content
	if *printFile {
		fmt.Println("#### Input yaml:")
		fmt.Println(string(data))
	}

	// Pretty print the nested ManagementGroup struct
	fmt.Println("\n#### Processed:")
	mg.print(0)

	// Print flattened array of []ManagementGroup
	flatMgs := flattenMg(mg)
	fmt.Println("\n#### Flattened array:")
	for _, m := range flatMgs {
		fmt.Printf("%+v\n", m)
	}

	// Marshalled to json
	fmt.Println("\n#### Marshalled json:")
	jsonF, _ := json.MarshalIndent(mg, "", "  ")
	fmt.Println(string(jsonF))
}
