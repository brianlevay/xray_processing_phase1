package processImgs

import (
	"math"
	"sort"
)

// Regression using the Y values (i) as the independent variable //

func FindCoreAxis(proc *ImgProcessor, Iraw [][]uint16) (float64, float64) {
	flag := -999.0
	ColMid, RowMid := centerOfMassBetweenEdges(proc, Iraw, flag)
	beta, alpha := medianOfSegments(proc, RowMid, ColMid, flag)
	offsetProj := (beta*proc.Yc + alpha) - proc.Xc
	offsetAct := offsetProj / proc.ProjMult
	theta := math.Atan(beta) * (180.0 / math.Pi)
	if theta > proc.MaxTheta {
		return proc.MaxTheta, offsetAct
	} else if theta < -proc.MaxTheta {
		return -proc.MaxTheta, offsetAct
	} else {
		return theta, offsetAct
	}
}

// This can be written to touch each pixel only once, but I'm not sure if it's worth summing
// masses in intervals that will be discarded //

func centerOfMassBetweenEdges(proc *ImgProcessor, Iraw [][]uint16, flag float64) ([]float64, []float64) {
	var leftEdge, rightEdge, leftMax, rightMax, largestGap int
	var mass, msum, xmsum float64
	RowMid := make([]float64, proc.Height)
	ColMid := make([]float64, proc.Height)

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
			ColMid[i] = xmsum / msum
		} else {
			ColMid[i] = flag
		}
		RowMid[i] = proc.Yd[i]
	}
	return ColMid, RowMid
}

func medianOfSegments(proc *ImgProcessor, X []float64, Y []float64, flag float64) (float64, float64) {
	var start, stop int
	var betaSeg, alphaSeg float64
	stepSize := int(float64(proc.Height) / float64(proc.Segments))
	betas := make([]float64, 0, proc.Segments)
	alphas := make([]float64, 0, proc.Segments)
	for k := 0; k < proc.Segments; k++ {
		start = k * stepSize
		stop = (k + 1) * stepSize
		if stop > proc.Height {
			stop = proc.Height
		}
		betaSeg, alphaSeg = linearRegression(X[start:stop], Y[start:stop], flag)
		if betaSeg != flag {
			betas = append(betas, betaSeg)
			alphas = append(alphas, alphaSeg)
		}
	}
	betaMedian := findMedianFlt(betas)
	alphaMedian := findMedianFlt(alphas)
	return betaMedian, alphaMedian
}

func linearRegression(X []float64, Y []float64, flag float64) (float64, float64) {
	var nsum, xsum, ysum, xxsum, xysum float64
	beta := flag
	alpha := flag
	nPts := len(X)
	for i := 0; i < nPts; i++ {
		if Y[i] != flag {
			nsum += 1.0
			xsum += X[i]
			ysum += Y[i]
			xxsum += X[i] * X[i]
			xysum += X[i] * Y[i]
		}
	}
	// Regression doesn't have any significant meaning for <3 points //
	if nsum >= 3.0 {
		xave := xsum / nsum
		yave := ysum / nsum
		beta = (xysum - (1.0/nsum)*xsum*ysum) / (xxsum - (1.0/nsum)*xsum*xsum)
		alpha = yave - beta*xave
	}
	return beta, alpha
}

func findMedianFlt(X []float64) float64 {
	sort.Float64s(X)
	nPts := len(X)
	midInt := int(float64(nPts-1) / 2.0)
	if (nPts - 1 - midInt) > (midInt - 0) {
		return (X[midInt] + X[midInt+1]) / 2.0
	}
	return X[midInt]
}
