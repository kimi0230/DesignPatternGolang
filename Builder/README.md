### Builder 建造者模式:
> 從複雜物件的佈局中抽取出生成程序, 以便用同一個生成程序製造各種不同的物件佈局

將一個複雜物件的構建與他的表示分離, 使得同樣的構建過程可以建立不同的表示

建造者的建造流程是在指揮者(Director)中，指揮者在用戶通知他現在具體的建造者是誰後，
建造出對應的產品，建造者中實現了產品的建造細節

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Builder.png?raw=true)
來源: 程杰-大話設計模式

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Builder 

```go
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

```