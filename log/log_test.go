package log

import (
	"fmt"
	"testing"
)

func TestLogInit(t *testing.T) {

	LogInit()
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
		a := fmt.Sprintf("测试%d", i)
		fmt.Println(a)
	}
}
