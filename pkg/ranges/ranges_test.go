package ranges_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/xyu-io/pcapshark/pkg/ranges"
)

func TestRangeGaps(t *testing.T) {
	testCases := []struct {
		total    string // start_index:end_index
		ranges   string // start_index:length  [end_index = start_index + length]
		expected string // start_index:length  [start_index = pre_end; length = end_start - start_index]
	}{
		{"0:0", "", "0:0"},
		{"0:10", "", "0:10"},

		{"0:10", "0:10", ""},
		{"0:10", "0:0", "0:10"},

		{"0:10", "1:9", "0:1"}, // 0 1 1 1 1 1 1 1 1 1
		{"0:10", "0:9", "9:1"}, // 0 1 2 3 4 5 6 7 8 9 10

		{"0:10", "1:1 8:1", "0:1 2:6 9:1"}, // 0 1 1 0 0 0 0 0 1 1 0
		{"0:10", "1:1 2:5 8:1", "0:1 9:1"}, // 0 1 1 1 1 1 1 1 0 0
		{"0:10", "1:1 2:8 8:2", "0:1"},
		{"0:10", "0:4 2:8 8:2", ""},

		// handle empty ranges
		{"0:12", "4:4 8:0", "0:4 8:4"},
		{"0:12", "0:0 4:4", "0:4 8:4"},
		{"0:12", "0:0 4:4 8:0", "0:4 8:4"},
		{"0:12", "0:0 0:0 4:4 8:0 8:0", "0:4 8:4"},
		{"0:12", "8:0", "0:12"},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("%v_%v_%v", tC.total, tC.ranges, tC.expected), func(t *testing.T) {
			actual := ranges.Gaps(ranges.RangeFromString(tC.total), ranges.SliceFromString(tC.ranges))
			if !reflect.DeepEqual(ranges.SliceFromString(tC.expected), actual) {
				t.Errorf("expected %v, got %v", tC.expected, actual)
			}
		})
	}
}
