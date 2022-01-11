### Interpreter 享元模式：
> 針對標的語言定義文法, 以及可解讀這個語句的解譯器

給定一個語言, 定義它的文法的一種表示, 並定義一個解譯器, 這個解譯器使用該表示來解釋語言中的句子

如果一種特定類型的問題發生的頻率夠高, 那麼可能就值得將該問題的各個實體表達為一個簡單語言中的句子. 這樣可以建構一個解譯器,
該解譯器透過解釋這些句子來解決問題.

![UML](https://github.com/kimi0230/DesignPatternGolang/blob/master/UML/Interpreter.png?raw=true)

### Example
https://github.com/kimi0230/DesignPatternGolang/tree/master/Interpreter 

```go
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

```