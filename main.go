package main

// #cgo LDFLAGS: -L./bindings/result/lib -lvalhalla_go
// #include <stdio.h>
// #include <stdlib.h>
// #include "bindings/valhalla_go.h"
import "C"
import (
	"errors"
	"unsafe"
)

type ValhallaActor struct {
	ptr unsafe.Pointer
}

func NewValhallaActor(configPath string) *ValhallaActor {
	cs := C.CString(configPath)
	actor := unsafe.Pointer(C.actor_init(cs))
	C.free(unsafe.Pointer(cs))
	return &ValhallaActor{ptr: actor}
}

func (actor *ValhallaActor) Isochrone(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_isochrone((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
	resp := C.GoString(cresp)
	C.free(unsafe.Pointer(cresp))
	switch isError {
	case 0:
		return resp, nil
	case 1:
		return "", errors.New(resp)
	default:
		panic("Invalid error code from valhalla C binding")
	}
}

func main() {
	request := `
		{"locations":[{"lat":40.744014,"lon":-73.990508}],"costing":"pedestrian","contours":[{"time":15.0,"color":"ff0000"}]}
	`
	actor := NewValhallaActor("test_config/config.json")
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
}
