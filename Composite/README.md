### Composite 組合模式： 
> 將物件組織成樹狀結構, 部分-全體 層級關係, 讓外界以一致性的方式對待個別物件和整體物件

將物件組合成**樹形結構**以表示 部分-整體的層次結構. 組合模式使得用戶對單個物件和組合物件的使用具有一致性

* 何時使用組合模式
需求中是表現部分與整體曾是的結構時, 以及你希望用戶可以忽略組合物件與單個物件的不同, 統一地使用組合結構中的所有物件時, 就應該考慮組合模式.

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Composite.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Composite 

```go
package composite

import (
	"fmt"
	"strings"
)

// Company : Component
type Company interface {
	Add(Company)
	Remove(Company)
	Display(int)
	LineOfDuty()
}

type RealCompany struct {
	name string
}

// ConcreteCompany : Composite
type ConcreteCompany struct {
	RealCompany
	children []Company
}

func NewConcreteCompany(name string) *ConcreteCompany {
	return &ConcreteCompany{RealCompany{name}, []Company{}}
}

func (c *ConcreteCompany) Add(com Company) {
	if c == nil {
		return
	}
	c.children = append(c.children, com)
}

func (c *ConcreteCompany) Remove(com Company) {
	if c == nil {
		return
	}
	for i, v := range c.children {
		if v == com {
			c.children = append(c.children[:i], c.children[i+1:]...)
			return
		}
	}
	return
}

func (c *ConcreteCompany) Display(depth int) {
	if c == nil {
		return
	}
	fmt.Println(strings.Repeat("  ", depth-1), "|--", c.name)
	for _, v := range c.children {
		v.Display(depth + 2)
	}
}

func (c *ConcreteCompany) LineOfDuty() {
	if c == nil {
		return
	}
	for _, v := range c.children {
		v.LineOfDuty()
	}
}

// HR
type HRDepartment struct {
	RealCompany
}

func NewHRDepartment(name string) *HRDepartment {
	return &HRDepartment{RealCompany{name}}
}

func (h *HRDepartment) Add(c Company)    {}
func (h *HRDepartment) Remove(c Company) {}
func (h *HRDepartment) Display(depth int) {
	if h == nil {
		return
	}
	fmt.Println(strings.Repeat("  ", depth-1), "|--", h.name)
}

func (h *HRDepartment) LineOfDuty() {
	if h == nil {
		return
	}
	fmt.Println(h.name, "員工招聘教育訓練管理")
}

// 財務部
type FinanceDepartment struct {
	RealCompany
}

func NewFinanceDepartment(name string) *FinanceDepartment {
	return &FinanceDepartment{RealCompany{name}}
}

func (f *FinanceDepartment) Add(c Company)    {}
func (f *FinanceDepartment) Remove(c Company) {}
func (f *FinanceDepartment) Display(depth int) {
	if f == nil {
		return
	}
	fmt.Println(strings.Repeat("  ", depth-1), "|--", f.name)
}

func (f *FinanceDepartment) LineOfDuty() {
	if f == nil {
		return
	}
	fmt.Println(f.name, "公司財務收支管理")
}
```