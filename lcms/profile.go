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

func OpenProfileFromMem(profdata []byte) *Profile {
	data := unsafe.Pointer(&profdata[0])
	dataLen := C.cmsUInt32Number(len(profdata))
	return &Profile{prof: C.cmsOpenProfileFromMem(data, dataLen)}
}

func Create_sRGBProfile() *Profile {
	return &Profile{prof: C.cmsCreate_sRGBProfile()}
}

func (prof *Profile) CloseProfile() {
	if prof.prof != nil {
		C.cmsCloseProfile(prof.prof)
	}
}
