# pcapshark
pcap file parser  implement by gin api

## note
- This project is based on github.com/wader/fq project modifications

### use
#### 1、run app and listen on port 8081
```shell
go run app.go
```
#### 2、post to upload pcap file. exp:
```shell
curl --location 'http://127.0.0.1:8081/upload' \
--header 'User-Agent: Reqable/2.30.3' \
--form 'upload=@"/C:/Users/xuanyu/Desktop/surica.pcap"'
```

#### 3、get json response data. exp:
```json
{
  "header": {
    "magic": "little_endian",
    "network": "ethernet",
    "sigfigs": 0,
    "snaplen": 262144,
    "thiszone": 0,
    "version_major": 2,
    "version_minor": 4
  },
  "packets_0_1": {
    "index_0": {
      "incl_len": 59,
      "orig_len": 59,
      "packet": {
        "destination": "fc:a0:5a:03:af:86",
        "ether_type": "ipv4",
        "payload": {
          "destination_ip": "10.52.2.133",
          "dont_fragment": true,
          "dscp": 0,
          "ecn": 0,
          "fragment_offset": 0,
          "header_checksum": 0,
          "identification": 14557,
          "ihl": 5,
          "more_fragments": false,
          "payload": {
            "ack": true,
            "acknowledgment_number": 709279170,
            "checksum": 12155,
            "cwr": false,
            "data_offset": 5,
            "destination_port": 8531,
            "ece": false,
            "fin": false,
            "ns": false,
            "payload": "0\r\n\r\n",
            "psh": true,
            "reserved": 0,
            "rst": false,
            "sequence_number": 2679815854,
            "source_port": 2598,
            "syn": false,
            "urg": false,
            "urgent_pointer": 0,
            "window_size": 32768
          },
          "protocol": "tcp",
          "reserved": 0,
          "source_ip": "10.52.24.111",
          "total_length": 45,
          "ttl": 128,
          "version": 4
        },
        "source": "34:cf:f6:fe:64:ec",
        "type_name": "ether8023_frame"
      },
      "ts_sec": 1740039511,
      "ts_usec": 465810,
      "type_name": "packet"
    }
  }
}
```
