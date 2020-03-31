// Code generated by mockery v1.0.0. DO NOT EDIT.

package automock

import mock "github.com/stretchr/testify/mock"
import model "github.com/kyma-incubator/compass/components/gateway/internal/auditlog/model"

// AuditlogMessageFactory is an autogenerated mock type for the AuditlogMessageFactory type
type AuditlogMessageFactory struct {
	mock.Mock
}

// CreateConfigurationChange provides a mock function with given fields:
func (_m *AuditlogMessageFactory) CreateConfigurationChange() model.ConfigurationChange {
	ret := _m.Called()

	var r0 model.ConfigurationChange
	if rf, ok := ret.Get(0).(func() model.ConfigurationChange); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.ConfigurationChange)
	}

	return r0
}

// CreateSecurityEvent provides a mock function with given fields:
func (_m *AuditlogMessageFactory) CreateSecurityEvent() model.SecurityEvent {
	ret := _m.Called()

	var r0 model.SecurityEvent
	if rf, ok := ret.Get(0).(func() model.SecurityEvent); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(model.SecurityEvent)
	}

	return r0
}
