package main

import "fyne.io/fyne/v2"

type Sizable interface {
	Size() fyne.Size
}

type LabelImpl struct {
	label Sizable
}

func (l *LabelImpl) GetSize() int {
	size := l.label.Size()
	return int(size.Width * size.Height)
}
