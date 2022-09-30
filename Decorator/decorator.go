/*
Decorator 裝飾模式： 將一個物件包裝起來, 以提供新的行為
將額外權責動態附加於物件身上, 不必衍生子類別即可彈性擴增功能
動態低給一個物件加入一些額外的職責, 就增加功能來說, 裝飾模式比產生子類別更為靈活
所有 decorator 在操作時需使用函數, 裝飾者順序很重要, 比如加密資料和過濾辭彙都可以是資料持久化前的裝飾功能
*/

package decorator

import "fmt"

// Person : concrete component
type Person struct {
	Name string
}

func (p *Person) show() {
	if p == nil {
		return
	}
	fmt.Println("裝扮的", p.Name)
}

// 服飾類別 Decorator
// Finery : Decorator
type Finery interface {
	show()
}

type Decorator struct {
	Finery
}

func (d *Decorator) SetDecorator(component Finery) {
	if d == nil {
		return
	}
	d.Finery = component
}

func (d *Decorator) show() {
	if d == nil {
		return
	}
	if d.Finery != nil {
		d.Finery.show()
	}
}

// 具體服飾類別 Concrete Decorator

type TShirts struct {
	Decorator
}

func (t *TShirts) show() {
	if t == nil {
		return
	}
	fmt.Println("大T恤")
	t.Decorator.show()
}

type BigTrouser struct {
	Decorator
}

func (b *BigTrouser) show() {
	if b == nil {
		return
	}
	fmt.Println("垮褲")
	b.Decorator.show()
}

type Sneakers struct {
	Decorator
}

func (s *Sneakers) show() {
	if s == nil {
		return
	}
	s.Decorator.show()
	fmt.Println("破球鞋")
}
