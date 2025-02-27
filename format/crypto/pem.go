package crypto

import (
	"embed"

	"github.com/xyu-io/pcapshark/pkg/interp"
)

//go:embed pem.jq
var pemFS embed.FS

func init() {
	interp.RegisterFS(pemFS)
}
