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
	Contrast     string  `json:"Contrast"`
	Low          float64 `json:"Low"`
	Mid          float64 `json:"Mid"`
	High         float64 `json:"High"`
	FolderName   string  `json:"FolderName"`
	FileAppend   string  `json:"FileAppend"`
}
