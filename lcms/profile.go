package lcms

/*
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

func OpenProfileFromFile(filename string) *Profile {
	csfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(csfilename))
	csmode := C.CString("r")
	defer C.free(unsafe.Pointer(csmode))
	return &Profile{prof: C.cmsOpenProfileFromFile(csfilename, csmode)}
}

func Create_sRGBProfile() *Profile {
	return &Profile{prof: C.cmsCreate_sRGBProfile()}
}

func (prof *Profile) CloseProfile() {
	C.cmsCloseProfile(prof.prof)
}
