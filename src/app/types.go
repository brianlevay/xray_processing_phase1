package main

import (
	"sync"
)

type Configuration struct {
	Mtx          sync.Mutex `json:"-"`
	SrcHeight    float64    `json:"SrcHeight"`
	CoreHeight   float64    `json:"CoreHeight"`
	Motion       float64    `json:"Motion"`
	HeightPxDet  int        `json:"HeightPxDet"`
	WidthPxDet   int        `json:"WidthPxDet"`
	CmPerPx      float64    `json:"CmPerPx"`
	Bits         int        `json:"Bits"`
	Nbins        int        `json:"Nbins"`
	HeightPxHist int        `json:"HeightPxHist"`
	WidthPxHist  int        `json:"WidthPxHist"`
	R            uint8      `json:"R"`
	G            uint8      `json:"G"`
	B            uint8      `json:"B"`
	ThreshFrac   float64    `json:"ThreshFrac"`
	GapMinFrac   float64    `json:"GapMinFrac"`
	GapMaxFrac   float64    `json:"GapMaxFrac"`
	Nmass        float64    `json:"Nmass"`
	FilterSteps  int        `json:"FilterSteps"`
	MaxTheta     float64    `json:"MaxTheta"`
	Tmin         float64    `json:"Tmin"`
	Tedge        float64    `json:"Tedge"`
	BorderPx     int        `json:"BorderPx"`
	ScaleWidth   float64    `json:"ScaleWidth"`
	RoiWidth     float64    `json:"RoiWidth"`
}

type HistogramSet struct {
	Cfg    *Configuration
	Mtx    sync.Mutex
	Set    []*Histogram
	Merged *Histogram
	Image  []byte
}

type ImgProcessor struct {
	Cfg          *Configuration `json:"-"`
	CoreType     string         `json:"CoreType"`
	CoreDiameter float64        `json:"CoreDiameter"`
	AxisMethod   string         `json:"AxisMethod"`
	AxisAngle    float64        `json:"AxisAngle"`
	AxisOffset   float64        `json:"AxisOffset"`
	IlowFrac     float64        `json:"IlowFrac"`
	IpeakFrac    float64        `json:"IpeakFrac"`
	IhighFrac    float64        `json:"IhighFrac"`
	FolderName   string         `json:"FolderName"`
	FileAppend   string         `json:"FileAppend"`
	ProjMult     float64        `json:"-"`
	ImaxInFlt    float64        `json:"-"`
	ImaxInInt    uint16         `json:"-"`
	ImaxOutFlt   float64        `json:"-"`
	ImaxOutInt   uint16         `json:"-"`
	IthreshInt   uint16         `json:"-"`
	PxGapMin     int            `json:"-"`
	PxGapMax     int            `json:"-"`
	MassTable    []float64      `json:"-"`
	Xd           []float64      `json:"-"`
	Yd           []float64      `json:"-"`
	Xc           float64        `json:"-"`
	Yc           float64        `json:"-"`
	MurhotTable  []float64      `json:"-"`
	Tref         float64        `json:"-"`
	Omin         float64        `json:"-"`
	Opeak        float64        `json:"-"`
	Omax         float64        `json:"-"`
	Lstep        float64        `json:"-"`
	IcontTable   []uint16       `json:"-"`
	Iscale       [][]uint16     `json:"-"`
}

type Histogram struct {
	Step float64
	Cts  []float64
}

type TModel struct {
	CoreType  string
	R         float64
	Cos0      float64
	Sin0      float64
	Xra       float64
	Yra       float64
	Zra       float64
	Xrs       float64
	Yrs       float64
	Zrs       float64
	DelZr     float64
	DelZr2    float64
	C         float64
	XrStep    float64
	XrMin     float64
	XrMax     float64
	Txz2Table []float64
}
