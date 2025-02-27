package all

// Package all registers all builtin formats with the default registry

import (
	_ "github.com/xyu-io/pcapshark/format/asn1"
	_ "github.com/xyu-io/pcapshark/format/bits"
	_ "github.com/xyu-io/pcapshark/format/bzip2"
	_ "github.com/xyu-io/pcapshark/format/crypto"
	_ "github.com/xyu-io/pcapshark/format/dns"
	_ "github.com/xyu-io/pcapshark/format/gzip"
	_ "github.com/xyu-io/pcapshark/format/inet"
	_ "github.com/xyu-io/pcapshark/format/json"
	_ "github.com/xyu-io/pcapshark/format/markdown"
	_ "github.com/xyu-io/pcapshark/format/math"
	_ "github.com/xyu-io/pcapshark/format/pcap"
	_ "github.com/xyu-io/pcapshark/format/protobuf"
	_ "github.com/xyu-io/pcapshark/format/text"
	_ "github.com/xyu-io/pcapshark/format/tls"
	_ "github.com/xyu-io/pcapshark/format/xml"
	_ "github.com/xyu-io/pcapshark/format/yaml"
	_ "github.com/xyu-io/pcapshark/format/zip"
)
