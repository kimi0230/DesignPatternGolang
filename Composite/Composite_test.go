package composite

import (
	"fmt"
	"testing"
)

func TestComposite(t *testing.T) {
	root := NewConcreteCompany("總公司")
	root.Add(NewHRDepartment("總公司人力資源部"))
	root.Add(NewFinanceDepartment("總公司財務部"))

	comp := NewConcreteCompany("三重分公司")
	comp.Add(NewHRDepartment("三重分公司人力資源部"))
	comp.Add(NewFinanceDepartment("三重分公司財務部"))
	root.Add(comp)

	comp1 := NewConcreteCompany("內湖分公司")
	comp1.Add(NewHRDepartment("內湖分公司人力資源部"))
	comp1.Add(NewFinanceDepartment("內湖分公司財務部"))
	comp.Add(comp1)

	comp2 := NewConcreteCompany("汐止分公司")
	comp2.Add(NewHRDepartment("汐止分公司人力資源部"))
	comp2.Add(NewFinanceDepartment("汐止分公司財務部"))
	comp.Add(comp2)

	fmt.Println("結構圖")
	root.Display(1)

	fmt.Println("職責")
	root.LineOfDuty()
}
