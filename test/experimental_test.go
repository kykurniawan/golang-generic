package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperimentalMin[T constraints.Ordered](first, second T) T {
	if first < second {
		return first
	}

	return second
}

func TestExperimentalMin(t *testing.T) {
	assert.Equal(t, 100, ExperimentalMin(100, 200))
	assert.Equal(t, float64(100), ExperimentalMin(float64(100), float64(200)))
}

func TestExperimentalMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Rizky",
	}

	second := map[string]string{
		"Name": "Rizky",
	}

	assert.True(t, maps.Equal(first, second))
}

func TestExperimentalSlices(t *testing.T) {
	first := []string{"Rizky"}
	second := []string{"Rizky"}

	assert.True(t, slices.Equal(first, second))
}
