// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	gitlab "github.com/xanzy/go-gitlab"

	mock "github.com/stretchr/testify/mock"
)

// PipelinesService is an autogenerated mock type for the PipelinesService type
type PipelinesService struct {
	mock.Mock
}

// GetPipeline provides a mock function with given fields: pid, pipeline, options
func (_m *PipelinesService) GetPipeline(pid interface{}, pipeline int, options ...gitlab.RequestOptionFunc) (*gitlab.Pipeline, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, pipeline)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *gitlab.Pipeline
	if rf, ok := ret.Get(0).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Pipeline); ok {
		r0 = rf(pid, pipeline, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gitlab.Pipeline)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, int, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, pipeline, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, int, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, pipeline, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ListProjectPipelines provides a mock function with given fields: pid, opt, options
func (_m *PipelinesService) ListProjectPipelines(pid interface{}, opt *gitlab.ListProjectPipelinesOptions, options ...gitlab.RequestOptionFunc) ([]*gitlab.PipelineInfo, *gitlab.Response, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, pid, opt)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 []*gitlab.PipelineInfo
	if rf, ok := ret.Get(0).(func(interface{}, *gitlab.ListProjectPipelinesOptions, ...gitlab.RequestOptionFunc) []*gitlab.PipelineInfo); ok {
		r0 = rf(pid, opt, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*gitlab.PipelineInfo)
		}
	}

	var r1 *gitlab.Response
	if rf, ok := ret.Get(1).(func(interface{}, *gitlab.ListProjectPipelinesOptions, ...gitlab.RequestOptionFunc) *gitlab.Response); ok {
		r1 = rf(pid, opt, options...)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*gitlab.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(interface{}, *gitlab.ListProjectPipelinesOptions, ...gitlab.RequestOptionFunc) error); ok {
		r2 = rf(pid, opt, options...)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}