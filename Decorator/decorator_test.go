package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	tests := []struct {
		name       string
		personName string
		decorator1 interface{}
		decorator2 interface{}
	}{
		{
			name:       "TestStrategy 正常收費",
			personName: "kimi",
			decorator1: new(TShirts),
			decorator2: new(BigTrouser),
		},
	}

	for _, tt := range tests {
		person := &Person{tt.personName}
		// person.show()

		ts := new(TShirts)
		ts.SetDecorator(person)
		// ts.show()

		bt := new(BigTrouser)
		bt.SetDecorator(ts)
		bt.show()

	}
}
