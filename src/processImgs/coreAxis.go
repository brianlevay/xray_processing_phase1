package processImgs

import (
	"math"
)

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	// Configuration Variables //
	Ithresh := uint16(0.8 * proc.ImaxInFlt)
	maxXdist := 3.0 * (proc.CmPerPxProj / proc.CmPerPxAct)
	maxTheta := 5.0

	beta, alpha := regressionFromEdges(proc, Iraw, Ithresh, maxXdist)
	offsetProj := axisAdjustment(proc, Iraw, Ithresh, beta, alpha)
	offsetAct := offsetProj * (proc.CmPerPxProj / proc.CmPerPxAct)
	theta := math.Atan(beta) * (180.0 / math.Pi)
	if math.Abs(theta) > maxTheta {
		return 0.0, 0.0
	} else {
		return theta, offsetAct
	}
}

func regressionFromEdges(proc *ImgProcessor, Iraw [][]uint16, Ithresh uint16, maxXdist float64) (float64, float64) {
	var leftEdge, rightEdge, leftMax, rightMax, maxGap int
	var w, wtsum, xsum, ysum, xxsum, xysum, Xmid, Ymid float64

	for i := 0; i < proc.Height; i++ {
		leftEdge, rightEdge, leftMax, rightMax, maxGap = 0, 0, 0, 0, 0
		for j := 0; j < (proc.Width - 1); j++ {
			if (Iraw[i][j] > Ithresh) && (Iraw[i][j+1] <= Ithresh) {
				leftEdge = j
			}
			if (Iraw[i][j] <= Ithresh) && (Iraw[i][j+1] > Ithresh) {
				rightEdge = j + 1
			}
			if (rightEdge - leftEdge) >= maxGap {
				maxGap = (rightEdge - leftEdge)
				leftMax = leftEdge
				rightMax = rightEdge
			}
		}
		Xmid = (proc.Xd[leftMax] + proc.Xd[rightMax]) / 2.0
		Ymid = proc.Yd[i]

		// Don't include values that are too far from the center of the image //
		// Regression using the Y values (i) as the independent variable //
		if (Xmid >= (proc.Xc - maxXdist)) && (Xmid <= (proc.Xc + maxXdist)) {
			w = proc.WtsGapTable[maxGap]
			wtsum += w
			xsum += w * Ymid
			ysum += w * Xmid
			xxsum += w * (Ymid * Ymid)
			xysum += w * (Ymid * Xmid)
		}
	}
	if wtsum == 0.0 {
		return 0.0, 0.0
	}
	beta := (xysum - (1/wtsum)*xsum*ysum) / (xxsum - (1/wtsum)*xsum*xsum)
	xave := xsum / wtsum
	yave := ysum / wtsum
	alpha := yave - beta*xave
	return beta, alpha
}

func axisAdjustment(proc *ImgProcessor, Iraw [][]uint16, Ithresh uint16, beta float64, alpha float64) float64 {
	var jline, jlow, jhigh, jmin int
	var Imin uint16
	var Xline, diffSum, diffCt float64

	RPx := int((0.5 * proc.CoreDiameter) / proc.CmPerPxProj)

	for i := 0; i < proc.Height; i++ {
		Xline = beta*proc.Yd[i] + alpha
		jline = int(Xline / proc.CmPerPxAct)
		Imin = Iraw[i][jline]
		if Imin <= Ithresh {
			jlow = jline - RPx
			jhigh = jline + RPx + 1
			jmin = jline
			for j := jlow; j < jhigh; j++ {
				if (Iraw[i][j] <= Ithresh) && (Iraw[j][j] < Imin) {
					jmin = j
					Imin = Iraw[i][j]
				}
			}
			diffSum += proc.Xd[jmin] - proc.Xd[jline]
			diffCt += 1.0
		}
	}
	diffAve := diffSum / diffCt
	offsetProj := (beta*proc.Yc + alpha + diffAve) - proc.Xc
	return offsetProj
}
