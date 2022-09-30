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
