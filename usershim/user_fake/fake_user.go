// Code generated by counterfeiter. DO NOT EDIT.
package user_fake

import (
	"os/user"
	"sync"

	"code.cloudfoundry.org/goshims/usershim"
)

type FakeUser struct {
	CurrentStub        func() (*user.User, error)
	currentMutex       sync.RWMutex
	currentArgsForCall []struct {
	}
	currentReturns struct {
		result1 *user.User
		result2 error
	}
	currentReturnsOnCall map[int]struct {
		result1 *user.User
		result2 error
	}
	LookupStub        func(string) (*user.User, error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		arg1 string
	}
	lookupReturns struct {
		result1 *user.User
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 *user.User
		result2 error
	}
	LookupGroupStub        func(string) (*user.Group, error)
	lookupGroupMutex       sync.RWMutex
	lookupGroupArgsForCall []struct {
		arg1 string
	}
	lookupGroupReturns struct {
		result1 *user.Group
		result2 error
	}
	lookupGroupReturnsOnCall map[int]struct {
		result1 *user.Group
		result2 error
	}
	LookupGroupIdStub        func(string) (*user.Group, error)
	lookupGroupIdMutex       sync.RWMutex
	lookupGroupIdArgsForCall []struct {
		arg1 string
	}
	lookupGroupIdReturns struct {
		result1 *user.Group
		result2 error
	}
	lookupGroupIdReturnsOnCall map[int]struct {
		result1 *user.Group
		result2 error
	}
	LookupIdStub        func(string) (*user.User, error)
	lookupIdMutex       sync.RWMutex
	lookupIdArgsForCall []struct {
		arg1 string
	}
	lookupIdReturns struct {
		result1 *user.User
		result2 error
	}
	lookupIdReturnsOnCall map[int]struct {
		result1 *user.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUser) Current() (*user.User, error) {
	fake.currentMutex.Lock()
	ret, specificReturn := fake.currentReturnsOnCall[len(fake.currentArgsForCall)]
	fake.currentArgsForCall = append(fake.currentArgsForCall, struct {
	}{})
	stub := fake.CurrentStub
	fakeReturns := fake.currentReturns
	fake.recordInvocation("Current", []interface{}{})
	fake.currentMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) CurrentCallCount() int {
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	return len(fake.currentArgsForCall)
}

func (fake *FakeUser) CurrentCalls(stub func() (*user.User, error)) {
	fake.currentMutex.Lock()
	defer fake.currentMutex.Unlock()
	fake.CurrentStub = stub
}

func (fake *FakeUser) CurrentReturns(result1 *user.User, result2 error) {
	fake.currentMutex.Lock()
	defer fake.currentMutex.Unlock()
	fake.CurrentStub = nil
	fake.currentReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) CurrentReturnsOnCall(i int, result1 *user.User, result2 error) {
	fake.currentMutex.Lock()
	defer fake.currentMutex.Unlock()
	fake.CurrentStub = nil
	if fake.currentReturnsOnCall == nil {
		fake.currentReturnsOnCall = make(map[int]struct {
			result1 *user.User
			result2 error
		})
	}
	fake.currentReturnsOnCall[i] = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Lookup(arg1 string) (*user.User, error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LookupStub
	fakeReturns := fake.lookupReturns
	fake.recordInvocation("Lookup", []interface{}{arg1})
	fake.lookupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeUser) LookupCalls(stub func(string) (*user.User, error)) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = stub
}

func (fake *FakeUser) LookupArgsForCall(i int) string {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	argsForCall := fake.lookupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUser) LookupReturns(result1 *user.User, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupReturnsOnCall(i int, result1 *user.User, result2 error) {
	fake.lookupMutex.Lock()
	defer fake.lookupMutex.Unlock()
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 *user.User
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroup(arg1 string) (*user.Group, error) {
	fake.lookupGroupMutex.Lock()
	ret, specificReturn := fake.lookupGroupReturnsOnCall[len(fake.lookupGroupArgsForCall)]
	fake.lookupGroupArgsForCall = append(fake.lookupGroupArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LookupGroupStub
	fakeReturns := fake.lookupGroupReturns
	fake.recordInvocation("LookupGroup", []interface{}{arg1})
	fake.lookupGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) LookupGroupCallCount() int {
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	return len(fake.lookupGroupArgsForCall)
}

func (fake *FakeUser) LookupGroupCalls(stub func(string) (*user.Group, error)) {
	fake.lookupGroupMutex.Lock()
	defer fake.lookupGroupMutex.Unlock()
	fake.LookupGroupStub = stub
}

func (fake *FakeUser) LookupGroupArgsForCall(i int) string {
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	argsForCall := fake.lookupGroupArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUser) LookupGroupReturns(result1 *user.Group, result2 error) {
	fake.lookupGroupMutex.Lock()
	defer fake.lookupGroupMutex.Unlock()
	fake.LookupGroupStub = nil
	fake.lookupGroupReturns = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroupReturnsOnCall(i int, result1 *user.Group, result2 error) {
	fake.lookupGroupMutex.Lock()
	defer fake.lookupGroupMutex.Unlock()
	fake.LookupGroupStub = nil
	if fake.lookupGroupReturnsOnCall == nil {
		fake.lookupGroupReturnsOnCall = make(map[int]struct {
			result1 *user.Group
			result2 error
		})
	}
	fake.lookupGroupReturnsOnCall[i] = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroupId(arg1 string) (*user.Group, error) {
	fake.lookupGroupIdMutex.Lock()
	ret, specificReturn := fake.lookupGroupIdReturnsOnCall[len(fake.lookupGroupIdArgsForCall)]
	fake.lookupGroupIdArgsForCall = append(fake.lookupGroupIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LookupGroupIdStub
	fakeReturns := fake.lookupGroupIdReturns
	fake.recordInvocation("LookupGroupId", []interface{}{arg1})
	fake.lookupGroupIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) LookupGroupIdCallCount() int {
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	return len(fake.lookupGroupIdArgsForCall)
}

func (fake *FakeUser) LookupGroupIdCalls(stub func(string) (*user.Group, error)) {
	fake.lookupGroupIdMutex.Lock()
	defer fake.lookupGroupIdMutex.Unlock()
	fake.LookupGroupIdStub = stub
}

func (fake *FakeUser) LookupGroupIdArgsForCall(i int) string {
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	argsForCall := fake.lookupGroupIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUser) LookupGroupIdReturns(result1 *user.Group, result2 error) {
	fake.lookupGroupIdMutex.Lock()
	defer fake.lookupGroupIdMutex.Unlock()
	fake.LookupGroupIdStub = nil
	fake.lookupGroupIdReturns = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupGroupIdReturnsOnCall(i int, result1 *user.Group, result2 error) {
	fake.lookupGroupIdMutex.Lock()
	defer fake.lookupGroupIdMutex.Unlock()
	fake.LookupGroupIdStub = nil
	if fake.lookupGroupIdReturnsOnCall == nil {
		fake.lookupGroupIdReturnsOnCall = make(map[int]struct {
			result1 *user.Group
			result2 error
		})
	}
	fake.lookupGroupIdReturnsOnCall[i] = struct {
		result1 *user.Group
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupId(arg1 string) (*user.User, error) {
	fake.lookupIdMutex.Lock()
	ret, specificReturn := fake.lookupIdReturnsOnCall[len(fake.lookupIdArgsForCall)]
	fake.lookupIdArgsForCall = append(fake.lookupIdArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LookupIdStub
	fakeReturns := fake.lookupIdReturns
	fake.recordInvocation("LookupId", []interface{}{arg1})
	fake.lookupIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeUser) LookupIdCallCount() int {
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	return len(fake.lookupIdArgsForCall)
}

func (fake *FakeUser) LookupIdCalls(stub func(string) (*user.User, error)) {
	fake.lookupIdMutex.Lock()
	defer fake.lookupIdMutex.Unlock()
	fake.LookupIdStub = stub
}

func (fake *FakeUser) LookupIdArgsForCall(i int) string {
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	argsForCall := fake.lookupIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeUser) LookupIdReturns(result1 *user.User, result2 error) {
	fake.lookupIdMutex.Lock()
	defer fake.lookupIdMutex.Unlock()
	fake.LookupIdStub = nil
	fake.lookupIdReturns = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) LookupIdReturnsOnCall(i int, result1 *user.User, result2 error) {
	fake.lookupIdMutex.Lock()
	defer fake.lookupIdMutex.Unlock()
	fake.LookupIdStub = nil
	if fake.lookupIdReturnsOnCall == nil {
		fake.lookupIdReturnsOnCall = make(map[int]struct {
			result1 *user.User
			result2 error
		})
	}
	fake.lookupIdReturnsOnCall[i] = struct {
		result1 *user.User
		result2 error
	}{result1, result2}
}

func (fake *FakeUser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.currentMutex.RLock()
	defer fake.currentMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.lookupGroupMutex.RLock()
	defer fake.lookupGroupMutex.RUnlock()
	fake.lookupGroupIdMutex.RLock()
	defer fake.lookupGroupIdMutex.RUnlock()
	fake.lookupIdMutex.RLock()
	defer fake.lookupIdMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUser) recordInvocation(key string, args []interface{}) {
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

var _ usershim.User = new(FakeUser)
