package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) <= 1e-9
}

func TestAdd(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name string
		a    float64
		b    float64
		nums []float64
		want float64
	}{
		{
			name: "Two positive numbers",
			a:    3,
			b:    2,
			want: 5,
		},
		{
			name: "Zero subtracted with any number",
			a:    0,
			b:    -2,
			want: -2,
		},
		{
			name: "Positive and negative number",
			a:    -5,
			b:    4,
			want: -1,
		},
		{
			name: "Two negative numbers",
			a:    -8,
			b:    -3,
			want: -11,
		},
		{
			name: "Two fractional numbers",
			a:    3.3,
			b:    3.2,
			want: 6.5,
		},
		{
			name: "More than two negative numbers",
			a:    -8,
			b:    -3,
			nums: []float64{-1, -2, -3},
			want: -17,
		},

		{
			name: "More than two fractional numbers",
			a:    3.3,
			b:    3.2,
			nums: []float64{2.1},
			want: 8.6,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Add(tc.a, tc.b, tc.nums...)
			if !closeEnough(tc.want, got) {
				t.Errorf("%s : Add(%f, %f, %v) : want %f, got %f",
					tc.name, tc.a, tc.b, tc.nums, tc.want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		t.Run("random_add", func(t *testing.T) {
			randA := rand.Float64() * float64(rand.Intn(10))
			randB := rand.Float64() * float64(rand.Intn(10))
			want := randA + randB
			got := calculator.Add(randA, randB)
			if want != got {
				t.Errorf("(Add random numbers): Add(%f, %f) : want %f, got %f",
					randA, randB, want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name string
		a    float64
		b    float64
		nums []float64
		want float64
	}{
		{
			name: "Two positive numbers",
			a:    5,
			b:    2,
			want: 3,
		},
		{
			name: "Zero subtracted with any number",
			a:    0,
			b:    2,
			want: -2,
		},
		{
			name: "Positive number subtracted from negative number",
			a:    -5,
			b:    4,
			want: -9,
		},
		{
			name: "Two negative numbers",
			a:    -8,
			b:    -3,
			want: -5,
		},
		{
			name: "Two fractional numbers",
			a:    3.3,
			b:    3.2,
			want: 0.1,
		},
		{
			name: "More than two negative numbers",
			a:    -8,
			b:    -3,
			nums: []float64{-1, -5, -10, -4},
			want: 15,
		},
		{
			name: "More than 2 fractions",
			a:    3.3,
			b:    3.2,
			nums: []float64{5.5, 1.1, 2.2},
			want: -8.7,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Subtract(tc.a, tc.b, tc.nums...)
			if !closeEnough(tc.want, got) {
				t.Errorf("%s : Subtract(%f, %f, %v) : want %f, got %f",
					tc.name, tc.a, tc.b, tc.nums, tc.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name string
		a    float64
		b    float64
		nums []float64
		want float64
	}{
		{
			name: "Two positive numbers",
			a:    5,
			b:    2,
			want: 10,
		},
		{
			name: "Zero multiplied with any number",
			a:    0,
			b:    2,
			want: 0,
		},
		{
			name: "Positive and negative number",
			a:    -5,
			b:    4,
			want: -20,
		},
		{
			name: "Two negative numbers",
			a:    -8,
			b:    -3,
			want: 24,
		},
		{
			name: "Two fractional numbers",
			a:    1.3,
			b:    3.1,
			want: 4.03,
		},
		{
			name: "More than 2 numbers",
			a:    2,
			b:    3,
			nums: []float64{2, 4},
			want: 48,
		},
		{
			name: "More than two fractional numbers",
			a:    5.5,
			b:    1.1,
			nums: []float64{2.2},
			want: 13.31,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got := calculator.Multiply(tc.a, tc.b, tc.nums...)
			if !closeEnough(tc.want, got) {
				t.Errorf("%s: Multiply(%f, %f, %v) : want %f, got %f",
					tc.name, tc.a, tc.b, tc.nums, tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		a           float64
		b           float64
		nums        []float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Two positive numbers",
			a:           10,
			b:           2,
			want:        5,
			errExpected: false,
		},
		{
			name:        "Zero divided with any number",
			a:           0,
			b:           2,
			want:        0,
			errExpected: false,
		},
		{
			name:        "Positive and negative number",
			a:           -5,
			b:           4,
			want:        -1.25,
			errExpected: false,
		},
		{
			name:        "Two negative numbers",
			a:           -8,
			b:           -2,
			want:        4,
			errExpected: false,
		},
		{
			name:        "Two fractional numbers",
			a:           3.3,
			b:           1.2,
			want:        2.75,
			errExpected: false,
		},
		{
			name:        "Divide by zero",
			a:           2,
			b:           0,
			want:        0,
			errExpected: true,
		},
		{
			name:        "More than 2 numbers",
			a:           40,
			b:           5,
			nums:        []float64{2, 2},
			want:        2,
			errExpected: false,
		},
		{
			name:        "More than 2 fractions",
			a:           5.5,
			b:           1.1,
			nums:        []float64{0.5, 0.2, 2.5},
			want:        20,
			errExpected: false,
		},
		{
			name:        "Divide by 0 in slice",
			a:           20,
			b:           2,
			nums:        []float64{2, 0},
			want:        0,
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Divide(tc.a, tc.b, tc.nums...)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Errorf("%s: Divide (%f, %f, %v) : Unexpected error status: %v",
					tc.name, tc.a, tc.b, tc.nums, err)
			}
			if !tc.errExpected && !closeEnough(tc.want, got) {
				t.Errorf("%s: Divide(%f, %f, %v) : want %f, got %f",
					tc.name, tc.a, tc.b, tc.nums, tc.want, got)
			}
		})
	}
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		a           float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Squareroot of positive number",
			a:           16,
			want:        4,
			errExpected: false,
		},
		{
			name:        "Squareroot of 0",
			a:           0,
			want:        0,
			errExpected: false,
		},
		{
			name:        "Squareroot of negative number",
			a:           -1,
			want:        0,
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.SquareRoot(tc.a)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Errorf("%s: Squareroot(%f) : Unexpected error received",
					tc.name, tc.a)
			}
			if !tc.errExpected && !closeEnough(tc.want, got) {
				t.Errorf("%s: Squareroot(%f) : want %f, got %f",
					tc.name, tc.a, tc.want, got)
			}
		})
	}
}

func TestEvaluate(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name        string
		input       string
		want        float64
		errExpected bool
	}{
		{
			name:        "Addition of two numbers with no spaces",
			input:       "2+2",
			want:        4,
			errExpected: false,
		},
		{
			name:        "Subtraction with embedded space",
			input:       "10     -     2",
			want:        8,
			errExpected: false,
		},
		{
			name:        "Multiplication with fraction",
			input:       "    4 *1.5",
			want:        6,
			errExpected: false,
		},
		{
			name:        "Division of two fraction",
			input:       "    4.4 /2.2",
			want:        2.0,
			errExpected: false,
		},
		{
			name:        "Unsupported operator",
			input:       "2%5",
			want:        0,
			errExpected: true,
		},
		{
			name:        "Single operand",
			input:       "2+",
			want:        0,
			errExpected: true,
		},
		{
			name:        "More than 2 operands",
			input:       "2+2+3",
			want:        0,
			errExpected: true,
		},
		{
			name:        "No operator ",
			input:       "2 2",
			want:        0,
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			got, err := calculator.Evaluate(tc.input)
			errReceived := err != nil
			if tc.errExpected != errReceived {
				t.Fatalf("%s: Evaluate(%s) : Unexpected error received",
					tc.name, tc.input)
			}
			if !tc.errExpected && !closeEnough(tc.want, got) {
				t.Errorf("%s: Evaluate(%s) : want %f, got %f",
					tc.name, tc.input, tc.want, got)
			}
		})
	}
}
