package dns

import (
	"github.com/xyu-io/pcapshark/format"
	"github.com/xyu-io/pcapshark/pkg/decode"
	"github.com/xyu-io/pcapshark/pkg/interp"
)

func init() {
	interp.RegisterFormat(
		format.DNS_TCP,
		&decode.Format{
			Description: "DNS packet (TCP)",
			Groups:      []*decode.Group{format.TCP_Stream},
			DecodeFn:    dnsTCPDecode,
		})
}

func dnsTCPDecode(d *decode.D) any {
	var tsi format.TCP_Stream_In
	if d.ArgAs(&tsi) {
		tsi.MustIsPort(d.Fatalf, format.TCPPortDomain)
	}
	return dnsDecode(d, true)
}
