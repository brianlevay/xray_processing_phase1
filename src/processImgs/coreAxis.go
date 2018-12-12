package processImgs

import (
	"math"
)

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	// Configuration Variables //
	maxXdist := 3.0

	var leftEdge, rightEdge, leftMax, rightMax, maxGap int
	var nsum, xsum, ysum, xxsum, xysum, Xmid, Ymid float64

	for i := 0; i < proc.Height; i++ {
		leftEdge, rightEdge, leftMax, rightMax, maxGap = 0, 0, 0, 0, 0
		for j := 0; j < (proc.Width - 1); j++ {
			if (Iraw[i][j] > proc.IthreshInt) && (Iraw[i][j+1] <= proc.IthreshInt) {
				leftEdge = j
			}
			if (Iraw[i][j] <= proc.IthreshInt) && (Iraw[i][j+1] > proc.IthreshInt) {
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
			nsum += 1.0
			xsum += Ymid
			ysum += Xmid
			xxsum += (Ymid * Ymid)
			xysum += (Ymid * Xmid)
		}
	}
	if nsum == 0.0 {
		return 0.0, 0.0
	}
	beta := (xysum - (1.0/nsum)*xsum*ysum) / (xxsum - (1.0/nsum)*xsum*xsum)
	xave := xsum / nsum
	yave := ysum / nsum
	alpha := yave - beta*xave

	offsetProj := (beta*proc.Yc + alpha) - proc.Xc
	offsetAct := offsetProj / proc.ProjMult
	theta := math.Atan(beta) * (180.0 / math.Pi)
	if math.Abs(theta) > proc.MaxTheta {
		return 0.0, 0.0
	} else {
		return theta, offsetAct
	}
}
