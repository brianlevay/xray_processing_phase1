package processImgs

import (
	"math"
)

func CalculateTModel(proc *ImgProcessor, theta float64, offset float64) [][]float64 {
	var Xrd, Yrd, DelXr, DelYr float64
	var k int

	t := newTModel(proc, theta, offset)
	nTxz2 := len(t.Txz2Table)

	tmodel := make([][]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		tmodel[i] = make([]float64, proc.Width)
		for j := 0; j < proc.Width; j++ {
			tmodel[i][j] = 0.0
			Xrd = proc.X[j]*t.cos0 - proc.Y[i]*t.sin0
			Yrd = proc.X[j]*t.sin0 - proc.Y[i]*t.cos0
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
