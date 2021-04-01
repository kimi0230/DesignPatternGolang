package visitor

import "testing"

func TestDemo(t *testing.T) {
	o := new(ObjectStructurePerson)
	o.Attach(&Man{"男人"})
	o.Attach(&Woman{"女人"})

	// 成功
	conreteVA := &Success{"成功"}
	// 失敗
	conreteVB := &Failing{"失敗"}
	o.Accept(conreteVA)
	o.Accept(conreteVB)
}
