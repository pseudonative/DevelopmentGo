package calculator

import "errors"

type Operation interface {
	Execute(a, b int) (int, error)
}

type AddOperation struct{}

func (ao AddOperation) Execute(a, b int) (int, error) {
	return a + b, nil
}

type SubtractOperation struct{}

func (so SubtractOperation) Execute(a, b int) (int, error) {
	return a - b, nil
}

type MultiplyOperation struct{}

func (mo MultiplyOperation) Execute(a, b int) (int, error) {
	return a * b, nil
}

type DivideOperation struct{}

func (do DivideOperation) Execute(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}
