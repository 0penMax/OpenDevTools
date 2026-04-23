package QR

import (
	"errors"
	"io"

	qrg "github.com/piglig/go-qr"
)

type OutputType string

const (
	Png OutputType = "png"
	Svg OutputType = "svg"
)

func Generate(data []byte, outputType OutputType, w io.Writer) error {
	qr, err := qrg.EncodeBinary(data, qrg.Medium)
	if err != nil {
		return err
	}

	config := qrg.NewQrCodeImgConfig(10, 4) // scale=10px, border=4 modules
	switch outputType {
	case Png:
		err = qr.WriteAsPNG(config, w)
		if err != nil {
			return err
		}
	case Svg:
		err = qr.WriteAsSVG(config, w)
		if err != nil {
			return err
		}
	default:
		return errors.New("invalid output type")
	}

	return nil
	//_ = qr.PNG(config, "hello.png")
	//_ = qr.SVG(config, "hello.svg")
	//
	//enc := qrcode.NewQRCodeWriter()
	//img, err := enc.Encode(string(data), gozxing.BarcodeFormat_QR_CODE, 250, 250, nil)
	//if err != nil {
	//	return err
	//}
	//// *BitMatrix implements the image.Image interface,
	//// so it is able to be passed to png.Encode directly.
	//return png.Encode(w, img)
}
