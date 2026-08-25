package operations

import "errors"

func Add(numA, numB float64) float64 {
	return numA + numB
}

func Sub(numA, numB float64) float64 {
	return numA - numB
}

func Mul(numA, numB float64) float64 {
	return numA * numB
}

func Div(numA, numB float64) (float64, error) {
	if numB == 0 {
		return 0, errors.New("Cannot divide the number by Zero")
	}

	return numA / numB, nil
}
