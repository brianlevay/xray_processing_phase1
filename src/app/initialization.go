package main

import (
	"errors"
	"math"
)

func (hset *HistogramSet) Initialize(cfg *Configuration) error {
	if (hset.Cfg.HeightPxHist == 0) || (hset.Cfg.WidthPxHist == 0) || (hset.Cfg.Bits < 0) || (hset.Cfg.Nbins <= 0) {
		return errors.New("Invalid configuration values for histogram")
	}
	return nil
}

func (proc *ImgProcessor) Initialize(cfg *Configuration) error {
	proc.Lstep = 0.001
	if (proc.Cfg.SrcHeight == 0.0) || (proc.CoreDiameter == 0.0) || (proc.Cfg.SrcHeight < (proc.Cfg.CoreHeight + proc.CoreDiameter)) {
		return errors.New("Invalid measurement geometry")
	}
	if (proc.Cfg.HeightPxDet == 0) || (proc.Cfg.WidthPxDet == 0) || (proc.Cfg.CmPerPx == 0.0) || (proc.Cfg.Bits < 0) {
		return errors.New("Invalid configuration values for detector and/or input data")
	}
	proc.ProjMult = 1.0 * (proc.Cfg.SrcHeight / (proc.Cfg.SrcHeight - proc.Cfg.CoreHeight - (proc.CoreDiameter / 2.0)))
	proc.ImaxInFlt = math.Pow(2, float64(proc.Cfg.Bits)) - 1.0
	proc.ImaxOutFlt = math.Pow(2, 16.0) - 1.0
	proc.ImaxInInt = uint16(proc.ImaxInFlt)
	proc.ImaxOutInt = uint16(proc.ImaxOutFlt)
	proc.IthreshInt = uint16(proc.Cfg.ThreshFrac * proc.ImaxInFlt)
	proc.PxGapMin = cmCoreToPx(proc, (proc.Cfg.GapMinFrac * proc.CoreDiameter))
	proc.PxGapMax = cmCoreToPx(proc, (proc.Cfg.GapMaxFrac * proc.CoreDiameter))
	proc.Omin = math.Log(proc.ImaxInFlt+1.0) - math.Log((proc.IhighFrac*proc.ImaxInFlt)+1.0)
	proc.Opeak = math.Log(proc.ImaxInFlt+1.0) - math.Log((proc.IpeakFrac*proc.ImaxInFlt)+1.0)
	proc.Omax = math.Log(proc.ImaxInFlt+1.0) - math.Log((proc.IlowFrac*proc.ImaxInFlt)+1.0)
	proc.Tref = proc.CoreDiameter
	if proc.CoreType == "HR" {
		proc.Tref = (proc.CoreDiameter / 2.0)
	}
	proc.CalculateMassTable()
	proc.CalculateXYd()
	proc.CalculateMurhotTable()
	proc.CalculateIcontTable()
	proc.CreateScaleBars()
	return nil
}

func newHistogram(bits int, nbins int) *Histogram {
	hist := new(Histogram)
	hist.Step = math.Pow(2, float64(bits)) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Cts = append(hist.Cts, 0)
	}
	return hist
}

func newTModel(proc *ImgProcessor, theta float64, offset float64) *TModel {
	t := new(TModel)
	t.CoreType = proc.CoreType
	t.R = (proc.CoreDiameter / 2.0)
	thetaR := theta * (math.Pi / 180.0)
	t.Cos0 = math.Cos(thetaR)
	t.Sin0 = math.Sin(thetaR)
	t.Xra = (proc.Xc+offset)*t.Cos0 - proc.Yc*t.Sin0
	t.Yra = proc.Xc*t.Sin0 - proc.Yc*t.Cos0
	t.Zra = (proc.Cfg.CoreHeight + t.R)
	t.Xrs = proc.Xc*t.Cos0 - proc.Yc*t.Sin0
	t.Yrs = proc.Xc*t.Sin0 - proc.Yc*t.Cos0
	t.Zrs = proc.Cfg.SrcHeight
	t.DelZr = (0.0 - t.Zrs)
	t.DelZr2 = t.DelZr * t.DelZr
	t.C = t.Xrs*t.Xrs - 2*t.Xrs*t.Xra + t.Xra*t.Xra
	t.C += t.Zrs*t.Zrs - 2*t.Zrs*t.Zra + t.Zra*t.Zra - t.R*t.R
	t.XrStep = proc.Cfg.CmPerPx / 2.0
	t.XrMin = t.DelZr*((t.Xra-t.R-t.Xrs)/(t.Zra+t.R-t.Zrs)) + t.Xrs
	t.XrMax = t.DelZr*((t.Xra+t.R-t.Xrs)/(t.Zra+t.R-t.Zrs)) + t.Xrs
	t.CalculateTxz2Table()
	return t
}
