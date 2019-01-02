package main

import (
	"github.com/alandtsang/godemo/mybox/log/logger"

	"fmt"
	"os"
)

func main() {
	logPath := "/home/zengxl/1.log"
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
	if err != nil {
		fmt.Println("Failed to open log file:", err)
		return
	}
	defer logFile.Close()

	log := logger.Init(true, logFile)
	defer log.Close()
	log.Info("This will log to the log file and the system log")
}
