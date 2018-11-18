package histogram

import (
	"image"
	"math"
	"sync"
)

func (hset *HistogramSet) ProcessImage(img *image.Image, wg *sync.WaitGroup) {
	hist := newHistogram(hset.Bits, hset.Nbins)
	x_min := (*img).Bounds().Min.X
	y_min := (*img).Bounds().Min.Y
	x_max := (*img).Bounds().Max.X
	y_max := (*img).Bounds().Max.Y
	var i_act float64
	var i_int int
	for x := x_min; x < x_max; x++ {
		for y := y_min; y < y_max; y++ {
			px := (*img).At(x, y)
			r, _, _, _ := px.RGBA()
			i_act = (float64(r) - hist.Min) / hist.Step
			i_int = int(math.Floor(i_act))
			if (i_int >= 0) && (i_int < hset.Nbins) {
				hist.Cts[i_int] += 1
			} else if i_int == hset.Nbins {
				hist.Cts[i_int-1] += 1
			}
		}
	}
	hset.Mtx.Lock()
	hset.Set = append(hset.Set, hist)
	hset.Mtx.Unlock()
	wg.Done()
	return
}
