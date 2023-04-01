const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("/static/js/chessgui"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)
    });