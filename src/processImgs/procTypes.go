package processImgs

import ()

type ImgProcessor struct {
	CoreType     string     `json:"CoreType"`
	CoreDiameter float64    `json:"CoreDiameter"`
	SrcHeight    float64    `json:"SrcHeight"`
	CoreHeight   float64    `json:"CoreHeight"`
	Motion       float64    `json:"Motion"`
	AxisMethod   string     `json:"AxisMethod"`
	AxisAngle    float64    `json:"AxisAngle"`
	AxisOffset   float64    `json:"AxisOffset"`
	Ilow         float64    `json:"Ilow"`
	Ipeak        float64    `json:"Ipeak"`
	Ihigh        float64    `json:"Ihigh"`
	FolderName   string     `json:"FolderName"`
	FileAppend   string     `json:"FileAppend"`
	Height       int        `json:"-"`
	Width        int        `json:"-"`
	Bits         int        `json:"-"`
	CmPerPx      float64    `json:"-"`
	ProjMult     float64    `json:"-"`
	ImaxInFlt    float64    `json:"-"`
	ImaxInInt    uint16     `json:"-"`
	ImaxOutFlt   float64    `json:"-"`
	ImaxOutInt   uint16     `json:"-"`
	IthreshInt   uint16     `json:"-"`
	PxGapMin     int        `json:"-"`
	PxGapMax     int        `json:"-"`
	Nmass        float64    `json:"-"`
	MassTable    []float64  `json:"-"`
	FilterSteps  int        `json:"-"`
	MaxTheta     float64    `json:"-"`
	Xd           []float64  `json:"-"`
	Yd           []float64  `json:"-"`
	Xc           float64    `json:"-"`
	Yc           float64    `json:"-"`
	MurhotTable  []float64  `json:"-"`
	Tref         float64    `json:"-"`
	Tmin         float64    `json:"-"`
	Omin         float64    `json:"-"`
	Omax         float64    `json:"-"`
	Lstep        float64    `json:"-"`
	IcontTable   []uint16   `json:"-"`
	BorderPx     int        `json:"-"`
	ScaleWidth   float64    `json:"-"`
	RoiWidth     float64    `json:"-"`
	Iscale       [][]uint16 `json:"-"`
}
