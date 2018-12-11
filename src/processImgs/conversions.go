package processImgs

import ()

func cmCoreToPx(proc *ImgProcessor, cm float64) int {
	return int((cm * proc.ProjMult) / proc.CmPerPx)
}

func pxToCmCore(proc *ImgProcessor, px int) float64 {
    return (float64(px) * proc.CmPerPx) / proc.ProjMult
}
