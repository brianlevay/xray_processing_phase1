package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize() {
	// Configuration Variables //
	proc.CmPerPx = 0.0099
	proc.Tmin = 0.5
	proc.Lstep = 0.001

	proc.ImaxInFlt = math.Pow(2, float64(proc.Bits)) - 1.0
	proc.ImaxOutFlt = math.Pow(2, 16.0) - 1.0
	proc.ImaxInInt = uint16(proc.ImaxInFlt)
	proc.ImaxOutInt = uint16(proc.ImaxOutFlt)
	proc.Omin = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Omax = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ilow+1.0)
	proc.Tref = proc.CoreDiameter
	if proc.CoreType == "HR" {
		proc.Tref = (proc.CoreDiameter / 2.0)
	}
	proc.ProjMult = 1.0 * (proc.SrcHeight / (proc.SrcHeight - proc.CoreHeight - (proc.CoreDiameter / 2.0)))
	proc.CalculateXYd()
	proc.CalculateWtsGapTable()
	proc.CalculateMurhotTable()
	proc.CalculateIcontTable()
	proc.CreateScaleBars()
	return
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

func (proc *ImgProcessor) CalculateWtsGapTable() {
	// Configuration Variables //
	deltaGapMin := 3.0
	deltaGapFlat := 0.5
	deltaGapMax := 0.5

	// This uses a simplistic relationship between pixel width and core width, //
	// not accounting for 3D projection of a cylinder //
	var gap float64
	ptHigh1 := proc.CoreDiameter
	ptHigh0 := ptHigh1 + deltaGapMax
	ptLow1 := ptHigh1 - deltaGapFlat
	ptLow0 := ptLow1 - deltaGapMin

	wtsGap := make([]float64, proc.Width)
	for k := 0; k < proc.Width; k++ {
		gap = pxToCmCore(proc, k)
		if (gap <= ptLow0) || (gap >= ptHigh0) {
			wtsGap[k] = 0.0
		} else if (gap >= ptLow1) && (gap <= ptHigh1) {
			wtsGap[k] = 1.0
		} else if (gap > ptLow0) && (gap < ptLow1) {
			wtsGap[k] = ((1.0-0.0)/(ptLow1-ptLow0))*(gap-ptLow0) + 0.0
		} else {
			wtsGap[k] = ((0.0-1.0)/(ptHigh0-ptHigh1))*(gap-ptHigh1) + 1.0
		}
	}
	proc.WtsGapTable = wtsGap
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
