package chap01

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMultiplication(t *testing.T) {
	five := &Dollar{Ammount: 5}
	five.Times(2)
	assert.Equal(t, 10, five.Ammount, "they should be equal")
}
