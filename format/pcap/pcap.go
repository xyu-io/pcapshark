package pcap

// https://wiki.wireshark.org/Development/LibpcapFileFormat
// TODO: tshark seems to not support sll2 in pcap, confusing

import (
	"embed"
	"github.com/xyu-io/pcapshark/format"
	"github.com/xyu-io/pcapshark/format/inet/flowsdecoder"
	"github.com/xyu-io/pcapshark/pkg/decode"
	"github.com/xyu-io/pcapshark/pkg/interp"
	"github.com/xyu-io/pcapshark/pkg/scalar"
)

//go:embed pcap.md
var pcapFS embed.FS

var pcapLinkFrameGroup decode.Group
var pcapTCPStreamGroup decode.Group
var pcapIPv4PacketGroup decode.Group

// writing application writes 0xa1b2c3d4 in native endian
const (
	// timestamp is seconds + microseconds
	bigEndian    = 0xa1b2c3d4
	littleEndian = 0xd4c3b2a1

	// timestamp is seconds + nanoseconds
	bigEndianNS    = 0xa1b23c4d
	littleEndianNS = 0x4d3cb2a1
)

var endianMap = scalar.UintMapSymStr{
	bigEndian:      "big_endian",
	littleEndian:   "little_endian",
	bigEndianNS:    "big_endian_ns",
	littleEndianNS: "little_endian_ns",
}

func init() {
	interp.RegisterFormat(
		format.PCAP,
		&decode.Format{
			Description: "PCAP packet capture",
			Groups:      []*decode.Group{format.Probe},
			Dependencies: []decode.Dependency{
				{Groups: []*decode.Group{format.Link_Frame}, Out: &pcapLinkFrameGroup},
				{Groups: []*decode.Group{format.TCP_Stream}, Out: &pcapTCPStreamGroup},
				{Groups: []*decode.Group{format.IPv4Packet}, Out: &pcapIPv4PacketGroup},
			},
			DecodeFn: decodePcap,
		})
	interp.RegisterFS(pcapFS)
}

func decodePcap(d *decode.D) any {
	var endian decode.Endian
	linkType := 0
	timestampUNSStr := "ts_usec"

	d.FieldStruct("header", func(d *decode.D) {
		magic := d.FieldU32("magic", d.UintAssert( //pcap文件标识 目前为“d4 c3 b2 a1”
			bigEndian,
			littleEndian,
			bigEndianNS,
			littleEndianNS,
		), endianMap, scalar.UintHex)

		switch magic {
		case bigEndian:
			endian = decode.BigEndian
		case littleEndian:
			endian = decode.LittleEndian
		case bigEndianNS:
			endian = decode.BigEndian
			timestampUNSStr = "ts_nsec"
		case littleEndianNS:
			endian = decode.LittleEndian
			timestampUNSStr = "ts_nsec"
		}

		d.Endian = endian

		d.FieldU16("version_major") // 2字节 主版本号
		d.FieldU16("version_minor") // 2字节 次版本号
		d.FieldS32("thiszone")      // 时区修正     并未使用
		d.FieldU32("sigfigs")       // 精确时间戳   并未使用
		d.FieldU32("snaplen")       // 抓包最大长度 如果要抓全，设为0x0000ffff
		linkType = int(d.FieldU32("network", format.LinkTypeMap))
	})

	d.Endian = endian
	fd := flowsdecoder.New(flowsdecoder.DecoderOptions{CheckTCPOptions: false})

	d.FieldArray("packets", func(d *decode.D) {
		for !d.End() {
			//todo stdout->yaml d.FieldStruct("packet", func(d *decode.D) {
			d.FieldStruct("", func(d *decode.D) {
				d.FieldValueStr("type_name", "packet")
				d.FieldU32("ts_sec")
				d.FieldU32(timestampUNSStr)
				inclLen := d.FieldU32("incl_len")
				origLen := d.FieldU32("orig_len")

				// "incl_len: the number of bytes of packet data actually captured and saved in the file. This value should never become larger than orig_len or the snaplen value of the global header"
				// "orig_len: the length of the packet as it appeared on the network when it was captured. If incl_len and orig_len differ, the actually saved packet size was limited by snaplen."

				// TODO: incl_len seems to be larger than snaplen in real pcap files
				// if inclLen > snapLen {
				// 	d.Errorf("incl_len %d > snaplen %d", inclLen, snapLen)
				// }

				if inclLen > origLen {
					d.Errorf("incl_len %d > orig_len %d", inclLen, origLen)
				}

				bs := d.ReadAllBits(d.BitBufRange(d.Pos(), int64(inclLen)*8))

				if fn, ok := linkToDecodeFn[linkType]; ok {
					// TODO: report decode errors
					_ = fn(fd, bs)
				}

				d.FieldFormatOrRawLen(
					"packet", //  "packet"
					int64(inclLen)*8,
					&pcapLinkFrameGroup,
					format.Link_Frame_In{
						Type:           linkType,
						IsLittleEndian: d.Endian == decode.LittleEndian,
					},
				)
			})
		}
	})
	fd.Flush()

	fieldFlows(d, fd, pcapTCPStreamGroup, pcapIPv4PacketGroup)

	return nil
}
