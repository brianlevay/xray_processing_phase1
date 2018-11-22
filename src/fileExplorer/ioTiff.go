package fileExplorer

import (
	tiff "golang.org/x/image/tiff"
	"image"
	"os"
	"path"
)

func OpenTiff(root string, filename string) (image.Image, error) {
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
		return img, nil
	}
}

func SaveTiff(img image.Image, root string, filename string) error {
    _, errExist := os.Stat(root)
    if os.IsNotExist(errExist) {
        errDir := os.Mkdir(root, os.ModePerm)
        if errDir != nil {
            return errDir
        }
    }
    pathtofile := path.Join(root, filename)
    outfile, errF := os.Create(pathtofile)
    if errF != nil {
        return errF
    } else {
        defer outfile.Close()
        errE := tiff.Encode(outfile, img, nil)
        if errE != nil {
            return errE
        }
        return nil
    }
}