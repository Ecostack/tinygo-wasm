//go:build js && wasm

package main

import (
	"crypto/rsa"
	"fmt"
	"math/big"
	"syscall/js"
)

func main() {
	fmt.Println("main is called")
	js.Global().Set("deriveKeyBySeed", deriveKeyBySeedWrapper())
	<-make(chan bool)
}

func deriveKeyBySeedWrapper() js.Func {
	jsonFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		//if len(args) != 1 {
		//	return "Invalid no of arguments passed"
		//}
		//inputJSON := args[0].String()
		//fmt.Printf("input %s\n", inputJSON)
		return deriveKeyBySeed()
		//pretty, err := deriveKeyBySeed()
		//if err != nil {
		//	fmt.Printf("unable to convert to json %s\n", err)
		//	return err.Error()
		//}
		//return pretty
	})
	return jsonFunc
}

func deriveKeyBySeed() rsa.PrivateKey {
	fmt.Println("some logs internal")
	return rsa.PrivateKey{
		PublicKey:   rsa.PublicKey{},
		D:           big.NewInt(123),
		Primes:      []*big.Int{big.NewInt(345), big.NewInt(234)},
		Precomputed: rsa.PrecomputedValues{},
	}

	//return 1234
	//return js.ValueOf(12345)
}
