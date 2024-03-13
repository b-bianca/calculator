package calculator

import (
	"errors"
	"testing"
)

func TestCalculator_Calculate(t *testing.T) {
	tests := []struct {
		name           string
		operator       Operator
		expectedResult float64
		expectedError  error
	}{
		{
			name:           "Sum test",
			operator:       Operator{ValueA: 2, ValueB: 3, Operator: "+"},
			expectedResult: 5,
			expectedError:  nil,
		},
		{
			name:           "Subtraction test",
			operator:       Operator{ValueA: 10, ValueB: 3, Operator: "-"},
			expectedResult: 7,
			expectedError:  nil,
		},
		{
			name:           "Multiplication test",
			operator:       Operator{ValueA: 4, ValueB: 5, Operator: "*"},
			expectedResult: 20,
			expectedError:  nil,
		},
		{
			name:           "Division test",
			operator:       Operator{ValueA: 10, ValueB: 2, Operator: "/"},
			expectedResult: 5,
			expectedError:  nil,
		},
		{
			name:           "Division by zero test",
			operator:       Operator{ValueA: 10, ValueB: 0, Operator: "/"},
			expectedResult: 0,
			expectedError:  errors.New("division by zero"),
		},
		{
			name:           "Invalid operator test",
			operator:       Operator{ValueA: 10, ValueB: 2, Operator: "%"},
			expectedResult: 0,
			expectedError:  errors.New("invalid operator"),
		},
	}

	calc := NewCalculator("Test calculator", "v1.0")

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resultado, err := calc.Calculate(&test.operator)

			if (err != nil && test.expectedError == nil) || (err == nil && test.expectedError != nil) || (err != nil && test.expectedError != nil && err.Error() != test.expectedError.Error()) {
				t.Errorf("Unexpected error: expected %v, received %v", test.expectedError, err)
			}

			if resultado.Result != test.expectedResult {
				t.Errorf("Expected result %f, received %f", test.expectedResult, resultado.Result)
			}
		})
	}
}
