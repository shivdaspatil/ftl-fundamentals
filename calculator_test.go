package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func almostEqual(a, b float64) bool {
	const float64EqualityThreshold = 1e-9
	return math.Abs(a-b) <= float64EqualityThreshold
}

func TestAdd(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		a    float64
		b    float64
	}{
		{
			name: "Two positive numbers",
			want: 5,
			a:    3,
			b:    2,
		},
		{
			name: "Zero subtracted with any number",
			want: -2,
			a:    0,
			b:    -2,
		},
		{
			name: "Positive and negative number",
			want: -1,
			a:    -5,
			b:    4,
		},
		{
			name: "Two negative numbers",
			want: -11,
			a:    -8,
			b:    -3,
		},
		{
			name: "Two fractional numbers",
			want: 6.5,
			a:    3.3,
			b:    3.2,
		},
	}

	for _, tc := range tcs {
		got := calculator.Add(tc.a, tc.b)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test Add(%f, %f) (%s) failed: want %f, got %f",
				tc.a, tc.b, tc.name, tc.want, got)
		}
	}
}

func TestAddN(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		nums []float64
	}{
		{
			name: "Two positive numbers",
			want: 3,
			nums: []float64{1, 2},
		},
		{
			name: "Single number in slice ",
			want: 5,
			nums: []float64{5},
		},
		{
			name: "Empty Slice ",
			want: 0,
			nums: []float64{},
		},
		{
			name: "More than 2 numbers",
			want: 10,
			nums: []float64{1, 2, 3, 4},
		},
		{
			name: "More than 2 fractions",
			want: 10.6,
			nums: []float64{1.1, 2.2, 3.1, 4.2},
		},
	}

	for _, tc := range tcs {
		got := calculator.AddN(tc.nums)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test AddN(%v) (%s) failed: want %f, got %f",
				tc.nums, tc.name, tc.want, got)
		}
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i <= 100; i++ {
		randA := rand.Float64() * float64(rand.Intn(10))
		randB := rand.Float64() * float64(rand.Intn(10))
		want := randA + randB
		got := calculator.Add(randA, randB)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		a    float64
		b    float64
	}{
		{
			name: "Two positive numbers",
			want: 3,
			a:    5,
			b:    2,
		},
		{
			name: "Zero subtracted with any number",
			want: -2,
			a:    0,
			b:    2,
		},
		{
			name: "Positive and negative number",
			want: -9,
			a:    -5,
			b:    4,
		},
		{
			name: "Two negative numbers",
			want: -5,
			a:    -8,
			b:    -3,
		},
		{
			name: "Two fractional numbers",
			want: 0.1,
			a:    3.3,
			b:    3.2,
		},
	}

	for _, tc := range tcs {
		got := calculator.Subtract(tc.a, tc.b)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test Subtract(%f, %f) (%s) failed: want %f, got %f",
				tc.a, tc.b, tc.name, tc.want, got)
		}
	}
}

func TestSubtractN(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		nums []float64
	}{
		{
			name: "Two positive numbers",
			want: -1,
			nums: []float64{1, 2},
		},
		{
			name: "Single number in slice ",
			want: 5,
			nums: []float64{5},
		},
		{
			name: "Empty Slice ",
			want: 0,
			nums: []float64{},
		},
		{
			name: "More than 2 numbers",
			want: -8,
			nums: []float64{1, 2, 3, 4},
		},
		{
			name: "More than 2 fractions",
			want: 2.2,
			nums: []float64{5.5, 1.1, 2.2},
		},
	}

	for _, tc := range tcs {
		got := calculator.SubtractN(tc.nums)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test SubtractN(%v) (%s) failed: want %f, got %f",
				tc.nums, tc.name, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		a    float64
		b    float64
	}{
		{
			name: "Two positive numbers",
			want: 10,
			a:    5,
			b:    2,
		},
		{
			name: "Zero multiplied with any number",
			want: 0,
			a:    0,
			b:    2,
		},
		{
			name: "Positive and negative number",
			want: -20,
			a:    -5,
			b:    4,
		},
		{
			name: "Two negative numbers",
			want: 24,
			a:    -8,
			b:    -3,
		},
		{
			name: "Two fractional numbers",
			want: 4.03,
			a:    1.3,
			b:    3.1,
		},
	}

	for _, tc := range tcs {
		got := calculator.Multiply(tc.a, tc.b)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test Multiply(%f, %f) (%s) failed: want %f, got %f",
				tc.a, tc.b, tc.name, tc.want, got)
		}
	}
}

func TestMultiplyN(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name string
		want float64
		nums []float64
	}{
		{
			name: "Two positive numbers",
			want: 10,
			nums: []float64{5, 2},
		},
		{
			name: "Single number in slice ",
			want: 5,
			nums: []float64{5},
		},
		{
			name: "Empty Slice ",
			want: 0,
			nums: []float64{},
		},
		{
			name: "More than 2 numbers",
			want: 24,
			nums: []float64{1, 2, 3, 4},
		},
		{
			name: "More than 2 fractions",
			want: 13.31,
			nums: []float64{5.5, 1.1, 2.2},
		},
	}

	for _, tc := range tcs {
		got := calculator.MultiplyN(tc.nums)
		if !almostEqual(tc.want, got) {
			t.Errorf("Test MultiplyN(%v) (%s) failed: want %f, got %f",
				tc.nums, tc.name, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name        string
		want        float64
		a           float64
		b           float64
		errExpected bool
	}{
		{
			name:        "Two positive numbers",
			want:        5,
			a:           10,
			b:           2,
			errExpected: false,
		},
		{
			name:        "Zero divided with any number",
			want:        0,
			a:           0,
			b:           2,
			errExpected: false,
		},
		{
			name:        "Positive and negative number",
			want:        -1.25,
			a:           -5,
			b:           4,
			errExpected: false,
		},
		{
			name:        "Two negative numbers",
			want:        4,
			a:           -8,
			b:           -2,
			errExpected: false,
		},
		{
			name:        "Two fractional numbers",
			want:        2.75,
			a:           3.3,
			b:           1.2,
			errExpected: false,
		},
		{
			name:        "Divide by zero",
			want:        0,
			a:           2,
			b:           0,
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Errorf("Test Divide (%f, %f) (%s) failed: Unexpected error received",
				tc.a, tc.b, tc.name)
		} else {
			if !tc.errExpected && !almostEqual(tc.want, got) {
				t.Errorf("Test Divide (%f, %f) (%s) failed: want %f, got %f",
					tc.a, tc.b, tc.name, tc.want, got)
			}
		}
	}
}

func TestDivideN(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name        string
		want        float64
		nums        []float64
		errExpected bool
	}{
		{
			name:        "Two positive numbers",
			want:        5,
			nums:        []float64{10, 2},
			errExpected: false,
		},
		{
			name:        "Single number in slice ",
			want:        5,
			nums:        []float64{5},
			errExpected: false,
		},
		{
			name:        "Empty Slice ",
			want:        0,
			nums:        []float64{},
			errExpected: false,
		},
		{
			name:        "More than 2 numbers",
			want:        1,
			nums:        []float64{20, 5, 2, 2},
			errExpected: false,
		},
		{
			name:        "More than 2 fractions",
			want:        2,
			nums:        []float64{5.5, 1.1, 2.5},
			errExpected: false,
		},
		{
			name:        "More than 2 fractions starting with 0",
			want:        0,
			nums:        []float64{0, 1.1, 2.5},
			errExpected: false,
		},
		{
			name:        "More than 2 fractions",
			want:        0,
			nums:        []float64{5.5, 0, 2.2},
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		got, err := calculator.DivideN(tc.nums)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Errorf("Test DivideN (%v) (%s) failed: Unexpected error received",
				tc.nums, tc.name)
		} else {
			if !tc.errExpected && !almostEqual(tc.want, got) {
				t.Errorf("Test DivideN (%v) (%s) failed: want %f, got %f",
					tc.nums, tc.name, tc.want, got)
			}
		}
	}
}

func TestSquareRoot(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name        string
		want        float64
		a           float64
		errExpected bool
	}{
		{
			name:        "Squareroot of positive number",
			want:        4,
			a:           16,
			errExpected: false,
		},
		{
			name:        "Squareroot of 0",
			want:        0,
			a:           0,
			errExpected: false,
		},
		{
			name:        "Squareroot of negative number",
			want:        0,
			a:           -1,
			errExpected: true,
		},
	}

	for _, tc := range tcs {
		got, err := calculator.SquareRoot(tc.a)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Errorf("Test Squareroot (%f) (%s) failed: Unexpected error received",
				tc.a, tc.name)
		} else {
			if !tc.errExpected && !almostEqual(tc.want, got) {
				t.Errorf("Test Squareroot (%f) (%s) failed: want %f, got %f",
					tc.a, tc.name, tc.want, got)
			}
		}
	}
}

func TestComputeString(t *testing.T) {
	t.Parallel()
	var tcs = []struct {
		name        string
		input       string
		want        float64
		errExpected bool
	}{
		{
			name:        "Addition of two numbers with no spaces",
			input:       "2*2",
			want:        4,
			errExpected: false,
		},
		{
			name:        "Subtraction with leading space",
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
			name:        "No operator",
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
	}

	for _, tc := range tcs {
		got, err := calculator.ComputeString(tc.input)
		errReceived := err != nil
		if tc.errExpected != errReceived {
			t.Errorf("Test ComputeString (%s) (%s) failed: Unexpected error received",
				tc.input, tc.name)
		} else {
			if !tc.errExpected && !almostEqual(tc.want, got) {
				t.Errorf("Test ComputeString (%s) (%s) failed: want %f, got %f",
					tc.input, tc.name, tc.want, got)
			}
		}
	}
}
