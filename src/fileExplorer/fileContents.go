package fileExplorer

import (
	"sync"
)

type FileContents struct {
	Mtx       sync.Mutex `json:"-"`
	Root      string     `json:"Root"`
	DirNames  []string   `json:"Dirs"`
	FileNames []string   `json:"Files"`
	Selected  []string   `json:"Selected"`
	Error     error      `json:"Error"`
}
