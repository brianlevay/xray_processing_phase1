package processImgs

import ()

type ImgProcessor struct {
	CoreType     string  `json:"CoreType"`
	CoreDiameter float64 `json:"CoreDiameter"`
	SrcHeight    float64 `json:"SrcHeight"`
	CoreHeight   float64 `json:"CoreHeight"`
	Motion       float64 `json:"Motion"`
	AxisMethod   string  `json:"AxisMethod"`
	AxisAngle    float64 `json:"AxisAngle"`
	AxisOffset   float64 `json:"AxisOffset"`
	Ilow         float64 `json:"Ilow"`
	Ipeak        float64 `json:"Ipeak"`
	Ihigh        float64 `json:"Ihigh"`
	FolderName   string  `json:"FolderName"`
	FileAppend   string  `json:"FileAppend"`
	Bits         int     `json:"-"`
	Imax         float64 `json:"-"`
	Pxcm         float64 `json:"-"`
	Tmin         float64 `json:"-"`
	Omin         float64 `json:"-"`
	Opeak        float64 `json:"-"`
	Omax         float64 `json:"-"`
	Xpeak        float64 `json:"-"`
	N            float64 `json:"-"`
	W            float64 `json:"-"`
}
