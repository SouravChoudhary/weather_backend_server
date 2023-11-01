// Code generated by MockGen. DO NOT EDIT.
// Source: weather_history.go

// Package service is a generated GoMock package.
package service

import (
	reflect "reflect"
	weather_history "weatherapp/db/weather_history"

	gomock "github.com/golang/mock/gomock"
)

// MockWeatherHistory is a mock of WeatherHistory interface.
type MockWeatherHistory struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherHistoryMockRecorder
}

// MockWeatherHistoryMockRecorder is the mock recorder for MockWeatherHistory.
type MockWeatherHistoryMockRecorder struct {
	mock *MockWeatherHistory
}

// NewMockWeatherHistory creates a new mock instance.
func NewMockWeatherHistory(ctrl *gomock.Controller) *MockWeatherHistory {
	mock := &MockWeatherHistory{ctrl: ctrl}
	mock.recorder = &MockWeatherHistoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeatherHistory) EXPECT() *MockWeatherHistoryMockRecorder {
	return m.recorder
}

// BulkDeleteRecords mocks base method.
func (m *MockWeatherHistory) BulkDeleteRecords(recordIDs []int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteRecords", recordIDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteRecords indicates an expected call of BulkDeleteRecords.
func (mr *MockWeatherHistoryMockRecorder) BulkDeleteRecords(recordIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteRecords", reflect.TypeOf((*MockWeatherHistory)(nil).BulkDeleteRecords), recordIDs)
}

// CreateRecord mocks base method.
func (m *MockWeatherHistory) CreateRecord(userID int, cityName, weatherData string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateRecord", userID, cityName, weatherData)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateRecord indicates an expected call of CreateRecord.
func (mr *MockWeatherHistoryMockRecorder) CreateRecord(userID, cityName, weatherData interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateRecord", reflect.TypeOf((*MockWeatherHistory)(nil).CreateRecord), userID, cityName, weatherData)
}

// DeleteRecord mocks base method.
func (m *MockWeatherHistory) DeleteRecord(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteRecord", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteRecord indicates an expected call of DeleteRecord.
func (mr *MockWeatherHistoryMockRecorder) DeleteRecord(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteRecord", reflect.TypeOf((*MockWeatherHistory)(nil).DeleteRecord), id)
}

// GetPaginatedRecords mocks base method.
func (m *MockWeatherHistory) GetPaginatedRecords(userID, recordLen, index int) ([]*weather_history.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaginatedRecords", userID, recordLen, index)
	ret0, _ := ret[0].([]*weather_history.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaginatedRecords indicates an expected call of GetPaginatedRecords.
func (mr *MockWeatherHistoryMockRecorder) GetPaginatedRecords(userID, recordLen, index interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaginatedRecords", reflect.TypeOf((*MockWeatherHistory)(nil).GetPaginatedRecords), userID, recordLen, index)
}

// GetRecordByID mocks base method.
func (m *MockWeatherHistory) GetRecordByID(id int) (*weather_history.Record, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRecordByID", id)
	ret0, _ := ret[0].(*weather_history.Record)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordByID indicates an expected call of GetRecordByID.
func (mr *MockWeatherHistoryMockRecorder) GetRecordByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRecordByID", reflect.TypeOf((*MockWeatherHistory)(nil).GetRecordByID), id)
}