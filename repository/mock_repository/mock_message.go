// Code generated by MockGen. DO NOT EDIT.
// Source: message.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	uuid "github.com/gofrs/uuid"
	gomock "github.com/golang/mock/gomock"
	model "github.com/traPtitech/traQ/model"
	repository "github.com/traPtitech/traQ/repository"
	reflect "reflect"
)

// MockMessageRepository is a mock of MessageRepository interface
type MockMessageRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMessageRepositoryMockRecorder
}

// MockMessageRepositoryMockRecorder is the mock recorder for MockMessageRepository
type MockMessageRepositoryMockRecorder struct {
	mock *MockMessageRepository
}

// NewMockMessageRepository creates a new mock instance
func NewMockMessageRepository(ctrl *gomock.Controller) *MockMessageRepository {
	mock := &MockMessageRepository{ctrl: ctrl}
	mock.recorder = &MockMessageRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMessageRepository) EXPECT() *MockMessageRepositoryMockRecorder {
	return m.recorder
}

// CreateMessage mocks base method
func (m *MockMessageRepository) CreateMessage(userID, channelID uuid.UUID, text string) (*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateMessage", userID, channelID, text)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateMessage indicates an expected call of CreateMessage
func (mr *MockMessageRepositoryMockRecorder) CreateMessage(userID, channelID, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateMessage", reflect.TypeOf((*MockMessageRepository)(nil).CreateMessage), userID, channelID, text)
}

// UpdateMessage mocks base method
func (m *MockMessageRepository) UpdateMessage(messageID uuid.UUID, text string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMessage", messageID, text)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMessage indicates an expected call of UpdateMessage
func (mr *MockMessageRepositoryMockRecorder) UpdateMessage(messageID, text interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMessage", reflect.TypeOf((*MockMessageRepository)(nil).UpdateMessage), messageID, text)
}

// DeleteMessage mocks base method
func (m *MockMessageRepository) DeleteMessage(messageID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteMessage", messageID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteMessage indicates an expected call of DeleteMessage
func (mr *MockMessageRepositoryMockRecorder) DeleteMessage(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteMessage", reflect.TypeOf((*MockMessageRepository)(nil).DeleteMessage), messageID)
}

// GetMessageByID mocks base method
func (m *MockMessageRepository) GetMessageByID(messageID uuid.UUID) (*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessageByID", messageID)
	ret0, _ := ret[0].(*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMessageByID indicates an expected call of GetMessageByID
func (mr *MockMessageRepositoryMockRecorder) GetMessageByID(messageID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessageByID", reflect.TypeOf((*MockMessageRepository)(nil).GetMessageByID), messageID)
}

// GetMessages mocks base method
func (m *MockMessageRepository) GetMessages(query repository.MessagesQuery) ([]*model.Message, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessages", query)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMessages indicates an expected call of GetMessages
func (mr *MockMessageRepositoryMockRecorder) GetMessages(query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessages", reflect.TypeOf((*MockMessageRepository)(nil).GetMessages), query)
}

// SetMessageUnread mocks base method
func (m *MockMessageRepository) SetMessageUnread(userID, messageID uuid.UUID, noticeable bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMessageUnread", userID, messageID, noticeable)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMessageUnread indicates an expected call of SetMessageUnread
func (mr *MockMessageRepositoryMockRecorder) SetMessageUnread(userID, messageID, noticeable interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessageUnread", reflect.TypeOf((*MockMessageRepository)(nil).SetMessageUnread), userID, messageID, noticeable)
}

// GetUnreadMessagesByUserID mocks base method
func (m *MockMessageRepository) GetUnreadMessagesByUserID(userID uuid.UUID) ([]*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnreadMessagesByUserID", userID)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnreadMessagesByUserID indicates an expected call of GetUnreadMessagesByUserID
func (mr *MockMessageRepositoryMockRecorder) GetUnreadMessagesByUserID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnreadMessagesByUserID", reflect.TypeOf((*MockMessageRepository)(nil).GetUnreadMessagesByUserID), userID)
}

// DeleteUnreadsByChannelID mocks base method
func (m *MockMessageRepository) DeleteUnreadsByChannelID(channelID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUnreadsByChannelID", channelID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUnreadsByChannelID indicates an expected call of DeleteUnreadsByChannelID
func (mr *MockMessageRepositoryMockRecorder) DeleteUnreadsByChannelID(channelID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUnreadsByChannelID", reflect.TypeOf((*MockMessageRepository)(nil).DeleteUnreadsByChannelID), channelID, userID)
}

// GetUserUnreadChannels mocks base method
func (m *MockMessageRepository) GetUserUnreadChannels(userID uuid.UUID) ([]*repository.UserUnreadChannel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserUnreadChannels", userID)
	ret0, _ := ret[0].([]*repository.UserUnreadChannel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserUnreadChannels indicates an expected call of GetUserUnreadChannels
func (mr *MockMessageRepositoryMockRecorder) GetUserUnreadChannels(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserUnreadChannels", reflect.TypeOf((*MockMessageRepository)(nil).GetUserUnreadChannels), userID)
}

// GetChannelLatestMessagesByUserID mocks base method
func (m *MockMessageRepository) GetChannelLatestMessagesByUserID(userID uuid.UUID, limit int, subscribeOnly bool) ([]*model.Message, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChannelLatestMessagesByUserID", userID, limit, subscribeOnly)
	ret0, _ := ret[0].([]*model.Message)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChannelLatestMessagesByUserID indicates an expected call of GetChannelLatestMessagesByUserID
func (mr *MockMessageRepositoryMockRecorder) GetChannelLatestMessagesByUserID(userID, limit, subscribeOnly interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChannelLatestMessagesByUserID", reflect.TypeOf((*MockMessageRepository)(nil).GetChannelLatestMessagesByUserID), userID, limit, subscribeOnly)
}

// AddStampToMessage mocks base method
func (m *MockMessageRepository) AddStampToMessage(messageID, stampID, userID uuid.UUID, count int) (*model.MessageStamp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddStampToMessage", messageID, stampID, userID, count)
	ret0, _ := ret[0].(*model.MessageStamp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddStampToMessage indicates an expected call of AddStampToMessage
func (mr *MockMessageRepositoryMockRecorder) AddStampToMessage(messageID, stampID, userID, count interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddStampToMessage", reflect.TypeOf((*MockMessageRepository)(nil).AddStampToMessage), messageID, stampID, userID, count)
}

// RemoveStampFromMessage mocks base method
func (m *MockMessageRepository) RemoveStampFromMessage(messageID, stampID, userID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveStampFromMessage", messageID, stampID, userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveStampFromMessage indicates an expected call of RemoveStampFromMessage
func (mr *MockMessageRepositoryMockRecorder) RemoveStampFromMessage(messageID, stampID, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveStampFromMessage", reflect.TypeOf((*MockMessageRepository)(nil).RemoveStampFromMessage), messageID, stampID, userID)
}
