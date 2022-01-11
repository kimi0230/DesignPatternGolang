### Flyweight 享元模式：
> 以共享機制有效地支援一大堆小規模的物件

運用共用技術有效地支援大量細粒度的物件

造成大量記憶體消耗時, 就應該考慮使用; 或者是物件的大多數狀態是外部狀態, 如果刪除物件的外部狀態, 那麼可以用相對較少的共用物件取代很多組物件, 此時可以考慮使用享元模式.
像是圍棋棋子物件可以減少到只有兩個實體.
開發一個可供多人註冊的部落格網站

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Flyweight.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Flyweight 

```go
package flyweight

import "fmt"

type Flyweight interface {
	Operation(int)
}

type ConcreteFlyweight struct {
	name string
}

func (c *ConcreteFlyweight) Operation(extrinsicState int) {
	if c == nil {
		return
	}
	fmt.Println("具體 Flyweight", extrinsicState)
}

type UnsharedConcreteFlyweight struct {
	name string
}

func (c *UnsharedConcreteFlyweight) Operation(extrinsicState int) {
	if c == nil {
		return
	}
	fmt.Println("不共用的具體 Flyweight", extrinsicState)
}

type FlyweightFactory struct {
	flyweights map[string]Flyweight
}

func NewFlyweightFactory() *FlyweightFactory {
	fwfactory := FlyweightFactory{make(map[string]Flyweight)}
	fwfactory.flyweights["X"] = &ConcreteFlyweight{"X"}
	fwfactory.flyweights["Y"] = &ConcreteFlyweight{"Y"}
	fwfactory.flyweights["Z"] = &ConcreteFlyweight{"Z"}
	return &fwfactory
}

func (f *FlyweightFactory) Flyweight(name string) Flyweight {
	if f == nil {
		return nil
	}
	if _, ok := f.flyweights[name]; !ok {
		f.flyweights[name] = &ConcreteFlyweight{name}
	}
	return f.flyweights[name]
}

```