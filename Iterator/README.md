### Iterator 迭代器模式:
> 無需知曉聚合物件的內部細節, 即可依序存取內含的每一個元素

提供一種方法依序存取一個聚合物件中各個元素, 而又不暴露該物件的內部表示

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Iterator.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Iterator 

```go
package iterator

type Iterator interface {
	First() interface{}
	Next() interface{}
	IsDone() bool
	CurrentItem() interface{}
}

type ConcreteIterator struct {
	// 定義一個具體聚集物件
	aggregate *ConcreteAggregate
	current   int
}

func NewConcreteIterator(aggregate *ConcreteAggregate) *ConcreteIterator {
	return &ConcreteIterator{aggregate, 0}
}

func (c *ConcreteIterator) First() interface{} {
	return c.aggregate.items[0]
}

func (c *ConcreteIterator) Next() interface{} {
	c.current++
	if c.current < c.aggregate.Count() {
		return c.aggregate.items[c.current]
	}
	return nil
}

func (c *ConcreteIterator) IsDone() bool {
	if c.current >= c.aggregate.Count() {
		return true
	}
	return false
}

func (c *ConcreteIterator) CurrentItem() interface{} {
	if c == nil {
		return nil
	}
	return c.aggregate.items[c.current]
}

type Aggregate interface {
	CreateIterator()
}

type Object struct {
	Name string
}

type ConcreteAggregate struct {
	items []Object
}

func NewConcreteAggregate() *ConcreteAggregate {
	return new(ConcreteAggregate)
}

func (c *ConcreteAggregate) Count() int {
	return len(c.items)
}

func (c *ConcreteAggregate) GetThis(index int) Object {
	if c == nil {
		return Object{}
	}

	return c.items[index]
}

// SetThis : Insert 到指定 index
func (c *ConcreteAggregate) SetThis(index int, val Object) {
	if c == nil {
		return
	}
	// 先要一組記憶體空間
	tmpObject := Object{}
	c.items = append(c.items, tmpObject)
	copy(c.items[index+1:], c.items[index:])
	c.items[index] = val
}

```