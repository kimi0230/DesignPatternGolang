package chainofresponsibility

import "testing"

func TestChainofresponsibility(t *testing.T) {
	h1 := NewConcreteHandler1("handler 1")
	h2 := NewConcreteHandler2("handler 2")
	h3 := NewConcreteHandler3("handler 3")

	h1.SetSuccessor(h2)
	h2.SetSuccessor(h3)

	requests := []int{2, 5, 14, 22, 18, 3, 27, 20}
	for _, v := range requests {
		h1.HandleRequest(v)
	}
}
