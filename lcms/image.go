package lcms

import (
	"fmt"
	"image"
	"image/color"
)

func ImageTransformByProfile(src_image image.Image, src_prof, dst_prof *Profile) (image.Image, error) {
	var dst_image image.Image
	rect := src_image.Bounds()
	width := rect.Dx()
	height := rect.Dy()
	colorModel := src_image.ColorModel()
	switch colorModel {
	case color.YCbCrModel:
		transform := CreateTransform(src_prof, TYPE_YCbCr_8, dst_prof, TYPE_YCbCr_8)
		src_ycbcr := src_image.(*image.YCbCr) // type assertions
		subsampleratio := src_ycbcr.SubsampleRatio
		dst_ycbcr := image.NewYCbCr(rect, subsampleratio)
		src_pix := make([]uint8, width*height*3)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				yi := src_ycbcr.YOffset(x, y)
				ci := src_ycbcr.COffset(x, y)
				src_pix[(x+y*width)*3] = src_ycbcr.Y[yi]
				src_pix[(x+y*width)*3+1] = src_ycbcr.Cb[ci]
				src_pix[(x+y*width)*3+2] = src_ycbcr.Cr[ci]
			}
		}
		len_pix := len(src_pix)
		dst_pix := make([]uint8, len_pix)
		transform.DoTransform(src_pix, dst_pix, len_pix)
		for y := 0; y < height; y++ {
			for x := 0; x < width; x++ {
				yi := dst_ycbcr.YOffset(x, y)
				ci := dst_ycbcr.COffset(x, y)
				dst_ycbcr.Y[yi] = dst_pix[(x+y*width)*3]
				dst_ycbcr.Cb[ci] = dst_pix[(x+y*width)*3+1]
				dst_ycbcr.Cr[ci] = dst_pix[(x+y*width)*3+2]
			}
		}
		dst_image = image.Image(dst_ycbcr)
	case color.RGBAModel:
		transform := CreateTransform(src_prof, TYPE_RGBA_8, dst_prof, TYPE_RGBA_8)
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
