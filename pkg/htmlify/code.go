package htmlify

import (
	"fmt"
	"regexp"
	"strings"
)

type Converter struct {
	HashComments bool
	Wraps        []*WrapDefinition
}

func NewConverter(wraps []*WrapDefinition, hashComments bool) *Converter {
	return &Converter{
		Wraps:        wraps,
		HashComments: hashComments,
	}
}

func (c *Converter) IsSingleLineComment(line string) bool {
	if c.HashComments {
		m, _ := regexp.MatchString("^\\s*#", line)
		return m
	}
	m, _ := regexp.MatchString("^\\s*//", line)
	return m
}

func (c *Converter) StartsMultiLineComment(line string) bool {
	if c.HashComments {
		return false
	}
	m, _ := regexp.MatchString("\\b/\\*\\b", line)
	return m
}

func (c *Converter) EndsMultiLineComment(line string) bool {
	if c.HashComments {
		return false
	}
	m, _ := regexp.MatchString("\\b\\*/\\b", line)
	return m
}

func (c *Converter) Commentify(line string) string {
	return fmt.Sprintf("<span class=\"code-comment\">%s</span>", line)
}

func (c *Converter) EscapeWhiteSpace(line string) string {
	nbline := strings.ReplaceAll(line, " ", "&nbsp;")
	nbline = strings.ReplaceAll(nbline, "\t", "&nbsp;&nbsp;&nbsp;&nbsp;")
	return nbline
}

func (c *Converter) CodeToHTML(text string) string {
	out := &strings.Builder{}
	lines := strings.Split(text, "\n")
	inComment := false
	for i, line := range lines {
		if i > 0 {
			out.WriteString("<br>\n")
		}
		startsMulti := c.StartsMultiLineComment(line)
		stopsMulti := c.EndsMultiLineComment(line)
		isSingle := c.IsSingleLineComment(line)
		isComment := inComment || isSingle || startsMulti
		if startsMulti {
			inComment = true
		}
		if stopsMulti {
			inComment = false
		}
		line = c.EscapeWhiteSpace(line)
		if isComment {
			line = c.Commentify(line)
			out.WriteString(line)
			continue
		}
		for _, w := range c.Wraps {
			line = w.WrapInString(line)
		}
		out.WriteString(line)
	}
	return out.String()
}

func GolangToHTML(text string) string {
	wraps := []*WrapDefinition{
		NewWrapDefinition("\\b(func|import|package)\\b", "golang-top-level-keyword", ""),
		NewWrapDefinition("\\b(break|case|chan|const|continue|default|defer|else|for|go|goto|if|interface|map|range|return|select|struct|switch|type|var)\\b", "golang-control-keyword", ""),
		NewWrapDefinition("\\b(string|int64|int63|int|bool|byte|float64|float32|float|any|interface\\{\\})\\b", "golang-variable-type", ""),
		NewWrapDefinition("\\b(([a-z]+\\.)?\\w+)(\\()", "golang-function", "{{1}}$3"),
		NewWrapDefinition("\\b(true|false|\\d+(\\.\\d+)?)\\b", "golang-primitive-value", "{{1}}"),
		NewWrapDefinition("([{}()]+)", "golang-brace", "{{1}}"),
	}
	converter := NewConverter(wraps, false)
	return converter.CodeToHTML(text)
}

func PythonToHTML(text string) string {
	wraps := []*WrapDefinition{
		NewWrapDefinition("\\b(def|import|class|from)\\b", "python-top-level-keyword", ""),
		NewWrapDefinition("\\b(break|case|continue|default|else|elif|for|if|range|return|select|switch|type|var|in|while|self|None)\\b", "python-control-keyword", ""),
		NewWrapDefinition("\\b(string|int|bool|float|List|Set|Dict)\\b", "python-variable-type", ""),
		NewWrapDefinition("\\b(\\w+)(\\()", "python-function", "{{1}}$2"),
	}
	converter := NewConverter(wraps, true)
	return converter.CodeToHTML(text)
}

func JavaToHTML(text string) string {
	wraps := []*WrapDefinition{
		NewWrapDefinition("\\b(public|static|private|protected|import|package)\\b", "java-top-level-keyword", ""),
		NewWrapDefinition("\\b(break|case|continue|default|else|for|if|interface|return|switch|type)\\b", "java-control-keyword", ""),
		NewWrapDefinition("\\b(String|int|long|short|bool|byte|float|double)\\b", "java-variable-type", ""),
		NewWrapDefinition("\\b(\\w+)(\\()", "golajavang-function", "{{1}}$2"),
	}
	converter := NewConverter(wraps, false)
	return converter.CodeToHTML(text)
}
