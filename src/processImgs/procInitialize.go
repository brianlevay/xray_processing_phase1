package processImgs

import (
	"math"
)

func (proc *ImgProcessor) Initialize(cfg map[string]float64) error {
	// Basic Detector Setup //
	proc.Height = int(cfg["Height"])
	proc.Width = int(cfg["Width"])
	proc.CmPerPx = cfg["CmPerPx"]

	// Raw Image Size //
	proc.Bits = int(cfg["Bits"])

	// Configuration Variables for CoreAxis //
	threshFrac := cfg["ThreshFrac"]
	gapMinFrac := cfg["GapMinFrac"]
	gapMaxFrac := cfg["GapMaxFrac"]
	proc.Nmass = cfg["Nmass"]
	proc.FilterSteps = int(cfg["FilterSteps"])
	proc.MaxTheta = cfg["MaxTheta"]

	// Configuration Variables for Scales //
	proc.BorderPx = int(cfg["BorderPx"])
	proc.ScaleWidth = cfg["ScaleWidth"]
	proc.RoiWidth = cfg["RoiWidth"]

	// Configuration variables not currently set in file //
	proc.Tmin = 0.5
	proc.Lstep = 0.001

	// Check for validity of data //
	errValid := proc.Validation()
	if errValid != nil {
		return errValid
	}

	// Calculations //
	proc.ProjMult = 1.0 * (proc.SrcHeight / (proc.SrcHeight - proc.CoreHeight - (proc.CoreDiameter / 2.0)))
	proc.ImaxInFlt = math.Pow(2, float64(proc.Bits)) - 1.0
	proc.ImaxOutFlt = math.Pow(2, 16.0) - 1.0
	proc.ImaxInInt = uint16(proc.ImaxInFlt)
	proc.ImaxOutInt = uint16(proc.ImaxOutFlt)
	proc.IthreshInt = uint16(threshFrac * proc.ImaxInFlt)
	proc.PxGapMin = cmCoreToPx(proc, (gapMinFrac * proc.CoreDiameter))
	proc.PxGapMax = cmCoreToPx(proc, (gapMaxFrac * proc.CoreDiameter))
	proc.Omin = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ihigh+1.0)
	proc.Omax = math.Log(proc.ImaxInFlt+1.0) - math.Log(proc.Ilow+1.0)
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
