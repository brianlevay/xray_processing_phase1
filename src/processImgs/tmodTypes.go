package processImgs

import ()

type TModel struct {
	CoreType  string
	R         float64
	Cos0      float64
	Sin0      float64
	Xra       float64
	Yra       float64
	Zra       float64
	Xrs       float64
	Yrs       float64
	Zrs       float64
	DelZr     float64
	DelZr2    float64
	C         float64
	XrStep    float64
	XrMin     float64
	XrMax     float64
	Txz2Table []float64
}
