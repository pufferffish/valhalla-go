package main

// #cgo LDFLAGS: -L./bindings/result/lib -lvalhalla_go
// #include <stdio.h>
// #include <stdlib.h>
// #include "bindings/valhalla_go.h"
import "C"
import "unsafe"

type ValhallaActor struct {
	ptr unsafe.Pointer
}

func NewValhallaActor(configPath string) *ValhallaActor {
	cs := C.CString(configPath)
	actor := unsafe.Pointer(C.actor_init(cs))
	C.free(unsafe.Pointer(cs))
	return &ValhallaActor{ptr: actor}
}

func (actor *ValhallaActor) Isochrone(request string) string {
	cs := C.CString(request)
	cresp := C.actor_isochrone((C.Actor)(actor.ptr), cs)
	resp := C.GoString(cresp)
	C.free(unsafe.Pointer(cresp))
	return resp
}

func main() {
	request := `
		{"locations":[{"lat":40.744014,"lon":-73.990508}],"costing":"pedestrian","contours":[{"time":15.0,"color":"ff0000"}]}
	`
	actor := NewValhallaActor("test_config/config.json")
	println(actor.Isochrone(request))
}
