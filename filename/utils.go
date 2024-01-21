package filename

import "path"

func GetFnameWithoutExtension(fname string) string {
	return fname[:len(fname)-len(path.Ext(fname))]
}
