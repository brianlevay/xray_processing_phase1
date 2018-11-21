package processImgs

import ()

type ImgProcessor struct {
	Bits         float64 `json:"Bits"`
	Low          float64 `json:"Low"`
	Mid          float64 `json:"Mid"`
	High         float64 `json:"High"`
	Method       string  `json:"Method"`
	SrcHeight    float64 `json:"SrcHeight"`
	CoreDiameter float64 `json:"CoreDiameter"`
	CoreType     string  `json:"CoreType"`
	Contrast     string  `json:"Contrast"`
	ROISize      float64 `json:"ROISize"`
	IncludeScale bool    `json:"IncludeScale"`
	IncludeROI   bool    `json:"IncludeROI"`
	FolderName   string  `json:"FolderName"`
	FileAppend   string  `json:"FileAppend"`
}
