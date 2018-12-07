package processImgs

import (
	"math"
)

func PrimaryCalcs(proc *ImgProcessor, Iraw [][]float64, tmodel [][]float64) [][]float64 {
	var murhot, murhotref, X, P, SX, SP, Y, Iproc float64
	Iout := make([][]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		Iout[i] = make([]float64, proc.Width)
		for j := 0; j < proc.Width; j++ {
			// Initial calculation
			murhot = math.Log(proc.ImaxIn+1.0) - math.Log(Iraw[i][j]+1.0)

			// Thickness compensation
			murhotref = murhot
			if tmodel[i][j] >= proc.Tmin {
				murhotref = murhot * (proc.Tref / tmodel[i][j])
			}

			// Contrast adjustment and rescaling
			X = (murhotref - proc.Omin) / (proc.Omax - proc.Omin)
			if X < 0.0 {
				X = 0.0
			}
			if X > 1.0 {
				X = 1.0
			}
			P = math.Pow(X, proc.N)
			SX = 0.5*math.Sin(math.Pi*(X-0.5)) + 0.5
			SP = 0.5*math.Sin(math.Pi*(P-0.5)) + 0.5
			Y = proc.W*SP + (1-proc.W)*SX
			Iproc = proc.ImaxOut * (1 - Y)

			// Drawing the scale bars
			Iout[i][j] = Iproc
			if proc.Iscale[i][j] == 0 {
				Iout[i][j] = 0.0
			} else if proc.Iscale[i][j] == 2 {
				Iout[i][j] = proc.ImaxOut
			}

			// Drawing the modelled edges of the core
			if (tmodel[i][j] < proc.Tmin) && (tmodel[i][j] > 0.0) {
				Iout[i][j] = 0.0
			}
		}
	}
	return Iout
}
