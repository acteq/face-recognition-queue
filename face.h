#ifndef FACE_H
#define FACE_H

#include "amcomdef.h"

#ifdef __cplusplus
extern "C" {
#endif

MHandle initEngine(const char*, const char*);
ASF_FaceFeature extract(const char *, MHandle);

#ifdef __cplusplus    
}
#endif

#endif