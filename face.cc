#include "arcsoft_face_sdk.h"
#include "asvloffscreen.h"
#include "merror.h"
// #include <iostream>  
// #include <string>
#include <stdio.h>
#include <stdlib.h>
// #include <string.h>
// #include <time.h>
#include<opencv.hpp>
#include "face.h"

using namespace std;

MHandle initEngine (const char* APPID, const char* SDKKEY) {
    MHandle handle = NULL;

    MRESULT res = MOK;
	ASF_ActiveFileInfo activeFileInfo = { 0 };
	res = ASFGetActiveFileInfo(&activeFileInfo);
	if (res != MOK){
		printf("ASFGetActiveFileInfo fail: %d\n", int(res));
        return NULL;
	}

	//SDK版本信息
	// const ASF_VERSION version = ASFGetVersion();
	
	res = ASFOnlineActivation(MPChar(APPID), MPChar(SDKKEY));
	if (MOK != res && MERR_ASF_ALREADY_ACTIVATED != res){
		printf("ASFOnlineActivation fail: %d\n", int(res));
        return NULL;
    }

	//初始化引擎
	MInt32 mask = ASF_FACE_DETECT | ASF_FACERECOGNITION | ASF_AGE | ASF_GENDER | ASF_FACE3DANGLE | ASF_LIVENESS | ASF_IR_LIVENESS;
	res = ASFInitEngine(ASF_DETECT_MODE_IMAGE, ASF_OP_0_ONLY, 32, 4, mask, &handle);
	if (res != MOK){
		printf("ASFInitEngine fail: %d\n", int(res));
        return NULL;
    }
	return handle;
}


ASF_FaceFeature extract(const char * filePath, MHandle handle) {
    cv::Mat originalImg = cv::imread(filePath, cv::IMREAD_UNCHANGED);
	
    //图像裁剪，宽度做四字节对齐
    int width = originalImg.cols - originalImg.cols%4;
    int height = originalImg.rows; //区域大小
    
    cv::Mat img(originalImg, cv::Rect(0, 0, width, height)); //设置源图像ROI

    int widthStep=(img.cols*img.elemSize()+3)/4*4;

    //图像数据以结构体形式传入，对更高精度的图像兼容性更好
    ASVLOFFSCREEN offscreen = { 0 };
    offscreen.u32PixelArrayFormat = ASVL_PAF_RGB24_B8G8R8;
    offscreen.i32Width = img.cols;
    offscreen.i32Height = img.rows;
    offscreen.pi32Pitch[0] = widthStep;
    offscreen.ppu8Plane[0] = (MUInt8*)img.data;

    ASF_MultiFaceInfo detectedFaces = { 0 };
    ASF_FaceFeature feature = { 0 };
    
    MRESULT res = ASFDetectFacesEx(handle, &offscreen, &detectedFaces);
    if (MOK != res){
        printf("ASFDetectFacesEx failed: %d\n", int(res));
    } else{
        // 打印人脸检测结果
        for (int i = 0; i < detectedFaces.faceNum; i++){
            printf("Face Id: %d\n", detectedFaces.faceID[i]);
            printf("Face Orient: %d\n", detectedFaces.faceOrient[i]);
            printf("Face Rect: (%d %d %d %d)\n",
            detectedFaces.faceRect[i].left, detectedFaces.faceRect[i].top,
            detectedFaces.faceRect[i].right,
            detectedFaces.faceRect[i].bottom);
        }

        ASF_SingleFaceInfo singleDetectedFaces = { 0 };
        singleDetectedFaces.faceRect.left = detectedFaces.faceRect[0].left;
        singleDetectedFaces.faceRect.top = detectedFaces.faceRect[0].top;
        singleDetectedFaces.faceRect.right = detectedFaces.faceRect[0].right;
        singleDetectedFaces.faceRect.bottom = detectedFaces.faceRect[0].bottom;
        singleDetectedFaces.faceOrient = detectedFaces.faceOrient[0];
        //图像数据以结构体形式传入，对更高精度的图像兼容性更好

        MRESULT res = ASFFaceFeatureExtractEx(handle, &offscreen, &singleDetectedFaces, &feature);
        if (MOK != res) {
            printf("ASFFaceFeatureExtractEx failed: %d\n", int(res));
        } 
    }
    

    return feature;
}

