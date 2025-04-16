package common_test

import (
	"math"
	"testing"

	"github.com/AlexFBP/backphish/common"
)

func TestDigits(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		// Casos básicos
		{3, 1},
		{10, 2},
		{15, 2},
		{99, 2},
		{150, 3},
		{200, 3},

		// Casos límite
		{0, 1},    // El número 0 tiene 1 dígito
		{-1, 1},   // Números negativos (si se consideran válidos)
		{-123, 3}, // Números negativos con más dígitos

		// Números grandes
		{1000000, 7},        // Un millón
		{123456789, 9},      // Nueve dígitos
		{math.MaxInt32, 10}, // Máximo valor de un int32
		{math.MaxInt64, 19}, // Máximo valor de un int64

		// Números pequeños negativos extremos
		{math.MinInt32, 10}, // Mínimo valor de un int32 (valor absoluto)
		{math.MinInt64, 19}, // Mínimo valor de un int64 (valor absoluto)
	}

	for _, test := range tests {
		result := common.Digits(test.input)
		if result != test.expected {
			t.Errorf("Digits(%d) = %d; want %d", test.input, result, test.expected)
		}
	}
}
