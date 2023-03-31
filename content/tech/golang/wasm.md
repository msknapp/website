---
title: Web Assembly (Go)
draft: false
weight: 49
lastmod: 2023-03-29
summary: A brief summary of how to use web assembly with Golang.
---

This contains a very brief summary of the steps to make and use a Golang Web Assembly artifact.

# Write Main
The main class/package needs to use the syscall/js package, and never stop.  Use this main class as an example:
```
package main

import (
    "strconv"
    "syscall/js"
)

func main() {
    // Add functions as needed.
    js.Global().Set("myFunction", myFunction())

    // Prevent the program from stopping.
    <-make(chan bool)
}

func myFunction() js.Func {  
    return js.FuncOf(func(this js.Value, args []js.Value) any {
        if len(args) != 1 {
            return "Invalid no of arguments passed"
        }
        input := args[0].String()

        // Incorporate business logic here.
        // You can call on functions in other packages too.
        output := input

        return output
    })
}
```

You should also follow the usual steps of creating a golang module and adding dependencies to it as needed.

# Build WASM

1. Compile with `GOOS=js GOARCH=wasm go build -o app.wasm wasm/main.go`
2. Copy the artifact somewhere into your website's static content.
3. Copy the Golang wasm javascript glue somewhere into your website's static content: `cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" /path/to/website/static`

# Update Web Page

Copy this into your html head section:
```
<script src="wasm_exec.js"></script>
<script>
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject).then((result) => {
        go.run(result.instance);
    });
</script>
```

In your page's html and javascript, update the code to utilize the functions that your wasm application defined.

For more information, refer to [Golang Bot's Web Assembly Using Go](https://golangbot.com/webassembly-using-go/).