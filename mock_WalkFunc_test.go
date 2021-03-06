// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package pathlib

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// MockWalkFunc is an autogenerated mock type for the WalkFunc type
type MockWalkFunc struct {
	mock.Mock
}

// Execute provides a mock function with given fields: path, info, err
func (_m *MockWalkFunc) Execute(path *Path, info os.FileInfo, err error) error {
	ret := _m.Called(path, info, err)

	var r0 error
	if rf, ok := ret.Get(0).(func(*Path, os.FileInfo, error) error); ok {
		r0 = rf(path, info, err)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
