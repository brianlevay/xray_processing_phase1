package processImgs

import (
	"math"
)

func Tmodel(proc *ImgProcessor, Iraw [][]float64, theta float64, offset float64) [][]float64 {
	height := len(Iraw)
	width := len(Iraw[0])
	tmodel := make([][]float64, height)
	for i := 0; i < height; i++ {
		tmodel[i] = make([]float64, width)
		for j := 0; j < width; j++ {

		}
	}
	return tmodel
}
