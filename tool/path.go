package tool

import (
	"path"
	"runtime"
)

func GetCurrentAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath[0 : len(abPath)-5]
}
