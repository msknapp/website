package main

import (
	"math"
	"math/big"
	"strconv"
	"syscall/js"

	"github.com/msknapp/website/pkg/factorial"
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

func Factorial() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputString := args[0].String()
		input, e := strconv.Atoi(inputString)
		if e != nil {
			return e.Error()
		}
		return factorial.Factorial(int32(input))
	})
}

func IsPrime() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid no of arguments passed"
		}
		inputString := args[0].String()
		i := &big.Int{}
		i, ok := i.SetString(inputString, 10)
		if !ok {
			return "The input is not an integer"
		}
		if i.Sign() < 1 {
			return false
		}
		if i.Sub(i, big.NewInt(math.MaxInt64)).Sign() < 1 {
			return i.ProbablyPrime(2)
		}
		return i.ProbablyPrime(12)
	})
}

func Permutations() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
		args[0].S
	})
}

func Combinations() js.Func {
	return js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 2 {
			return "Invalid no of arguments passed"
		}
	})
}
