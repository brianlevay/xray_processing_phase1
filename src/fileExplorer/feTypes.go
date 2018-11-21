package fileExplorer

import (
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
	ProcessImage(root string, filename string, wg *sync.WaitGroup)
}
