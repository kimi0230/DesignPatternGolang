### Bridge 橋接模式：
> 將實作體系與抽象體系分離開來, 讓兩者能各自更動各自演進

將抽象部分與他的實現部分分離, 使他們都可以獨立低變化

合成/聚合復用原則: 盡量使用合成/聚合, 不要使用類別繼承
聚合表示一種弱的擁有關係, 表現是的是 A 物件可以包涵 B 物件,
但 B 物件不是 A 物件的一部分;
ex: 雁群/雁, 在雁群類別中有雁陣列物件
合成則是一種強的擁有關係, 表現了嚴格的部分和整體的關係, 部分和整體的生命週期一樣
ex: 鳥/翅膀, 在鳥類別中初始化時, 實體翅膀同時產生

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Bridge.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Bridge

```go
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
```