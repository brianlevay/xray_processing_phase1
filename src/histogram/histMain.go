package histogram

import (
	fe "fileExplorer"
	"image"
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

func (hset *HistogramSet) MergeHistograms() {
	hist := newHistogram(hset.Bits, hset.Nbins)
	nhists := len(hset.Set)
	for i := 0; i < nhists; i++ {
		for b := 0; b < hset.Nbins; b++ {
			hist.Cts[b] += hset.Set[i].Cts[b]
		}
	}
	hset.Merged = hist
}
