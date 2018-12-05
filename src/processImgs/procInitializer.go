package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize() {
	bits := 14
	cmpx := 0.0099
	tmin := 0.5

	proc.Bits = bits
	proc.ImaxIn = math.Pow(2, float64(bits)) - 1.0
	proc.ImaxOut = math.Pow(2, 16.0) - 1.0
	proc.CmPx = cmpx
	proc.Tmin = tmin

	proc.Omin = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Opeak = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ipeak+1.0)
	proc.Omax = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ilow+1.0)

	proc.Xpeak = (proc.Opeak - proc.Omin) / (proc.Omax - proc.Omin)
	proc.N = math.Log(0.5) / math.Log(proc.Xpeak)
	proc.W = 1.0 - math.Pow(math.Abs(2.0*(proc.Xpeak-0.5)), 2)
}
