package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/msknapp/website/pkg/htmlify"
)

func main() {
	language := flag.String("l", "golang", "the language")
	flag.Parse()
	bts, _ := ioutil.ReadAll(os.Stdin)
	text := string(bts)
	var out string
	switch *language {
	case "golang", "go":
		out = htmlify.GolangToHTML(text)
	case "java":
		out = htmlify.JavaToHTML(text)
	case "python":
		out = htmlify.PythonToHTML(text)
	}
	fmt.Print(out)
}
