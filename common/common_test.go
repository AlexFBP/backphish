package common_test

import (
	"testing"

	"github.com/AlexFBP/backphish/common"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{3, 1},
		{10, 2},
		{15, 2},
		{99, 2},
		{150, 3},
		{200, 3},
	}

	for _, test := range tests {
		result := common.Digits(test.input)
		if result != test.expected {
			t.Errorf("Digits(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
