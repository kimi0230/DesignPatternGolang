### Visitor 訪問者模式: 
> 定義能逐一施行於物件結構裡各個元素的操作, 讓你不必修改作用對象的類別介面, 就能定義新的操作

表示一個作用於某物件結構中的各元素之操作. 他使你可以在不改變各元素之類別的前提下, 定義作用於這些元素的新操作

訪問者增加具體的 Element 是困難的, 但增加依賴於複雜物件結構的構建的操作就變得容易. 
僅需要增加一個新的訪問者, 即可在一個物件結構上定義一個新操作

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Visitor.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Visitor 

```go
package visitor

import "fmt"

type Visitor interface {
	VisitConcreteElementA(ConcreteElementA)
	VisitConcreteElementB(ConcreteElementB)
}

type ConcreteVisitorA struct {
}

func (c *ConcreteVisitorA) VisitConcreteElementA(ele ConcreteElementA) {
	if c == nil {
		return
	}
	ele.OperatorA()
}

func (c *ConcreteVisitorA) VisitConcreteElementB(ele ConcreteElementB) {
	if c == nil {
		return
	}
	ele.OperatorB()
}

type ConcreteVisitorB struct {
}

func (c *ConcreteVisitorB) VisitConcreteElementA(ele ConcreteElementA) {
	if c == nil {
		return
	}
	ele.OperatorA()
}

func (c *ConcreteVisitorB) VisitConcreteElementB(ele ConcreteElementB) {
	if c == nil {
		return
	}
	ele.OperatorB()
}

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(visitor Visitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementA(*c)
}

func (c *ConcreteElementA) OperatorA() {
	if c == nil {
		return
	}
	fmt.Println("OperatorA")
}

type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(visitor Visitor) {
	if c == nil {
		return
	}
	visitor.VisitConcreteElementB(*c)
}

func (c *ConcreteElementB) OperatorB() {
	if c == nil {
		return
	}
	fmt.Println("OperatorB")
}

type ObjectStructure struct {
	elemetes []Element
}

func (o *ObjectStructure) Attach(ele Element) {
	if o == nil || ele == nil {
		return
	}
	o.elemetes = append(o.elemetes, ele)
}

func (o *ObjectStructure) Detach(ele Element) {
	if o == nil || ele == nil {
		return
	}
	for i, val := range o.elemetes {
		if val == ele {
			o.elemetes = append(o.elemetes[:i], o.elemetes[i+1:]...)
			break
		}
	}
}

func (o *ObjectStructure) Accept(v Visitor) {
	if o == nil {
		return
	}
	for _, val := range o.elemetes {
		val.Accept(v)
	}
}
```