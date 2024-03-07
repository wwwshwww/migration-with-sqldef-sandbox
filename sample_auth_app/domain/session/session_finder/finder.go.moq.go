// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package session_finder

import (
	"example_app/sample_auth_app/domain/session/session"
	"sync"
)

// Ensure, that FinderMock does implement Finder.
// If this is not the case, regenerate this file with moq.
var _ Finder = &FinderMock{}

// FinderMock is a mock implementation of Finder.
//
//	func TestSomethingThatUsesFinder(t *testing.T) {
//
//		// make and configure a mocked Finder
//		mockedFinder := &FinderMock{
//			FindFunc: func(fo FilteringOptions, so SortingOptions) ([]session.ID, error) {
//				panic("mock out the Find method")
//			},
//		}
//
//		// use mockedFinder in code that requires Finder
//		// and then make assertions.
//
//	}
type FinderMock struct {
	// FindFunc mocks the Find method.
	FindFunc func(fo FilteringOptions, so SortingOptions) ([]session.ID, error)

	// calls tracks calls to the methods.
	calls struct {
		// Find holds details about calls to the Find method.
		Find []struct {
			// Fo is the fo argument value.
			Fo FilteringOptions
			// So is the so argument value.
			So SortingOptions
		}
	}
	lockFind sync.RWMutex
}

// Find calls FindFunc.
func (mock *FinderMock) Find(fo FilteringOptions, so SortingOptions) ([]session.ID, error) {
	if mock.FindFunc == nil {
		panic("FinderMock.FindFunc: method is nil but Finder.Find was just called")
	}
	callInfo := struct {
		Fo FilteringOptions
		So SortingOptions
	}{
		Fo: fo,
		So: so,
	}
	mock.lockFind.Lock()
	mock.calls.Find = append(mock.calls.Find, callInfo)
	mock.lockFind.Unlock()
	return mock.FindFunc(fo, so)
}

// FindCalls gets all the calls that were made to Find.
// Check the length with:
//
//	len(mockedFinder.FindCalls())
func (mock *FinderMock) FindCalls() []struct {
	Fo FilteringOptions
	So SortingOptions
} {
	var calls []struct {
		Fo FilteringOptions
		So SortingOptions
	}
	mock.lockFind.RLock()
	calls = mock.calls.Find
	mock.lockFind.RUnlock()
	return calls
}
