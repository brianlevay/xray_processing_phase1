package histogram

import (
	"math"
)

func newHistogram(bits int, nbins int) *Histogram {
	hist := new(Histogram)
	hist.Min = 0
	hist.Max = math.Pow(2, float64(bits)) - 1
	hist.Step = (hist.Max - hist.Min) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Bins = append(hist.Bins, float64(i)*hist.Step+hist.Min)
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
