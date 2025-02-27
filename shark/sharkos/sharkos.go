package sharkos

import (
	"github.com/xyu-io/pcapshark/pkg/interp"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
)

type SharkOS struct {
	interp.Interp
	interruptChan chan struct{}
	pcapChan      chan string
}

func NewMyOS(outChan chan string) *SharkOS {

	return &SharkOS{
		pcapChan: outChan,
	}
}

func (m *SharkOS) Platform() interp.Platform {
	//TODO implement me
	panic("implement me 1")
}

// Stdin 本案例不需要命令行输入
func (m *SharkOS) Stdin() interp.Input {
	return nil
}

func (m *SharkOS) Stdout() interp.Output {
	return stdoutOutput{m, io.Discard}
}

func (m *SharkOS) Stderr() interp.Output {
	return stdoutOutput{m, io.Discard}
}

func (m *SharkOS) InterruptChan() chan struct{} {
	return m.interruptChan
}

func (m *SharkOS) Args() []string {
	return []string{}
}

func (m *SharkOS) Environ() []string {
	return os.Environ()
}

func (m *SharkOS) ConfigDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	fqDir := filepath.Join(configDir, "fq")

	if runtime.GOOS != "darwin" {
		return fqDir, nil
	}

	// this is to support fallback to ~/.config on macOS/darwin
	if _, err := os.Stat(fqDir); err == nil {
		return fqDir, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(homeDir, ".config", "fq"), nil
}

func (m *SharkOS) FS() fs.FS {
	return oSFS{}
}

func (m *SharkOS) Readline(opts interp.ReadlineOpts) (string, error) {
	return "", nil
}

func (m *SharkOS) History() ([]string, error) {
	return []string{}, nil
}
