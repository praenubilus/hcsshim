package main

import (
	/*
		#include <stdlib.h>
	*/
	"C"
	"fmt"
	"unsafe"
	"os"
)

//export get_rootfs_layer_hashes
func get_rootfs_layer_hashes(cImage *C.char, cTag *C.char, destination *C.char, cSize *C.int, cResults ***C.char) {
	image := C.GoString(cImage)
	tag := C.GoString(cTag)
	layers, err := getRootfsLayerHashes(image, tag, C.GoString(destination))
	fmt.Println(err)
	if err != nil {
		os.Exit(1)
	}
	print("\n Image:[" + image + ":" + tag + "]\n")

	results := C.malloc(C.size_t(len(layers)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	// *[1 << 30]*C.char is a a pointer to an array of size 1 << 30, of *C.char values. The size is
	// arbitrary, and only represents an upper bound that needs to be valid on the host system.
	a := (*[1<<30 - 1]*C.char)(results)

	for idx, layer := range layers {
		a[idx] = C.CString(layer)
	}
	*cSize = C.int(len(layers))
	*cResults = (**C.char)(results)
}

//export clean_up
func clean_up(cSize C.int, cResults **C.char) {
	fmt.Printf("\nsize is %v is\n", cSize)
	for i := 0; i < int(cSize); i++ {
		layer := *(**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(cResults)) + unsafe.Sizeof(*cResults)*uintptr(i)))
		// fmt.Printf("\nvalue is %s \n", C.GoString(layer))
		C.free(unsafe.Pointer(layer))
	}
	C.free(unsafe.Pointer(cResults))
}
