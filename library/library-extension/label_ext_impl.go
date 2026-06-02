package library_extension

import (
	"fmt"

	"github.com/johndoe/library/library"
)

type PointerObject interface {
	library.Sizable
	PointerX() int
	PointerY() int
}

type LabelExtImpl struct {
	library.LabelImpl
}

func (l *LabelExtImpl) GetPosition() string {
	if ptr, ok := l.Label().(PointerObject); ok {
		return fmt.Sprintf("(%d,%d)", ptr.PointerX(), ptr.PointerY())
	}
	return "(0,0)"
}
