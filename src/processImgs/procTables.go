package processImgs

import (
	"math"
)

func (proc *ImgProcessor) CalculateMassTable() {
	var massBase float64
	nvals := int(proc.ImaxInInt + 1)
	masses := make([]float64, nvals)
	for k := 0; k < nvals; k++ {
		massBase = float64(proc.ImaxInInt-uint16(k)) / float64(proc.ImaxInFlt)
		masses[k] = math.Pow(massBase, proc.Nmass)
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
