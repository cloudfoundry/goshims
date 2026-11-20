package execshim

import (
	"code.cloudfoundry.org/goshims/osshim"
	"io"
	"syscall"
)

//go:generate counterfeiter -o exec_fake/fake_cmd.go . Cmd

type Cmd interface {
	Start() error
	SetStdout(io.Writer)
	SetStderr(io.Writer)
	SetStdin(io.Reader)
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Wait() error
	SetEnv([]string)
	Run() error
	CombinedOutput() ([]byte, error)
	Pid() int
	SysProcAttr() *syscall.SysProcAttr
	Process() osshim.Process
}
