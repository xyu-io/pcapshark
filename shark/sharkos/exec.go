package sharkos

import (
	"context"
	"fmt"
	"github.com/xyu-io/pcapshark/pkg/interp"
	"os"
)

func Exec(fileName string, out chan string) error {
	pcapBytes, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	mos := NewMyOS(out)
	defer close(out)
	i, err := interp.New(mos, interp.DefaultRegistry)
	if err != nil {
		return err
	}
	iter, err := i.SetPcapByte(pcapBytes).EvalFuncWithoutArg(context.Background(), map[string]any{
		"args": []any{
			"",
			"-d",
			"pcap",
			"da",
			fileName, // 上面pcapBytes失败时候使用 - 内置原始文件加载方式
		},
		"version": "1.0.0",
	}, "_main", interp.EvalOpts{Output: mos.Stdout()})
	if err != nil {
		return err
	}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		fmt.Println(v)
	}
	return nil
}

func BytesExec(ctx context.Context, pcapBytes []byte, out chan string) error {
	mos := NewMyOS(out)
	defer close(out)

	i, err := interp.New(mos, interp.DefaultRegistry)
	if err != nil {
		return err
	}
	iter, err := i.SetPcapByte(pcapBytes).EvalFuncWithoutArg(ctx, map[string]any{
		"args": []any{
			"",
			"-d",
			"pcap",
			"da",
			"", // 上面pcapBytes失败时候使用 - 内置原始文件加载方式
		},
		"version": "1.0.0",
	}, "_main", interp.EvalOpts{Output: mos.Stdout()})
	if err != nil {
		return err
	}
	for {
		_, ok := iter.Next()
		if !ok {
			ctx.Done()
			break
		}
	}
	return nil
}

func DefaultExec(fileName string, out chan string) error {
	mos := NewMyOS(out)
	i, err := interp.New(mos, interp.DefaultRegistry)
	if err != nil {
		return err
	}
	defer close(out)
	iter, err := i.EvalFuncWithoutArg(context.Background(), map[string]any{
		"args": []any{
			"",
			"-d",
			"pcap",
			"da",
			fileName, // 上面pcapBytes失败时候使用 - 内置原始文件加载方式
		},
		"version": "1.0.0",
	}, "_main", interp.EvalOpts{
		//Filename: fileName,
		Output: mos.Stdout()})
	if err != nil {
		return err
	}
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		fmt.Println(v)
	}

	return nil
}
