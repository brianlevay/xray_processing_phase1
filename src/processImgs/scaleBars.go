package processImgs

import (
	"math"
)

func (proc *ImgProcessor) CreateScaleBars() {
	cmPxHeight := cmCoreToPx(proc, 1.0)
	cmPxWidth := cmCoreToPx(proc, proc.ScaleWidth)
	roiPxHeight := cmCoreToPx(proc, proc.Motion)
	roiPxWidth := cmCoreToPx(proc, proc.RoiWidth)
	roiIstart := int(float64(proc.Height-roiPxHeight) / 2.0)
	iBorder := []int{0, (proc.Height - 1)}
	jBorder := []int{0, (proc.BorderPx + cmPxWidth + proc.BorderPx - 1)}
	iCms := []int{(iBorder[0] + proc.BorderPx), (iBorder[1] - proc.BorderPx)}
	jCms := []int{(jBorder[0] + proc.BorderPx), (jBorder[1] - proc.BorderPx)}
	iLn := []int{roiIstart, (roiIstart + roiPxHeight - 1)}
	jLn := []int{(jBorder[1] + 1), (jBorder[1] + roiPxWidth)}

	Iscale := make([][]uint16, proc.Height)
	for i := 0; i < proc.Height; i++ {
		Iscale[i] = make([]uint16, proc.Width)
		for j := 0; j < proc.Width; j++ {
			Iscale[i][j] = 1
			if isInside(i, j, iBorder, jBorder) {
				Iscale[i][j] = 0
			}
			if isInside(i, j, iCms, jCms) {
				if evenCm(i, cmPxHeight) {
					Iscale[i][j] = 0
				} else {
					Iscale[i][j] = 2
				}
			}
			if isInside(i, j, iLn, jLn) {
				Iscale[i][j] = 0
			}
		}
	}
	proc.Iscale = Iscale
	return
}

func isInside(i int, j int, iBounds []int, jBounds []int) bool {
	if (i >= iBounds[0]) && (i <= iBounds[1]) && (j >= jBounds[0]) && (j <= jBounds[1]) {
		return true
	} else {
		return false
	}
}

func evenCm(i int, cmPxHeight int) bool {
	x := math.Floor(float64(i) / float64(cmPxHeight))
	rem := math.Mod(x, 2.0)
	if rem == 0.0 {
		return true
	} else {
		return false
	}
}
