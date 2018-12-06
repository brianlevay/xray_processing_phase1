package processImgs

import (
	"math"
)

func Tmodel(proc *ImgProcessor, i int, j int, theta float64, offset float64) float64 {
	r := (proc.CoreDiameter / 2.0)
	xc := (float64(proc.Width) * proc.CmPx) / 2.0
	yc := (float64(proc.Height) * proc.CmPx) / 2.0
	thetaRad := theta * (math.Pi / 180.0)
	xra, _, zra := rotate((xc + offset), yc, (proc.CoreHeight + r), thetaRad)
	xrs, yrs, zrs := rotate(xc, yc, proc.SrcHeight, thetaRad)

	x := float64(j)*proc.CmPx + (proc.CmPx / 2.0)
	y := float64(i)*proc.CmPx + (proc.CmPx / 2.0)
	xrp, yrp, zrp := rotate(x, y, 0.0, thetaRad)

	delX := xrp - xrs
	delY := yrp - yrs
	delZ := zrp - zrs
	dist := math.Sqrt(math.Pow(delX, 2) + math.Pow(delY, 2) + math.Pow(delZ, 2))
	if dist == 0.0 {
		return 0.0
	}
	ux := delX / dist
	uz := delZ / dist

	A := math.Pow(ux, 2) + math.Pow(uz, 2)
	B := 2*ux*(xrs-xra) + 2*uz*(zrs-zra)
	C := math.Pow(xrs, 2) - 2*xrs*xra + math.Pow(xra, 2) + math.Pow(zrs, 2) - 2*zrs*zra + math.Pow(zra, 2) - math.Pow(r, 2)
	det := math.Pow(B, 2) - 4*A*C
	if det <= 0.0 {
		return 0.0
	}
	tc1 := (-B - math.Sqrt(det)) / (2 * A)
	tc2 := (-B + math.Sqrt(det)) / (2 * A)
	th := (zra - zrs) / uz
	if (tc1 <= 0.0) || (tc2 <= 0.0) || (th <= 0.0) {
		return 0.0
	}
	if proc.CoreType == "HR" {
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

func rotate(x float64, y float64, z float64, thetaR float64) (float64, float64, float64) {
	xr := x*math.Cos(thetaR) - y*math.Sin(thetaR)
	yr := x*math.Sin(thetaR) - y*math.Cos(thetaR)
	zr := z
	return xr, yr, zr
}
