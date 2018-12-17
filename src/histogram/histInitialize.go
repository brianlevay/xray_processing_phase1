package histogram

import (
	"errors"
	"math"
)

func (hset *HistogramSet) Initialize(cfg map[string]float64) error {
	// Configuration variables
	hset.Bits = int(cfg["Bits"])
	hset.Nbins = int(cfg["Nbins"])
	hset.HeightPxHist = int(cfg["HeightPxHist"])
	hset.WidthPxHist = int(cfg["WidthPxHist"])
	hset.R = uint8(cfg["R"])
	hset.G = uint8(cfg["G"])
	hset.B = uint8(cfg["B"])

	if (hset.HeightPxHist == 0) || (hset.WidthPxHist == 0) || (hset.Bits < 0) || (hset.Nbins <= 0) {
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
