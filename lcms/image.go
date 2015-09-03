package lcms

import (
	"fmt"
	"image"
	"image/color"
)

func ImageTransformByProfile(src_image image.Image, src_prof Profile, dst_prof Profile) (image.Image, error) {
	var dst_image image.Image
	transform := CreateTransform(&src_prof, TYPE_RGBA_8, &dst_prof, TYPE_RGBA_8)
	rect := src_image.Bounds()

	colorModel := src_image.ColorModel()
	switch colorModel {
	case color.RGBAModel:
		src_rgba64 := src_image.(*image.RGBA64) // type assertions
		dst_rgba64 := image.NewRGBA(rect)
		src_pix := src_rgba64.Pix
		dst_pix := dst_rgba64.Pix
		len_pix := len(src_pix)
		transform.DoTransform(src_pix, dst_pix, len_pix)
		dst_image = image.Image(dst_rgba64)
	default:
		return nil, fmt.Errorf("ImageTransformByProfile: Unsupported ColorModel(%d)", colorModel)
	}
	return dst_image, nil
}
