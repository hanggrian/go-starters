package library

import "fyne.io/fyne/v2"

type Sizable interface {
	Size() fyne.Size
}

type LabelImpl struct {
	label Sizable
}

func NewLabelImpl(label Sizable) LabelImpl {
	return LabelImpl{label: label}
}

func (l *LabelImpl) Label() Sizable {
	return l.label
}

func (l *LabelImpl) GetSize() int {
	size := l.label.Size()
	return int(size.Width * size.Height)
}
