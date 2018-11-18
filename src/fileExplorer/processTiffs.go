package fileExplorer

import (
	tiff "golang.org/x/image/tiff"
	"log"
	"os"
	"path"
	"sync"
)

func ProcessTiffs(contents *FileContents, proc Processor) {
	var wg sync.WaitGroup
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
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
				wg.Add(1)
				go proc.ProcessImage(&img, &wg)
			}
		}
	}
	wg.Wait()
	return
}
