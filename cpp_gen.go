//go:build ignore
// +build ignore

package main

import (
	"os"
	"text/template"
)

type Instance struct {
	Functions []string
}

func writeTemplate(instance *Instance, templatePath, outputPath string) {
	tl, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}

	out, err := os.OpenFile(outputPath, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		panic(err)
	}

	if err = tl.Execute(out, *instance); err != nil {
		panic(err)
	}
}

func main() {
	instance := Instance{
		Functions: []string{"route", "locate", "optimized_route", "matrix", "isochrone", "trace_route", "trace_attributes", "height", "transit_available", "expansion", "centroid", "status"},
	}
	writeTemplate(&instance, "templates/valhalla_go.templ.cpp", "bindings/valhalla_go.cpp")
	writeTemplate(&instance, "templates/valhalla_go.templ.h", "bindings/valhalla_go.h")
}
