//go:build ignore
// +build ignore

package main

import (
	"os"
	"strings"
	"text/template"
	"unicode"
)

type Instance struct {
	Functions map[string]string
}

func ConvertName(name string) string {
	convertedName := ""
	for _, word := range strings.Split(name, "_") {
		r := []rune(word)
		convertedName += string(append([]rune{unicode.ToUpper(r[0])}, r[1:]...))
	}
	return convertedName
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
	fn := []string{"route", "locate", "optimized_route", "matrix", "isochrone", "trace_route", "trace_attributes", "height", "transit_available", "expansion", "centroid", "status"}
	instance := Instance{
		Functions: map[string]string{},
	}
	for _, v := range fn {
		instance.Functions[v] = ConvertName(v)
	}

	writeTemplate(&instance, "templates/valhalla_go.templ.cpp", "bindings/valhalla_go.cpp")
	writeTemplate(&instance, "templates/valhalla_go.templ.h", "bindings/valhalla_go.h")
	writeTemplate(&instance, "templates/valhalla.templ.go", "valhalla.go")
}
