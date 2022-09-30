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
