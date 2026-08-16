package osshim

import "os"

type ProcessShim struct {
	*os.Process
}

func (p *ProcessShim) Signal(sig os.Signal) error {
	return p.Signal(sig)
}
