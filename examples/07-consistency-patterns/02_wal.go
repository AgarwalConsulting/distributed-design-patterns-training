package main

import (
	"bufio"
	"fmt"
	"os"
)

const logFile = "write_ahead_log.txt"
const dataFile = "data_store.txt"

func writeToLog(data string) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(data + "\n")
	writer.Flush()
}

func applyToDataStore(data string) {
	file, err := os.OpenFile(dataFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println("Error opening data file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString(data + "\n")
	writer.Flush()
}

func recoverFromLog() {
	file, err := os.Open(logFile)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		applyToDataStore(scanner.Text())
	}
}

func write(data string) {
	writeToLog(data)
	applyToDataStore(data)
}

func main() {
	// Recovery from log at startup
	recoverFromLog()

	// Example write operation
	write("This is a test entry")
}
