package histogram

import (
	"bytes"
	"image"
	"image/png"
	"log"
)

func (hset *HistogramSet) DrawImage() {
	var index, top, k int
	var value float64

	// Avoids the last bin to prevent scaling to saturated values //
	maxVal := 0.0
	for i := 0; i < (hset.Nbins - 1); i++ {
		if hset.Merged.Cts[i] > maxVal {
			maxVal = hset.Merged.Cts[i]
		}
	}
	bins_per_px := float64(hset.Nbins) / float64(hset.WidthPxHist)
	cts_per_px := maxVal / float64(hset.HeightPxHist)
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{hset.WidthPxHist, hset.HeightPxHist}
	rgba := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	for j := 0; j < hset.WidthPxHist; j++ {
		index = int(float64(j) * bins_per_px)
		value = hset.Merged.Cts[index]
		top = int((maxVal - value) / cts_per_px)
		for i := 0; i < hset.HeightPxHist; i++ {
			k = (i * rgba.Stride) + j*4
			if i < top {
				rgba.Pix[k] = 255
				rgba.Pix[k+1] = 255
				rgba.Pix[k+2] = 255
			} else {
				rgba.Pix[k] = hset.R
				rgba.Pix[k+1] = hset.G
				rgba.Pix[k+2] = hset.B
			}
			rgba.Pix[k+3] = 255
		}
	}
	buffer := new(bytes.Buffer)
	errE := png.Encode(buffer, rgba)
	if errE != nil {
		log.Println("Unable to encode image.")
	}
	hset.Image = buffer.Bytes()
	return
}
