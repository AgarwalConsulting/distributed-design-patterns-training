package main

import (
	"fmt"
	"os"
)

type LogEntry struct {
	Action string
}

func WriteLog(entry LogEntry) error {
	// Open or create log file
	file, err := os.OpenFile("wal.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write log entry
	_, err = file.WriteString(fmt.Sprintf("%s\n", entry.Action))
	return err
}

func ApplyChanges(action string) {
	// Simulate applying some changes
	fmt.Printf("Applying: %s\n", action)
}

func main() {
	// Example: Write logs and apply changes
	action := "UserLoggedIn"
	entry := LogEntry{Action: action}

	// Step 1: Write to WAL
	if err := WriteLog(entry); err != nil {
		fmt.Printf("Error writing to log: %v\n", err)
		return
	}

	// Step 2: Apply changes (e.g., database)
	ApplyChanges(action)
}
