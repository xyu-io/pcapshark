package interp

import (
	"fmt"
	"io"
	"math/big"
	"strconv"
	"unicode/utf8"

	"github.com/xyu-io/pcapshark/internal/mathx"
	"github.com/xyu-io/pcapshark/pkg/bitio"
	"github.com/xyu-io/pcapshark/pkg/scalar"
)

func previewValue(v any, df scalar.DisplayFormat, opts *Options) string {
	switch vv := v.(type) {
	case bool:
		if vv {
			return "true"
		}
		return "false"
	case int:
		// TODO: DisplayFormat is weird
		return mathx.PadFormatInt(int64(vv), df.FormatBase(), true, 0)
	case int64:
		return mathx.PadFormatInt(vv, df.FormatBase(), true, 0)
	case uint64:
		return mathx.PadFormatUint(vv, df.FormatBase(), true, 0)
	case float64:
		// TODO: float32? better truncated to significant digits?
		return strconv.FormatFloat(vv, 'g', -1, 64)
	case string:
		runeLength := utf8.RuneCountInString(vv)
		if opts.StringTruncate != 0 && runeLength > opts.StringTruncate {
			runes := []rune(vv)
			vv = string(runes[0:opts.StringTruncate])
		}
		return strconv.Quote(vv)
	case nil:
		return "null"
	case bitio.Reader:
		r := bitio.NewIOReader(vv)
		all, err := io.ReadAll(r)
		if err != nil {
			return "raw bits"
		}

		return strconv.Quote(string(all))
	case Binary:
		r := vv.JQValueToGoJQ()
		if _, ok := r.(string); ok {
			return r.(string)
		}
		return "raw binary data"
	case *big.Int:
		return mathx.PadFormatBigInt(vv, df.FormatBase(), true, 0)
	case map[string]any:
		return "{}"
	case []any:
		return "[]"

	default:
		panic(fmt.Sprintf("unreachable %v (%T)", v, v))
	}
}
