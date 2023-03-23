package test

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first

	return d.First
}

func TestData(t *testing.T) {
	dataString := Data[string]{
		First:  "Rizky",
		Second: "Kurniawan",
	}

	dataString.SayHello("Rizky")

	fmt.Println(dataString)

	dataInt := Data[int]{
		First:  1,
		Second: 2,
	}

	dataInt.ChangeFirst(30)

	fmt.Println(dataInt)
}
