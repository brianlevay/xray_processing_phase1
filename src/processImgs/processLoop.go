package processImgs

import ()

func ProcessByPixel(proc *ImgProcessor, Iraw [][]float64, theta float64, offset float64) [][]float64 {
	var tmodel, murhot, murhotref float64
	height := len(Iraw)
	width := len(Iraw[0])
	Iproc := make([][]float64, height)
	for i := 0; i < height; i++ {
		Iproc[i] = make([]float64, width)
		for j := 0; j < width; j++ {
			tmodel = Tmodel(proc, i, j, height, width, theta, offset)
			murhot = MuRhoT(proc, Iraw[i][j])
			murhotref = Compensation(proc, murhot, tmodel)
			Iproc[i][j] = ContrastAdjustment(proc, murhotref)
		}
	}
	return Iproc
}
