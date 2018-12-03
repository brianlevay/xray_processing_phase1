package processImgs

import (
	"math"
)

func ContrastAdjustment(murhotref [][]float64, low float64, peak float64, high float64, bits int) [][]float64 {
	var X, P, SX, SP, Y float64

	Imax := math.Pow(2, float64(bits)) - 1.0
	Xpeak := (peak - low) / (high - low)
	n := math.Log(0.5) / math.Log(Xpeak)
	w := 1.0 - math.Pow(math.Abs(2.0*(Xpeak-0.5)), 2)

	height := len(murhotref)
	width := len(murhotref[0])
	Iproc := make([][]float64, height)
	for i := 0; i < height; i++ {
		Iproc[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			X = (murhotref[i][j] - low) / (high - low)
			if X < 0.0 {
				X = 0.0
			}
			if X > 1.0 {
				X = 1.0
			}
			P = math.Pow(X, n)
			SX = 0.5*math.Sin(math.Pi*(X-0.5)) + 0.5
			SP = 0.5*math.Sin(math.Pi*(P-0.5)) + 0.5
			Y = w*SP + (1-w)*SX
			Iproc[i][j] = Imax * (1 - Y)
		}
	}
	return Iproc
}
