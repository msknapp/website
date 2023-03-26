package htmlify

import (
	"fmt"
	"regexp"
)

type WrapDefinition struct {
	Expression  *regexp.Regexp
	ReplaceWith string
}

func NewWrapDefinition(exp, clazz string) *WrapDefinition {
	expr := regexp.MustCompile(exp)
	rep := fmt.Sprintf("<span class=\"%s\">$1</span>", clazz)
	return &WrapDefinition{
		Expression:  expr,
		ReplaceWith: rep,
	}
}

func (w *WrapDefinition) WrapInString(text string) string {
	return w.Expression.ReplaceAllString(text, w.ReplaceWith)
}
