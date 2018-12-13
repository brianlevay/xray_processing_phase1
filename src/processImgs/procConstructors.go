package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize() {
	// Configuration Variables for Calculations //
	proc.Height = 1550
	proc.Width = 1032
	proc.Bits = 14
	proc.CmPerPx = 0.0099
	proc.Tmin = 0.5
	proc.Lstep = 0.001

	threshFrac := 0.9
	gapMinFrac := 0.7
	gapMaxFrac := 1.2
	proc.MaxTheta = 5.0

	// Configuration Variables for Scales //
	proc.BorderPx = 2
	proc.ScaleWidth = 0.2
	proc.RoiWidth = 0.1

	// Calculated Values //
	proc.ProjMult = 1.0 * (proc.SrcHeight / (proc.SrcHeight - proc.CoreHeight - (proc.CoreDiameter / 2.0)))

	proc.ImaxInFlt = math.Pow(2, float64(proc.Bits)) - 1.0
	proc.ImaxOutFlt = math.Pow(2, 16.0) - 1.0
	proc.ImaxInInt = uint16(proc.ImaxInFlt)
	proc.ImaxOutInt = uint16(proc.ImaxOutFlt)

	proc.IthreshInt = uint16(threshFrac * proc.ImaxInFlt)
	proc.PxGapMin = cmCoreToPx(proc, (gapMinFrac * proc.CoreDiameter))
	proc.PxGapMax = cmCoreToPx(proc, (gapMaxFrac * proc.CoreDiameter))

	proc.Omin = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Omax = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ilow+1.0)
	proc.Tref = proc.CoreDiameter
	if proc.CoreType == "HR" {
		proc.Tref = (proc.CoreDiameter / 2.0)
	}

	proc.CalculateMassTable()
	proc.CalculateXYd()
	proc.CalculateMurhotTable()
	proc.CalculateIcontTable()
	proc.CreateScaleBars()
	return
}

func (proc *ImgProcessor) CalculateMassTable() {
	nvals := int(proc.ImaxInInt + 1)
	masses := make([]float64, nvals)
	for k := 0; k < nvals; k++ {
		masses[k] = float64(proc.ImaxInInt-uint16(k)) / float64(proc.ImaxInFlt)
	}
	proc.MassTable = masses
}

func (proc *ImgProcessor) CalculateXYd() {
	Yc := (float64(proc.Height) * proc.CmPerPx) / 2.0
	Xc := (float64(proc.Width) * proc.CmPerPx) / 2.0
	Yd := make([]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		Yd[i] = float64(i)*proc.CmPerPx + (proc.CmPerPx / 2.0)
	}
	Xd := make([]float64, proc.Width)
	for j := 0; j < proc.Width; j++ {
		Xd[j] = float64(j)*proc.CmPerPx + (proc.CmPerPx / 2.0)
	}
	proc.Xc = Xc
	proc.Yc = Yc
	proc.Xd = Xd
	proc.Yd = Yd
}

func (proc *ImgProcessor) CalculateMurhotTable() {
	nvals := int(proc.ImaxInInt + 1)
	murhots := make([]float64, nvals)
	for k := 0; k < nvals; k++ {
		murhots[k] = math.Log(proc.ImaxInFlt+1.0) - math.Log(float64(k)+1.0)
	}
	proc.MurhotTable = murhots
}

func (proc *ImgProcessor) CalculateIcontTable() {
	var L, P, SF, SP, Y float64
	Opeak := math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ipeak+1.0)
	Lpeak := (Opeak - proc.Omin) / (proc.Omax - proc.Omin)
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
}
