package tool

import (
	"path"
	"runtime"
	"strings"
)

func GetCurrentAbPath() string {
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}
	var theInd = strings.Index(abPath, "gotest")
	return abPath[0 : theInd+6]
}
