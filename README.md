# go-graph

## SYNOPSIS

go-graph   [-s | -silent]  [-h | -help]

## Description

Generates a go dependency graph as html and open the system html-browser with this file.

go-graph uses go mod graph to generate a graph file for the current project.
Then it generates the HTML-file based on the graph file and save this file
as go_dependencies.html. At last it opens the browser
with the HTML-file, if not the parameter -s|-silent is set.

## Parameters

```
  -s | -silent       Don't open the browser, only generate the html file.
  -h | -help         Show help
```

## Build

```sh
go mod tidy
go build
copy the binary go-graph to /usr/local/bin or an other directory in your search path
```

## Usage

Cd into the base directory of your go project and call go-graph
