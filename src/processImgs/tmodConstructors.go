package processImgs

import (
	"math"
)

func newTModel(proc *ImgProcessor, theta float64, offset float64) *TModel {
	t := new(TModel)
	t.CoreType = proc.CoreType
	t.R = (proc.CoreDiameter / 2.0)
	thetaR := theta * (math.Pi / 180.0)
	t.Cos0 = math.Cos(thetaR)
	t.Sin0 = math.Sin(thetaR)
	t.Xra = (proc.Xc+offset)*t.Cos0 - proc.Yc*t.Sin0
	t.Yra = proc.Xc*t.Sin0 - proc.Yc*t.Cos0
	t.Zra = (proc.CoreHeight + t.R)
	t.Xrs = proc.Xc*t.Cos0 - proc.Yc*t.Sin0
	t.Yrs = proc.Xc*t.Sin0 - proc.Yc*t.Cos0
	t.Zrs = proc.SrcHeight
	t.DelZr = (0.0 - t.Zrs)
	t.DelZr2 = t.DelZr * t.DelZr
	t.C = t.Xrs*t.Xrs - 2*t.Xrs*t.Xra + t.Xra*t.Xra
	t.C += t.Zrs*t.Zrs - 2*t.Zrs*t.Zra + t.Zra*t.Zra - t.R*t.R
	t.XrStep = proc.CmPerPx / 2.0
	t.XrMin = t.DelZr*((t.Xra-t.R-t.Xrs)/(t.Zra+t.R-t.Zrs)) + t.Xrs
	t.XrMax = t.DelZr*((t.Xra+t.R-t.Xrs)/(t.Zra+t.R-t.Zrs)) + t.Xrs
	t.CalculateTxz2Table()
	return t
}

func (t *TModel) CalculateTxz2Table() {
	var Txz, Xrd, DelXr, dist, uXr, uZr, th, A, B, det, tc1, tc2 float64
	nVals := int((t.XrMax-t.XrMin)/t.XrStep) + 1

	t.Txz2Table = make([]float64, nVals)
	for k := 0; k < nVals; k++ {
		Txz = 0.0
		Xrd = float64(k)*t.XrStep + t.XrMin
		DelXr = Xrd - t.Xrs
		dist = math.Max(math.Sqrt((DelXr*DelXr)+t.DelZr2), 0.1)
		uXr = DelXr / dist
		uZr = t.DelZr / dist
		th = (t.Zra - t.Zrs) / uZr
		A = uXr*uXr + uZr*uZr
		B = 2*uXr*(t.Xrs-t.Xra) + 2*uZr*(t.Zrs-t.Zra)
		det = B*B - 4*A*t.C

		if det > 0.0 {
			tc1 = (-B - math.Sqrt(det)) / (2 * A)
			tc2 = (-B + math.Sqrt(det)) / (2 * A)
			if t.CoreType == "HR" {
				if th < tc1 {
					Txz = tc2 - tc1
				} else if (tc1 < th) && (th < tc2) {
					Txz = tc2 - th
				} else {
					Txz = 0.0
				}
			} else {
				Txz = tc2 - tc1
			}
		}
		t.Txz2Table[k] = Txz * Txz
	}
}
