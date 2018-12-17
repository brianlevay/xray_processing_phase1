package main

import (
	"bytes"
	fe "fileExplorer"
	"image"
	"image/png"
	"log"
	"math"
	"sync"
)

func ImageHistogram(contents *fe.FileContents, hset *HistogramSet) {
	var wg sync.WaitGroup
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
		wg.Add(1)
		go hset.ProcessImage(contents.Root, contents.Selected[i], &wg)
	}
	wg.Wait()
	hset.MergeHistograms()
	hset.DrawImage()
	return
}

func (hset *HistogramSet) ProcessImage(root string, filename string, wg *sync.WaitGroup) {
	var pxVal uint16
	var k, h_int int
	var h_act float64
	img, errImg := fe.OpenTiff(root, filename)
	if errImg != nil {
		log.Println("Error opening " + filename)
		wg.Done()
		return
	}
	gray16, ok := img.(*image.Gray16)
	if ok == false {
		log.Println(filename + " not Gray16 format")
		wg.Done()
		return
	}
	hist := newHistogram(hset.Cfg.Bits, hset.Cfg.Nbins)
	// x_max and y_max aren't the maximum values, they're max+1 //
	x_min := gray16.Rect.Min.X
	x_max := gray16.Rect.Max.X
	y_min := gray16.Rect.Min.Y
	y_max := gray16.Rect.Max.Y
	width := x_max - x_min
	height := y_max - y_min

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k = (i-y_min)*gray16.Stride + (j-x_min)*2
			pxVal = uint16(gray16.Pix[k+0])<<8 | uint16(gray16.Pix[k+1])
			h_act = float64(pxVal) / hist.Step
			h_int = int(math.Floor(h_act))
			if (h_int >= 0) && (h_int < hset.Cfg.Nbins) {
				hist.Cts[h_int] += 1
			} else if h_int >= hset.Cfg.Nbins {
				hist.Cts[h_int-1] += 1
			}
		}
	}
	hset.Mtx.Lock()
	hset.Set = append(hset.Set, hist)
	hset.Mtx.Unlock()
	wg.Done()
	return
}

func (hset *HistogramSet) MergeHistograms() {
	hist := newHistogram(hset.Cfg.Bits, hset.Cfg.Nbins)
	nhists := len(hset.Set)
	for i := 0; i < nhists; i++ {
		for b := 0; b < hset.Cfg.Nbins; b++ {
			hist.Cts[b] += hset.Set[i].Cts[b]
		}
	}
	hset.Merged = hist
}

func (hset *HistogramSet) DrawImage() {
	var index, top, k int
	var value float64
	// Avoids the last bin to prevent scaling to saturated values //
	maxVal := 0.0
	for i := 0; i < (hset.Cfg.Nbins - 1); i++ {
		if hset.Merged.Cts[i] > maxVal {
			maxVal = hset.Merged.Cts[i]
		}
	}
	bins_per_px := float64(hset.Cfg.Nbins) / float64(hset.Cfg.WidthPxHist)
	cts_per_px := maxVal / float64(hset.Cfg.HeightPxHist)
	topLeft := image.Point{0, 0}
	bottomRight := image.Point{hset.Cfg.WidthPxHist, hset.Cfg.HeightPxHist}
	rgba := image.NewRGBA(image.Rectangle{topLeft, bottomRight})

	for j := 0; j < hset.Cfg.WidthPxHist; j++ {
		index = int(float64(j) * bins_per_px)
		value = hset.Merged.Cts[index]
		top = int((maxVal - value) / cts_per_px)
		for i := 0; i < hset.Cfg.HeightPxHist; i++ {
			k = (i * rgba.Stride) + j*4
			if i < top {
				rgba.Pix[k] = 255
				rgba.Pix[k+1] = 255
				rgba.Pix[k+2] = 255
			} else {
				rgba.Pix[k] = hset.Cfg.R
				rgba.Pix[k+1] = hset.Cfg.G
				rgba.Pix[k+2] = hset.Cfg.B
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
