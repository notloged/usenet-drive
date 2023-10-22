// Code generated by MockGen. DO NOT EDIT.
// Source: ./connectionpool.go

// Package connectionpool is a generated GoMock package.
package connectionpool

import (
	io "io"
	reflect "reflect"
	time "time"

	nntp "github.com/chrisfarms/nntp"
	gomock "github.com/golang/mock/gomock"
)

// MockNntpConnection is a mock of NntpConnection interface.
type MockNntpConnection struct {
	ctrl     *gomock.Controller
	recorder *MockNntpConnectionMockRecorder
}

// MockNntpConnectionMockRecorder is the mock recorder for MockNntpConnection.
type MockNntpConnectionMockRecorder struct {
	mock *MockNntpConnection
}

// NewMockNntpConnection creates a new mock instance.
func NewMockNntpConnection(ctrl *gomock.Controller) *MockNntpConnection {
	mock := &MockNntpConnection{ctrl: ctrl}
	mock.recorder = &MockNntpConnectionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNntpConnection) EXPECT() *MockNntpConnectionMockRecorder {
	return m.recorder
}

// Article mocks base method.
func (m *MockNntpConnection) Article(id string) (*nntp.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Article", id)
	ret0, _ := ret[0].(*nntp.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Article indicates an expected call of Article.
func (mr *MockNntpConnectionMockRecorder) Article(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Article", reflect.TypeOf((*MockNntpConnection)(nil).Article), id)
}

// ArticleText mocks base method.
func (m *MockNntpConnection) ArticleText(id string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ArticleText", id)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArticleText indicates an expected call of ArticleText.
func (mr *MockNntpConnectionMockRecorder) ArticleText(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArticleText", reflect.TypeOf((*MockNntpConnection)(nil).ArticleText), id)
}

// Authenticate mocks base method.
func (m *MockNntpConnection) Authenticate(username, password string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", username, password)
	ret0, _ := ret[0].(error)
	return ret0
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockNntpConnectionMockRecorder) Authenticate(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockNntpConnection)(nil).Authenticate), username, password)
}

// Body mocks base method.
func (m *MockNntpConnection) Body(id string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body", id)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Body indicates an expected call of Body.
func (mr *MockNntpConnectionMockRecorder) Body(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockNntpConnection)(nil).Body), id)
}

// Capabilities mocks base method.
func (m *MockNntpConnection) Capabilities() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capabilities")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Capabilities indicates an expected call of Capabilities.
func (mr *MockNntpConnectionMockRecorder) Capabilities() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capabilities", reflect.TypeOf((*MockNntpConnection)(nil).Capabilities))
}

// Date mocks base method.
func (m *MockNntpConnection) Date() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Date")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Date indicates an expected call of Date.
func (mr *MockNntpConnectionMockRecorder) Date() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Date", reflect.TypeOf((*MockNntpConnection)(nil).Date))
}

// Group mocks base method.
func (m *MockNntpConnection) Group(group string) (int, int, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group", group)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(int)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Group indicates an expected call of Group.
func (mr *MockNntpConnectionMockRecorder) Group(group interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockNntpConnection)(nil).Group), group)
}

// Head mocks base method.
func (m *MockNntpConnection) Head(id string) (*nntp.Article, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head", id)
	ret0, _ := ret[0].(*nntp.Article)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockNntpConnectionMockRecorder) Head(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockNntpConnection)(nil).Head), id)
}

// HeadText mocks base method.
func (m *MockNntpConnection) HeadText(id string) (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HeadText", id)
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HeadText indicates an expected call of HeadText.
func (mr *MockNntpConnectionMockRecorder) HeadText(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HeadText", reflect.TypeOf((*MockNntpConnection)(nil).HeadText), id)
}

// Help mocks base method.
func (m *MockNntpConnection) Help() (io.Reader, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Help")
	ret0, _ := ret[0].(io.Reader)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Help indicates an expected call of Help.
func (mr *MockNntpConnectionMockRecorder) Help() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Help", reflect.TypeOf((*MockNntpConnection)(nil).Help))
}

// Last mocks base method.
func (m *MockNntpConnection) Last() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Last")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Last indicates an expected call of Last.
func (mr *MockNntpConnectionMockRecorder) Last() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Last", reflect.TypeOf((*MockNntpConnection)(nil).Last))
}

// List mocks base method.
func (m *MockNntpConnection) List(a ...string) ([]string, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a_2 := range a {
		varargs = append(varargs, a_2)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockNntpConnectionMockRecorder) List(a ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockNntpConnection)(nil).List), a...)
}

// ModeReader mocks base method.
func (m *MockNntpConnection) ModeReader() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ModeReader")
	ret0, _ := ret[0].(error)
	return ret0
}

// ModeReader indicates an expected call of ModeReader.
func (mr *MockNntpConnectionMockRecorder) ModeReader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ModeReader", reflect.TypeOf((*MockNntpConnection)(nil).ModeReader))
}

// NewGroups mocks base method.
func (m *MockNntpConnection) NewGroups(since time.Time) ([]*nntp.Group, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewGroups", since)
	ret0, _ := ret[0].([]*nntp.Group)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewGroups indicates an expected call of NewGroups.
func (mr *MockNntpConnectionMockRecorder) NewGroups(since interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewGroups", reflect.TypeOf((*MockNntpConnection)(nil).NewGroups), since)
}

// NewNews mocks base method.
func (m *MockNntpConnection) NewNews(group string, since time.Time) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewNews", group, since)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewNews indicates an expected call of NewNews.
func (mr *MockNntpConnectionMockRecorder) NewNews(group, since interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewNews", reflect.TypeOf((*MockNntpConnection)(nil).NewNews), group, since)
}

// Next mocks base method.
func (m *MockNntpConnection) Next() (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Next indicates an expected call of Next.
func (mr *MockNntpConnectionMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockNntpConnection)(nil).Next))
}

// Overview mocks base method.
func (m *MockNntpConnection) Overview(begin, end int) ([]nntp.MessageOverview, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Overview", begin, end)
	ret0, _ := ret[0].([]nntp.MessageOverview)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Overview indicates an expected call of Overview.
func (mr *MockNntpConnectionMockRecorder) Overview(begin, end interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Overview", reflect.TypeOf((*MockNntpConnection)(nil).Overview), begin, end)
}

// Post mocks base method.
func (m *MockNntpConnection) Post(a *nntp.Article) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Post", a)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post.
func (mr *MockNntpConnectionMockRecorder) Post(a interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockNntpConnection)(nil).Post), a)
}

// Quit mocks base method.
func (m *MockNntpConnection) Quit() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Quit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Quit indicates an expected call of Quit.
func (mr *MockNntpConnectionMockRecorder) Quit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Quit", reflect.TypeOf((*MockNntpConnection)(nil).Quit))
}

// RawPost mocks base method.
func (m *MockNntpConnection) RawPost(r io.Reader) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawPost", r)
	ret0, _ := ret[0].(error)
	return ret0
}

// RawPost indicates an expected call of RawPost.
func (mr *MockNntpConnectionMockRecorder) RawPost(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawPost", reflect.TypeOf((*MockNntpConnection)(nil).RawPost), r)
}

// Stat mocks base method.
func (m *MockNntpConnection) Stat(id string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", id)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Stat indicates an expected call of Stat.
func (mr *MockNntpConnectionMockRecorder) Stat(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockNntpConnection)(nil).Stat), id)
}

// MockUsenetConnectionPool is a mock of UsenetConnectionPool interface.
type MockUsenetConnectionPool struct {
	ctrl     *gomock.Controller
	recorder *MockUsenetConnectionPoolMockRecorder
}

// MockUsenetConnectionPoolMockRecorder is the mock recorder for MockUsenetConnectionPool.
type MockUsenetConnectionPoolMockRecorder struct {
	mock *MockUsenetConnectionPool
}

// NewMockUsenetConnectionPool creates a new mock instance.
func NewMockUsenetConnectionPool(ctrl *gomock.Controller) *MockUsenetConnectionPool {
	mock := &MockUsenetConnectionPool{ctrl: ctrl}
	mock.recorder = &MockUsenetConnectionPoolMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsenetConnectionPool) EXPECT() *MockUsenetConnectionPoolMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockUsenetConnectionPool) Close(c NntpConnection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockUsenetConnectionPoolMockRecorder) Close(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockUsenetConnectionPool)(nil).Close), c)
}

// Free mocks base method.
func (m *MockUsenetConnectionPool) Free(c NntpConnection) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Free", c)
	ret0, _ := ret[0].(error)
	return ret0
}

// Free indicates an expected call of Free.
func (mr *MockUsenetConnectionPoolMockRecorder) Free(c interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Free", reflect.TypeOf((*MockUsenetConnectionPool)(nil).Free), c)
}

// Get mocks base method.
func (m *MockUsenetConnectionPool) Get() (NntpConnection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(NntpConnection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockUsenetConnectionPoolMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockUsenetConnectionPool)(nil).Get))
}

// GetActiveConnections mocks base method.
func (m *MockUsenetConnectionPool) GetActiveConnections() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActiveConnections")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetActiveConnections indicates an expected call of GetActiveConnections.
func (mr *MockUsenetConnectionPoolMockRecorder) GetActiveConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActiveConnections", reflect.TypeOf((*MockUsenetConnectionPool)(nil).GetActiveConnections))
}

// GetFreeConnections mocks base method.
func (m *MockUsenetConnectionPool) GetFreeConnections() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFreeConnections")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetFreeConnections indicates an expected call of GetFreeConnections.
func (mr *MockUsenetConnectionPoolMockRecorder) GetFreeConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFreeConnections", reflect.TypeOf((*MockUsenetConnectionPool)(nil).GetFreeConnections))
}

// GetMaxConnections mocks base method.
func (m *MockUsenetConnectionPool) GetMaxConnections() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxConnections")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetMaxConnections indicates an expected call of GetMaxConnections.
func (mr *MockUsenetConnectionPoolMockRecorder) GetMaxConnections() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxConnections", reflect.TypeOf((*MockUsenetConnectionPool)(nil).GetMaxConnections))
}