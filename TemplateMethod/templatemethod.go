/*
	Template Method 範本方式模式: 由次類別決定如何實現一個演算法中的步驟
	定義一個操作中的演算法的骨架, 而將一些步驟延遲到子類別中.
	範本方式使得子類別可以不改變一個演算法的結構即可重定義該演算法的某些特定步驟
*/
package templatemethod

import "fmt"

type AbstractClass interface {
	// 一些抽象行為, 放到子類別去實現
	PrimitiveOperation1()
	PrimitiveOperation2()
}

type template struct {
	abstractClass AbstractClass
}

// TemplateMethod : 範本function 給出了邏輯的骨架
// 而邏輯的組成是一些相應的抽象操作
// 他們都推遲到子類別實現
func (t *template) TemplateMethod() {
	if t == nil {
		return
	}
	t.abstractClass.PrimitiveOperation1()
	t.abstractClass.PrimitiveOperation2()
}

type concreteClassA struct {
	template
}

func NewClassA() *concreteClassA {
	classA := concreteClassA{}
	return &classA
}

func (a *concreteClassA) PrimitiveOperation1() {
	fmt.Println("具體類別A方法1實現")
}
func (a *concreteClassA) PrimitiveOperation2() {
	fmt.Println("具體類別A方法2實現")
}

type concreteClassB struct {
	template
}

func NewClassB() *concreteClassB {
	classB := concreteClassB{}
	return &classB
}

func (b *concreteClassB) PrimitiveOperation1() {
	fmt.Println("具體類別B方法1實現")
}
func (b *concreteClassB) PrimitiveOperation2() {
	fmt.Println("具體類別B方法2實現")
}
