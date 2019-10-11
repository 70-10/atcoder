package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogic(t *testing.T) {
	tests := []struct {
		n       int
		k       int
		answers []int
		expect  []string
	}{
		{n: 6, k: 3, answers: []int{3, 1, 3, 2}, expect: []string{"No", "No", "Yes", "No", "No", "No"}},
		{n: 6, k: 5, answers: []int{3, 1, 3, 2}, expect: []string{"Yes", "Yes", "Yes", "Yes", "Yes", "Yes"}},
		{n: 10, k: 13, answers: []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5, 8, 9, 7, 9}, expect: []string{"No", "No", "No", "No", "Yes", "No", "No", "No", "Yes", "No"}},
	}

	for _, test := range tests {
		assert.Equal(t, test.expect, logic(test.n, test.k, test.answers))
	}
}
