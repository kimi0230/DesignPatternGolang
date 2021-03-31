package interpreter

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	context := new(PlayContext)
	context.SetText("O 2 E 4 G 4 A 3 E 4 G 4 D 3 E 4 G 4 A 4 O 3 C 1 O 2 A 4 G 1 C 4 E 4 D 3")
	// expression := NewScale()
	// expression.Interpret(context)

	fmt.Println(context.Text())
	for len(context.Text()) > 1 {
		str := context.Text()[0:1]
		var expression IExpression
		switch str {
		case "O":
			expression = NewScale()
		case "C", "D", "E", "F", "G", "A", "B", "P":
			// fmt.Println("dd")
			expression = NewNote()
		}
		// expression := NewScale()
		expression.Interpret(context, expression)
	}

}
