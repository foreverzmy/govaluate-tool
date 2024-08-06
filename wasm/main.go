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

	functions := convertJsValueToFuncMap(args[1])
	tokens, _ := ParseTokens(expression, functions)

	parser := NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		return js.ValueOf(map[string]string{"error": err.Error()})
	}

	code := ast.Generate()

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
