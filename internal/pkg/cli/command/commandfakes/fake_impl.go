/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package commandfakes

import (
	"os"
	"os/exec"
	"sync"
)

type FakeImpl struct {
	CmdPidStub        func(*exec.Cmd) uint32
	cmdPidMutex       sync.RWMutex
	cmdPidArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdPidReturns struct {
		result1 uint32
	}
	cmdPidReturnsOnCall map[int]struct {
		result1 uint32
	}
	CmdStartStub        func(*exec.Cmd) error
	cmdStartMutex       sync.RWMutex
	cmdStartArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdStartReturns struct {
		result1 error
	}
	cmdStartReturnsOnCall map[int]struct {
		result1 error
	}
	CmdWaitStub        func(*exec.Cmd) error
	cmdWaitMutex       sync.RWMutex
	cmdWaitArgsForCall []struct {
		arg1 *exec.Cmd
	}
	cmdWaitReturns struct {
		result1 error
	}
	cmdWaitReturnsOnCall map[int]struct {
		result1 error
	}
	CommandStub        func(string, ...string) *exec.Cmd
	commandMutex       sync.RWMutex
	commandArgsForCall []struct {
		arg1 string
		arg2 []string
	}
	commandReturns struct {
		result1 *exec.Cmd
	}
	commandReturnsOnCall map[int]struct {
		result1 *exec.Cmd
	}
	GetHomeDirectoryStub        func(int) (string, error)
	getHomeDirectoryMutex       sync.RWMutex
	getHomeDirectoryArgsForCall []struct {
		arg1 int
	}
	getHomeDirectoryReturns struct {
		result1 string
		result2 error
	}
	getHomeDirectoryReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	NotifyStub        func(chan<- os.Signal, ...os.Signal)
	notifyMutex       sync.RWMutex
	notifyArgsForCall []struct {
		arg1 chan<- os.Signal
		arg2 []os.Signal
	}
	SignalStub        func(*exec.Cmd, os.Signal) error
	signalMutex       sync.RWMutex
	signalArgsForCall []struct {
		arg1 *exec.Cmd
		arg2 os.Signal
	}
	signalReturns struct {
		result1 error
	}
	signalReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) CmdPid(arg1 *exec.Cmd) uint32 {
	fake.cmdPidMutex.Lock()
	ret, specificReturn := fake.cmdPidReturnsOnCall[len(fake.cmdPidArgsForCall)]
	fake.cmdPidArgsForCall = append(fake.cmdPidArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdPidStub
	fakeReturns := fake.cmdPidReturns
	fake.recordInvocation("CmdPid", []interface{}{arg1})
	fake.cmdPidMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdPidCallCount() int {
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	return len(fake.cmdPidArgsForCall)
}

func (fake *FakeImpl) CmdPidCalls(stub func(*exec.Cmd) uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = stub
}

func (fake *FakeImpl) CmdPidArgsForCall(i int) *exec.Cmd {
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	argsForCall := fake.cmdPidArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdPidReturns(result1 uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = nil
	fake.cmdPidReturns = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeImpl) CmdPidReturnsOnCall(i int, result1 uint32) {
	fake.cmdPidMutex.Lock()
	defer fake.cmdPidMutex.Unlock()
	fake.CmdPidStub = nil
	if fake.cmdPidReturnsOnCall == nil {
		fake.cmdPidReturnsOnCall = make(map[int]struct {
			result1 uint32
		})
	}
	fake.cmdPidReturnsOnCall[i] = struct {
		result1 uint32
	}{result1}
}

func (fake *FakeImpl) CmdStart(arg1 *exec.Cmd) error {
	fake.cmdStartMutex.Lock()
	ret, specificReturn := fake.cmdStartReturnsOnCall[len(fake.cmdStartArgsForCall)]
	fake.cmdStartArgsForCall = append(fake.cmdStartArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdStartStub
	fakeReturns := fake.cmdStartReturns
	fake.recordInvocation("CmdStart", []interface{}{arg1})
	fake.cmdStartMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdStartCallCount() int {
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	return len(fake.cmdStartArgsForCall)
}

func (fake *FakeImpl) CmdStartCalls(stub func(*exec.Cmd) error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = stub
}

func (fake *FakeImpl) CmdStartArgsForCall(i int) *exec.Cmd {
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	argsForCall := fake.cmdStartArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdStartReturns(result1 error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = nil
	fake.cmdStartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdStartReturnsOnCall(i int, result1 error) {
	fake.cmdStartMutex.Lock()
	defer fake.cmdStartMutex.Unlock()
	fake.CmdStartStub = nil
	if fake.cmdStartReturnsOnCall == nil {
		fake.cmdStartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cmdStartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdWait(arg1 *exec.Cmd) error {
	fake.cmdWaitMutex.Lock()
	ret, specificReturn := fake.cmdWaitReturnsOnCall[len(fake.cmdWaitArgsForCall)]
	fake.cmdWaitArgsForCall = append(fake.cmdWaitArgsForCall, struct {
		arg1 *exec.Cmd
	}{arg1})
	stub := fake.CmdWaitStub
	fakeReturns := fake.cmdWaitReturns
	fake.recordInvocation("CmdWait", []interface{}{arg1})
	fake.cmdWaitMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CmdWaitCallCount() int {
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	return len(fake.cmdWaitArgsForCall)
}

func (fake *FakeImpl) CmdWaitCalls(stub func(*exec.Cmd) error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = stub
}

func (fake *FakeImpl) CmdWaitArgsForCall(i int) *exec.Cmd {
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	argsForCall := fake.cmdWaitArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CmdWaitReturns(result1 error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = nil
	fake.cmdWaitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CmdWaitReturnsOnCall(i int, result1 error) {
	fake.cmdWaitMutex.Lock()
	defer fake.cmdWaitMutex.Unlock()
	fake.CmdWaitStub = nil
	if fake.cmdWaitReturnsOnCall == nil {
		fake.cmdWaitReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cmdWaitReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Command(arg1 string, arg2 ...string) *exec.Cmd {
	fake.commandMutex.Lock()
	ret, specificReturn := fake.commandReturnsOnCall[len(fake.commandArgsForCall)]
	fake.commandArgsForCall = append(fake.commandArgsForCall, struct {
		arg1 string
		arg2 []string
	}{arg1, arg2})
	stub := fake.CommandStub
	fakeReturns := fake.commandReturns
	fake.recordInvocation("Command", []interface{}{arg1, arg2})
	fake.commandMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CommandCallCount() int {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	return len(fake.commandArgsForCall)
}

func (fake *FakeImpl) CommandCalls(stub func(string, ...string) *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = stub
}

func (fake *FakeImpl) CommandArgsForCall(i int) (string, []string) {
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	argsForCall := fake.commandArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) CommandReturns(result1 *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	fake.commandReturns = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeImpl) CommandReturnsOnCall(i int, result1 *exec.Cmd) {
	fake.commandMutex.Lock()
	defer fake.commandMutex.Unlock()
	fake.CommandStub = nil
	if fake.commandReturnsOnCall == nil {
		fake.commandReturnsOnCall = make(map[int]struct {
			result1 *exec.Cmd
		})
	}
	fake.commandReturnsOnCall[i] = struct {
		result1 *exec.Cmd
	}{result1}
}

func (fake *FakeImpl) GetHomeDirectory(arg1 int) (string, error) {
	fake.getHomeDirectoryMutex.Lock()
	ret, specificReturn := fake.getHomeDirectoryReturnsOnCall[len(fake.getHomeDirectoryArgsForCall)]
	fake.getHomeDirectoryArgsForCall = append(fake.getHomeDirectoryArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.GetHomeDirectoryStub
	fakeReturns := fake.getHomeDirectoryReturns
	fake.recordInvocation("GetHomeDirectory", []interface{}{arg1})
	fake.getHomeDirectoryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetHomeDirectoryCallCount() int {
	fake.getHomeDirectoryMutex.RLock()
	defer fake.getHomeDirectoryMutex.RUnlock()
	return len(fake.getHomeDirectoryArgsForCall)
}

func (fake *FakeImpl) GetHomeDirectoryCalls(stub func(int) (string, error)) {
	fake.getHomeDirectoryMutex.Lock()
	defer fake.getHomeDirectoryMutex.Unlock()
	fake.GetHomeDirectoryStub = stub
}

func (fake *FakeImpl) GetHomeDirectoryArgsForCall(i int) int {
	fake.getHomeDirectoryMutex.RLock()
	defer fake.getHomeDirectoryMutex.RUnlock()
	argsForCall := fake.getHomeDirectoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetHomeDirectoryReturns(result1 string, result2 error) {
	fake.getHomeDirectoryMutex.Lock()
	defer fake.getHomeDirectoryMutex.Unlock()
	fake.GetHomeDirectoryStub = nil
	fake.getHomeDirectoryReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetHomeDirectoryReturnsOnCall(i int, result1 string, result2 error) {
	fake.getHomeDirectoryMutex.Lock()
	defer fake.getHomeDirectoryMutex.Unlock()
	fake.GetHomeDirectoryStub = nil
	if fake.getHomeDirectoryReturnsOnCall == nil {
		fake.getHomeDirectoryReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getHomeDirectoryReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Notify(arg1 chan<- os.Signal, arg2 ...os.Signal) {
	fake.notifyMutex.Lock()
	fake.notifyArgsForCall = append(fake.notifyArgsForCall, struct {
		arg1 chan<- os.Signal
		arg2 []os.Signal
	}{arg1, arg2})
	stub := fake.NotifyStub
	fake.recordInvocation("Notify", []interface{}{arg1, arg2})
	fake.notifyMutex.Unlock()
	if stub != nil {
		fake.NotifyStub(arg1, arg2...)
	}
}

func (fake *FakeImpl) NotifyCallCount() int {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	return len(fake.notifyArgsForCall)
}

func (fake *FakeImpl) NotifyCalls(stub func(chan<- os.Signal, ...os.Signal)) {
	fake.notifyMutex.Lock()
	defer fake.notifyMutex.Unlock()
	fake.NotifyStub = stub
}

func (fake *FakeImpl) NotifyArgsForCall(i int) (chan<- os.Signal, []os.Signal) {
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	argsForCall := fake.notifyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) Signal(arg1 *exec.Cmd, arg2 os.Signal) error {
	fake.signalMutex.Lock()
	ret, specificReturn := fake.signalReturnsOnCall[len(fake.signalArgsForCall)]
	fake.signalArgsForCall = append(fake.signalArgsForCall, struct {
		arg1 *exec.Cmd
		arg2 os.Signal
	}{arg1, arg2})
	stub := fake.SignalStub
	fakeReturns := fake.signalReturns
	fake.recordInvocation("Signal", []interface{}{arg1, arg2})
	fake.signalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) SignalCallCount() int {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	return len(fake.signalArgsForCall)
}

func (fake *FakeImpl) SignalCalls(stub func(*exec.Cmd, os.Signal) error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = stub
}

func (fake *FakeImpl) SignalArgsForCall(i int) (*exec.Cmd, os.Signal) {
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	argsForCall := fake.signalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) SignalReturns(result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	fake.signalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) SignalReturnsOnCall(i int, result1 error) {
	fake.signalMutex.Lock()
	defer fake.signalMutex.Unlock()
	fake.SignalStub = nil
	if fake.signalReturnsOnCall == nil {
		fake.signalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.signalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cmdPidMutex.RLock()
	defer fake.cmdPidMutex.RUnlock()
	fake.cmdStartMutex.RLock()
	defer fake.cmdStartMutex.RUnlock()
	fake.cmdWaitMutex.RLock()
	defer fake.cmdWaitMutex.RUnlock()
	fake.commandMutex.RLock()
	defer fake.commandMutex.RUnlock()
	fake.getHomeDirectoryMutex.RLock()
	defer fake.getHomeDirectoryMutex.RUnlock()
	fake.notifyMutex.RLock()
	defer fake.notifyMutex.RUnlock()
	fake.signalMutex.RLock()
	defer fake.signalMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
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
