package main

import (
	"flag"
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	logo := flag.String("logo", "", "PNG file with a logo to put in the middle")
	out := flag.String("outfile", "out.png", "filename for the output PNG file")
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Printf("Missing URL argument")
		return
	}
	qrc, err := qrcode.New(args[0])
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	options := []standard.ImageOption{
		standard.WithBuiltinImageEncoder(standard.PNG_FORMAT),
	}
	if *logo != "" {
		options = append(options,
			standard.WithLogoImageFilePNG(*logo),
			standard.WithLogoSizeMultiplier(3),
			standard.WithQRWidth(21))
	}

	w, err := standard.New(*out, options...)
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}
