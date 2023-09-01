package tool

import (
	"fmt"
	"testing"
)

func TestGetCurrentAbPath(t *testing.T) {
	path := GetCurrentAbPath()
	fmt.Println(path)
}
