package processImgs

import (
	"math"
)

// Regression using the Y values (i) as the independent variable //

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	Xmid, Ymid := centerOfMassBetweenEdges(proc, Iraw)
	beta, alpha := linearRegression(Ymid, Xmid)
	offsetProj := (beta*proc.Yc + alpha) - proc.Xc
	offsetAct := offsetProj / proc.ProjMult
	theta := math.Atan(beta) * (180.0 / math.Pi)
	if math.Abs(theta) > proc.MaxTheta {
		return 0.0, offsetAct
	} else {
		return theta, offsetAct
	}
}

// This can be written to touch each pixel only once, but I'm not sure if it's worth summing
// masses in intervals that will be discarded //

func centerOfMassBetweenEdges(proc *ImgProcessor, Iraw [][]uint16) ([]float64, []float64) {
	var leftEdge, rightEdge, leftMax, rightMax, largestGap int
	var mass, msum, xmsum float64
	Xmid := make([]float64, proc.Height)
	Ymid := make([]float64, proc.Height)

	for i := 0; i < proc.Height; i++ {
		leftEdge, rightEdge, leftMax, rightMax, largestGap = 0, 0, 0, 0, 0
		for j := 0; j < (proc.Width - 1); j++ {
			if (Iraw[i][j] > proc.IthreshInt) && (Iraw[i][j+1] <= proc.IthreshInt) {
				leftEdge = j
			}
			if (Iraw[i][j] <= proc.IthreshInt) && (Iraw[i][j+1] > proc.IthreshInt) {
				rightEdge = j + 1
			}
			if (rightEdge - leftEdge) >= largestGap {
				largestGap = (rightEdge - leftEdge)
				leftMax = leftEdge
				rightMax = rightEdge
			}
		}
		msum, xmsum = 0.0, 0.0
		if (largestGap >= proc.PxGapMin) && (largestGap <= proc.PxGapMax) {
			for j := leftMax; j < (rightMax + 1); j++ {
				mass = proc.MassTable[Iraw[i][j]]
				msum += mass
				xmsum += proc.Xd[j] * mass
			}
			Xmid[i] = xmsum / msum
			Ymid[i] = proc.Yd[i]
		} else {
			Xmid[i] = -1.0
			Ymid[i] = -1.0
		}
	}
	return Xmid, Ymid
}

func linearRegression(X []float64, Y []float64) (float64, float64) {
	var nsum, xsum, ysum, xxsum, xysum float64
	nPts := len(X)
	for i := 0; i < nPts; i++ {
		if Y[i] != -1.0 {
			nsum += 1.0
			xsum += X[i]
			ysum += Y[i]
			xxsum += X[i] * X[i]
			xysum += X[i] * Y[i]
		}
	}
	beta := (xysum - (1.0/nsum)*xsum*ysum) / (xxsum - (1.0/nsum)*xsum*xsum)
	xave := xsum / nsum
	yave := ysum / nsum
	alpha := yave - beta*xave
	return beta, alpha
}
