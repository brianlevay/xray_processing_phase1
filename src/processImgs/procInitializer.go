package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize() {
	bits := 14
	pxcm := 0.0099
	tmin := 0.5

	proc.Bits = bits
	proc.Imax = math.Pow(2, float64(bits)) - 1.0
	proc.Pxcm = pxcm
	proc.Tmin = tmin

	proc.Omin = math.Log(proc.Imax+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Opeak = math.Log(proc.Imax+1.0) - math.Log(proc.Ipeak+1.0)
	proc.Omax = math.Log(proc.Imax+1.0) - math.Log(proc.Ilow+1.0)

	proc.Xpeak = (proc.Opeak - proc.Omin) / (proc.Omax - proc.Omin)
	proc.N = math.Log(0.5) / math.Log(proc.Xpeak)
	proc.W = 1.0 - math.Pow(math.Abs(2.0*(proc.Xpeak-0.5)), 2)
}
