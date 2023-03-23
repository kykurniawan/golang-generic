package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number3 interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

func Min3[T Number2](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestTypeInference(t *testing.T) {
	assert.Equal(t, 21, Min3(21, 23))
	assert.Equal(t, int64(21), Min3(int64(21), int64(23)))
}
