package lcms

/*
#include <lcms2.h>
*/
import "C"

type CMSType int

const (
	TYPE_RGB_8   CMSType = C.TYPE_RGB_8
	TYPE_RGBA_8  CMSType = C.TYPE_RGBA_8
	TYPE_YCbCr_8 CMSType = C.TYPE_YCbCr_8
)
