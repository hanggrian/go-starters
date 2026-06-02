package main

import (
	"reflect"
	"testing"

	"fyne.io/fyne/v2"
	"go.uber.org/mock/gomock"
)

type MockSizable struct {
	ctrl     *gomock.Controller
	recorder *MockSizableMockRecorder
}

type MockSizableMockRecorder struct {
	mock *MockSizable
}

func NewMockSizable(ctrl *gomock.Controller) *MockSizable {
	mock := &MockSizable{ctrl: ctrl}
	mock.recorder = &MockSizableMockRecorder{mock}
	return mock
}

func (m *MockSizable) EXPECT() *MockSizableMockRecorder {
	return m.recorder
}

func (m *MockSizable) Size() fyne.Size {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(fyne.Size)
	return ret0
}

func (mr *MockSizableMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(
		mr.mock,
		"Size",
		reflect.TypeOf((*MockSizable)(nil).Size),
	)
}

func TestLabelImpl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLabel := NewMockSizable(ctrl)
	mockLabel.EXPECT().Size().Return(fyne.NewSize(2, 4))

	impl := LabelImpl{label: mockLabel}
	if result := impl.GetSize(); result != 8 {
		t.Errorf("expected 8, got %d", result)
	}
}
