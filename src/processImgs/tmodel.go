package processImgs

import (
	"math"
)

func TModel(proc *ImgProcessor, theta float64, offset float64) [][]float64 {
	var xrp, yrp, delX, delY, dist, ux, uz, th, A, B, det, tc1, tc2 float64

	thetaR := theta * (math.Pi / 180.0)
	cos0 := math.Cos(thetaR)
	sin0 := math.Sin(thetaR)

	r := (proc.CoreDiameter / 2.0)
	xra := (proc.Xc+offset)*cos0 - proc.Yc*sin0
	zra := (proc.CoreHeight + r)
	xrs := proc.Xc*cos0 - proc.Yc*sin0
	yrs := proc.Xc*sin0 - proc.Yc*cos0
	zrs := proc.SrcHeight
	delZ := (0.0 - zrs)
	delZsq := delZ * delZ
	C := xrs*xrs - 2*xrs*xra + xra*xra + zrs*zrs - 2*zrs*zra + zra*zra - r*r

	tmodel := make([][]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		tmodel[i] = make([]float64, proc.Width)
		for j := 0; j < proc.Width; j++ {
			xrp = proc.X[j]*cos0 - proc.Y[i]*sin0
			yrp = proc.X[j]*sin0 - proc.Y[i]*cos0
			delX = xrp - xrs
			delY = yrp - yrs
			dist = math.Max(math.Sqrt((delX*delX)+(delY*delY)+delZsq), 0.1)
			ux = delX / dist
			uz = delZ / dist

			th = (zra - zrs) / uz
			A = ux*ux + uz*uz
			B = 2*ux*(xrs-xra) + 2*uz*(zrs-zra)
			det = B*B - 4*A*C

			if det <= 0.0 {
				tmodel[i][j] = 0.0
			} else {
				tc1 = (-B - math.Sqrt(det)) / (2 * A)
				tc2 = (-B + math.Sqrt(det)) / (2 * A)
				if proc.CoreType == "HR" {
					if th < tc1 {
						tmodel[i][j] = tc2 - tc1
					} else if (tc1 < th) && (th < tc2) {
						tmodel[i][j] = tc2 - th
					} else {
						tmodel[i][j] = 0.0
					}
				} else {
					tmodel[i][j] = tc2 - tc1
				}
			}
		}
	}
	return tmodel
}
