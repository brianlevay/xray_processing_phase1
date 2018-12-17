package main

import (
	"strings"
)

func stringToSlice(valString string) []string {
	replacer := strings.NewReplacer("[", "", "]", "", "\"", "")
	cleaned := replacer.Replace(valString)
	values := strings.Split(cleaned, ",")
	if (len(values) == 1) && (strings.Compare(values[0], "") == 0) {
		return []string{}
	}
	return values
}

func cmCoreToPx(proc *ImgProcessor, cm float64) int {
	return int((cm * proc.ProjMult) / proc.Cfg.CmPerPx)
}

func pxToCmCore(proc *ImgProcessor, px int) float64 {
	return (float64(px) * proc.Cfg.CmPerPx) / proc.ProjMult
}
