package htmlify

import (
	"fmt"
	"regexp"
	"strings"
)

type WrapDefinition struct {
	Expression  *regexp.Regexp
	ReplaceWith string
}

func createReplaceWith(clazz, replacePattern string) string {
	open := fmt.Sprintf("<span class=\"%s\">$", clazz)
	cls := "</span>"
	out := strings.Replace(replacePattern, "{{", open, 1)
	out = strings.Replace(out, "}}", cls, 1)
	return out
}

func NewWrapDefinition(exp, clazz, replacePattern string) *WrapDefinition {
	expr := regexp.MustCompile(exp)
	if replacePattern == "" {
		replacePattern = "{{1}}"
	}
	rep := createReplaceWith(clazz, replacePattern)

	return &WrapDefinition{
		Expression:  expr,
		ReplaceWith: rep,
	}
}

func (w *WrapDefinition) WrapInString(text string) string {
	return w.Expression.ReplaceAllString(text, w.ReplaceWith)
}
