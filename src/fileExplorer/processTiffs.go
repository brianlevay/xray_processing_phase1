package fileExplorer

import (
	"sync"
)

func ProcessTiffs(contents *FileContents, proc Processor) {
	var wg sync.WaitGroup
	nfiles := len(contents.Selected)
	for i := 0; i < nfiles; i++ {
		wg.Add(1)
		go proc.ProcessImage(contents.Root, contents.Selected[i], &wg)
	}
	wg.Wait()
	return
}
