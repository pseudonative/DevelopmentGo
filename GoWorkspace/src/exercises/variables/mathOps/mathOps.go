package mathops

var OperationCount int

var secretFactor = 2

func Double(value int) int {
	OperationCount++
	return value * secretFactor
}
