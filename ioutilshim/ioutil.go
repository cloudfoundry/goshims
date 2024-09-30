// This file was generated by counterfeiter
// with command: counterfeiter -p -o /Users/pivotal/workspace/local-volume-release/src/code.cloudfoundry.org/goshims/ioutilshim io/ioutil
// BUT THEN WE MODIFIED IT SO MAKE SURE TO COPY THOSE MODIFICATIONS FORWARDS
package ioutilshim

import (
	"code.cloudfoundry.org/goshims/osshim"
	"io"
	"os"
)

//go:generate counterfeiter -o ioutil_fake/fake_ioutil.go . Ioutil
type Ioutil interface {
	ReadAll(r io.Reader) ([]byte, error)
	ReadFile(filename string) ([]byte, error)
	WriteFile(filename string, data []byte, perm os.FileMode) error
	ReadDir(dirname string) ([]os.FileInfo, error)
	NopCloser(r io.Reader) io.ReadCloser
	TempFile(dir, prefix string) (f osshim.File, err error)
	TempDir(dir, prefix string) (name string, err error)
}
