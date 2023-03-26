package main

import (
	"syscall/js"

	"github.com/msknapp/website/pkg/htmlify"
)

func main() {
	js.Global().Set("golangToHTML", golangToHTML())
	js.Global().Set("pythonToHTML", pythonToHTML())
	js.Global().Set("javaToHTML", javaToHTML())
	<-make(chan bool)
}

func golangToHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputText := args[0].String()
		return htmlify.GolangToHTML(inputText)
	})
}

func pythonToHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputText := args[0].String()
		return htmlify.PythonToHTML(inputText)
	})
}

func javaToHTML() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputText := args[0].String()
		return htmlify.JavaToHTML(inputText)
	})
}
