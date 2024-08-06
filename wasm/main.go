package main

import (
	"syscall/js"

	. "github.com/piex/govaluate-tool/parser"
)

func main() {
	console := js.Global().Get("console")
	console.Call("log", "WASM Go Initialized")
	c := make(chan struct{}, 0)
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	global := js.Global()
	globalWasmFunc := global.Get("globalWasmFunc")

	if globalWasmFunc.Type() == js.TypeUndefined || globalWasmFunc.Type() == js.TypeNull {
		// 如果 globalWasmFunc 不存在，则创建一个新的对象
		globalWasmFunc = js.Global().Get("Object").New()
		global.Set("globalWasmFunc", globalWasmFunc)
	}

	globalWasmFunc.Set("format", js.FuncOf(format))
}

func format(this js.Value, args []js.Value) interface{} {
	expression := args[0].String()

	tokens, _ := ParseTokens(expression, map[string]ExpressionFunction{
		"mapGet": {
			Name:       "mapGet",
			Parameters: []string{},
			ReturnType: "",
		},
		"isNil": {
			Name:       "isNil",
			Parameters: []string{},
			ReturnType: "",
		},
		"getValue": {
			Name:       "getValue",
			Parameters: []string{},
			ReturnType: "",
		},
		"StrLen": {
			Name:       "StrLen",
			Parameters: []string{},
			ReturnType: "",
		},
		"getAbUidStr": {
			Name:       "getAbUidStr",
			Parameters: []string{},
			ReturnType: "",
		},
	})

	parser := NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return js.ValueOf(map[string]string{"error": err.Error()})
	}

	code := ast.Generate()

	return js.ValueOf(code)
}
