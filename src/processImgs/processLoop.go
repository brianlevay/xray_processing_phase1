package processImgs

import ()

func ProcessByPixel(proc *ImgProcessor, Iraw [][]float64, theta float64, offset float64) [][]float64 {
	var tmodel, murhot, murhotref, Iproc float64
	Iout := make([][]float64, proc.Height)
	for i := 0; i < proc.Height; i++ {
		Iout[i] = make([]float64, proc.Width)
		for j := 0; j < proc.Width; j++ {
			tmodel = Tmodel(proc, i, j, theta, offset)
			murhot = MuRhoT(proc, Iraw[i][j])
			murhotref = Compensation(proc, murhot, tmodel)
			Iproc = ContrastAdjustment(proc, murhotref)
			Iout[i][j] = AddScaleBars(proc, i, j, Iproc)
		}
	}
	return Iout
}
