package main

import (
	"fmt"

	calculator "github.com/b-bianca/calculator/internal"
)

func main() {

	operationQueue := []calculator.Operator{
		{ValueA: 2, ValueB: 3, Operator: "+"},
		{ValueA: 14, ValueB: 8, Operator: "-"},
		{ValueA: 5, ValueB: 6, Operator: "*"},
		{ValueA: 2147483647, ValueB: 2, Operator: "+"},
		{ValueA: 18, ValueB: 3, Operator: "/"},
	}

	calc := calculator.NewCalculator("Test calculator", "v1.0")

	results := make([]float64, len(operationQueue))

	for i := 0; i < len(operationQueue); i++ {
		o := operationQueue[i]
		result, err := calc.Calculate(&o)
		if err != nil {
			fmt.Printf("Calculate failed: %v\n", err)
			continue
		}
		fmt.Println(result.ValueA, result.Operator, result.ValueB, result.Result)

		results[i] = result.Result

		calculator.PrintRemainingOperations(operationQueue, i)
	}

	fmt.Println("Total results list:")
	for _, r := range results {
		fmt.Println(r)
	}
}
