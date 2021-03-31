package interpreter

import (
	"fmt"
	"strconv"
	"strings"
)

// PlayContext : 演奏內容 Context
type PlayContext struct {
	text string
}

func (c *PlayContext) Text() string {
	if c == nil {
		return ""
	}
	return c.text
}
func (c *PlayContext) SetText(str string) {
	if c == nil {
		return
	}
	c.text = str
}

// 運算式類別 AbstractExpression
type IExpression interface {
	Interpret(*PlayContext, IExpression)
	Execute(key string, val int)
}

type Expression struct {
}

func (e *Expression) Interpret(context *PlayContext, ie IExpression) {
	// fmt.Println("context.Text() ", context.Text())
	if len(context.Text()) == 1 {
		return
	}
	playKey := context.Text()[0:1]
	context.SetText(context.Text()[2:])
	spaceIndex := strings.Index(context.Text(), " ")
	// fmt.Println("spaceIndex ", spaceIndex)
	if spaceIndex < 0 {
		spaceIndex = 0
	}
	playVal, _ := strconv.Atoi(string(context.Text()[0:spaceIndex]))
	// fmt.Println("playVal", string(context.Text()[0]))
	context.SetText(context.Text()[strings.Index(context.Text(), " ")+1:])

	ie.Execute(playKey, playVal)
}

func (e *Expression) Execute(key string, val int) {
}

type Note struct {
	Expression
}

func NewNote() *Note {
	return &Note{Expression{}}
}

func (n *Note) Execute(key string, val int) {
	note := ""
	switch key {
	case "C":
		note = "1"
	case "D":
		note = "2"
	case "E":
		note = "3"
	case "F":
		note = "4"
	case "G":
		note = "5"
	case "A":
		note = "6"
	case "B":
		note = "7"
	}
	fmt.Printf("%s", note)
}

type Scale struct {
	Expression
}

func NewScale() *Scale {
	return &Scale{Expression{}}
}

func (s *Scale) Execute(key string, val int) {
	scale := ""
	switch val {
	case 1:
		scale = "低音"
	case 2:
		scale = "中音"
	case 3:
		scale = "高音"
	}
	fmt.Printf("%s", scale)
}
