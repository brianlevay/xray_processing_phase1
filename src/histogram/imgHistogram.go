package histogram

import (
	fe "fileExplorer"
	tiff "golang.org/x/image/tiff"
	"image"
	"log"
	"math"
	"os"
	"path"
)

func ImageHistogram(contents *fe.FileContents, bits int, nbins int) *Histogram {
	hist := newHistogram(bits, nbins)
	for i := 0; i < len(contents.Selected); i++ {
		pathtofile := path.Join(contents.Root, contents.Selected[i])
		infile, errF := os.Open(pathtofile)
		if errF != nil {
			log.Println(errF)
		} else {
			defer infile.Close()
			img, errD := tiff.Decode(infile)
			if errD != nil {
			    log.Println(errD)
			} else {
			    countPixels(&img, hist)
			}
		}
	}
	return hist
}

func newHistogram(bits int, nbins int) *Histogram {
    hist := new(Histogram)
	hist.Min = 0
	hist.Max = math.Pow(2, float64(bits))
	hist.Step = (hist.Max - hist.Min) / float64(nbins)
	for i := 0; i < nbins; i++ {
		hist.Bins = append(hist.Bins, float64(i)*hist.Step + hist.Min)
		hist.Cts = append(hist.Cts, 0)
	}
	return hist
}

func countPixels(img *image.Image, hist *Histogram) {
    x_min := (*img).Bounds().Min.X
    y_min := (*img).Bounds().Min.Y
    x_max := (*img).Bounds().Max.X
    y_max := (*img).Bounds().Max.Y
    nbins := len(hist.Bins)
    var i_act float64
    var i_int int
    for x := x_min; x < x_max; x++ {
        for y := y_min; y < y_max; y++ {
            px := (*img).At(x,y)
            r,_,_,_ := px.RGBA()
            i_act = (float64(r) - hist.Min) / hist.Step
            i_int = int(math.Floor(i_act))
            if (i_int >= 0) && (i_int <= nbins) {
                hist.Cts[i_int] += 1
            }
        }
    }
    return
}