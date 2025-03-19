package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/jpeg"
	"io"
	"log"
	"os"

	"github.com/livesense-inc/go-lcms/lcms"
)

var (
	//go:embed default.icc
	iccProfile []byte

	//go:embed cmyk.jpg
	cmykImage []byte
)

func main() {
	r := bytes.NewReader(cmykImage)
	w, err := os.OpenFile("rgb.jpg", os.O_RDWR|os.O_CREATE, 0o644)
	if err != nil {
		log.Fatal(err)
	}
	if err := convertColor(r, w); err != nil {
		log.Fatal(err)
	}
}

func convertColor(r io.Reader, w io.Writer) error {
	srcProf := lcms.OpenProfileFromMem(iccProfile)
	if srcProf == nil {
		return fmt.Errorf("failed to open a source profile")
	}
	defer srcProf.CloseProfile()

	dstProf := lcms.CreateSRGBProfile()
	if dstProf == nil {
		return fmt.Errorf("failed to open a destination profile")
	}
	defer dstProf.CloseProfile()

	t := lcms.CreateTransform(srcProf, lcms.TYPE_CMYK_8, dstProf, lcms.TYPE_RGBA_8)
	if t == nil {
		return fmt.Errorf("failed to create a transform object")
	}
	defer t.DeleteTransform()

	img, format, err := image.Decode(r)
	if err != nil {
		return nil
	}
	if format != "jpeg" {
		return fmt.Errorf("not jpeg image")
	}

	switch src := img.(type) {
	case *image.CMYK:
		dst := image.NewRGBA(src.Bounds())
		t.DoTransform(src.Pix, dst.Pix, len(src.Pix))
		for i := range dst.Pix {
			if (i+1)%4 == 0 {
				dst.Pix[i] = 255 // Alpha
			}
		}
		bw := bufio.NewWriter(w)
		if err := jpeg.Encode(bw, dst, &jpeg.Options{Quality: jpeg.DefaultQuality}); err != nil {
			return err
		}
	default:
		return fmt.Errorf("not CMYK image")
	}

	return nil
}
