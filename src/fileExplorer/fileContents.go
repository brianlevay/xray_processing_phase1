package fileExplorer

import ()

type FileContents struct {
	Root      string   `json:"Root"`
	DirNames  []string `json:"Dirs"`
	FileNames []string `json:"Files"`
	Selected  []string `json:"Selected"`
	Extension string   `json:"-"`
	Error     error    `json:"Error"`
}
