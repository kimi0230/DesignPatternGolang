package factorymethod

import (
	"testing"
)

func TestSampleFactory(t *testing.T) {
	tests := []struct {
		name string
		oper string
		numA float32
		numB float32
		want float32
	}{
		{
			name: "SampleFactory Add",
			oper: "+",
			numA: 10,
			numB: 120,
			want: 130,
		},
		{
			name: "SampleFactory Sub",
			oper: "-",
			numA: 100,
			numB: 110,
			want: -10,
		},
		{
			name: "SampleFactory Mul",
			oper: "*",
			numA: 10,
			numB: 110,
			want: 1100,
		},
		{
			name: "SampleFactory Div",
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
