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
	"time"

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
	if err := convertColor(r, w, iccProfile); err != nil {
		log.Fatal(err)
	}
}

func convertColor(r io.Reader, w io.Writer, srcICCProfData []byte) error {
	srcProf, err := lcms.OpenProfileFromMem(srcICCProfData)
	if err != nil {
		return err
	}
	defer srcProf.CloseProfile()

	dstProf, err := lcms.CreateSRGBProfile()
	if err != nil {
		return err
	}
	defer dstProf.CloseProfile()

	now := time.Now()
	t, err := lcms.CreateTransform(srcProf, lcms.TYPE_CMYK_8, dstProf, lcms.TYPE_RGBA_8)
	if err != nil {
		return err
	}
	defer t.DeleteTransform()
	fmt.Printf("transform object creation: %s\n", time.Since(now))

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
		t.DoTransform(src.Pix, dst.Pix, len(src.Pix)/4)
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
