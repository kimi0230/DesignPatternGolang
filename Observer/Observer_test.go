package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {
	var s = new(ConcreteSubject)
	var kimi = NewObserver("kimi", s, 1)
	var yellow = NewObserver("yellow", s, 1)
	s.Attach(kimi)
	s.Attach(yellow)
	s.SetState(1)
	s.Notify()

	s.SetState(0)
	s.Notify()

	s.Detach(yellow)
	s.Notify()
	s.Attach(kimi)
}
