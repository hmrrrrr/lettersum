package main

import (
	"fmt"
	"syscall/js"
)

func getJsDoc() js.Value {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		fmt.Println("Unable to get document")
	}
	return jsDoc
}

func convertToSum(input string) (sum int32) {
	for i := 0; i < len(input); i++ {
		if (int32(input[i])-96 >= 0) && int32(input[i])-96 < 26 {
			sum += int32(input[i]) - 96
		}
	}
	return sum
}

func main() {
	fmt.Println("running..!")
	js.Global().Set("convertToSum", convertToSumWrapper())
	<-make(chan bool)
}

func convertToSumWrapper() js.Func {
	ctsFunc := js.FuncOf(func(this js.Value, args []js.Value) any {
		if len(args) != 1 {
			return "Invalid args"
		}
		input := args[0].String()
		fmt.Printf("input %s\n", input)
		jsDoc := getJsDoc()
		jsOutputArea := jsDoc.Call("getElementById", "textoutput")
		jsOutputArea.Set("value", convertToSum(input))
		return nil
	})
	return ctsFunc
}
