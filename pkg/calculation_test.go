package calculation_test

import (
	"testing"

	"github.com/lollmark/calculator_go/pkg"
)

func TestCalc(t *testing.T) {
	successCases := []struct {
		name       string
		expression string
		expected   float64
	}{
		{"Simple Addition", "7+3", 10},
		{"Subtraction with Zero", "15-0", 15},
		{"Multiplication by One", "9*1", 9},
		{"Division Resulting in Fraction", "7/4", 1.75},
		{"Complex Parentheses", "(10-3)*2", 14},
		{"Mixed Operations with Parentheses", "(8/4)+(5*2)", 12},
		{"Floating Point Multiplication", "3.2*2.5", 8},
	}

	for _, tc := range successCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calculation.Calc(tc.expression)
			if err != nil {
				t.Fatalf("Expected success but got error: %v", err)
			}
			if result != tc.expected {
				t.Fatalf("Expected %f but got %f", tc.expected, result)
			}
		})
	}

	failCases := []struct {
		name       string
		expression string
	}{
		{"Alphabet Character", "4+b"},
		{"Zero Division", "18/0"},
		{"Unclosed Parentheses", "(6+9"},
		{"Empty Input", ""},
		{"Multiple Consecutive Operators", "7**2"},
	}

	for _, tc := range failCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := calculation.Calc(tc.expression)
			if err == nil {
				t.Fatalf("Expected error but got success for expression: %s", tc.expression)
			}
		})
	}
}
