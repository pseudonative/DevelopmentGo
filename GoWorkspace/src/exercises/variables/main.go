package main

import (
	"fmt"

	"github.com/pseudonative/mathOps/calculator"
	concurrentprocessor "github.com/pseudonative/mathOps/concurrentProcessor"
	mathops "github.com/pseudonative/mathOps/mathOps"
)

func performOperation(op calculator.Operation, a, b int) {
	result, err := op.Execute(a, b)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("Result", result)
}

func main() {
	value := 11
	doubled := mathops.Double(value)
	fmt.Printf("Double of %d is %d\n", value, doubled)
	fmt.Printf("Operations performed: %d\n", mathops.OperationCount)

	a, b := 72, 9
	addOp := calculator.AddOperation{}
	performOperation(addOp, a, b)

	subtractOp := calculator.SubtractOperation{}
	performOperation(subtractOp, a, b)

	multiplyOp := calculator.MultiplyOperation{}
	performOperation(multiplyOp, a, b)

	divideOp := calculator.DivideOperation{}
	performOperation(divideOp, a, b)
	fmt.Println("############################")
	fmt.Println("############################")
	fmt.Println("############################")

	entries := []concurrentprocessor.DataEntry{1, 2, 3, 4, 5}
	successChan := make(chan concurrentprocessor.Result)
	errorChan := make(chan error)

	for _, entry := range entries {
		go concurrentprocessor.ProcessData(entry, successChan, errorChan)
	}
	for i := 0; i < len(entries); i++ {
		select {
		case successMsg := <-successChan:
			fmt.Printf("Success processing entry %d: %s\n", successMsg.Entry, successMsg.Output)
		case errMsg := <-errorChan:
			fmt.Println("Error:", errMsg)
		}
	}
}
