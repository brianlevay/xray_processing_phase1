package histogram

import (
	"math"
)

func newHistogram(bits int, nbins int) *Histogram {
	hist := new(Histogram)
	hist.Step = math.Pow(2, float64(bits)) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Cts = append(hist.Cts, 0)
	}
	return hist
}

func newHistogramSet(bits int, nbins int) *HistogramSet {
	histSet := new(HistogramSet)
	histSet.Bits = bits
	histSet.Nbins = nbins
	return histSet
}
