package library_extension

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"go.uber.org/mock/gomock"

	"github.com/johndoe/library/library"
)

type MockPointerObject struct {
	ctrl     *gomock.Controller
	recorder *MockPointerObjectMockRecorder
}

type MockPointerObjectMockRecorder struct {
	mock *MockPointerObject
}

func NewMockPointerObject(ctrl *gomock.Controller) *MockPointerObject {
	mock := &MockPointerObject{ctrl: ctrl}
	mock.recorder = &MockPointerObjectMockRecorder{mock}
	return mock
}

func (m *MockPointerObject) EXPECT() *MockPointerObjectMockRecorder {
	return m.recorder
}

func (m *MockPointerObject) Size() fyne.Size {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(fyne.Size)
	return ret0
}

func (mr *MockPointerObjectMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockPointerObject)(nil).Size))
}

func (m *MockPointerObject) PointerX() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PointerX")
	ret0, _ := ret[0].(int)
	return ret0
}

func (mr *MockPointerObjectMockRecorder) PointerX() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PointerX", reflect.TypeOf((*MockPointerObject)(nil).PointerX))
}

func (m *MockPointerObject) PointerY() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PointerY")
	ret0, _ := ret[0].(int)
	return ret0
}

func (mr *MockPointerObjectMockRecorder) PointerY() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PointerY", reflect.TypeOf((*MockPointerObject)(nil).PointerY))
}

func TestLabelExtImpl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLabel := NewMockPointerObject(ctrl)
	mockLabel.EXPECT().PointerX().Return(0)
	mockLabel.EXPECT().PointerY().Return(1)

	impl := LabelExtImpl{LabelImpl: library.NewLabelImpl(mockLabel)}
	if result := impl.GetPosition(); result != "(0,1)" {
		t.Errorf("expected (0,1), got %s", result)
	}
}
