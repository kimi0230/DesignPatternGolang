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
