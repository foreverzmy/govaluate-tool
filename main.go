package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

var (
	inputFile  = "demo/input.txt"
	tokenFile  = "demo/token.json"
	astFile    = "demo/ast.txt"
	outputFile = "demo/output.txt"
)

func main() {
	expression, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tokens, _ := parseTokens(string(expression), map[string]ExpressionFunction{
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
		"getAbUidInt64": {
			Name:       "getAbUidInt64",
			Parameters: []string{},
			ReturnType: "",
		},
		"StrLen": {
			Name:       "StrLen",
			Parameters: []string{},
			ReturnType: "",
		},
	})

	jsonStr, _ := json.MarshalIndent(tokens, "", "  ")
	jsonStr = bytes.ReplaceAll(jsonStr, []byte(`\u0026`), []byte(`&`))

	// 将修改后的内容写入新文件
	err = os.WriteFile(tokenFile, jsonStr, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	parser := NewParser(tokens)
	ast, err := parser.Parse()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	astJson, _ := json.MarshalIndent(ast, "", "  ")

	// 将修改后的内容写入新文件
	err = os.WriteFile(astFile, astJson, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	code := generate(ast, 0)
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}
