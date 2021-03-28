package mediator

import "testing"

func TestMediator(t *testing.T) {
	m := NewConcreteMediator()
	c1 := NewConcreteColleague1(m)
	c2 := NewConcreteColleague2(m)

	m.AddColleague(c1)
	m.AddColleague(c2)

	c1.Send("吃過飯了沒")
	c2.Send("沒有呢, 你打算請客?")
}
