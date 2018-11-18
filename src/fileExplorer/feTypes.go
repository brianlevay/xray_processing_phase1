package fileExplorer

import (
	"image"
	"sync"
)

type FileContents struct {
	Root      string   `json:"Root"`
	DirNames  []string `json:"Dirs"`
	FileNames []string `json:"Files"`
	Selected  []string `json:"Selected"`
	Extension string   `json:"-"`
}

type Processor interface {
	ProcessImage(img *image.Image, wg *sync.WaitGroup)
}
