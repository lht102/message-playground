// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jobworker "github.com/lht102/message-playground/jobworker"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

type Service_Expecter struct {
	mock *mock.Mock
}

func (_m *Service) EXPECT() *Service_Expecter {
	return &Service_Expecter{mock: &_m.Mock}
}

// CreateJob provides a mock function with given fields: ctx, createJobCmd
func (_m *Service) CreateJob(ctx context.Context, createJobCmd jobworker.CreateJobCommand) (jobworker.Job, error) {
	ret := _m.Called(ctx, createJobCmd)

	var r0 jobworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, jobworker.CreateJobCommand) jobworker.Job); ok {
		r0 = rf(ctx, createJobCmd)
	} else {
		r0 = ret.Get(0).(jobworker.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, jobworker.CreateJobCommand) error); ok {
		r1 = rf(ctx, createJobCmd)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type Service_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//  - ctx context.Context
//  - createJobCmd jobworker.CreateJobCommand
func (_e *Service_Expecter) CreateJob(ctx interface{}, createJobCmd interface{}) *Service_CreateJob_Call {
	return &Service_CreateJob_Call{Call: _e.mock.On("CreateJob", ctx, createJobCmd)}
}

func (_c *Service_CreateJob_Call) Run(run func(ctx context.Context, createJobCmd jobworker.CreateJobCommand)) *Service_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobworker.CreateJobCommand))
	})
	return _c
}

func (_c *Service_CreateJob_Call) Return(_a0 jobworker.Job, _a1 error) *Service_CreateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetJob provides a mock function with given fields: ctx, _a1
func (_m *Service) GetJob(ctx context.Context, _a1 uuid.UUID) (jobworker.Job, error) {
	ret := _m.Called(ctx, _a1)

	var r0 jobworker.Job
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) jobworker.Job); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Get(0).(jobworker.Job)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Service_GetJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJob'
type Service_GetJob_Call struct {
	*mock.Call
}

// GetJob is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 uuid.UUID
func (_e *Service_Expecter) GetJob(ctx interface{}, _a1 interface{}) *Service_GetJob_Call {
	return &Service_GetJob_Call{Call: _e.mock.On("GetJob", ctx, _a1)}
}

func (_c *Service_GetJob_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *Service_GetJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *Service_GetJob_Call) Return(_a0 jobworker.Job, _a1 error) *Service_GetJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewService interface {
	mock.TestingT
	Cleanup(func())
}

// NewService creates a new instance of Service. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewService(t mockConstructorTestingTNewService) *Service {
	mock := &Service{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
