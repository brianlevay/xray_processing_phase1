package fileExplorer

import (
	tiff "golang.org/x/image/tiff"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"sync"
)

func (contents *FileContents) UpdateDir(rootDir string) {
	var dirNames []string = []string{}
	var fileNames []string = []string{}
	files, err := ioutil.ReadDir(rootDir)
	if err != nil {
		log.Println(err)
		return
	}
	for _, file := range files {
		if file.IsDir() == true {
			dirNames = append(dirNames, file.Name())
		} else {
			if filepath.Ext(file.Name()) == contents.Extension {
				fileNames = append(fileNames, file.Name())
			}
		}
	}
	contents.Root = rootDir
	contents.DirNames = dirNames
	contents.FileNames = fileNames
	contents.Selected = []string{}
	return
}

func (contents *FileContents) ProcessTiffs(proc Processer) {
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
