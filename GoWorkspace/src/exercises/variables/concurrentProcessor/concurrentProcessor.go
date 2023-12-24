package concurrentprocessor

import "fmt"

type DataEntry int

type Result struct {
	Entry  DataEntry
	Output string
}

func ProcessData(entry DataEntry, successChan chan<- Result, errorChan chan<- error) {
	if entry%2 == 0 {
		successChan <- Result{Entry: entry, Output: fmt.Sprintf("Processed entry: %d", entry)}
	} else {
		errorChan <- fmt.Errorf("error processing entry; %d", entry)
	}
}
