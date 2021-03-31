package interpreter

import (
	"fmt"
	"testing"
)

func TestDemo(t *testing.T) {
	context := new(PlayContext)
	context.SetText("O 2 T 1500 E 4 G 4 A 3 E 4 T 500 G 4 D 3 E 4 G 4 A 4 O 3 C 1 O 2 A 4 G 1 C 4 E 4 D 3")
	// 中音 慢速 3 5 6 3 中速 5 2 3 5 6 高音 1 中音 6 5 1 3 2
	fmt.Println(context.Text())
	for len(context.Text()) > 1 {
		str := context.Text()[0:1]
		var expression IExpression
		switch str {
		case "O":
			expression = NewScale()
		case "T":
			expression = NewSpeed()
		case "C", "D", "E", "F", "G", "A", "B", "P":
			// fmt.Println("dd")
			expression = NewNote()
		}
		expression.Interpret(context, expression)
	}
}
