package processImgs

import (
	"math"
)

func ContrastAdjustment(proc *ImgProcessor, murhotref [][]float64) [][]float64 {
	var X, P, SX, SP, Y float64
	height := len(murhotref)
	width := len(murhotref[0])
	Iproc := make([][]float64, height)
	for i := 0; i < height; i++ {
		Iproc[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			X = (murhotref[i][j] - proc.Omin) / (proc.Omax - proc.Omin)
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
			Iproc[i][j] = proc.Imax * (1 - Y)
		}
	}
	return Iproc
}
