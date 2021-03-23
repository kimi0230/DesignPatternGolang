package memento

import "fmt"

// GameRole: 遊戲角色, 發起人  Originator
type GameRole struct {
	Vit int
	Atk int
	Def int
}

func (g *GameRole) InitState() {
	if g == nil {
		return
	}
	g.Vit = 100
	g.Atk = 100
	g.Def = 100
}

func (g *GameRole) Fight() {
	if g == nil {
		return
	}
	g.Vit = 0
	g.Atk = 0
	g.Def = 0
}

func (g *GameRole) StateDisplay() {
	if g == nil {
		return
	}
	fmt.Println("當前狀態：")
	fmt.Println("生命力 : ", g.Vit)
	fmt.Println("攻擊力 :", g.Atk)
	fmt.Println("防禦力 :", g.Def)
}

// 保存角色狀態
func (g *GameRole) SaveState() RoleStateMemento {
	if g == nil {
		return RoleStateMemento{}
	}
	return RoleStateMemento{g.Vit, g.Atk, g.Def}
}
func (g *GameRole) RecoveryState(memento RoleStateMemento) {
	if g == nil {
		return
	}
	g.Vit = memento.Vitality()
	g.Atk = memento.Attack()
	g.Def = memento.Defense()
}

// RoleStateMemento : 角色狀態儲存箱, 備忘錄 Memento
type RoleStateMemento struct {
	vit int
	atk int
	def int
}

func (r *RoleStateMemento) Vitality() int {
	if r == nil {
		return 0
	}
	return r.vit
}

func (r *RoleStateMemento) Attack() int {
	if r == nil {
		return 0
	}
	return r.atk
}

func (r *RoleStateMemento) Defense() int {
	if r == nil {
		return 0
	}
	return r.def
}

// 角色狀態管理員 Caretaker
type RoleStateCaretaker struct {
	memento RoleStateMemento
}

func (rc *RoleStateCaretaker) Memento() RoleStateMemento {
	if rc == nil {
		return RoleStateMemento{}
	}
	return rc.memento
}

func (rc *RoleStateCaretaker) SetMemento(m RoleStateMemento) {
	if rc == nil {
		return
	}
	rc.memento = m
}
