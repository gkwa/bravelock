package bravelock

type HardcodedStrategy struct {
	Filename string
}

func (s *HardcodedStrategy) GetFilename() string {
	return s.Filename
}
