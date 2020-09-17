package main

/*
#cgo CFLAGS:  -I./ArcSoft_ArcFace_Linux_x64_V3.0/inc -I/usr/include -I/usr/local/include
#cgo CXXFLAGS:  -std=c++11  -I./ArcSoft_ArcFace_Linux_x64_V3.0/inc -I/usr/include -I/usr/local/include -I/usr/include/opencv
#cgo LDFLAGS:  -lstdc++ -L${SRCDIR}/ArcSoft_ArcFace_Linux_x64_V3.0/lib/linux_x64 -larcsoft_face -larcsoft_face_engine 
#cgo LDFLAGS:  -L/usr/lib/x86_64-linux-gnu -l:libopencv_core.a 
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_imgproc.a 
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_highgui.a  
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_flann.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_photo.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_features2d.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_contrib.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_legacy.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_ts.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_objdetect.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_ocl.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_stitching.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_superres.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_calib3d.a 
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libopencv_gpu.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libpng.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libjpeg.a
#cgo LDFLAGS:      -L/usr/lib/x86_64-linux-gnu  -l:libtiff.a
#cgo LDFLAGS:      -Wl,-rpath=./
#include "face.h"
*/
import "C"

import (
	// "gocv.io/x/gocv"
	// "log"
	"unsafe"
)

type Handle unsafe.Pointer

func initEngine() Handle {
	APPID := "D617np8jyKt1jN9gMr7ENbTJ1XvRWsQAsvdAehMKvqzr"
	SDKKEY := "BTeyKLCYBzfrNLQduGUwfyRZkDQU4RB2b3UJZjjuppLU"

	return Handle(C.initEngine(C.CString(APPID), C.CString(SDKKEY)))
}

func extract(filePath string, handle Handle ) []byte{

	// handle := C.initEngine(C.CString(APPID), C.CString(SDKKEY))
	feature := C.extract(C.CString(filePath), unsafe.Pointer(handle))
	
	if feature.featureSize > 0 {
		return C.GoBytes(unsafe.Pointer(feature.feature), feature.featureSize)
	}else{
		return nil
	}
}

type EnginePool struct {
    cached chan Handle
    size   int
}

func NewEnginePool(size int) *EnginePool {
    x := EnginePool{}
    x.cached = make(chan Handle, size)
    x.size = size
    return &x
}

func (c *EnginePool) Get() Handle {
    var res Handle
    select {
    case res = <-c.cached:
        // fmt.Println("---get--")
    default:
        // fmt.Println("---create one--")
        res = Handle(initEngine())
    }

    return res
}

func (p *EnginePool) Put(c Handle) {
    select {
    case p.cached <- c:
        // fmt.Println("---put--")
    default:
        c = nil
        // fmt.Println("---destroy--")
    }
}