package main

import ()

func PrimaryCalcs(proc *ImgProcessor, Iraw [][]uint16, tmodel [][]float64) [][]uint16 {
	var murhot, murhotref, L float64
	var Lindex uint16
	var Icont uint16

	Iout := make([][]uint16, proc.Cfg.HeightPxDet)
	for i := 0; i < proc.Cfg.HeightPxDet; i++ {
		Iout[i] = make([]uint16, proc.Cfg.WidthPxDet)
		for j := 0; j < proc.Cfg.WidthPxDet; j++ {
			// Initial calculation
			murhot = proc.MurhotTable[Iraw[i][j]]

			// Thickness compensation
			murhotref = murhot
			if tmodel[i][j] >= proc.Cfg.Tmin {
				murhotref = murhot * (proc.Tref / tmodel[i][j])
			}

			// Contrast adjustment and rescaling
			L = (murhotref - proc.Omin) / (proc.Omax - proc.Omin)
			if L < 0.0 {
				L = 0.0
			}
			if L > 1.0 {
				L = 1.0
			}
			Lindex = uint16(L / proc.Lstep)
			Icont = proc.IcontTable[Lindex]

			// Drawing the scale bars
			Iout[i][j] = Icont
			if proc.Iscale[i][j] == 0 {
				Iout[i][j] = 0
			} else if proc.Iscale[i][j] == 2 {
				Iout[i][j] = proc.ImaxOutInt
			}

			// Drawing the modelled edges of the core
			if (tmodel[i][j] < proc.Cfg.Tedge) && (tmodel[i][j] > 0.0) {
				Iout[i][j] = 0
			}
		}
	}
	return Iout
}
