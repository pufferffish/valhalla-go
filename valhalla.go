package valhalla

// #cgo LDFLAGS: -L./result/lib -lvalhalla_go
// #include <stdio.h>
// #include <stdlib.h>
// #include "./bindings/valhalla_go.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Actor struct {
	ptr unsafe.Pointer
}

func NewActorFromFile(configPath string) (*Actor, error) {
	var isError uint8 = 0
	cs := C.CString(configPath)
	resp := C.actor_init_from_file(cs, (*C.char)(unsafe.Pointer(&isError)))
	C.free(unsafe.Pointer(cs))
	switch isError {
	case 0:
		return &Actor{ptr: unsafe.Pointer(resp)}, nil
	case 1:
		err := C.GoString((*C.char)(resp))
		C.free(unsafe.Pointer(resp))
		return nil, errors.New(err)
	default:
		panic("Invalid error code from valhalla C binding")
	}
}

func NewActorFromConfig(config *Config) (*Actor, error) {
	var isError uint8 = 0
	cs := C.CString(config.String())
	resp := C.actor_init_from_config(cs, (*C.char)(unsafe.Pointer(&isError)))
	C.free(unsafe.Pointer(cs))
	switch isError {
	case 0:
		return &Actor{ptr: unsafe.Pointer(resp)}, nil
	case 1:
		err := C.GoString((*C.char)(resp))
		C.free(unsafe.Pointer(resp))
		return nil, errors.New(err)
	default:
		panic("Invalid error code from valhalla C binding")
	}
}


func (actor *Actor) Centroid(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_centroid((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Expansion(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_expansion((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Height(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_height((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Isochrone(request string) (string, error) {
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

func (actor *Actor) Locate(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_locate((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Matrix(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_matrix((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) OptimizedRoute(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_optimized_route((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Route(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_route((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) Status(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_status((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) TraceAttributes(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_trace_attributes((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) TraceRoute(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_trace_route((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

func (actor *Actor) TransitAvailable(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_transit_available((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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

