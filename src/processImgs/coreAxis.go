package processImgs

import (
	"math"
)

// Regression using the Y values (i) as the independent variable //

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	flag := -999.0
	nStd := 2.0
	colMid, rowMid := centerOfMassBetweenEdges(proc, Iraw, flag)
	beta, alpha := iterativeRegression(proc, rowMid, colMid, nStd, flag)
	offset := ((beta*proc.Yc + alpha) - proc.Xc) / proc.ProjMult
	theta := math.Atan(beta) * (180.0 / math.Pi)
	if theta > proc.MaxTheta {
		return proc.MaxTheta, offset
	} else if theta < -proc.MaxTheta {
		return -proc.MaxTheta, offset
	} else {
		return theta, offset
	}
}

// This can be written to touch each pixel only once, but I'm not sure if it's worth summing
// masses in intervals that will be discarded //

func centerOfMassBetweenEdges(proc *ImgProcessor, Iraw [][]uint16, flag float64) ([]float64, []float64) {
	var leftEdge, rightEdge, leftMax, rightMax, largestGap int
	var mass, msum, xmsum float64
	rowMid := make([]float64, proc.Height)
	colMid := make([]float64, proc.Height)

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
			colMid[i] = xmsum / msum
		} else {
			colMid[i] = flag
		}
		rowMid[i] = proc.Yd[i]
	}
	return colMid, rowMid
}

func iterativeRegression(proc *ImgProcessor, X []float64, Y []float64, nStd float64, flag float64) (float64, float64) {
	beta, alpha := linearRegression(X, Y, flag)
	filterData(X, Y, beta, alpha, nStd, flag)
	for k := 0; k < proc.FilterSteps; k++ {
		beta, alpha = linearRegression(X, Y, flag)
		filterData(X, Y, beta, alpha, nStd, flag)
	}
	return beta, alpha
}

func linearRegression(X []float64, Y []float64, flag float64) (float64, float64) {
	var nsum, xsum, ysum, xxsum, xysum float64
	nPts := len(X)
	for k := 0; k < nPts; k++ {
		if Y[k] != flag {
			nsum += 1.0
			xsum += X[k]
			ysum += Y[k]
			xxsum += X[k] * X[k]
			xysum += X[k] * Y[k]
		}
	}
	if nsum < 2.0 {
		return flag, flag
	}
	xave := xsum / nsum
	yave := ysum / nsum
	covariance := (xysum - (1.0/nsum)*xsum*ysum)
	variance := (xxsum - (1.0/nsum)*xsum*xsum)
	beta := covariance / variance
	alpha := yave - beta*xave
	return beta, alpha
}

func filterData(X []float64, Y []float64, beta float64, alpha float64, nStd float64, flag float64) {
	var nsum, xsum, xxsum, res float64
	nPts := len(X)
	for k := 0; k < nPts; k++ {
		if Y[k] != flag {
			res = Y[k] - beta*X[k] - alpha
			nsum += 1.0
			xsum += res
			xxsum += res * res
		}
	}
	if nsum < 2.0 {
		return
	}
	variance := (xxsum - (1.0/nsum)*xsum*xsum)
	resStd := math.Sqrt(variance / (nsum - 1.0))
	for k := 0; k < nPts; k++ {
		if Y[k] != flag {
			res = math.Abs(Y[k] - beta*X[k] - alpha)
			if res >= nStd*resStd {
				Y[k] = flag
			}
		}
	}
	return
}
