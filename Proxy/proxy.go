/*
Proxy 代理模式： 將物件包裝起來, 以控制對此物件的存取
為其他物件提供一種代理, 以控制對這個物件的存取
代理和代理的對象接口一致，客户端不知道代理对象
*/
package proxy

import "fmt"

// IGiveGift : Subject
type IGiveGift interface {
	GiveDolls()
	GiveFlowers()
	GiveChocolate()
}

type SchoolGirl struct {
	name string
}

func (g *SchoolGirl) Name() string {
	if g == nil {
		return ""
	}
	return g.name
}

func (g *SchoolGirl) SetName(name string) {
	if g == nil {
		return
	}
	g.name = name
}

// Pursuit : real subject
type Pursuit struct {
	girl SchoolGirl
}

func (p *Pursuit) GiveDolls() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你洋娃娃")
}

func (p *Pursuit) GiveFlowers() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你鮮花")
}

func (p *Pursuit) GiveChocolate() {
	if p == nil {
		return
	}
	fmt.Println(p.girl.name, "送你巧克力")
}

// Proxy Proxy
type Proxy struct {
	pursuit Pursuit
}

func NewProxy(girl SchoolGirl) *Proxy {
	pursuit := Pursuit{girl}
	return &Proxy{pursuit}
}

func (p *Proxy) GiveDolls() {
	if p == nil {
		return
	}
	p.pursuit.GiveDolls()
}
func (p *Proxy) GiveFlowers() {
	if p == nil {
		return
	}
	p.pursuit.GiveFlowers()
}

func (p *Proxy) GiveChocolate() {
	if p == nil {
		return
	}
	p.pursuit.GiveChocolate()
}
