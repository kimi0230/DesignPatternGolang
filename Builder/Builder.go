package builder

import "fmt"

// 產品類別, 由多個零件組成
type Product struct {
	parts []string
}

func (p *Product) Add(part string) {
	p.parts = append(p.parts, part)
}

func (p *Product) Show() {
	fmt.Println(p.parts)
}

type Builder interface {
	BuildPartA()
	BuildPartB()
	GetResult() *Product
}

// 具體建造者 實現Builder interface 構造和裝配各個零件
type ConcreteBuilder1 struct {
	product Product
}

func (c *ConcreteBuilder1) BuildPartA() {
	c.product.Add("零件A")
}

func (c *ConcreteBuilder1) BuildPartB() {
	c.product.Add("零件B")
}

func (c *ConcreteBuilder1) GetResult() *Product {
	return &c.product
}

// 具體建造者
type ConcreteBuilder2 struct {
	product Product
}

func (c *ConcreteBuilder2) BuildPartA() {
	c.product.Add("零件X")
}

func (c *ConcreteBuilder2) BuildPartB() {
	c.product.Add("零件Y")
}

func (c *ConcreteBuilder2) GetResult() *Product {
	return &c.product
}

// 是建構一個使用 Builder介面的物件
type Director struct {
	b Builder
}

func (d *Director) Construct() {
	d.b.BuildPartA()
	d.b.BuildPartB()
}
