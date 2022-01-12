package main

import (
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
	Children    *[]ManagementGroup `yaml:"children"`
	Parent      string             `yaml:"parent"`
}

// For "pretty printing" of ManagementGroup struct :)
func printManagementGroup(v ManagementGroup, level int) {
	fmt.Printf("%s%s (%s)\n", strings.Repeat(" ", level), v.DisplayName, v.Id)
	if v.Parent != "" {
		fmt.Printf("%sParent: %s\n", strings.Repeat(" ", level), v.Parent)
	}
	if v.Children != nil {
		level += 5
		fmt.Printf("%sChildren: \n", strings.Repeat(" ", level-2))
		for _, c := range *v.Children {
			printManagementGroup(c, level) // Recurse!
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

	// Print file content
	if *printFile {
		fmt.Println("#### Input yaml:")
		fmt.Println(string(data))
	}

	// Print unmarshalled data
	fmt.Println("#### Processed:")
	printManagementGroup(mg, 0)
}
