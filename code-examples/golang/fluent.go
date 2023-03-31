package golang

import "fmt"

type MyFluentThing struct {
	Value int
	Text  string
}

func NewFluentThing() *MyFluentThing {
	return &MyFluentThing{}
}
func (t *MyFluentThing) WithValue(v int) *MyFluentThing {
	t.Value = v
	return t
}

func (t *MyFluentThing) WithText(s string) *MyFluentThing {
	t.Text = s
	return t
}

func Run() {
	x := NewFluentThing().WithText("foo").WithValue(3)
	fmt.Printf("%s, %d", x.Text, x.Value)
}
