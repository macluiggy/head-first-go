package calc

import (
	"fmt"
	"reflect"
)

func Add(a, b float64) (float64, error) {
	if reflect.TypeOf(a).String() != "float64" {
		return 0, fmt.Errorf("a is not float64")
	}
	if reflect.TypeOf(b).String() != "float64" {
		return 0, fmt.Errorf("b is not float64")
	}
	return a + b, nil
}

func Substract(a, b float64) (float64, error) {
	if reflect.TypeOf(a).String() != "float64" {
		return 0, fmt.Errorf("a is not float64")
	}
	if reflect.TypeOf(b).String() != "float64" {
		return 0, fmt.Errorf("b is not float64")
	}
	return a - b, nil
}
