package histogram

import (
	"sync"
)

type Histogram struct {
	Step float64   `json:"Step"`
	Cts  []float64 `json:"Cts"`
}

type HistogramSet struct {
	Mtx    sync.Mutex   `json:"-"`
	Bits   int          `json:"-"`
	Nbins  int          `json:"-"`
	Set    []*Histogram `json:"-"`
	Merged *Histogram   `json:"-"`
	Width  int          `json:"Width"`
	Height int          `json:"Height"`
	R      uint8        `json:"R"`
	G      uint8        `json:"G"`
	B      uint8        `json:"B"`
	Image  []byte       `json:"-"`
}
