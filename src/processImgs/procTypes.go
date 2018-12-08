package processImgs

import ()

type ImgProcessor struct {
	Height       int        `json:"Height"`
	Width        int        `json:"Width"`
	Bits         int        `json:"Bits"`
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
	ImaxInFlt    float64    `json:"-"`
	ImaxInInt    uint16     `json:"-"`
	ImaxOutFlt   float64    `json:"-"`
	ImaxOutInt   uint16     `json:"-"`
	CmPx         float64    `json:"-"`
	X            []float64  `json:"-"`
	Y            []float64  `json:"-"`
	Xc           float64    `json:"-"`
	Yc           float64    `json:"-"`
	Tref         float64    `json:"-"`
	Tmin         float64    `json:"-"`
	Omin         float64    `json:"-"`
	Omax         float64    `json:"-"`
	Lstep        float64    `json:"-"`
	MurhotTable  []float64  `json:"-"`
	IcontTable   []uint16   `json:"-"`
	Iscale       [][]uint16 `json:"-"`
}
