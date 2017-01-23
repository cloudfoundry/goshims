package osshim

import (
	"os"
)

type FileShim struct {
	Delegate *os.File
}

func (f *FileShim) Fd() uintptr {
	return f.Fd()
}

func (f *FileShim) Close() error {
	return f.Close()
}

func (f *FileShim) Stat() (os.FileInfo, error) {
	return f.Stat()
}

func (f *FileShim) Read(b []byte) (n int, err error) {
	return f.Read(b)
}

func (f *FileShim) ReadAt(b []byte, off int64) (n int, err error) {
	return f.ReadAt(b, off)
}

func (f *FileShim) Write(b []byte) (n int, err error) {
	return f.Write(b)
}

func (f *FileShim) WriteAt(b []byte, off int64) (n int, err error) {
	return f.WriteAt(b, off)
}

func (f *FileShim) Seek(offset int64, whence int) (ret int64, err error) {
	return f.Seek(offset, whence)
}

func (f *FileShim) WriteString(s string) (n int, err error) {
	return f.WriteString(s)
}

func (f *FileShim) Chdir() error {
	return f.Chdir()
}