package calculator_test

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"calculator"
)

func TestAdd(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		input []float64
		want  float64
	}{
		{[]float64{2, 2, 2}, 6},
	}

	for _, tc := range testcases {
		var sb strings.Builder
		sb.WriteString("Add(")
		for i, v := range tc.input {
			if i > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		sb.WriteString(")")

		t.Run(sb.String(), func(t *testing.T) {
			got := calculator.Add(tc.input[0], tc.input[1], tc.input[2:]...)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		input []float64
		want  float64
	}{
		{[]float64{12, 4, 8}, 0},
	}

	for _, tc := range testcases {
		var sb strings.Builder
		sb.WriteString("Sub(")
		for i, v := range tc.input {
			if i > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		sb.WriteString(")")

		t.Run(sb.String(), func(t *testing.T) {
			got := calculator.Subtract(tc.input[0], tc.input[1], tc.input[2:]...)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		input []float64
		want  float64
	}{
		{[]float64{3, 3, 3}, 27},
	}

	for _, tc := range testcases {
		var sb strings.Builder
		sb.WriteString("Mul(")
		for i, v := range tc.input {
			if i > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		sb.WriteString(")")

		t.Run(sb.String(), func(t *testing.T) {
			got := calculator.Multiply(tc.input[0], tc.input[1], tc.input[2:]...)
			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		input       []float64
		want        float64
		errExpected bool
	}{
		{[]float64{12, 4, 3}, 1, false},
	}

	for _, tc := range testcases {
		var sb strings.Builder
		sb.WriteString("Divide(")
		for i, v := range tc.input {
			if i > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		sb.WriteString(")")

		t.Run(sb.String(), func(t *testing.T) {
			got, err := calculator.Divide(tc.input[0], tc.input[1], tc.input[2:]...)
			if tc.errExpected && err == nil {
				t.Errorf("error expected")
			}

			if !tc.errExpected && err != nil {
				t.Error(err)
			}

			if tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < 100; i++ {
		input := make([]float64, 2, 3)
		input[0], input[1] = r.Float64()*100, r.Float64()*100

		n := r.Intn(cap(input) - len(input) + 1)
		for j := 0; j < n; j++ {
			input = append(input, r.Float64()*100)
		}

		var want float64
		for _, v := range input {
			want += v
		}

		var sb strings.Builder
		sb.WriteString("Add(")
		for i, v := range input {
			if i > 0 {
				sb.WriteString(", ")
			}

			sb.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		}
		sb.WriteString(")")

		t.Run(sb.String(), func(t *testing.T) {
			got := calculator.Add(input[0], input[1], input[2:]...)
			if want != got {
				t.Errorf("want %f, got %f", want, got)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	testcases := []struct {
		a           float64
		want        float64
		errExpected bool
	}{
		{-4, 0, true},
		{4, 2, false},
	}

	for _, tc := range testcases {
		t.Run(fmt.Sprintf("Sqrt(%f)", tc.a), func(t *testing.T) {
			got, err := calculator.Sqrt(tc.a)
			if tc.errExpected && err == nil {
				t.Errorf("error expected")
			}

			if !tc.errExpected && err != nil {
				t.Error(err)
			}

			if err == nil && tc.want != got {
				t.Errorf("want %f, got %f", tc.want, got)
			}
		})
	}
}
