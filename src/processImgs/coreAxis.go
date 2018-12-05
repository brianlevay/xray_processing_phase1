package processImgs

import (
	"math"
)

func FindCoreAxis(proc *ImgProcessor, Iraw [][]float64) (float64, float64) {
	height := len(Iraw)
	width := len(Iraw[0])

	Ithresh := 0.8 * proc.Imax
	maxTheta := 5.0
	minWidthPx := int(0.8 * (proc.CoreDiameter / proc.CmPx))
	minPts := int(0.5 * float64(height))

	var edgeVals, edgeJs []int
	iMid := []int{}
	jMid := []int{}
	for i := 0; i < height; i++ {
		edgeVals = []int{}
		edgeJs = []int{}
		for j := 0; j < (width - 1); j++ {
			if (Iraw[i][j] > Ithresh) && (Iraw[i][j+1] <= Ithresh) {
				edgeVals = append(edgeVals, 1)
				edgeJs = append(edgeJs, j)
			}
			if (Iraw[i][j] <= Ithresh) && (Iraw[i][j+1] > Ithresh) {
				edgeVals = append(edgeVals, -1)
				edgeJs = append(edgeJs, j)
			}
		}
		jLeft, jRight := farthestEdges(edgeVals, edgeJs)
		if (jRight - jLeft) >= minWidthPx {
			iMid = append(iMid, i)
			jMid = append(jMid, int(float64(jRight+jLeft)/2.0))
		}
	}
	theta, offset := axisFit(proc, minPts, iMid, jMid)
	if math.Abs(theta) <= maxTheta {
		return theta, offset
	} else {
		return 0.0, 0.0
	}
}

func farthestEdges(edgeVals []int, edgeJs []int) (int, int) {
	delta, maxDelta, jLeft, jRight := 0, 0, 0, 0
	nEdges := len(edgeVals)
	for k := 0; k < (nEdges - 1); k++ {
		if (edgeVals[k] == 1) && (edgeVals[k+1] == -1) {
			delta = edgeJs[k+1] - edgeJs[k]
			if delta > maxDelta {
				maxDelta = delta
				jLeft = edgeJs[k]
				jRight = edgeJs[k+1]
			}
		}
	}
	return jLeft, jRight
}

func axisFit(proc *ImgProcessor, minPts int, iMid []int, jMid []int) (float64, float64) {
	nPts := len(iMid)
	if nPts < minPts {
		return 0.0, 0.0
	}
	//beta, alpha := linearRegressionInts(iMid, jMid)
	return 0.0, 0.0
}

func linearRegressionInts(xVals []int, yVals []int) (float64, float64) {
	xsum, ysum, xxsum, xysum := 0.0, 0.0, 0.0, 0.0
	nPts := len(xVals)
	nFlt := float64(nPts)
	for k := 0; k < nPts; k++ {
		xsum += float64(xVals[k])
		ysum += float64(yVals[k])
		xxsum += (float64(xVals[k]) * float64(xVals[k]))
		xysum += (float64(xVals[k]) * float64(yVals[k]))
	}
	beta := (xysum - (1/nFlt)*xsum*ysum) / (xxsum - (1/nFlt)*xsum*xsum)
	xave := xsum / nFlt
	yave := ysum / nFlt
	alpha := yave - beta*xave
	return beta, alpha
}
