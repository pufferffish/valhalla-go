package main

import (
	"fmt"
	"valhalla-go/valhalla"
)

func main() {
	request := `
		{"locations":[{"lat":40.744014,"lon":-73.990508}],"costing":"pedestrian","contours":[{"time":15.0,"color":"ff0000"}]}
	`
	actor, err := valhalla.NewActorFromFile("test_config/config.json")
	if err != nil {
		panic(err.Error())
	}

	resp, err := actor.Isochrone(request)
	if err != nil {
		println(err.Error())
	} else {
		println(resp)
	}

	resp, err = actor.Isochrone("}")
	if err != nil {
		println(err.Error())
	} else {
		println(resp)
	}

	_, err = valhalla.NewActorFromFile("waewaewe")
	if err != nil {
		println(err.Error())
	} else {
		panic("expected error")
	}

	conf := valhalla.DefaultConfig()
	conf.SetTileExtractPath("/goofy/ahh")
	_, err = valhalla.NewActorFromConfig(conf)
	if err != nil {
		panic(err)
	}

	fmt.Println(valhalla.DefaultConfig())
}
