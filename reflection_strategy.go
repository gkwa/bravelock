package bravelock

import (
	"path"
	"runtime"
)

type ReflectionStrategy struct{}

func (s *ReflectionStrategy) GetFilename() string {
	_, filename, _, _ := runtime.Caller(1)
	return path.Base(filename)
}
