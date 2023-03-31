package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomeFunction(t *testing.T) {
	tester := assert.New(t)
	tester.Nil(nil)
}
