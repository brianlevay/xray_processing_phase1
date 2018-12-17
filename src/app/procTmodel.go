package main

import (
	"math"
)

func CalculateTModel(proc *ImgProcessor, theta float64, offset float64) [][]float64 {
	var Xrd, Yrd, DelXr, DelYr float64
	var k int
	t := newTModel(proc, theta, offset)
	nTxz2 := len(t.Txz2Table)

	tmodel := make([][]float64, proc.Cfg.HeightPxDet)
	for i := 0; i < proc.Cfg.HeightPxDet; i++ {
		tmodel[i] = make([]float64, proc.Cfg.WidthPxDet)
		for j := 0; j < proc.Cfg.WidthPxDet; j++ {
			tmodel[i][j] = 0.0
			Xrd = proc.Xd[j]*t.Cos0 - proc.Yd[i]*t.Sin0
			Yrd = proc.Xd[j]*t.Sin0 - proc.Yd[i]*t.Cos0
			DelXr = (Xrd - t.Xrs)
			DelYr = (Yrd - t.Yrs)
			k = int((Xrd - t.XrMin) / t.XrStep)
			if (k >= 0) && (k < nTxz2) {
				tmodel[i][j] = math.Sqrt(t.Txz2Table[k] * (1.0 + ((DelYr * DelYr) / (DelXr*DelXr + t.DelZr2))))
			}
		}
	}
	return tmodel
}
