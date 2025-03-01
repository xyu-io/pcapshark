package bitio_test

import (
	"fmt"
	"testing"

	"github.com/xyu-io/pcapshark/pkg/bitio"
)

func TestReverseBytes64Panic(t *testing.T) {
	// TODO: check panic string
	defer func() { _ = recover() }()
	bitio.Read64([]byte{}, 0, 65)
	t.Error("should panic")
}

func TestReverseBytes64(t *testing.T) {
	testCases := []struct {
		nBits    int
		n        uint64
		expected uint64
	}{
		{nBits: 0, n: 0, expected: 0},
		{nBits: 8, n: 0x01, expected: 0x01},
		{nBits: 16, n: 0x0123, expected: 0x2301},
		{nBits: 24, n: 0x012345, expected: 0x452301},
		{nBits: 32, n: 0x01234567, expected: 0x67452301},
		{nBits: 40, n: 0x0123456789, expected: 0x8967452301},
		{nBits: 48, n: 0x0123456789ab, expected: 0xab8967452301},
		{nBits: 56, n: 0x0123456789abcd, expected: 0xcdab8967452301},
		{nBits: 64, n: 0x0123456789abcdef, expected: 0xefcdab8967452301},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%d %x %x", tC.nBits, tC.n, tC.expected), func(t *testing.T) {
			actual := bitio.ReverseBytes64(tC.nBits, tC.n)
			if tC.expected != actual {
				t.Errorf("expected %x, got %x", tC.expected, actual)
			}
		})
	}
}
