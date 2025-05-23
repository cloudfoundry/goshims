// Code generated by counterfeiter. DO NOT EDIT.
package exec_fake

import (
	"io"
	"sync"
	"syscall"

	"code.cloudfoundry.org/goshims/execshim"
)

type FakeCmd struct {
	CombinedOutputStub        func() ([]byte, error)
	combinedOutputMutex       sync.RWMutex
	combinedOutputArgsForCall []struct {
	}
	combinedOutputReturns struct {
		result1 []byte
		result2 error
	}
	combinedOutputReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	PidStub        func() int
	pidMutex       sync.RWMutex
	pidArgsForCall []struct {
	}
	pidReturns struct {
		result1 int
	}
	pidReturnsOnCall map[int]struct {
		result1 int
	}
	RunStub        func() error
	runMutex       sync.RWMutex
	runArgsForCall []struct {
	}
	runReturns struct {
		result1 error
	}
	runReturnsOnCall map[int]struct {
		result1 error
	}
	SetEnvStub        func([]string)
	setEnvMutex       sync.RWMutex
	setEnvArgsForCall []struct {
		arg1 []string
	}
	SetStderrStub        func(io.Writer)
	setStderrMutex       sync.RWMutex
	setStderrArgsForCall []struct {
		arg1 io.Writer
	}
	SetStdinStub        func(io.Reader)
	setStdinMutex       sync.RWMutex
	setStdinArgsForCall []struct {
		arg1 io.Reader
	}
	SetStdoutStub        func(io.Writer)
	setStdoutMutex       sync.RWMutex
	setStdoutArgsForCall []struct {
		arg1 io.Writer
	}
	StartStub        func() error
	startMutex       sync.RWMutex
	startArgsForCall []struct {
	}
	startReturns struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StderrPipeStub        func() (io.ReadCloser, error)
	stderrPipeMutex       sync.RWMutex
	stderrPipeArgsForCall []struct {
	}
	stderrPipeReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	stderrPipeReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	StdoutPipeStub        func() (io.ReadCloser, error)
	stdoutPipeMutex       sync.RWMutex
	stdoutPipeArgsForCall []struct {
	}
	stdoutPipeReturns struct {
		result1 io.ReadCloser
		result2 error
	}
	stdoutPipeReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 error
	}
	SysProcAttrStub        func() *syscall.SysProcAttr
	sysProcAttrMutex       sync.RWMutex
	sysProcAttrArgsForCall []struct {
	}
	sysProcAttrReturns struct {
		result1 *syscall.SysProcAttr
	}
	sysProcAttrReturnsOnCall map[int]struct {
		result1 *syscall.SysProcAttr
	}
	WaitStub        func() error
	waitMutex       sync.RWMutex
	waitArgsForCall []struct {
	}
	waitReturns struct {
		result1 error
	}
	waitReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCmd) CombinedOutput() ([]byte, error) {
	fake.combinedOutputMutex.Lock()
	ret, specificReturn := fake.combinedOutputReturnsOnCall[len(fake.combinedOutputArgsForCall)]
	fake.combinedOutputArgsForCall = append(fake.combinedOutputArgsForCall, struct {
	}{})
	stub := fake.CombinedOutputStub
	fakeReturns := fake.combinedOutputReturns
	fake.recordInvocation("CombinedOutput", []interface{}{})
	fake.combinedOutputMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCmd) CombinedOutputCallCount() int {
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	return len(fake.combinedOutputArgsForCall)
}

func (fake *FakeCmd) CombinedOutputCalls(stub func() ([]byte, error)) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = stub
}

func (fake *FakeCmd) CombinedOutputReturns(result1 []byte, result2 error) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = nil
	fake.combinedOutputReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) CombinedOutputReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.combinedOutputMutex.Lock()
	defer fake.combinedOutputMutex.Unlock()
	fake.CombinedOutputStub = nil
	if fake.combinedOutputReturnsOnCall == nil {
		fake.combinedOutputReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.combinedOutputReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) Pid() int {
	fake.pidMutex.Lock()
	ret, specificReturn := fake.pidReturnsOnCall[len(fake.pidArgsForCall)]
	fake.pidArgsForCall = append(fake.pidArgsForCall, struct {
	}{})
	stub := fake.PidStub
	fakeReturns := fake.pidReturns
	fake.recordInvocation("Pid", []interface{}{})
	fake.pidMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCmd) PidCallCount() int {
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	return len(fake.pidArgsForCall)
}

func (fake *FakeCmd) PidCalls(stub func() int) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = stub
}

func (fake *FakeCmd) PidReturns(result1 int) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	fake.pidReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeCmd) PidReturnsOnCall(i int, result1 int) {
	fake.pidMutex.Lock()
	defer fake.pidMutex.Unlock()
	fake.PidStub = nil
	if fake.pidReturnsOnCall == nil {
		fake.pidReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.pidReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeCmd) Run() error {
	fake.runMutex.Lock()
	ret, specificReturn := fake.runReturnsOnCall[len(fake.runArgsForCall)]
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
	}{})
	stub := fake.RunStub
	fakeReturns := fake.runReturns
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCmd) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeCmd) RunCalls(stub func() error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeCmd) RunReturns(result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	fake.runReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) RunReturnsOnCall(i int, result1 error) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = nil
	if fake.runReturnsOnCall == nil {
		fake.runReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.runReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) SetEnv(arg1 []string) {
	var arg1Copy []string
	if arg1 != nil {
		arg1Copy = make([]string, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.setEnvMutex.Lock()
	fake.setEnvArgsForCall = append(fake.setEnvArgsForCall, struct {
		arg1 []string
	}{arg1Copy})
	stub := fake.SetEnvStub
	fake.recordInvocation("SetEnv", []interface{}{arg1Copy})
	fake.setEnvMutex.Unlock()
	if stub != nil {
		fake.SetEnvStub(arg1)
	}
}

func (fake *FakeCmd) SetEnvCallCount() int {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	return len(fake.setEnvArgsForCall)
}

func (fake *FakeCmd) SetEnvCalls(stub func([]string)) {
	fake.setEnvMutex.Lock()
	defer fake.setEnvMutex.Unlock()
	fake.SetEnvStub = stub
}

func (fake *FakeCmd) SetEnvArgsForCall(i int) []string {
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	argsForCall := fake.setEnvArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCmd) SetStderr(arg1 io.Writer) {
	fake.setStderrMutex.Lock()
	fake.setStderrArgsForCall = append(fake.setStderrArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	stub := fake.SetStderrStub
	fake.recordInvocation("SetStderr", []interface{}{arg1})
	fake.setStderrMutex.Unlock()
	if stub != nil {
		fake.SetStderrStub(arg1)
	}
}

func (fake *FakeCmd) SetStderrCallCount() int {
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	return len(fake.setStderrArgsForCall)
}

func (fake *FakeCmd) SetStderrCalls(stub func(io.Writer)) {
	fake.setStderrMutex.Lock()
	defer fake.setStderrMutex.Unlock()
	fake.SetStderrStub = stub
}

func (fake *FakeCmd) SetStderrArgsForCall(i int) io.Writer {
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	argsForCall := fake.setStderrArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCmd) SetStdin(arg1 io.Reader) {
	fake.setStdinMutex.Lock()
	fake.setStdinArgsForCall = append(fake.setStdinArgsForCall, struct {
		arg1 io.Reader
	}{arg1})
	stub := fake.SetStdinStub
	fake.recordInvocation("SetStdin", []interface{}{arg1})
	fake.setStdinMutex.Unlock()
	if stub != nil {
		fake.SetStdinStub(arg1)
	}
}

func (fake *FakeCmd) SetStdinCallCount() int {
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	return len(fake.setStdinArgsForCall)
}

func (fake *FakeCmd) SetStdinCalls(stub func(io.Reader)) {
	fake.setStdinMutex.Lock()
	defer fake.setStdinMutex.Unlock()
	fake.SetStdinStub = stub
}

func (fake *FakeCmd) SetStdinArgsForCall(i int) io.Reader {
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	argsForCall := fake.setStdinArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCmd) SetStdout(arg1 io.Writer) {
	fake.setStdoutMutex.Lock()
	fake.setStdoutArgsForCall = append(fake.setStdoutArgsForCall, struct {
		arg1 io.Writer
	}{arg1})
	stub := fake.SetStdoutStub
	fake.recordInvocation("SetStdout", []interface{}{arg1})
	fake.setStdoutMutex.Unlock()
	if stub != nil {
		fake.SetStdoutStub(arg1)
	}
}

func (fake *FakeCmd) SetStdoutCallCount() int {
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	return len(fake.setStdoutArgsForCall)
}

func (fake *FakeCmd) SetStdoutCalls(stub func(io.Writer)) {
	fake.setStdoutMutex.Lock()
	defer fake.setStdoutMutex.Unlock()
	fake.SetStdoutStub = stub
}

func (fake *FakeCmd) SetStdoutArgsForCall(i int) io.Writer {
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	argsForCall := fake.setStdoutArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCmd) Start() error {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct {
	}{})
	stub := fake.StartStub
	fakeReturns := fake.startReturns
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCmd) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeCmd) StartCalls(stub func() error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = stub
}

func (fake *FakeCmd) StartReturns(result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) StartReturnsOnCall(i int, result1 error) {
	fake.startMutex.Lock()
	defer fake.startMutex.Unlock()
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) StderrPipe() (io.ReadCloser, error) {
	fake.stderrPipeMutex.Lock()
	ret, specificReturn := fake.stderrPipeReturnsOnCall[len(fake.stderrPipeArgsForCall)]
	fake.stderrPipeArgsForCall = append(fake.stderrPipeArgsForCall, struct {
	}{})
	stub := fake.StderrPipeStub
	fakeReturns := fake.stderrPipeReturns
	fake.recordInvocation("StderrPipe", []interface{}{})
	fake.stderrPipeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCmd) StderrPipeCallCount() int {
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	return len(fake.stderrPipeArgsForCall)
}

func (fake *FakeCmd) StderrPipeCalls(stub func() (io.ReadCloser, error)) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = stub
}

func (fake *FakeCmd) StderrPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = nil
	fake.stderrPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) StderrPipeReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.stderrPipeMutex.Lock()
	defer fake.stderrPipeMutex.Unlock()
	fake.StderrPipeStub = nil
	if fake.stderrPipeReturnsOnCall == nil {
		fake.stderrPipeReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.stderrPipeReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) StdoutPipe() (io.ReadCloser, error) {
	fake.stdoutPipeMutex.Lock()
	ret, specificReturn := fake.stdoutPipeReturnsOnCall[len(fake.stdoutPipeArgsForCall)]
	fake.stdoutPipeArgsForCall = append(fake.stdoutPipeArgsForCall, struct {
	}{})
	stub := fake.StdoutPipeStub
	fakeReturns := fake.stdoutPipeReturns
	fake.recordInvocation("StdoutPipe", []interface{}{})
	fake.stdoutPipeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCmd) StdoutPipeCallCount() int {
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	return len(fake.stdoutPipeArgsForCall)
}

func (fake *FakeCmd) StdoutPipeCalls(stub func() (io.ReadCloser, error)) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = stub
}

func (fake *FakeCmd) StdoutPipeReturns(result1 io.ReadCloser, result2 error) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = nil
	fake.stdoutPipeReturns = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) StdoutPipeReturnsOnCall(i int, result1 io.ReadCloser, result2 error) {
	fake.stdoutPipeMutex.Lock()
	defer fake.stdoutPipeMutex.Unlock()
	fake.StdoutPipeStub = nil
	if fake.stdoutPipeReturnsOnCall == nil {
		fake.stdoutPipeReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 error
		})
	}
	fake.stdoutPipeReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 error
	}{result1, result2}
}

func (fake *FakeCmd) SysProcAttr() *syscall.SysProcAttr {
	fake.sysProcAttrMutex.Lock()
	ret, specificReturn := fake.sysProcAttrReturnsOnCall[len(fake.sysProcAttrArgsForCall)]
	fake.sysProcAttrArgsForCall = append(fake.sysProcAttrArgsForCall, struct {
	}{})
	stub := fake.SysProcAttrStub
	fakeReturns := fake.sysProcAttrReturns
	fake.recordInvocation("SysProcAttr", []interface{}{})
	fake.sysProcAttrMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCmd) SysProcAttrCallCount() int {
	fake.sysProcAttrMutex.RLock()
	defer fake.sysProcAttrMutex.RUnlock()
	return len(fake.sysProcAttrArgsForCall)
}

func (fake *FakeCmd) SysProcAttrCalls(stub func() *syscall.SysProcAttr) {
	fake.sysProcAttrMutex.Lock()
	defer fake.sysProcAttrMutex.Unlock()
	fake.SysProcAttrStub = stub
}

func (fake *FakeCmd) SysProcAttrReturns(result1 *syscall.SysProcAttr) {
	fake.sysProcAttrMutex.Lock()
	defer fake.sysProcAttrMutex.Unlock()
	fake.SysProcAttrStub = nil
	fake.sysProcAttrReturns = struct {
		result1 *syscall.SysProcAttr
	}{result1}
}

func (fake *FakeCmd) SysProcAttrReturnsOnCall(i int, result1 *syscall.SysProcAttr) {
	fake.sysProcAttrMutex.Lock()
	defer fake.sysProcAttrMutex.Unlock()
	fake.SysProcAttrStub = nil
	if fake.sysProcAttrReturnsOnCall == nil {
		fake.sysProcAttrReturnsOnCall = make(map[int]struct {
			result1 *syscall.SysProcAttr
		})
	}
	fake.sysProcAttrReturnsOnCall[i] = struct {
		result1 *syscall.SysProcAttr
	}{result1}
}

func (fake *FakeCmd) Wait() error {
	fake.waitMutex.Lock()
	ret, specificReturn := fake.waitReturnsOnCall[len(fake.waitArgsForCall)]
	fake.waitArgsForCall = append(fake.waitArgsForCall, struct {
	}{})
	stub := fake.WaitStub
	fakeReturns := fake.waitReturns
	fake.recordInvocation("Wait", []interface{}{})
	fake.waitMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeCmd) WaitCallCount() int {
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	return len(fake.waitArgsForCall)
}

func (fake *FakeCmd) WaitCalls(stub func() error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = stub
}

func (fake *FakeCmd) WaitReturns(result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	fake.waitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) WaitReturnsOnCall(i int, result1 error) {
	fake.waitMutex.Lock()
	defer fake.waitMutex.Unlock()
	fake.WaitStub = nil
	if fake.waitReturnsOnCall == nil {
		fake.waitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.waitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCmd) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.combinedOutputMutex.RLock()
	defer fake.combinedOutputMutex.RUnlock()
	fake.pidMutex.RLock()
	defer fake.pidMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.setEnvMutex.RLock()
	defer fake.setEnvMutex.RUnlock()
	fake.setStderrMutex.RLock()
	defer fake.setStderrMutex.RUnlock()
	fake.setStdinMutex.RLock()
	defer fake.setStdinMutex.RUnlock()
	fake.setStdoutMutex.RLock()
	defer fake.setStdoutMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stderrPipeMutex.RLock()
	defer fake.stderrPipeMutex.RUnlock()
	fake.stdoutPipeMutex.RLock()
	defer fake.stdoutPipeMutex.RUnlock()
	fake.sysProcAttrMutex.RLock()
	defer fake.sysProcAttrMutex.RUnlock()
	fake.waitMutex.RLock()
	defer fake.waitMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCmd) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ execshim.Cmd = new(FakeCmd)
