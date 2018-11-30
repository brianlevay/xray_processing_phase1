package processImgs

import ()

func CopyFloat(origslice [][]float64) [][]float64 {
	height := len(origslice)
	width := len(origslice[0])
	newslice := make([][]float64, height)
	for i := 0; i < height; i++ {
		newslice[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			newslice[i][j] = origslice[i][j]
		}
	}
	return newslice
}
