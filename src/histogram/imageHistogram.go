package histogram

import (
	fe "fileExplorer"
)

func ImageHistogram(contents *fe.FileContents, bits int, nbins int) *Histogram {
	histSet := newHistogramSet(bits, nbins)
	fe.ProcessTiffs(contents, histSet)
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
