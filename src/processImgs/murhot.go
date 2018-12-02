package processImgs

import (
	"math"
)

func MuRhoT(Iraw [][]float64, bits int) [][]float64 {
	Imax := math.Pow(2, float64(bits)) - 1.0
	height := len(Iraw)
	width := len(Iraw[0])
	murhot := make([][]float64, height)
	for i := 0; i < height; i++ {
		murhot[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			murhot[i][j] = math.Log(Imax+1.0) - math.Log(Iraw[i][j]+1.0)
		}
	}
	return murhot
}

func MuRhoTbounds(Ilow float64, Ipeak float64, Ihigh float64, bits int) (float64, float64, float64) {
	Imax := math.Pow(2, float64(bits)) - 1.0
	Omin := math.Log(Imax+1.0) - math.Log(Ihigh+1.0)
	Opeak := math.Log(Imax+1.0) - math.Log(Ipeak+1.0)
	Omax := math.Log(Imax+1.0) - math.Log(Ilow+1.0)
	return Omin, Opeak, Omax
}
