package templatemethod

import (
	"testing"
)

func TestTemplate(t *testing.T) {
	a := NewClassA()
	a.abstractClass = a
	a.TemplateMethod()

	b := NewClassB()
	b.abstractClass = b
	b.TemplateMethod()
}
