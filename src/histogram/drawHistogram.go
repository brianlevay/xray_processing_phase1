package histogram

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"log"
)

func DrawHistogram(hist *Histogram, width int, height int) *bytes.Buffer {
	// Configuration Variables //
	blue := color.RGBA{66, 134, 244, 255}

	nbins := len(hist.Cts)
	maxcts := hist.MaxCts()
	bins_per_px := float64(nbins) / float64(width)
	cts_per_px := maxcts / float64(height)

	topLeft := image.Point{0, 0}
	bottomRight := image.Point{width, height}
	img := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	var index int
	var value float64
	var top int
	for x := 0; x < width; x++ {
		index = int(float64(x) * bins_per_px)
		value = hist.Cts[index]
		top = int((maxcts - value) / cts_per_px)
		for y := 0; y < height; y++ {
			if y < top {
				img.Set(x, y, color.White)
			} else {
				img.Set(x, y, blue)
			}
		}
	}

	buffer := new(bytes.Buffer)
	errE := png.Encode(buffer, img)
	if errE != nil {
		log.Println("Unable to encode image.")
	}
	return buffer
}
