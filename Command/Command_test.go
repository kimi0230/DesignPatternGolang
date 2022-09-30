package command

import "testing"

func TestCommand(t *testing.T) {
	invoker := new(Waiter)
	command1 := NewBakeMuttonCommand("烤羊肉")
	recever := new(Barbecuer)

	command1.SetReceiver(*recever)
	invoker.SetOrder(command1)

	command2 := NewBakeChickenWingCommand("烤雞翅")
	invoker.SetOrder(command2)
	invoker.CancelOrder(command1)
	invoker.Notify()
}
