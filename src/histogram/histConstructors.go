package histogram

import (
	"math"
)

func (hset *HistogramSet) Initialize() {
	// Configuration variables
	hset.Bits = 14
	hset.Nbins = 256
}

func newHistogram(bits int, nbins int) *Histogram {
	hist := new(Histogram)
	hist.Step = math.Pow(2, float64(bits)) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Cts = append(hist.Cts, 0)
	}
	return hist
}
