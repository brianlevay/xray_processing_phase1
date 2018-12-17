package processImgs

import ()

type ImgProcessor struct {
	CoreType     string     `json:"CoreType"`
	CoreDiameter float64    `json:"CoreDiameter"`
	AxisMethod   string     `json:"AxisMethod"`
	AxisAngle    float64    `json:"AxisAngle"`
	AxisOffset   float64    `json:"AxisOffset"`
	IlowFrac     float64    `json:"IlowFrac"`
	IpeakFrac    float64    `json:"IpeakFrac"`
	IhighFrac    float64    `json:"IhighFrac"`
	FolderName   string     `json:"FolderName"`
	FileAppend   string     `json:"FileAppend"`
	SrcHeight    float64    `json:"-"`
	CoreHeight   float64    `json:"-"`
	Motion       float64    `json:"-"`
	HeightPxDet  int        `json:"-"`
	WidthPxDet   int        `json:"-"`
	CmPerPx      float64    `json:"-"`
	Bits         int        `json:"-"`
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
	Tedge        float64    `json:"-"`
	Omin         float64    `json:"-"`
	Opeak        float64    `json:"-"`
	Omax         float64    `json:"-"`
	Lstep        float64    `json:"-"`
	IcontTable   []uint16   `json:"-"`
	BorderPx     int        `json:"-"`
	ScaleWidth   float64    `json:"-"`
	RoiWidth     float64    `json:"-"`
	Iscale       [][]uint16 `json:"-"`
}
