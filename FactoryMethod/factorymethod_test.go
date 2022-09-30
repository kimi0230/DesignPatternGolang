package factorymethod

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {
	tests := []struct {
		name string
		oper string
		numA float32
		numB float32
		want float32
	}{
		{
			name: "FactoryMethod Add",
			oper: "+",
			numA: 10,
			numB: 120,
			want: 130,
		},
		{
			name: "FactoryMethod Sub",
			oper: "-",
			numA: 100,
			numB: 110,
			want: -10,
		},
		{
			name: "FactoryMethod Mul",
			oper: "*",
			numA: 10,
			numB: 110,
			want: 1100,
		},
		{
			name: "FactoryMethod Div",
			oper: "/",
			numA: 10,
			numB: 5,
			want: 2,
		},
	}

	for _, tt := range tests {
		var of *OperationFactory
		add := of.createoperation(tt.oper)
		add.SetNumA(tt.numA)
		add.SetNumB(tt.numB)
		if val, ok := add.GetResult(); ok == true && val != tt.want {
			t.Error("Add Error", val)
		}
	}
}
