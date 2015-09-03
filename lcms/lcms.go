package lcms

/*
#cgo LDFLAGS: -llcms2

#include <stdlib.h>
#include <lcms2.h>

*/
import "C"

import (
	"fmt"
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

func Create_sRGBProfile() *Profile {
	return &Profile{prof: C.cmsCreate_sRGBProfile()}
}

func (prof *Profile) CloseProfile() {
	C.cmsCloseProfile(prof.prof)
}

func CreateTransform(src_prof *Profile, src_type CMSType, dst_prof *Profile, dst_type CMSType) *Transform {
	transform := C.cmsCreateTransform(
		src_prof.prof, C.cmsUInt32Number(src_type),
		dst_prof.prof, C.cmsUInt32Number(dst_type),
		C.INTENT_PERCEPTUAL, 0)
	return &Transform{trans: transform}
}

func (trans *Transform) DeleteTransform() {
	C.cmsDeleteTransform(trans.trans)
}

func (trans *Transform) DoTransform(inputBuffer []uint8, outputBuffer []uint8, length int) error {
	inputLen := len(inputBuffer)
	outputLen := len(outputBuffer)
	if inputLen < length {
		return fmt.Errorf("DoTransform: inputLen(%d) < length(%d)", inputLen, length)
	}
	if outputLen < length {
		return fmt.Errorf("DoTransform: outputLen(%d) < length(%d)", outputLen, length)
	}
	inputPtr := unsafe.Pointer(&inputBuffer)
	outputPtr := unsafe.Pointer(&outputBuffer)
	C.cmsDoTransform(trans.trans, inputPtr, outputPtr, C.cmsUInt32Number(length))
	return nil
}
