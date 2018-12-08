package histogram

import (
	fe "fileExplorer"
	"image"
	"log"
	"math"
	"sync"
)

func (hset *HistogramSet) ProcessImage(root string, filename string, wg *sync.WaitGroup) {
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
	hist := newHistogram(hset.Bits, hset.Nbins)

	// x_max and y_max aren't the maximum values, they're max+1 //
	x_min := gray16.Rect.Min.X
	x_max := gray16.Rect.Max.X
	y_min := gray16.Rect.Min.Y
	y_max := gray16.Rect.Max.Y
	width := x_max - x_min
	height := y_max - y_min

	var pxVal uint16
	var k, h_int int
	var h_act float64
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			k = (i-y_min)*gray16.Stride + (j-x_min)*2
			pxVal = uint16(gray16.Pix[k+0])<<8 | uint16(gray16.Pix[k+1])
			h_act = float64(pxVal) / hist.Step
			h_int = int(math.Floor(h_act))
			if (h_int >= 0) && (h_int < hset.Nbins) {
				hist.Cts[h_int] += 1
			} else if h_int >= hset.Nbins {
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

func (hist *Histogram) MaxCts() float64 {
	var maxVal float64 = 0
	for i := 0; i < (len(hist.Cts) - 1); i++ {
		if hist.Cts[i] > maxVal {
			maxVal = hist.Cts[i]
		}
	}
	return maxVal
}
