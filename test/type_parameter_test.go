package test

import (
	"fmt"
	"testing"
)

func Length[T any](param T) T {
	fmt.Println(param)

	return param
}

func TestSample(t *testing.T) {
	var resultString string = Length("Rizky")

	fmt.Println(resultString)

	var resultNumber int = Length(20)

	fmt.Println(resultNumber)
}
