/*
	Facade 外觀模式：
	為子系統中的一組介面提供一個一致的界面，此模式定義了一個高層介面，
	這個介面使得這一子系統更加容易使用（投資：基金，股票，房產）

	個人想法：Mediator 中介者模式、外觀模式：每個對像都保存一份中介者對象，
	在和其他對象交互時，通過中介者來完成，
	外觀模式：外觀中保存了一堆對象，這些對像或者是組成某個子系統的，
	將其封裝在外觀對像中，給客戶端一種只有一個對象的感覺，
	外觀模式是一結構型模式，
	中介者模式是一行為性模式
*/
package facade

import "fmt"

// 外觀類別, 她需要了解所有的子系統的方法或屬性, 並進行組合, 以備外界調用
type Facade struct {
	One   SubSystemOne
	Two   SubSystemTwo
	Three SubSystemThree
	Four  SubSystemFour
}

func NewFacade() *Facade {
	return &Facade{SubSystemOne{}, SubSystemTwo{}, SubSystemThree{}, SubSystemFour{}}
}

func (f *Facade) MethodA() {
	fmt.Println("方法組 A")
	f.One.MethodOne()
	f.Three.MethodThree()
}

func (f *Facade) MethodB() {
	fmt.Println("方法組 B")
	f.Two.MethodTwo()
	f.Four.MethodFour()
}

type SubSystemOne struct {
}

func (s *SubSystemOne) MethodOne() {
	fmt.Println("子系統方法一")
}

type SubSystemTwo struct {
}

func (s *SubSystemTwo) MethodTwo() {
	fmt.Println("子系統方法二")
}

type SubSystemThree struct {
}

func (s *SubSystemThree) MethodThree() {
	fmt.Println("子系統方法三")
}

type SubSystemFour struct {
}

func (s *SubSystemFour) MethodFour() {
	fmt.Println("子系統方法四")
}
