package shark

import (
	"fmt"
	_ "github.com/xyu-io/pcapshark/format/all"
	"github.com/xyu-io/pcapshark/shark/sharkos"
	"testing"
)

func TestMyOS(t *testing.T) {
	fileList := map[string]string{
		"normal": "E:\\my-git\\fq-master\\pcapshark\\testdata\\local.pcap",
		"stream": "E:\\my-git\\fq-master\\pcapshark\\testdata\\tmp.pcap",
	}

	for tp, fileName := range fileList {
		switch tp {
		case "stream":
			fmt.Println("---------------------------------stream---------------------------------")
			var out = make(chan string)
			go func() {
				for res := range out {
					fmt.Println(res)
				}
			}()
			err := sharkos.Exec(fileName, out)
			if err != nil {
				t.Log(err.Error())
				return
			}

		default:
			fmt.Println("---------------------------------default---------------------------------")
			var out = make(chan string)
			go func() {
				for res := range out {
					fmt.Println(res)
				}
			}()
			err := sharkos.DefaultExec(fileName, out)
			if err != nil {
				t.Log(err.Error())
				return
			}
		}
	}
}
