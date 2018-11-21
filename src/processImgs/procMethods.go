package processImgs

import (
	fe "fileExplorer"
	"image"
	"log"
	"path"
	"sync"
)

func (proc *ImgProcessor) ProcessImage(root string, filename string, wg *sync.WaitGroup) {
	img, errImg := fe.OpenTiff(root, filename)
	if errImg != nil {
		wg.Done()
		return
	}
	return
}
