package lcms

/*
#cgo LDFLAGS: -llcms2

#include <stdlib.h>
#include <lcms2.h>

*/
import "C"

import (
	"unsafe"
)

type Profile struct {
	prof C.cmsHPROFILE
}

type Transform struct {
	trans C.cmsHTRANSFORM
}

func OpenProfileFromFile(filename string) *Profile {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	csmode := C.CString("r")
	defer C.free(unsafe.Pointer(csmode))
	return &Profile{prof: C.cmsOpenProfileFromFile(csfilename, csmode)}
}

func cmsCreate_sRGBProfile() *Profile {
	return &Profile{prof: C.cmsCreate_sRGBProfile()}
}

func (prof *Profile) CloseProfile() {
	C.cmsCloseProfile(prof.prof)
}

func CreateTransform(src_prof Profile, dst_prof Profile) *Transform {
	transform := C.cmsCreateTransform(
		src_prof.prof, C.TYPE_BGR_8,
		dst_prof.prof, C.TYPE_BGR_8,
		C.INTENT_PERCEPTUAL, 0)
	return &Transform{trans: transform}
}

func (trans *Transform) DeleteTransform() {
	C.cmsDeleteTransform(trans.trans)
}
