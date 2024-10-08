// This file was generated by counterfeiter
// with command: counterfeiter -p -o /Users/pivotal/workspace/local-volume-release/src/code.cloudfoundry.org/goshims/osshim os
// BUT THEN WE MODIFIED IT SO MAKE SURE TO COPY THOSE MODIFICATIONS FORWARDS
package osshim

import (
	"os"
	"time"
)

type OsShim struct{}

func (sh *OsShim) FindProcess(pid int) (*os.Process, error) {
	return os.FindProcess(pid)
}

func (sh *OsShim) StartProcess(name string, argv []string, attr *os.ProcAttr) (*os.Process, error) {
	return os.StartProcess(name, argv, attr)
}

func (sh *OsShim) Hostname() (name string, err error) {
	return os.Hostname()
}

func (sh *OsShim) Expand(s string, mapping func(string) string) string {
	return os.Expand(s, mapping)
}

func (sh *OsShim) ExpandEnv(s string) string {
	return os.ExpandEnv(s)
}

func (sh *OsShim) Getenv(key string) string {
	return os.Getenv(key)
}

func (sh *OsShim) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func (sh *OsShim) Setenv(key string, value string) error {
	return os.Setenv(key, value)
}

func (sh *OsShim) Unsetenv(key string) error {
	return os.Unsetenv(key)
}

func (sh *OsShim) Clearenv() {
	os.Clearenv()
}

func (sh *OsShim) Environ() []string {
	return os.Environ()
}

func (sh *OsShim) NewSyscallError(syscall string, err error) error {
	return os.NewSyscallError(syscall, err)
}

func (sh *OsShim) IsExist(err error) bool {
	return os.IsExist(err)
}

func (sh *OsShim) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (sh *OsShim) IsPermission(err error) bool {
	return os.IsPermission(err)
}

func (sh *OsShim) Getpid() int {
	return os.Getpid()
}

func (sh *OsShim) Getppid() int {
	return os.Getppid()
}

func (sh *OsShim) Mkdir(name string, perm os.FileMode) error {
	return os.Mkdir(name, perm)
}

func (sh *OsShim) Chdir(dir string) error {
	return os.Chdir(dir)
}

func (sh *OsShim) Open(name string) (File, error) {
	o, err := os.Open(name)
	return &FileShim{Delegate: o}, err
}

func (sh *OsShim) Create(name string) (File, error) {
	o, err := os.Create(name)
	return &FileShim{Delegate: o}, err
}

func (sh *OsShim) Rename(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

func (sh *OsShim) NewFile(fd uintptr, name string) File {
	return &FileShim{Delegate: os.NewFile(fd, name)}
}

func (sh *OsShim) OpenFile(name string, flag int, perm os.FileMode) (File, error) {
	o, err := os.OpenFile(name, flag, perm)
	return &FileShim{Delegate: o}, err
}

func (sh *OsShim) Truncate(name string, size int64) error {
	return os.Truncate(name, size)
}

func (sh *OsShim) Remove(name string) error {
	return os.Remove(name)
}

func (sh *OsShim) Chmod(name string, mode os.FileMode) error {
	return os.Chmod(name, mode)
}

func (sh *OsShim) Chtimes(name string, atime time.Time, mtime time.Time) error {
	return os.Chtimes(name, atime, mtime)
}

func (sh *OsShim) Pipe() (File, File, error) {
	r, w, err := os.Pipe()
	return &FileShim{Delegate: r}, &FileShim{Delegate: w}, err
}

func (sh *OsShim) Link(oldname string, newname string) error {
	return os.Link(oldname, newname)
}

func (sh *OsShim) Symlink(oldname string, newname string) error {
	return os.Symlink(oldname, newname)
}

func (sh *OsShim) Readlink(name string) (string, error) {
	return os.Readlink(name)
}

func (sh *OsShim) ReadDir(name string) ([]os.DirEntry, error) {
	return os.ReadDir(name)
}

func (sh *OsShim) ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func (sh *OsShim) WriteFile(name string, data []byte, perm os.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (sh *OsShim) Chown(name string, uid int, gid int) error {
	return os.Chown(name, uid, gid)
}

func (sh *OsShim) Lchown(name string, uid int, gid int) error {
	return os.Lchown(name, uid, gid)
}

func (sh *OsShim) TempDir() string {
	return os.TempDir()
}

func (sh *OsShim) Stat(name string) (os.FileInfo, error) {
	return os.Stat(name)
}

func (sh *OsShim) Lstat(name string) (os.FileInfo, error) {
	return os.Lstat(name)
}

func (sh *OsShim) Getwd() (dir string, err error) {
	return os.Getwd()
}

func (sh *OsShim) MkdirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (sh *OsShim) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (sh *OsShim) IsPathSeparator(c uint8) bool {
	return os.IsPathSeparator(c)
}

func (sh *OsShim) Getuid() int {
	return os.Getuid()
}

func (sh *OsShim) Geteuid() int {
	return os.Geteuid()
}

func (sh *OsShim) Getgid() int {
	return os.Getgid()
}

func (sh *OsShim) Getegid() int {
	return os.Getegid()
}

func (sh *OsShim) Getgroups() ([]int, error) {
	return os.Getgroups()
}

func (sh *OsShim) Exit(code int) {
	os.Exit(code)
}

func (sh *OsShim) Getpagesize() int {
	return os.Getpagesize()
}

func (sh *OsShim) SameFile(fi1 os.FileInfo, fi2 os.FileInfo) bool {
	return os.SameFile(fi1, fi2)
}
