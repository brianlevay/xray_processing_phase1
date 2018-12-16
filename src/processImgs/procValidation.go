package processImgs

import (
	"errors"
)

func (proc *ImgProcessor) Validation() error {
	if (proc.SrcHeight == 0.0) || (proc.CoreDiameter == 0.0) || (proc.SrcHeight < (proc.CoreHeight + proc.CoreDiameter)) {
		return errors.New("Invalid measurement geometry")
	}
	if (proc.Height == 0) || (proc.Width == 0) || (proc.CmPerPx == 0.0) || (proc.Bits < 0) {
		return errors.New("Invalid configuration values for detector and/or input data")
	}
	return nil
}
