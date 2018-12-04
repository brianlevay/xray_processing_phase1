package processImgs

import ()

func FindCoreAxis(proc *ImgProcessor, Iraw [][]float64) (float64, float64) {
	thresh := 0.8

	iMid := []float64{}
	jMid := []float64{}
	Ithresh := proc.Imax * thresh
	height := len(Iraw)
	width := len(Iraw[0])

	for i := 0; i < height; i++ {
		for j := 0; j < (width - 1); j++ {
			if (Iraw[i][j] > Ithresh) && (Iraw[i][j+1] <= Ithresh) {

			}
			if (Iraw[i][j] <= Ithresh) && (Iraw[i][j+1] > Ithresh) {

			}
		}
	}
	return 0.0, 0.0
}
