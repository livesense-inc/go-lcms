package lcms

// #include <stdlib.h>
// #include <lcms2.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Profile struct {
	inner C.cmsHPROFILE
}

func (p *Profile) CloseProfile() {
	if p.inner == nil {
		return
	}
	C.cmsCloseProfile(p.inner)
}

func OpenProfileFromMem(d []byte) (*Profile, error) {
	data := unsafe.Pointer(&d[0])
	dataLen := C.cmsUInt32Number(len(d))
	p := C.cmsOpenProfileFromMem(data, dataLen)
	if p == nil {
		return nil, fmt.Errorf("failed to open a profile")
	}
	return &Profile{inner: p}, nil
}

func CreateSRGBProfile() (*Profile, error) {
	p := C.cmsCreate_sRGBProfile()
	if p == nil {
		return nil, fmt.Errorf("failed to open sRGB profile")
	}
	return &Profile{inner: p}, nil
}
