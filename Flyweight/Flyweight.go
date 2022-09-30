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
