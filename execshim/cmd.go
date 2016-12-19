package execshim

import (
	"io"
	"os/exec"
	"syscall"
)

//go:generate counterfeiter -o exec_fake/fake_cmd.go . Cmd

type Cmd interface {
	Start() error
	StdoutPipe() (io.ReadCloser, error)
	StderrPipe() (io.ReadCloser, error)
	Wait() error
	Run() error
	CombinedOutput() ([]byte, error)

	SysProcAttr() *syscall.SysProcAttr
}

type cmd struct {
	*exec.Cmd
}

func (c *cmd) Start() error {
	return c.Cmd.Start()
}

func (c *cmd) StdoutPipe() (io.ReadCloser, error) {
	return c.Cmd.StdoutPipe()
}

func (c *cmd) StderrPipe() (io.ReadCloser, error) {
	return c.Cmd.StderrPipe()
}

func (c *cmd) Wait() error {
	return c.Wait()
}

func (c *cmd) Run() error {
	return c.Run()
}

func (c *cmd) CombinedOutput() ([]byte, error) {
	return c.Cmd.CombinedOutput()
}

func (c *cmd) SysProcAttr() *syscall.SysProcAttr {
	return c.Cmd.SysProcAttr
}
