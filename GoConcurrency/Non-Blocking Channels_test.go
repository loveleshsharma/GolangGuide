package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCakeMaker(t *testing.T) {

	outputChan := make(chan string)
	expectedKind := "chocolate"
	expectedTimes := 1

	go cakeMaker(expectedKind,expectedTimes,outputChan)

	var actualKind string
	select {
		case actualKind = <- outputChan :
	}

	assert.Equal(t,expectedKind,actualKind)
}
