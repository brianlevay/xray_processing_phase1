package processImgs

import ()

func Compensation(proc *ImgProcessor, murhot [][]float64, tmodel [][]float64) [][]float64 {
	tref := proc.CoreDiameter
	if proc.CoreType == "HR" {
		tref = (tref / 2.0)
	}
	height := len(murhot)
	width := len(murhot[0])
	murhotref := make([][]float64, height)
	for i := 0; i < height; i++ {
		murhotref[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			if tmodel[i][j] >= proc.Tmin {
				murhotref[i][j] = murhot[i][j] * (tref / tmodel[i][j])
			} else {
				murhotref[i][j] = murhot[i][j]
			}
		}
	}
	return murhotref
}
