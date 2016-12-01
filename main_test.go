package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	assert.Equal(t, "7", fizzBuzz(7))
	assert.Equal(t, "Fizz", fizzBuzz(18))
	assert.Equal(t, "Buzz", fizzBuzz(20))
	assert.Equal(t, "FizzBuzz", fizzBuzz(45))
}
