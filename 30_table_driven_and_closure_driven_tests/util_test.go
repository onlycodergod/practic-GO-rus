package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseInt(t *testing.T) {
	tests := map[string]struct {
		value int
		want  int
	}{
		"positive value":     {value: 123, want: 321},
		"negative value":     {value: -679, want: -976},
		"zero value":         {value: 0, want: 0},
		"big positive value": {value: 123456789012345678, want: 0},
		"big negative value": {value: -123456789012345678, want: 0},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			res, err := ReverseInt(testCase.value)
			req.Equal(testCase.want, res)
			req.NoError(err)
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	tests := map[string]struct {
		value []int
		want  bool
	}{
		"contain 1":     {value: []int{1, 1, 2, 1}, want: true},
		"contain 2":     {value: []int{1, 3, 2, 1}, want: true},
		"contain 3":     {value: []int{2, -1, 3, -1}, want: true},
		"contain 4":     {value: []int{2, -1, 0, 0}, want: true},
		"not contain 1": {value: []int{1, 2, 3, 4}, want: false},
		"not contain 2": {value: []int{0, 2, 3, 4}, want: false},
	}

	for name, testCase := range tests {
		t.Run(name, func(t *testing.T) {
			req := require.New(t)
			res := ContainsDuplicate(testCase.value)
			req.Equal(testCase.want, res)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	palindromCase := func(value int, want bool) func(t *testing.T) {
		return func(t *testing.T) {
			req := require.New(t)
			res := IsPalindrome(value)
			req.Equal(want, res)
		}
	}

	t.Run("palingrom 1", palindromCase(121, true))
	t.Run("palingrom 2", palindromCase(1234321, true))
	t.Run("palingrom 3", palindromCase(1, true))
	t.Run("palingrom 4", palindromCase(0, true))
	t.Run("not palingrom 1", palindromCase(123, false))
	t.Run("not palingrom 2", palindromCase(-121, false))
}
