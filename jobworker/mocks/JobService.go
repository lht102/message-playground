// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	jobworker "github.com/lht102/message-playground/jobworker"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// JobService is an autogenerated mock type for the JobService type
type JobService struct {
	mock.Mock
}

type JobService_Expecter struct {
	mock *mock.Mock
}

func (_m *JobService) EXPECT() *JobService_Expecter {
	return &JobService_Expecter{mock: &_m.Mock}
}

// CreateJob provides a mock function with given fields: ctx, createJobCmd
func (_m *JobService) CreateJob(ctx context.Context, createJobCmd jobworker.CreateJobCommand) (jobworker.Job, error) {
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

// JobService_CreateJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateJob'
type JobService_CreateJob_Call struct {
	*mock.Call
}

// CreateJob is a helper method to define mock.On call
//  - ctx context.Context
//  - createJobCmd jobworker.CreateJobCommand
func (_e *JobService_Expecter) CreateJob(ctx interface{}, createJobCmd interface{}) *JobService_CreateJob_Call {
	return &JobService_CreateJob_Call{Call: _e.mock.On("CreateJob", ctx, createJobCmd)}
}

func (_c *JobService_CreateJob_Call) Run(run func(ctx context.Context, createJobCmd jobworker.CreateJobCommand)) *JobService_CreateJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(jobworker.CreateJobCommand))
	})
	return _c
}

func (_c *JobService_CreateJob_Call) Return(_a0 jobworker.Job, _a1 error) *JobService_CreateJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// ExecuteJob provides a mock function with given fields: ctx, _a1
func (_m *JobService) ExecuteJob(ctx context.Context, _a1 uuid.UUID) error {
	ret := _m.Called(ctx, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JobService_ExecuteJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteJob'
type JobService_ExecuteJob_Call struct {
	*mock.Call
}

// ExecuteJob is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 uuid.UUID
func (_e *JobService_Expecter) ExecuteJob(ctx interface{}, _a1 interface{}) *JobService_ExecuteJob_Call {
	return &JobService_ExecuteJob_Call{Call: _e.mock.On("ExecuteJob", ctx, _a1)}
}

func (_c *JobService_ExecuteJob_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *JobService_ExecuteJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *JobService_ExecuteJob_Call) Return(_a0 error) *JobService_ExecuteJob_Call {
	_c.Call.Return(_a0)
	return _c
}

// GetJob provides a mock function with given fields: ctx, _a1
func (_m *JobService) GetJob(ctx context.Context, _a1 uuid.UUID) (jobworker.Job, error) {
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

// JobService_GetJob_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetJob'
type JobService_GetJob_Call struct {
	*mock.Call
}

// GetJob is a helper method to define mock.On call
//  - ctx context.Context
//  - _a1 uuid.UUID
func (_e *JobService_Expecter) GetJob(ctx interface{}, _a1 interface{}) *JobService_GetJob_Call {
	return &JobService_GetJob_Call{Call: _e.mock.On("GetJob", ctx, _a1)}
}

func (_c *JobService_GetJob_Call) Run(run func(ctx context.Context, _a1 uuid.UUID)) *JobService_GetJob_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(uuid.UUID))
	})
	return _c
}

func (_c *JobService_GetJob_Call) Return(_a0 jobworker.Job, _a1 error) *JobService_GetJob_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type mockConstructorTestingTNewJobService interface {
	mock.TestingT
	Cleanup(func())
}

// NewJobService creates a new instance of JobService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewJobService(t mockConstructorTestingTNewJobService) *JobService {
	mock := &JobService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
