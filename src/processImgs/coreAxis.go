package processImgs

import (
	"math"
)

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	Ithresh := uint16(0.8 * proc.ImaxInFlt)
	maxTheta := 5.0
	minWidth := 0.8 * proc.CoreDiameter
	minPts := int(0.5 * float64(proc.Height))

	var leftEdge, rightEdge, leftMax, rightMax, maxGap float64
	nPts := 0
	Xmid := make([]float64, proc.Height)
	Ymid := make([]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		leftEdge, rightEdge, leftMax, rightMax, maxGap = 0.0, 0.0, 0.0, 0.0, 0.0
		for j := 0; j < (proc.Width - 1); j++ {
			if (Iraw[i][j] > Ithresh) && (Iraw[i][j+1] <= Ithresh) {
				leftEdge = proc.X[j]
			}
			if (Iraw[i][j] <= Ithresh) && (Iraw[i][j+1] > Ithresh) {
				rightEdge = proc.X[j+1]
			}
			if ((rightEdge - leftEdge) >= maxGap) && ((rightEdge - leftEdge) >= minWidth) {
				maxGap = (rightEdge - leftEdge)
				leftMax = leftEdge
				rightMax = rightEdge
			}
		}
		if maxGap > 0.0 {
			nPts += 1
			Xmid[i] = (leftMax + rightMax) / 2.0
			Ymid[i] = proc.Y[i]
		} else {
			Xmid[i] = -1.0
			Ymid[i] = -1.0
		}
	}
	if nPts < minPts {
		return 0.0, 0.0
	}
	theta, offset := axisFit(proc, Xmid, Ymid)
	if math.Abs(theta) > maxTheta {
		return 0.0, 0.0
	}
	return theta, offset
}

func axisFit(proc *ImgProcessor, Xmid []float64, Ymid []float64) (float64, float64) {
	// Regression using the Y values (i) as the independent variable //
	x := Ymid
	y := Xmid

	nFlt, xsum, ysum, xxsum, xysum := 0.0, 0.0, 0.0, 0.0, 0.0
	for i := 0; i < proc.Height; i++ {
		if y[i] != -1.0 {
			nFlt += 1
			xsum += x[i]
			ysum += y[i]
			xxsum += x[i] * x[i]
			xysum += x[i] * y[i]
		}
	}
	beta := (xysum - (1/nFlt)*xsum*ysum) / (xxsum - (1/nFlt)*xsum*xsum)
	xave := xsum / nFlt
	yave := ysum / nFlt
	alpha := yave - beta*xave

	theta := math.Atan(beta) * (180.0 / math.Pi)
	Xline := beta*proc.Yc + alpha
	Xoffset := Xline - proc.Xc
	return theta, Xoffset
}
