package lcms

// #include <stdlib.h>
// #include <lcms2.h>
import "C"

import (
	"fmt"
	"unsafe"
)

type Transform struct {
	inner C.cmsHTRANSFORM
}

func (t *Transform) DeleteTransform() {
	if t.inner == nil {
		return
	}
	C.cmsDeleteTransform(t.inner)
}

func (t *Transform) DoTransform(in []uint8, out []uint8, size int) error {
	C.cmsDoTransform(
		t.inner,
		unsafe.Pointer(&in[0]),
		unsafe.Pointer(&out[0]),
		C.cmsUInt32Number(size),
	)
	return nil
}

func CreateTransform(
	srcProf *Profile,
	srcType CMSType,
	dstProf *Profile,
	dstType CMSType,
) (*Transform, error) {
	t := C.cmsCreateTransformTHR(
		C.cmsCreateContext(nil, nil),
		srcProf.inner,
		C.cmsUInt32Number(srcType),
		dstProf.inner,
		C.cmsUInt32Number(dstType),
		C.INTENT_PERCEPTUAL,
		C.cmsUInt32Number(C.cmsFLAGS_NOCACHE),
	)
	if t == nil {
		return nil, fmt.Errorf("failed to create a transform object")
	}
	return &Transform{inner: t}, nil
}
