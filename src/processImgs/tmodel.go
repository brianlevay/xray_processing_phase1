package processImgs

import (
	"math"
)

func Tmodel(proc *ImgProcessor, i int, j int, height int, width int, theta float64, offset float64) float64 {
	r := (proc.CoreDiameter / 2.0)
	xc := (float64(width) * proc.CmPx) / 2.0
	yc := (float64(height) * proc.CmPx) / 2.0
	axis := []float64{(xc + offset), yc, (proc.CoreHeight + r)}
	src := []float64{xc, yc, proc.SrcHeight}
	thetaRad := radians(theta)
	axisR := rotate(axis, thetaRad)
	srcR := rotate(src, thetaRad)

	x := float64(j)*proc.CmPx + (proc.CmPx / 2.0)
	y := float64(i)*proc.CmPx + (proc.CmPx / 2.0)
	pt := []float64{x, y, 0}
	ptR := rotate(pt, thetaRad)
	tmodel := thickness(ptR, axisR, srcR, r, proc.CoreType)
	return tmodel
}

func radians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

func degrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

func rotate(pt []float64, thetaR float64) []float64 {
	xr := pt[0]*math.Cos(thetaR) - pt[1]*math.Sin(thetaR)
	yr := pt[0]*math.Sin(thetaR) - pt[1]*math.Cos(thetaR)
	zr := pt[2]
	return []float64{xr, yr, zr}
}

func unitVector(ptR []float64, srcR []float64) []float64 {
	delX := ptR[0] - srcR[0]
	delY := ptR[1] - srcR[1]
	delZ := ptR[2] - srcR[2]
	dist := math.Sqrt(math.Pow(delX, 2) + math.Pow(delY, 2) + math.Pow(delZ, 2))
	if dist == 0.0 {
		return []float64{0.0, 0.0, 0.0}
	}
	ux := delX / dist
	uy := delY / dist
	uz := delZ / dist
	return []float64{ux, uy, uz}
}

func thickness(ptR []float64, axisR []float64, srcR []float64, r float64, coreType string) float64 {
	unitR := unitVector(ptR, srcR)
	if unitR[2] >= 0.0 {
		return 0.0
	}
	A := math.Pow(unitR[0], 2) + math.Pow(unitR[2], 2)
	B := 2*unitR[0]*(srcR[0]-axisR[0]) + 2*unitR[2]*(srcR[2]-axisR[2])
	C := math.Pow(srcR[0], 2) - 2*srcR[0]*axisR[0] + math.Pow(axisR[0], 2) + math.Pow(srcR[2], 2) - 2*srcR[2]*axisR[2] + math.Pow(axisR[2], 2) - math.Pow(r, 2)
	det := math.Pow(B, 2) - 4*A*C
	if det <= 0.0 {
		return 0.0
	}
	tc1 := (-B - math.Sqrt(det)) / (2 * A)
	tc2 := (-B + math.Sqrt(det)) / (2 * A)
	th := (axisR[2] - srcR[2]) / unitR[2]
	if (tc1 <= 0.0) || (tc2 <= 0.0) || (th <= 0.0) {
		return 0.0
	}
	if coreType == "HR" {
		if th < tc1 {
			return tc2 - tc1
		} else if (tc1 < th) && (th < tc2) {
			return tc2 - th
		} else {
			return 0.0
		}
	}
	return tc2 - tc1
}
