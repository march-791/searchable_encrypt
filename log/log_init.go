package log

import (
	"fmt"
	"gotest/tool"
	"log"
	"os"
	"strings"
)

func LogInit() {
	absDir := tool.GetCurrentAbPath()
	fmt.Println(strings.Replace(absDir, "\\", "/", -1))
	logFile, err := os.OpenFile(absDir+"/log"+"/search.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("open log file failed, err:", err)
		return
	}
	log.SetOutput(logFile)
}
