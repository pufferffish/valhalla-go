package valhalla

// #cgo LDFLAGS: -L../bindings/result/lib -lvalhalla_go
// #include <stdio.h>
// #include <stdlib.h>
// #include "../bindings/valhalla_go.h"
import "C"
import (
	"errors"
	"unsafe"
)

type Actor struct {
	ptr unsafe.Pointer
}

func NewActor(configPath string) (*Actor, error) {
	var isError uint8 = 0
	cs := C.CString(configPath)
	resp := C.actor_init(cs, (*C.char)(unsafe.Pointer(&isError)))
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

{{ range $k, $v := .Functions }}
func (actor *Actor) {{$v}}(request string) (string, error) {
	var isError uint8 = 0
	cs := C.CString(request)
	cresp := C.actor_{{$k}}((C.Actor)(actor.ptr), cs, (*C.char)(unsafe.Pointer(&isError)))
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
{{ end }}