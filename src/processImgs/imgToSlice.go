package processImgs

import (
	"image"
	"image/color"
)

func Gray16ToUint16(img image.Image) [][]uint16 {
	x_min := img.Bounds().Min.X
	x_max := img.Bounds().Max.X
	y_min := img.Bounds().Min.Y
	y_max := img.Bounds().Max.Y
	width := x_max - x_min + 1
	height := y_max - y_min + 1

	slice := make([][]uint16, height)
	for i := 0; i < height; i++ {
		slice[i] = make([]uint16, width)
		for j := 0; j < width; j++ {
			px := img.At(j+y_min, i+x_min)
			r, _, _, _ := px.RGBA()
			slice[i][j] = uint16(r)
		}
	}
	return slice
}

func Uint16ToGray16(slice [][]uint16) *image.Gray16 {
	height := len(slice)
	width := len(slice[0])
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}
	img := image.NewGray16(image.Rectangle{topLeft, bottomRight})
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := color.Gray16{slice[i][j]}
			img.Set(j, i, c)
		}
	}
	return img
}
