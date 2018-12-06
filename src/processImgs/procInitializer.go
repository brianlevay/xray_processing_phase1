package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize() {
	proc.ImaxIn = math.Pow(2, float64(proc.Bits)) - 1.0
	proc.ImaxOut = math.Pow(2, 16.0) - 1.0

	proc.CmPx = 0.0099
	proc.Y = make([]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		proc.Y[i] = float64(i)*proc.CmPx + (proc.CmPx / 2.0)
	}
	proc.X = make([]float64, proc.Width)
	for j := 0; j < proc.Width; j++ {
		proc.X[j] = float64(j)*proc.CmPx + (proc.CmPx / 2.0)
	}
	proc.Xc = (float64(proc.Width) * proc.CmPx) / 2.0
	proc.Yc = (float64(proc.Height) * proc.CmPx) / 2.0

	proc.Tref = proc.CoreDiameter
	if proc.CoreType == "HR" {
		proc.Tref = (proc.CoreDiameter / 2.0)
	}
	proc.Tmin = 0.5

	proc.Omin = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Opeak = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ipeak+1.0)
	proc.Omax = math.Log(proc.ImaxIn+1.0) - math.Log(proc.Ilow+1.0)

	proc.Xpeak = (proc.Opeak - proc.Omin) / (proc.Omax - proc.Omin)
	proc.N = math.Log(0.5) / math.Log(proc.Xpeak)
	proc.W = 1.0 - math.Pow(math.Abs(2.0*(proc.Xpeak-0.5)), 2)
	return
}
