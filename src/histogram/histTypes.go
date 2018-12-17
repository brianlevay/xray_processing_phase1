package histogram

import (
	"sync"
)

type Histogram struct {
	Step float64
	Cts  []float64
}

type HistogramSet struct {
	Mtx          sync.Mutex
	Bits         int
	Nbins        int
	Set          []*Histogram
	Merged       *Histogram
	HeightPxHist int
	WidthPxHist  int
	R            uint8
	G            uint8
	B            uint8
	Image        []byte
}
