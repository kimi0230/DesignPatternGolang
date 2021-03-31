package interpreter

import "fmt"

type Context struct {
	input  string
	output string
}

func (c *Context) Input() string {
	if c == nil {
		return ""
	}
	return c.input
}
func (c *Context) SetInput(str string) {
	if c == nil {
		return
	}
	c.input = str
}

func (c *Context) Output() string {
	if c == nil {
		return ""
	}
	return c.output
}
func (c *Context) SetOutput(str string) {
	if c == nil {
		return
	}
	c.output = str
}

// 抽象運算式
type AbstractExpression interface {
	Interpret(*Context)
}

type TerminalExpression struct {
}

func (t *TerminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	fmt.Println("終端解譯器")
}

type NonterminalExpression struct {
}

func (t *NonterminalExpression) Interpret(context *Context) {
	if t == nil {
		return
	}
	fmt.Println("非終端解譯器")
}
