package visitor

import "testing"

func TestVisitor(t *testing.T) {
	o := new(ObjectStructure)
	o.Attach(new(ConcreteElementA))
	o.Attach(new(ConcreteElementB))

	conreteVA := new(ConcreteVisitorA)
	conreteVB := new(ConcreteVisitorB)
	o.Accept(conreteVA)
	o.Accept(conreteVB)
}
