package test

import (
	"fmt"
	"testing"
)

func MultipleTypeParameter[T1 any, T2 any](param1 T1, param2 T2) {
	fmt.Println(param1)
	fmt.Println(param2)
}

func TestMultipleTypeParameter(t *testing.T) {
	MultipleTypeParameter("Mawar", 10)
	MultipleTypeParameter(10, "Mawar")
}
