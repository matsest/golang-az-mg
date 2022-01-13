# Parsing YAML with Go

Sample code for parsing a YAML file with Go. Just for learning :)

## Learning...

- Unmarshal YAML files into a nested `ManagementGroup` [struct](https://golangdocs.com/structs-in-golang)
- Print nested struct
- Flatten a nested struct into an array
- Marshal nested struct to json with [encoding/json](https://go.dev/pkg/encoding/json/)
- Using [flag](https://pkg.go.dev/flag) for command-line flag parsing

## Input

To have an [example YAML input](mg.yml) a representation of [Azure Management Groups](https://docs.microsoft.com/en-us/azure/governance/management-groups/overview) is used.

## Usage

```bash

# run without building
go run main.go

# build
go build -o main main.go

# run binary and print help text
./main --help

# run with other file
./main -file <yaml file>

```

Related:

- [Working with YAML in Golang (ZetCode)](https://zetcode.com/golang/yaml/)
- [Golang Yaml Package blog post (GODocs)](https://golangdocs.com/golang-yaml-package)
- [Golang YAML package / API Reference](https://pkg.go.dev/gopkg.in/yaml.v3)
