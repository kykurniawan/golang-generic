package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindMin[T interface {
	int | int64 | float32 | float64
}](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestFindMin(t *testing.T) {
	assert.Equal(t, 20, FindMin(20, 40))
	assert.Equal(t, int64(20), FindMin(int64(20), int64(40)))
	assert.Equal(t, float64(20), FindMin(float64(20), float64(40)))
}

func GetFirst[T []E, E any](slice T) E {
	return slice[0]
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Rizky", "Kurniawan",
	}

	first := GetFirst(names)

	assert.Equal(t, "Rizky", first)
}
