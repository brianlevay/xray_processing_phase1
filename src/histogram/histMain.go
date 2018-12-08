package histogram

import (
	fe "fileExplorer"
	"sync"
)

func ImageHistogram(contents *fe.FileContents, bits int, nbins int) *Histogram {
	var wg sync.WaitGroup
	histSet := newHistogramSet(bits, nbins)
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
		wg.Add(1)
		go histSet.ProcessImage(contents.Root, contents.Selected[i], &wg)
	}
	wg.Wait()
	hist := mergeHistograms(histSet)
	return hist
}

func mergeHistograms(hset *HistogramSet) *Histogram {
	hist := newHistogram(hset.Bits, hset.Nbins)
	nhists := len(hset.Set)
	for i := 0; i < nhists; i++ {
		for b := 0; b < hset.Nbins; b++ {
			hist.Cts[b] += hset.Set[i].Cts[b]
		}
	}
	return hist
}
