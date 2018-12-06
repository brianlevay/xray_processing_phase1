package processImgs

import (
	"math"
)

func TModel(proc *ImgProcessor, theta float64, offset float64) [][]float64 {
	var xrp, yrp, zrp, delX, delY, delZ, dist, ux, uz, th, A, B float64
	r := (proc.CoreDiameter / 2.0)
	thetaRad := theta * (math.Pi / 180.0)
	xra, _, zra := rotate((proc.Xc + offset), proc.Yc, (proc.CoreHeight + r), thetaRad)
	xrs, yrs, zrs := rotate(proc.Xc, proc.Yc, proc.SrcHeight, thetaRad)
	C := math.Pow(xrs, 2) - 2*xrs*xra + math.Pow(xra, 2) + math.Pow(zrs, 2) - 2*zrs*zra + math.Pow(zra, 2) - math.Pow(r, 2)

	tmodel := make([][]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		tmodel[i] = make([]float64, proc.Width)
		for j := 0; j < proc.Width; j++ {
			xrp, yrp, zrp = rotate(proc.X[j], proc.Y[i], 0.0, thetaRad)
			delX = xrp - xrs
			delY = yrp - yrs
			delZ = zrp - zrs
			dist = math.Max(math.Sqrt(math.Pow(delX, 2)+math.Pow(delY, 2)+math.Pow(delZ, 2)), 0.1)
			ux = delX / dist
			uz = delZ / dist
			th = (zra - zrs) / uz
			A = math.Pow(ux, 2) + math.Pow(uz, 2)
			B = 2*ux*(xrs-xra) + 2*uz*(zrs-zra)
			tmodel[i][j] = thickness(proc, th, A, B, C)
		}
	}
	return tmodel
}

func thickness(proc *ImgProcessor, th float64, A float64, B float64, C float64) float64 {
	det := math.Pow(B, 2) - 4*A*C
	if det <= 0.0 {
		return 0.0
	}
	tc1 := (-B - math.Sqrt(det)) / (2 * A)
	tc2 := (-B + math.Sqrt(det)) / (2 * A)
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
