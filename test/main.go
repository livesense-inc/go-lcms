package main

// http://blog.golang.org/go-image-package
// http://d.hatena.ne.jp/taknb2nch/20131231/1388500659

import (
	"flag"
	"fmt"
	"github.com/yoya/go-lcms/lcms"
	"image"
	"image/jpeg"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "usage:main <imagefile> <proffile>")
}

func main() {
	flag.Parse()
	imagefile := flag.Arg(0)
	proffile := flag.Arg(1)
	ifd_image, err := os.Open(imagefile)
	if err != nil {
		fmt.Fprintln(os.Stderr, "not found:", err)
		usage()
		os.Exit(1)
	}
	/*
		ifd_prof, err := os.Open(proffile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "not found:"+proffile)
			os.Exit(1)
		}
	*/

	src_image, err := jpeg.Decode(ifd_image)
	if err != nil {
		fmt.Fprintln(os.Stderr, "jpeg.Decode failed:", err)
		os.Exit(1)
	}
	//
	lcms.ErrorReportingToStderr()
	var src_prof *lcms.Profile = lcms.OpenProfileFromFile(proffile)
	defer src_prof.CloseProfile()
	var dst_prof *lcms.Profile = lcms.Create_sRGBProfile()
	defer dst_prof.CloseProfile()
	dst_image, err := lcms.ImageTransformByProfile(src_image, src_prof, dst_prof)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ImageTransformByProfile failed:", err)
		os.Exit(1)
	}
	image_ycbcr := dst_image.(*image.YCbCr)
	opts := jpeg.Options{Quality: 86}
	jpeg.Encode(os.Stdout, image_ycbcr, &opts)
}
