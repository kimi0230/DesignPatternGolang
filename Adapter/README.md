### Adapter 轉接器模式:
> 將類別的介面轉換成外界所預期的另一種介面, 讓原先囿於見面不相容問題而無法協力合作的類別能夠兜在一起用

將一個類別的介面轉換成客戶希望的另外一個介面, Adapter 模式始原本由於介面不相容而不能一起工作的那些類別可以一起工作

與 Proxy 代理模式相似
* [Proxy 代理模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Proxy)
    使用介面, 客戶端不知道代理對象
* [ Adapter 轉接器模式:](https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter)
    可以是具體或抽象類別,或是介面, 對象的接口和客户端想要的不一樣, Adapter將接口封装一下, 改成客戶端想要的接口

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Adapter.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Adapter

```go
package adapter

import "fmt"

// Player : target
type Player interface {
	Attack()
	Defense()
}

// 前鋒
type Forwards struct {
	Name string
}

func (p *Forwards) Attack() {
	if p == nil {
		return
	}
	fmt.Printf("前鋒 %s 進攻 \n", p.Name)
}

func (p *Forwards) Defense() {
	if p == nil {
		return
	}
	fmt.Printf("前鋒 %s 防守 \n", p.Name)
}

func NewForwards(name string) Player {
	return &Forwards{name}
}

// 中鋒
type Center struct {
	Name string
}

func (p *Center) Attack() {
	if p == nil {
		return
	}
	fmt.Printf("中鋒 %s 進攻 \n", p.Name)
}

func (p *Center) Defense() {
	if p == nil {
		return
	}
	fmt.Printf("中鋒 %s 防守 \n", p.Name)
}

func NewCenter(name string) Player {
	return &Center{name}
}

// 後衛
type Guards struct {
	Name string
}

func (p *Guards) Attack() {
	if p == nil {
		return
	}
	fmt.Printf("後衛 %s 進攻 \n", p.Name)
}

func (p *Guards) Defense() {
	if p == nil {
		return
	}
	fmt.Printf("後衛 %s 防守 \n", p.Name)
}

func NewGuards(name string) Player {
	return &Guards{name}
}

// 外籍中鋒 : Adaptee
type ForeignCenter struct {
	Name string
}

func (p *ForeignCenter) AttackYo() {
	if p == nil {
		return
	}
	fmt.Printf("外籍中鋒 %s 進攻yo \n", p.Name)
}
func (p *ForeignCenter) DefenseYo() {
	if p == nil {
		return
	}
	fmt.Printf("外籍中鋒 %s 防守yo \n", p.Name)
}

// Translator : Adapter
type Translator struct {
	f ForeignCenter
}

func (t *Translator) Attack() {
	if t == nil {
		return
	}
	t.f.AttackYo()
}

func (t *Translator) Defense() {
	if t == nil {
		return
	}
	t.f.DefenseYo()
}

func NewTranslator(name string) Player {
	return &Translator{ForeignCenter{name}}
}
```