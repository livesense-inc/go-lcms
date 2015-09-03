package lcms

import (
	"image"
)

func ImageTransformByProfile(src_image *image.RGBA, src_prof Profile, dst_prof Profile) (*image.RGBA, error) {
	rect := src_image.Bounds()
	dst_image := image.NewRGBA(rect)
	//
	transform := CreateTransform(&src_prof, TYPE_RGBA_8, &dst_prof, TYPE_RGBA_8)
	//
	src_pix := src_image.Pix
	dst_pix := dst_image.Pix
	len_pix := len(src_pix)
	transform.DoTransform(src_pix, dst_pix, len_pix)
	return dst_image, nil
}
