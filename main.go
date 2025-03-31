package main

import (
	"fmt"
	"time"
)

type Logger struct{}

func (x Logger) debugLog(logMessage string, endLine byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] Caught a panic in debugLog:", r)
		}
	}()

	if logMessage == "" {
		panic("Log message cannot be empty")
	}

	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Print("[DEBUG " + t + "]: " + logMessage + string(endLine))
}

func (x Logger) errorLog(logMessage string, endLine byte) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] Caught a panic in errorLog:", r)
		}
	}()

	if logMessage == "" {
		panic("Log message cannot be empty")
	}

	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Print("[ERROR " + t + "]: " + logMessage + string(endLine))
}

func (x Logger) infoLog(logMessage string, endLine byte) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("[ERROR] Caught a panic in infoLog:", r)
		}
	}()

	if logMessage == "" {
		panic("Log message cannot be empty")
	}

	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Print("[INFO " + t + "]: " + logMessage + string(endLine))
}

func main() {
	var console = Logger{}

	console.debugLog("Hello World", '\n')
	console.errorLog("Hello World", '\n')
	console.infoLog("Hello World", '\n')

	console.debugLog("", '\n')
	console.errorLog("", '\n')
	console.infoLog("", '\n')
}
