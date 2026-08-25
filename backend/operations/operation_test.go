package operations

import (
	"testing"
)

func TestOperations(t *testing.T){
	testCase := [] struct {
		name	string
		numA	float64
		numB	float64
		operation func(float64, float64) float64
		expected float64
	}{
		{
			name:	"addition",
			numA:	10,
			numB:	5,
			operation: Add,
			expected: 15,
		},
		{
			name: "subtraction",
			numA:	10,
			numB:	5,
			operation: Sub,
			expected: 5,
		},
		{
			name: "multiplication",
			numA:	10,
			numB:	5,
			operation: Mul,
			expected: 50,
		},
	}

	testDivs := [] struct {
		name	string
		numA	float64
		numB	float64
		operation func(float64, float64) (float64, error)
		expected float64
		expectedError	bool
	}{
		{
			name:	"valid division",
			numA:	10,
			numB:	5,
			operation: Div,
			expected: 2,
			expectedError: false,
		},
		{
			name:	"invalid division",
			numA:	10,
			numB:	0,
			operation: Div,
			expected: 0,
			expectedError: true,
		},
	}

	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T){
			result := test.operation(test.numA, test.numB)
			if test.expected != result{
				t.Errorf(
					"expected %f, but got %f",
					test.expected,
					result,	
				)
			}
		})
	}

	for _, test := range testDivs{
		t.Run(test.name, func(t *testing.T){
			result, err := test.operation(test.numA, test.numB)
			if test.expectedError {
				if err == nil{
					t.Fatalf("Expected Error, but got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("expected no error, but got error: %s", err)
			}

			if result != test.expected {
				t.Errorf("expected %f, but got %f", test.expected, result)
			}
			
		})
	}
}