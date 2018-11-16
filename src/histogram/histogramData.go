package histogram

import (
	"sync"
)

type Histogram struct {
	Min  float64   `json:"Min"`
	Max  float64   `json:"Max"`
	Step float64   `json:"Step"`
	Bins []float64 `json:"Bins"`
	Cts  []float64 `json:"Cts"`
}

type HistogramSet struct {
	Set []*Histogram
	Mtx sync.Mutex
}
