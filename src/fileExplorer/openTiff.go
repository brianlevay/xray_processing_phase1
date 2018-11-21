package fileExplorer

import (
	tiff "golang.org/x/image/tiff"
	"image"
	"os"
	"path"
)

func OpenTiff(root string, filename string) (*image.Image, error) {
	pathtofile := path.Join(root, filename)
	infile, errF := os.Open(pathtofile)
	if errF != nil {
		return nil, errF
	} else {
		defer infile.Close()
		img, errD := tiff.Decode(infile)
		if errD != nil {
			return nil, errD
		}
		return &img, nil
	}
}
