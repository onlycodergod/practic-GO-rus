package util

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseIntPositiveValue(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(123)
	req.NoError(err)
}

func TestReverseIntNegativeValue(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(-649)
	req.NoError(err)
}

func TestReverseIntZeroValue(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(0)
	req.NoError(err)
}

func TestReverseIntBigPositiveValue(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(123456789012345678)
	req.NoError(err)
}

func TestReverseIntBigNegativeValue(t *testing.T) {
	req := require.New(t)
	_, err := ReverseInt(-123456789012345678)
	req.NoError(err)
}
