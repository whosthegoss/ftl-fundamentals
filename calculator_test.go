package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}

	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 6, b: 2, want: 4},
		{a: 1, b: 1, want: 0},
		{a: 7, b: 0, want: 7},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}

	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 6, b: 2, want: 12},
		{a: 1, b: 1, want: 1},
		{a: 7, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}

	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b        float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{a: 6, b: 2, want: 3, errExpected: false},
		{a: 1, b: 1, want: 1, errExpected: false},
		{a: 7, b: 0, want: 0, errExpected: true},
	}
	for _, tc := range testCases {
		got := calculator.Divide(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, errExpected %f got %f, errExpected: %f",
				tc.a, tc.b, tc.want, tc.errExpected, got)
		}

	}
}
