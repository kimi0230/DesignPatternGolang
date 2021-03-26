package command

import (
	"fmt"
	"time"
)

// Waiter : Invoker
type Waiter struct {
	orders []Command
}

func (w *Waiter) SetOrder(c Command) {
	if w == nil {
		return
	}
	w.orders = append(w.orders, c)
	fmt.Printf("增加訂單: %s, 時間: %s \n", c.ToString(), time.Now().UTC().String())
}

func (w *Waiter) CancelOrder(c Command) {
	if w == nil {
		return
	}
	for i := 0; i < len(w.orders); i++ {
		if w.orders[i] == c {
			w.orders = removeCommand(w.orders, i)
			break
		}
	}
	fmt.Printf("刪除訂單: %s, 時間: %s \n", c.ToString(), time.Now().UTC().String())
}
func removeCommand(slice []Command, s int) []Command {
	slice[s] = nil
	return append(slice[:s], slice[s+1:]...)
}
func (w *Waiter) Notify() {
	for _, cmd := range w.orders {
		cmd.Execute()
	}
}

// Command : Command
type Command interface {
	ToString() string
	Execute()
}

// BakeMuttonCommand : 烤羊肉命令 ConcreteCommand
type BakeMuttonCommand struct {
	name     string
	receiver Barbecuer
}

func NewBakeMuttonCommand(name string) *BakeMuttonCommand {
	return &BakeMuttonCommand{name, Barbecuer{}}
}
func (c *BakeMuttonCommand) SetReceiver(r Barbecuer) {
	if c == nil {
		return
	}
	c.receiver = r
}
func (c *BakeMuttonCommand) ToString() string {
	if c == nil {
		return ""
	}
	return c.name
}
func (c *BakeMuttonCommand) Execute() {
	c.receiver.BakeMutton()
}

// BakeChickenWingCommand : 烤雞翅命令 ConcreteCommand
type BakeChickenWingCommand struct {
	name     string
	receiver Barbecuer
}

func NewBakeChickenWingCommand(name string) *BakeChickenWingCommand {
	return &BakeChickenWingCommand{name, Barbecuer{}}
}
func (c *BakeChickenWingCommand) SetReceiver(r Barbecuer) {
	if c == nil {
		return
	}
	c.receiver = r
}
func (c *BakeChickenWingCommand) ToString() string {
	if c == nil {
		return ""
	}
	return c.name
}
func (c *BakeChickenWingCommand) Execute() {
	c.receiver.BakeChickenWing()
}

// Barbecuer :烤肉串者 Receiver
type Barbecuer struct{}

func (r *Barbecuer) BakeMutton() {
	if r == nil {
		return
	}
	fmt.Println("烤羊肉串")
}
func (r *Barbecuer) BakeChickenWing() {
	if r == nil {
		return
	}
	fmt.Println("烤雞翅")
}
