package processImgs

import (
	"image"
	"image/color"
)

func Gray16ToFloat(img image.Image) [][]float64 {
	x_min := img.Bounds().Min.X
	y_min := img.Bounds().Min.Y
	x_max := img.Bounds().Max.X
	y_max := img.Bounds().Max.Y
	width := x_max - x_min + 1
	height := y_max - y_min + 1
	
	slice := make([][]float64, height)
	for i := 0; i < height; i++ {
		slice[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			px := img.At(j + y_min, i + x_min)
			r, _, _, _ := px.RGBA()
			slice[i][j] = float64(r)
		}
	}
	return slice
}

func FloatToGray16(slice [][]float64) *image.Gray16 {
	height := len(slice)
	width := len(slice[0])
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}
	img := image.NewGray16(image.Rectangle{topLeft, bottomRight})
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			c := color.Gray16{uint16(slice[i][j])}
			img.Set(j, i, c)
		}
	}
	return img
}