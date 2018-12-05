package processImgs

import (
	"math"
)

func MuRhoT(proc *ImgProcessor, Iraw float64) float64 {
	return math.Log(proc.ImaxIn+1.0) - math.Log(Iraw+1.0)
}

func Compensation(proc *ImgProcessor, murhot float64, tmodel float64) float64 {
	tref := proc.CoreDiameter
	if proc.CoreType == "HR" {
		tref = (tref / 2.0)
	}
	if tmodel >= proc.Tmin {
		return murhot * (tref / tmodel)
	}
	return murhot
}

func ContrastAdjustment(proc *ImgProcessor, murhotref float64) float64 {
	X := (murhotref - proc.Omin) / (proc.Omax - proc.Omin)
	if X < 0.0 {
		X = 0.0
	}
	if X > 1.0 {
		X = 1.0
	}
	P := math.Pow(X, proc.N)
	SX := 0.5*math.Sin(math.Pi*(X-0.5)) + 0.5
	SP := 0.5*math.Sin(math.Pi*(P-0.5)) + 0.5
	Y := proc.W*SP + (1-proc.W)*SX
	return proc.ImaxOut * (1 - Y)
}
