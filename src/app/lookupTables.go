package main

import (
	"math"
)

func (proc *ImgProcessor) CalculateMassTable() {
	var massBase float64
	nvals := int(proc.ImaxInInt + 1)
	masses := make([]float64, nvals)
	for k := 0; k < nvals; k++ {
		massBase = float64(proc.ImaxInInt-uint16(k)) / float64(proc.ImaxInFlt)
		masses[k] = math.Pow(massBase, proc.Cfg.Nmass)
	}
	proc.MassTable = masses
	return
}

func (proc *ImgProcessor) CalculateXYd() {
	Yc := (float64(proc.Cfg.HeightPxDet) * proc.Cfg.CmPerPx) / 2.0
	Xc := (float64(proc.Cfg.WidthPxDet) * proc.Cfg.CmPerPx) / 2.0
	Yd := make([]float64, proc.Cfg.HeightPxDet)
	for i := 0; i < proc.Cfg.HeightPxDet; i++ {
		Yd[i] = float64(i)*proc.Cfg.CmPerPx + (proc.Cfg.CmPerPx / 2.0)
	}
	Xd := make([]float64, proc.Cfg.WidthPxDet)
	for j := 0; j < proc.Cfg.WidthPxDet; j++ {
		Xd[j] = float64(j)*proc.Cfg.CmPerPx + (proc.Cfg.CmPerPx / 2.0)
	}
	proc.Xc = Xc
	proc.Yc = Yc
	proc.Xd = Xd
	proc.Yd = Yd
	return
}

func (proc *ImgProcessor) CalculateMurhotTable() {
	nvals := int(proc.ImaxInInt + 1)
	murhots := make([]float64, nvals)
	for k := 0; k < nvals; k++ {
		murhots[k] = math.Log(proc.ImaxInFlt+1.0) - math.Log(float64(k)+1.0)
	}
	proc.MurhotTable = murhots
	return
}

func (proc *ImgProcessor) CalculateIcontTable() {
	var L, P, SF, SP, Y float64
	Lpeak := (proc.Opeak - proc.Omin) / (proc.Omax - proc.Omin)
	N := math.Log(0.5) / math.Log(Lpeak)
	W := 1.0 - math.Pow(math.Abs(2.0*(Lpeak-0.5)), 2.0)
	nvals := int((1.0-0.0)/proc.Lstep) + 1

	Icont := make([]uint16, nvals)
	for k := 0; k < nvals; k++ {
		L = float64(k) * proc.Lstep
		P = math.Pow(L, N)
		SF = 0.5*math.Sin(math.Pi*(L-0.5)) + 0.5
		SP = 0.5*math.Sin(math.Pi*(P-0.5)) + 0.5
		Y = W*SP + (1.0-W)*SF
		Icont[k] = uint16(proc.ImaxOutFlt * (1.0 - Y))
	}
	proc.IcontTable = Icont
	return
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
	return
}
