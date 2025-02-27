package math

import (
	"embed"

	"github.com/xyu-io/pcapshark/pkg/interp"
)

//go:embed radix.jq
var radixFS embed.FS

func init() {
	interp.RegisterFS(radixFS)
}
