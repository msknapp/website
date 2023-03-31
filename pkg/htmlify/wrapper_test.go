package htmlify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrapper(t *testing.T) {
	tester := assert.New(t)
	wr := NewWrapDefinition("\\b(\\w+)(\\()", "theclazz", "{{1}}$2")
	out := wr.WrapInString("foo bar(baz)")
	tester.Equal("foo <span class=\"theclazz\">bar</span>(baz)", out)

	wr = NewWrapDefinition("\\b(func|return)\\b", "control", "")
	out = wr.WrapInString("the func will return")
	tester.Equal("the <span class=\"control\">func</span> will <span class=\"control\">return</span>", out)
}
