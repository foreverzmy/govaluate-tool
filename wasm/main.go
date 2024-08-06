package main

import (
	"fmt"
	"syscall/js"

	. "github.com/piex/govaluate-tool/parser"
)

func main() {
	console := js.Global().Get("console")
	console.Call("log", "WASM Go Initialized version 0.0.1-alpha.8")
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
	console := js.Global().Get("console")
	expression := args[0].String()
	debug := false
	if len(args) == 3 {
		debug = args[2].Bool()
	}
	if debug {
		console.Call("log", fmt.Sprintf("expression: %s\n", expression))
	}

	functions := convertJsValueToFuncMap(args[1])
	if debug {
		console.Call("log", fmt.Sprintf("functions: %+v\n", functions))
	}

	tokens, err := ParseTokens(expression, functions)
	if err != nil {
		console.Call("error", fmt.Sprintf("tokens error: %e\n", err))
		return js.ValueOf(map[string]string{"error": err.Error()})
	}
	if debug {
		console.Call("log", fmt.Sprintf("tokens: %+v\n", tokens))
	}

	parser := NewParser(tokens)
	if debug {
		console.Call("log", fmt.Sprintf("NewParser: %+v\n", parser))
	}

	ast, err := parser.Parse()
	if err != nil {
		console.Call("error", fmt.Sprintf("parser error: %e\n", err))
		return js.ValueOf(map[string]string{"error": err.Error()})
	}
	if debug {
		console.Call("log", fmt.Sprintf("ast: %+v\n", ast))
	}

	code := ast.Generate()
	if debug {
		console.Call("log", fmt.Sprintf("code: %s\n", code))
	}

	return js.ValueOf(code)
}

func convertJsValueToFuncMap(functions js.Value) map[string]ExpressionFunction {
	funcMap := make(map[string]ExpressionFunction)
	keys := js.Global().Get("Object").Call("keys", functions)
	length := keys.Length()
	for i := 0; i < length; i++ {
		key := keys.Index(i).String()
		funcDef := functions.Get(key)
		parameters := make([]string, funcDef.Get("parameters").Length())
		for j := 0; j < funcDef.Get("parameters").Length(); j++ {
			parameters[j] = funcDef.Get("parameters").Index(j).String()
		}
		funcMap[key] = ExpressionFunction{
			Name:       key,
			Parameters: parameters,
			ReturnType: funcDef.Get("returnType").String(),
		}
	}
	return funcMap
}
