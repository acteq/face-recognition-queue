package main

/*
#cgo CFLAGS:  -I./ArcSoft_ArcFace_Linux_x64_V3.0/inc -I/usr/include -I/usr/local/include
#cgo CXXFLAGS: -std=c++11  -I./ArcSoft_ArcFace_Linux_x64_V3.0/inc -I/usr/include -I/usr/local/include -I/usr/local/include/opencv4 -I/usr/local/include/opencv4/opencv2
#cgo LDFLAGS: -L${SRCDIR}/ArcSoft_ArcFace_Linux_x64_V3.0/lib/linux_x64 -L/usr/local/lib -lopencv_core -lopencv_imgproc -lopencv_highgui -lopencv_imgcodecs -lstdc++ -larcsoft_face -larcsoft_face_engine -Wl,-rpath=./
#include "face.h"
*/
import "C"

import (
	// "gocv.io/x/gocv"
	// "log"
	// "strings"
)


func extract(filePath string) {
	APPID := "64p4MfjsFyVXQM21jbTmHQbCeDcv4wwgVVaQMZaAhmGD"
	SDKKEY := "8tGcaiv4BzcG232WmGKWTa1izjsXWSLTg8CKJsGvESX4"

	C.initEngine(C.CString(APPID), C.CString(SDKKEY))
}