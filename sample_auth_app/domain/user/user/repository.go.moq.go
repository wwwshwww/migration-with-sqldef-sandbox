// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package user

import (
	"sync"
)

// Ensure, that RepositoryMock does implement Repository.
// If this is not the case, regenerate this file with moq.
var _ Repository = &RepositoryMock{}

// RepositoryMock is a mock implementation of Repository.
//
//	func TestSomethingThatUsesRepository(t *testing.T) {
//
//		// make and configure a mocked Repository
//		mockedRepository := &RepositoryMock{
//			BulkDeleteFunc: func(iDs []ID) error {
//				panic("mock out the BulkDelete method")
//			},
//			BulkGetFunc: func(iDs []ID) ([]User, error) {
//				panic("mock out the BulkGet method")
//			},
//			BulkSaveFunc: func(users []User) error {
//				panic("mock out the BulkSave method")
//			},
//			GetFunc: func(iD ID) (User, error) {
//				panic("mock out the Get method")
//			},
//			SaveFunc: func(user User) error {
//				panic("mock out the Save method")
//			},
//		}
//
//		// use mockedRepository in code that requires Repository
//		// and then make assertions.
//
//	}
type RepositoryMock struct {
	// BulkDeleteFunc mocks the BulkDelete method.
	BulkDeleteFunc func(iDs []ID) error

	// BulkGetFunc mocks the BulkGet method.
	BulkGetFunc func(iDs []ID) ([]User, error)

	// BulkSaveFunc mocks the BulkSave method.
	BulkSaveFunc func(users []User) error

	// GetFunc mocks the Get method.
	GetFunc func(iD ID) (User, error)

	// SaveFunc mocks the Save method.
	SaveFunc func(user User) error

	// calls tracks calls to the methods.
	calls struct {
		// BulkDelete holds details about calls to the BulkDelete method.
		BulkDelete []struct {
			// IDs is the iDs argument value.
			IDs []ID
		}
		// BulkGet holds details about calls to the BulkGet method.
		BulkGet []struct {
			// IDs is the iDs argument value.
			IDs []ID
		}
		// BulkSave holds details about calls to the BulkSave method.
		BulkSave []struct {
			// Users is the users argument value.
			Users []User
		}
		// Get holds details about calls to the Get method.
		Get []struct {
			// ID is the iD argument value.
			ID ID
		}
		// Save holds details about calls to the Save method.
		Save []struct {
			// User is the user argument value.
			User User
		}
	}
	lockBulkDelete sync.RWMutex
	lockBulkGet    sync.RWMutex
	lockBulkSave   sync.RWMutex
	lockGet        sync.RWMutex
	lockSave       sync.RWMutex
}

// BulkDelete calls BulkDeleteFunc.
func (mock *RepositoryMock) BulkDelete(iDs []ID) error {
	if mock.BulkDeleteFunc == nil {
		panic("RepositoryMock.BulkDeleteFunc: method is nil but Repository.BulkDelete was just called")
	}
	callInfo := struct {
		IDs []ID
	}{
		IDs: iDs,
	}
	mock.lockBulkDelete.Lock()
	mock.calls.BulkDelete = append(mock.calls.BulkDelete, callInfo)
	mock.lockBulkDelete.Unlock()
	return mock.BulkDeleteFunc(iDs)
}

// BulkDeleteCalls gets all the calls that were made to BulkDelete.
// Check the length with:
//
//	len(mockedRepository.BulkDeleteCalls())
func (mock *RepositoryMock) BulkDeleteCalls() []struct {
	IDs []ID
} {
	var calls []struct {
		IDs []ID
	}
	mock.lockBulkDelete.RLock()
	calls = mock.calls.BulkDelete
	mock.lockBulkDelete.RUnlock()
	return calls
}

// BulkGet calls BulkGetFunc.
func (mock *RepositoryMock) BulkGet(iDs []ID) ([]User, error) {
	if mock.BulkGetFunc == nil {
		panic("RepositoryMock.BulkGetFunc: method is nil but Repository.BulkGet was just called")
	}
	callInfo := struct {
		IDs []ID
	}{
		IDs: iDs,
	}
	mock.lockBulkGet.Lock()
	mock.calls.BulkGet = append(mock.calls.BulkGet, callInfo)
	mock.lockBulkGet.Unlock()
	return mock.BulkGetFunc(iDs)
}

// BulkGetCalls gets all the calls that were made to BulkGet.
// Check the length with:
//
//	len(mockedRepository.BulkGetCalls())
func (mock *RepositoryMock) BulkGetCalls() []struct {
	IDs []ID
} {
	var calls []struct {
		IDs []ID
	}
	mock.lockBulkGet.RLock()
	calls = mock.calls.BulkGet
	mock.lockBulkGet.RUnlock()
	return calls
}

// BulkSave calls BulkSaveFunc.
func (mock *RepositoryMock) BulkSave(users []User) error {
	if mock.BulkSaveFunc == nil {
		panic("RepositoryMock.BulkSaveFunc: method is nil but Repository.BulkSave was just called")
	}
	callInfo := struct {
		Users []User
	}{
		Users: users,
	}
	mock.lockBulkSave.Lock()
	mock.calls.BulkSave = append(mock.calls.BulkSave, callInfo)
	mock.lockBulkSave.Unlock()
	return mock.BulkSaveFunc(users)
}

// BulkSaveCalls gets all the calls that were made to BulkSave.
// Check the length with:
//
//	len(mockedRepository.BulkSaveCalls())
func (mock *RepositoryMock) BulkSaveCalls() []struct {
	Users []User
} {
	var calls []struct {
		Users []User
	}
	mock.lockBulkSave.RLock()
	calls = mock.calls.BulkSave
	mock.lockBulkSave.RUnlock()
	return calls
}

// Get calls GetFunc.
func (mock *RepositoryMock) Get(iD ID) (User, error) {
	if mock.GetFunc == nil {
		panic("RepositoryMock.GetFunc: method is nil but Repository.Get was just called")
	}
	callInfo := struct {
		ID ID
	}{
		ID: iD,
	}
	mock.lockGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	mock.lockGet.Unlock()
	return mock.GetFunc(iD)
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//
//	len(mockedRepository.GetCalls())
func (mock *RepositoryMock) GetCalls() []struct {
	ID ID
} {
	var calls []struct {
		ID ID
	}
	mock.lockGet.RLock()
	calls = mock.calls.Get
	mock.lockGet.RUnlock()
	return calls
}

// Save calls SaveFunc.
func (mock *RepositoryMock) Save(user User) error {
	if mock.SaveFunc == nil {
		panic("RepositoryMock.SaveFunc: method is nil but Repository.Save was just called")
	}
	callInfo := struct {
		User User
	}{
		User: user,
	}
	mock.lockSave.Lock()
	mock.calls.Save = append(mock.calls.Save, callInfo)
	mock.lockSave.Unlock()
	return mock.SaveFunc(user)
}

// SaveCalls gets all the calls that were made to Save.
// Check the length with:
//
//	len(mockedRepository.SaveCalls())
func (mock *RepositoryMock) SaveCalls() []struct {
	User User
} {
	var calls []struct {
		User User
	}
	mock.lockSave.RLock()
	calls = mock.calls.Save
	mock.lockSave.RUnlock()
	return calls
}
