package filename

import (
	"path"
	"runtime"
)

type FilenameStrategy interface {
	GetFilename() string
}

type HardcodedStrategy struct {
	Filename string
}

func (s *HardcodedStrategy) GetFilename() string {
	return s.Filename
}

type FilenameFromGoPackageStrategy struct{}

func (s *FilenameFromGoPackageStrategy) GetFilename() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Base(filename)
}
