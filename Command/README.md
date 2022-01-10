### Command 命令模式:
> 將訊息封裝成物件, 以便能用各種不同訊息, 暫佇, 紀律, 復原等方式加以參數化處理

將一個請求封裝為一個物件, 讓你可以用不同的請求對客戶進行參數化; 對請求排隊或紀錄請求日誌, 以及支援可取消的操作

不要為程式碼加上基於猜測的, 實際上不需要的功能. 如果不清楚一個系統是否需要命令模式, 就不要急著實現它.
事實上, 在需要的時候透過重構實現這個模式並不難, 只有在真正需要如取消/恢復等功能時, 
把原本的程式碼重構為命令模式才有意義

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Command.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Command 

```go
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
```