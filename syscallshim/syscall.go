// Code generated by counterfeiter. DO NOT EDIT.
// with command: counterfeiter -p -o /Users/hjortj1/go/src/code.cloudfoundry.org/goshims/syscallshim syscall
// AND THEN WE REMOVED A MESS OF STUFF THAT DOESNT COMPILE, SO CAVEAT EMPTOR AND ALL THAT
package syscallshim

import "syscall"

//go:generate counterfeiter -o syscall_fake/fake_syscall.go . Syscall
type Syscall interface {
	ParseDirent(buf []byte, max int, names []string) (consumed int, count int, newnames []string)
	Getenv(key string) (value string, found bool)
	Setenv(key, value string) error
	Clearenv()
	Unsetenv(key string) error
	Environ() []string
	ForkExec(argv0 string, argv []string, attr *syscall.ProcAttr) (pid int, err error)
	StartProcess(argv0 string, argv []string, attr *syscall.ProcAttr) (pid int, handle uintptr, err error)
	Exec(argv0 string, argv []string, envv []string) (err error)
	CloseOnExec(fd int)
	SetNonblock(fd int, nonblocking bool) (err error)
	Close(fd int) error
	Dup(fd int) (int, error)
	Dup2(fd, newfd int) error
	Fstat(fd int, st *syscall.Stat_t) error
	Read(fd int, b []byte) (int, error)
	Write(fd int, b []byte) (int, error)
	Pread(fd int, b []byte, offset int64) (int, error)
	Pwrite(fd int, b []byte, offset int64) (int, error)
	Seek(fd int, offset int64, whence int) (int64, error)
	Pipe(fd []int) error
	FcntlFlock(fd uintptr, cmd int, lk *syscall.Flock_t) error
	ReadDirent(fd int, buf []byte) (int, error)
	Open(path string, openmode int, perm uint32) (fd int, err error)
	Mkdir(path string, perm uint32) error
	Stat(path string, st *syscall.Stat_t) error
	Lstat(path string, st *syscall.Stat_t) error
	Unlink(path string) error
	Rmdir(path string) error
	Chmod(path string, mode uint32) error
	Fchmod(fd int, mode uint32) error
	Chown(path string, uid, gid int) error
	Fchown(fd int, uid, gid int) error
	Lchown(path string, uid, gid int) error
	UtimesNano(path string, ts []syscall.Timespec) error
	Link(path, link string) error
	Rename(from, to string) error
	Truncate(path string, length int64) error
	Ftruncate(fd int, length int64) error
	Chdir(path string) error
	Fchdir(fd int) error
	Readlink(path string, buf []byte) (n int, err error)
	Symlink(path, link string) error
	Fsync(fd int) error
	Socket(proto, sotype, unused int) (fd int, err error)
	Bind(fd int, sa syscall.Sockaddr) error
	Listen(fd int, backlog int) error
	Accept(fd int) (newfd int, sa syscall.Sockaddr, err error)
	Getsockname(fd int) (sa syscall.Sockaddr, err error)
	Getpeername(fd int) (sa syscall.Sockaddr, err error)
	Connect(fd int, sa syscall.Sockaddr) error
	Recvfrom(fd int, p []byte, flags int) (n int, from syscall.Sockaddr, err error)
	Sendto(fd int, p []byte, flags int, to syscall.Sockaddr) error
	Recvmsg(fd int, p, oob []byte, flags int) (n, oobn, recvflags int, from syscall.Sockaddr, err error)
	Sendmsg(fd int, p, oob []byte, to syscall.Sockaddr, flags int) error
	SendmsgN(fd int, p, oob []byte, to syscall.Sockaddr, flags int) (n int, err error)
	GetsockoptInt(fd, level, opt int) (value int, err error)
	SetsockoptInt(fd, level, opt int, value int) error
	SetsockoptByte(fd, level, opt int, value byte) error
	SetsockoptLinger(fd, level, opt int, l *syscall.Linger) error
	Shutdown(fd int, how int) error
	SetsockoptICMPv6Filter(fd, level, opt int, filter *syscall.ICMPv6Filter) error
	SetsockoptIPMreq(fd, level, opt int, mreq *syscall.IPMreq) error
	SetsockoptIPv6Mreq(fd, level, opt int, mreq *syscall.IPv6Mreq) error
	SetsockoptInet4Addr(fd, level, opt int, value [4]byte) error
	SetsockoptString(fd, level, opt int, s string) error
	SetsockoptTimeval(fd, level, opt int, tv *syscall.Timeval) error
	Socketpair(domain, typ, proto int) (fd [2]int, err error)
	Getwd() (wd string, err error)
	CmsgLen(datalen int) int
	CmsgSpace(datalen int) int
	ParseSocketControlMessage(b []byte) ([]syscall.SocketControlMessage, error)
	UnixRights(fds ...int) []byte
	ParseUnixRights(m *syscall.SocketControlMessage) ([]int, error)
	ByteSliceFromString(s string) ([]byte, error)
	BytePtrFromString(s string) (*byte, error)
	Getpagesize() int
	Getgroups() (gids []int, err error)
	Setgroups(gids []int) (err error)
	Wait4(pid int, wstatus *syscall.WaitStatus, options int, rusage *syscall.Rusage) (wpid int, err error)
	GetsockoptInet4Addr(fd, level, opt int) (value [4]byte, err error)
	GetsockoptIPMreq(fd, level, opt int) (*syscall.IPMreq, error)
	GetsockoptIPv6Mreq(fd, level, opt int) (*syscall.IPv6Mreq, error)
	GetsockoptIPv6MTUInfo(fd, level, opt int) (*syscall.IPv6MTUInfo, error)
	GetsockoptICMPv6Filter(fd, level, opt int) (*syscall.ICMPv6Filter, error)
	Utimes(path string, tv []syscall.Timeval) (err error)
	Futimes(fd int, tv []syscall.Timeval) (err error)
	Mmap(fd int, offset int64, length int, prot int, flags int) (data []byte, err error)
	Munmap(b []byte) (err error)
	PtraceAttach(pid int) (err error)
	PtraceDetach(pid int) (err error)
	Kill(pid int, signum syscall.Signal) (err error)
	Gettimeofday(tv *syscall.Timeval) error
	Access(path string, mode uint32) (err error)
	Mknod(path string, mode uint32, dev int) (err error)
	Mkfifo(path string, mode uint32) (err error)
	Getpgrp() (pid int)
	Setuid(uid int) (err error)
	Setgid(gid int) (err error)
	Getrlimit(resource int, rlim *syscall.Rlimit) (err error)
	Setrlimit(resource int, rlim *syscall.Rlimit) (err error)
	Fstatfs(fd int, buf *syscall.Statfs_t) (err error)
	Statfs(path string, buf *syscall.Statfs_t) (err error)
	RawSyscall(trap, a1, a2, a3 uintptr) (r1, r2 uintptr, err syscall.Errno)
	RawSyscall6(trap, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)
	Getegid() int
	Geteuid() int
	Getgid() int
	Getppid() int
	Getpid() int
	Getuid() int
	Sendfile(outfd int, infd int, offset *int64, count int) (written int, err error)
	Exit(code int)
	NsecToTimeval(nsec int64) (tv syscall.Timeval)
	TimespecToNsec(ts syscall.Timespec) int64
	NsecToTimespec(nsec int64) (ts syscall.Timespec)
	TimevalToNsec(tv syscall.Timeval) int64
	Chroot(path string) (err error)
	Flock(fd int, how int) (err error)
	Getpgid(pid int) (pgid int, err error)
	Getpriority(which int, who int) (prio int, err error)
	Getrusage(who int, rusage *syscall.Rusage) (err error)
	Mlock(b []byte) (err error)
	Mlockall(flags int) (err error)
	Mprotect(b []byte, prot int) (err error)
	Munlock(b []byte) (err error)
	Munlockall() (err error)
	Setpgid(pid int, pgid int) (err error)
	Setpriority(which int, who int, prio int) (err error)
	Setregid(rgid int, egid int) (err error)
	Setreuid(ruid int, euid int) (err error)
	Setsid() (pid int, err error)
	Settimeofday(tp *syscall.Timeval) (err error)
	Umask(newmask int) (oldmask int)
	Faccessat(dirfd int, path string, mode uint32, flags int) (err error)
}
