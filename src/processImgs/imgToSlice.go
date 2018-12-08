package processImgs

import (
	"errors"
	"image"
)

func Gray16ToUint16(img image.Image) ([][]uint16, error) {
	gray16, ok := img.(*image.Gray16)
	if ok == false {
		return [][]uint16{}, errors.New("Image not Gray16 format")
	}
	// x_max and y_max aren't the maximum values, they're max+1 //
	x_min := gray16.Rect.Min.X
	x_max := gray16.Rect.Max.X
	y_min := gray16.Rect.Min.Y
	y_max := gray16.Rect.Max.Y
	width := x_max - x_min
	height := y_max - y_min

	var k int
	slice := make([][]uint16, height)
	for i := 0; i < height; i++ {
		slice[i] = make([]uint16, width)
		for j := 0; j < width; j++ {
			k = (i-y_min)*gray16.Stride + (j-x_min)*2
			slice[i][j] = uint16(gray16.Pix[k+0])<<8 | uint16(gray16.Pix[k+1])
		}
	}
	return slice, nil
}

func Uint16ToGray16(slice [][]uint16) *image.Gray16 {
	x_min := 0
	y_min := 0
	height := len(slice)
	width := len(slice[0])
	topLeft := image.Point{x_min, y_min}
	bottomRight := image.Point{width + x_min, height + y_min}
	gray16 := image.NewGray16(image.Rectangle{topLeft, bottomRight})

	var k int
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k = (i-y_min)*gray16.Stride + (j-x_min)*2
			gray16.Pix[k+0] = uint8(slice[i][j] >> 8)
			gray16.Pix[k+1] = uint8(slice[i][j])
		}
	}
	return gray16
}
