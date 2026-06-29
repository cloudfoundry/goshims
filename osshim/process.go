package osshim

import "os"

//go:generate counterfeiter -o os_fake/fake_process.go . Process
type Process interface {
	Signal(sig os.Signal) error
}
