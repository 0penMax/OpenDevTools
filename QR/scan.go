package QR

import (
	"bytes"
	"fmt"
	"image"
	"openDevTools/models"

	// import gif, jpeg, png
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	// import bmp, tiff, webp
	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/tiff"
	_ "golang.org/x/image/webp"

	"github.com/makiuchi-d/gozxing"
	multiQR "github.com/makiuchi-d/gozxing/multi/qrcode"
)

func Scan(b []byte) ([]models.ResultItem, error) {
	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, fmt.Errorf("failed to read image: %v", err)
	}

	source := gozxing.NewLuminanceSourceFromImage(img)
	bin := gozxing.NewHybridBinarizer(source)
	bbm, err := gozxing.NewBinaryBitmap(bin)

	if err != nil {
		return nil, fmt.Errorf("error during processing: %v", err)
	}

	qrReader := multiQR.NewQRCodeMultiReader()
	result, err := qrReader.DecodeMultiple(bbm, nil)
	if err != nil {
		return nil, fmt.Errorf("unable to decode QRCode: %v", err)
	}

	var results []models.ResultItem
	for i, element := range result {
		results = append(results, models.ResultItem{
			Name:  fmt.Sprintf("%d", i),
			Value: element.String(),
		})

	}
	return results, nil
}
