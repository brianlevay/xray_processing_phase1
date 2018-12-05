package processImgs

import (
	"math"
)

func MuRhoT(proc *ImgProcessor, Iraw [][]float64) [][]float64 {
	height := len(Iraw)
	width := len(Iraw[0])
	murhot := make([][]float64, height)
	for i := 0; i < height; i++ {
		murhot[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			murhot[i][j] = math.Log(proc.ImaxIn+1.0) - math.Log(Iraw[i][j]+1.0)
		}
	}
	return murhot
}
