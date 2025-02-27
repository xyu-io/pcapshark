package json

import (
	"embed"

	"github.com/xyu-io/pcapshark/pkg/interp"
)

//go:embed jq.jq
var jqFS embed.FS

func init() {
	interp.RegisterFS(jqFS)
}
