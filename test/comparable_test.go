package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	return value1 == value1
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame("Rizky", "Rizky"))
	assert.True(t, IsSame(true, true))
	assert.True(t, IsSame(20, 20))
}
