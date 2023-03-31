package htmlify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var text = `func MyFunc(s string) any {
	fmt.Printf("this is my function")
	return true
}`

func TestCode(t *testing.T) {
	tester := assert.New(t)
	out := GolangToHTML(text)
	tester.NotEmpty(out)
}
