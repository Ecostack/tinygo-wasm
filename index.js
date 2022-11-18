

const go = new Go();
const importObject = go.importObject;

WebAssembly.instantiateStreaming(fetch("wasm.wasm"),importObject).then(
    (results) => {
        go.run(results.instance);

        debugger;

        // const someVal = results.instance.exports.deriveKeyBySeed(2048, new Uint8Array([1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4]))
        // const someVal = deriveKeyBySeed(2048, new Uint8Array([1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4, 1, 2, 3, 4]))
        const someVal = deriveKeyBySeed()

        console.log(someVal)

        // Do something with the results!
    }
);