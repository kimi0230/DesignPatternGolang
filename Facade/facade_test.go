package facade

import "testing"

func TestFacade(t *testing.T) {
	f := NewFacade()
	f.MethodA()
	f.MethodB()
}
