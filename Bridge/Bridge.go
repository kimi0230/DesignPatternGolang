package bridge

import "fmt"

// HandsetSoft: 手機軟體 Implementor
type HandsetSoft interface {
	Run()
}

// HandsetGame : 手機遊戲 ConcreteImplementorA
type HandsetGame struct{}

func (h *HandsetGame) Run() {
	fmt.Println("執行手機遊戲")
}

// HandsetAddressList : 手機通訊錄 ConcreteImplementorB
type HandsetAddressList struct {
}

func (h *HandsetAddressList) Run() {
	fmt.Println("執行手機通訊錄")
}

// HandsetBrand : 手機品牌 Abstraction
type HandsetBrand interface {
	SetHandsetSoft(HandsetSoft)
	Run()
}

// HandsetBrandN : RefinedAbstraction
type HandsetBrandN struct {
	soft HandsetSoft
}

func (h *HandsetBrandN) SetHandsetSoft(software HandsetSoft) {
	if h == nil {
		return
	}
	h.soft = software
}

func (h *HandsetBrandN) Run() {
	if h == nil {
		return
	}
	h.soft.Run()
}

// HandsetBrandM : RefinedAbstraction
type HandsetBrandM struct {
	soft HandsetSoft
}

func (h *HandsetBrandM) SetHandsetSoft(software HandsetSoft) {
	if h == nil {
		return
	}
	h.soft = software
}

func (h *HandsetBrandM) Run() {
	if h == nil {
		return
	}
	h.soft.Run()
}
