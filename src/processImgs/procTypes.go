package processImgs

import ()

type ImgProcessor struct {
	Low          float64 `json:"Low"`
	Mid          float64 `json:"Mid"`
	High         float64 `json:"High"`
	SrcHeight    float64 `json:"SrcHeight"`
	CoreHeight   float64 `json:"CoreHeight"`
	CoreDiameter float64 `json:"CoreDiameter"`
	Motion       float64 `json:"Motion"`
	CoreType     string  `json:"CoreType"`
	Contrast     string  `json:"Contrast"`
	FolderName   string  `json:"FolderName"`
	FileAppend   string  `json:"FileAppend"`
}
