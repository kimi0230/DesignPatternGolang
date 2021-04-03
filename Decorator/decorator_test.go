package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {
	tests := []struct {
		name       string
		personName string
	}{
		{
			name:       "TestStrategy 正常收費",
			personName: "kimi",
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
