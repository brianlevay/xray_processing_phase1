package histogram

import (
	"sync"
)

type Histogram struct {
	Step float64   `json:"Step"`
	Cts  []float64 `json:"Cts"`
}

type HistogramSet struct {
	Mtx   sync.Mutex
	Bits  int
	Nbins int
	Set   []*Histogram
}
