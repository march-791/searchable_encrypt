package log

import (
	"fmt"
	"testing"
	"time"
)

func TestLogInit(t *testing.T) {

	startTime := time.Now()
	LogInit()
	for i := 0; i <= 1000; i++ {
		fmt.Println(i)
	}
	elapsedTime := time.Since(startTime) / time.Millisecond  // duration in ms
	fmt.Printf("Segment finished in %dms", int(elapsedTime)) //Segment finished in xxms
}
