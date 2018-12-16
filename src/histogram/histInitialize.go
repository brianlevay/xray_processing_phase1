package histogram

import (
	"errors"
	"math"
)

func (hset *HistogramSet) Initialize(cfg map[string]float64) error {
	// Configuration variables
	hset.Bits = int(cfg["Bits"])
	hset.Nbins = int(cfg["Nbins"])

	if (hset.Height == 0) || (hset.Width == 0) || (hset.Bits < 0) || (hset.Nbins <= 0) {
		return errors.New("Invalid configuration values for histogram")
	}
	return nil
}

func newHistogram(bits int, nbins int) *Histogram {
	hist := new(Histogram)
	hist.Step = math.Pow(2, float64(bits)) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Cts = append(hist.Cts, 0)
	}
	return hist
}
