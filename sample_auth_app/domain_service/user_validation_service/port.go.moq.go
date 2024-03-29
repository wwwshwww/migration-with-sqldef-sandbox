// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user_validation_service

import (
	"sync"
)

// Ensure, that PortMock does implement Port.
// If this is not the case, regenerate this file with moq.
var _ Port = &PortMock{}

// PortMock is a mock implementation of Port.
//
//	func TestSomethingThatUsesPort(t *testing.T) {
//
//		// make and configure a mocked Port
//		mockedPort := &PortMock{
//			CheckDuplicatedInExtSourceFunc: func(dupModels []dupModel) ([]bool, error) {
//				panic("mock out the CheckDuplicatedInExtSource method")
//			},
//		}
//
//		// use mockedPort in code that requires Port
//		// and then make assertions.
//
//	}
type PortMock struct {
	// CheckDuplicatedInExtSourceFunc mocks the CheckDuplicatedInExtSource method.
	CheckDuplicatedInExtSourceFunc func(dupModels []dupModel) ([]bool, error)

	// calls tracks calls to the methods.
	calls struct {
		// CheckDuplicatedInExtSource holds details about calls to the CheckDuplicatedInExtSource method.
		CheckDuplicatedInExtSource []struct {
			// DupModels is the dupModels argument value.
			DupModels []dupModel
		}
	}
	lockCheckDuplicatedInExtSource sync.RWMutex
}

// CheckDuplicatedInExtSource calls CheckDuplicatedInExtSourceFunc.
func (mock *PortMock) CheckDuplicatedInExtSource(dupModels []dupModel) ([]bool, error) {
	if mock.CheckDuplicatedInExtSourceFunc == nil {
		panic("PortMock.CheckDuplicatedInExtSourceFunc: method is nil but Port.CheckDuplicatedInExtSource was just called")
	}
	callInfo := struct {
		DupModels []dupModel
	}{
		DupModels: dupModels,
	}
	mock.lockCheckDuplicatedInExtSource.Lock()
	mock.calls.CheckDuplicatedInExtSource = append(mock.calls.CheckDuplicatedInExtSource, callInfo)
	mock.lockCheckDuplicatedInExtSource.Unlock()
	return mock.CheckDuplicatedInExtSourceFunc(dupModels)
}

// CheckDuplicatedInExtSourceCalls gets all the calls that were made to CheckDuplicatedInExtSource.
// Check the length with:
//
//	len(mockedPort.CheckDuplicatedInExtSourceCalls())
func (mock *PortMock) CheckDuplicatedInExtSourceCalls() []struct {
	DupModels []dupModel
} {
	var calls []struct {
		DupModels []dupModel
	}
	mock.lockCheckDuplicatedInExtSource.RLock()
	calls = mock.calls.CheckDuplicatedInExtSource
	mock.lockCheckDuplicatedInExtSource.RUnlock()
	return calls
}
