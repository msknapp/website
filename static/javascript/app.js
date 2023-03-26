
function runConversion(inputLanguage) {
    console.info("running");
    let s = document.getElementById("app-input").value;
    let out = "";
    if (inputLanguage == "golang") {
        out = golangToHTML(s);
    } else if (inputLanguage == "python") {
        out = pythonToHTML(s);
    } else {
        out = javaToHTML(s);
    }
    out = out.replaceAll(' ', "&nbsp;");
    out = out.replaceAll('\t', "&nbsp;&nbsp;&nbsp;&nbsp;");
    document.getElementById("app-output").value = out;
}

const go = new Go();
WebAssembly.instantiateStreaming(fetch("/wasm/app.wasm"), go.importObject).then((result) => {
    go.run(result.instance);
});