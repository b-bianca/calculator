package calculator

import (
	"errors"
	"fmt"
)

type Operator struct {
	ValueA   float64
	ValueB   float64
	Operator string
	Result   float64
}

type Calculator struct {
	Name    string
	Version string
}

// NewCalculator creates a new instance of Calculator
func NewCalculator(name, version string) *Calculator {
	return &Calculator{
		Name:    name,
		Version: version,
	}
}

// Calculate is responsible for calculating the result of a specified mathematical operation
func (c *Calculator) Calculate(o *Operator) (*Operator, error) {
	switch o.Operator {
	case "+":
		o.Result = c.Sum(o)
	case "-":
		o.Result = c.Sub(o)
	case "*":
		o.Result = c.Mult(o)
	case "/":
		if o.ValueB == 0 {
			return o, errors.New("division by zero")
		}
		o.Result = c.Div(o)
	default:
		return o, errors.New("invalid operator")
	}
	return o, nil
}

func (c *Calculator) Sum(o *Operator) float64 {
	return o.ValueA + o.ValueB
}

func (c *Calculator) Sub(o *Operator) float64 {
	return o.ValueA - o.ValueB
}

func (c *Calculator) Mult(o *Operator) float64 {
	return o.ValueA * o.ValueB
}

func (c *Calculator) Div(o *Operator) float64 {
	return o.ValueA / o.ValueB
}

// PrintRemainingOperations prints the remaining operations in the operation queue from a specific index currentIdx
func PrintRemainingOperations(operationQueue []Operator, currentIdx int) {
	fmt.Println("Queue of remaining operations:")
	for j := currentIdx + 1; j < len(operationQueue); j++ {
		fmt.Println(operationQueue[j].ValueA, operationQueue[j].Operator, operationQueue[j].ValueB)
	}

	if currentIdx+1 == len(operationQueue) {
		fmt.Println("There are no more operations to be processed")
	}
	fmt.Println()

}
