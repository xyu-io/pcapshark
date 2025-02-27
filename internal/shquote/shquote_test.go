package shquote_test

import (
	"reflect"
	"testing"

	"github.com/xyu-io/pcapshark/internal/shquote"
)

func TestSplit(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
	}{
		{``, nil},
		{`abbc`, []string{`abbc`}},
		{` abbc `, []string{`abbc`}},
		{`  a  bb  c  `, []string{`a`, `bb`, `c`}},
		{`\a`, []string{`a`}},
		{`a\bc`, []string{`abc`}},
		{`a bb c`, []string{`a`, `bb`, `c`}},
		{`"b b"`, []string{`b b`}},
		{`"b ' b"`, []string{`b ' b`}},
		{`"b \"b"`, []string{`b "b`}},
		{`'b b'`, []string{`b b`}},
		{`'b " b'`, []string{`b " b`}},
		{`'b \"b'`, []string{`b \"b`}},
		{`a'b'"c"`, []string{`abc`}},
		{`a "b"c"`, []string{`a`, `bc`}},
		{`a "b" c`, []string{`a`, `b`, `c`}},
		{`a"b"c`, []string{`abc`}},
		{`a'b'c`, []string{`abc`}},
		{`AB=dc abc`, []string{`AB=dc`, `abc`}},
		{`AB="d c" abc`, []string{`AB=d c`, `abc`}},
		{`AA="a a" BB=b abc`, []string{`AA=a a`, `BB=b`, `abc`}},
	}
	for _, tC := range testCases {
		t.Run(tC.input, func(t *testing.T) {
			actual := shquote.Split(tC.input)
			if !reflect.DeepEqual(tC.expected, actual) {
				t.Errorf("expected %#v, got %#v", tC.expected, actual)
			}
		})
	}
}
