package histogram

import (
	fe "fileExplorer"
	"fmt"
	tif "golang.org/x/image/tiff"
	"image"
	"math"
	"os"
)

func ImageHistogram(contents *fe.FileContents, bits int, nbins int) {
	var min float64 = 0
	var max float64 = math.Pow(2, float64(bits))
	var step float64 = (max - min) / float64(nbins)

	var bins []float64
	var cts []float64
	for i := 0; i < nbins; i++ {
		bins = append(bins, float64(i)*step+min)
		cts = append(cts, 0)
	}

	for i := 0; i < len(contents.Selected); i++ {
		infile, err := os.Open(os.Args[1])
	}

	fmt.Println(bins)
	return
}
