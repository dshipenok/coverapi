package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Multiply_ByZero_ExpectZero(t *testing.T) {
	r := multiply(0, 10)

	assert.EqualValues(t, 0, r)
}

func Test_Divide_DivideByZero_ExpectInfinity(t *testing.T) {
	r := divide(1, 0)

	assert.Equal(t, math.Inf(1), r)
}
