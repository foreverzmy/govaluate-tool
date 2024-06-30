package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	. "github.com/piex/govaluate-tool/parser"
)

func main() {
	runDemo("01")
	runDemo("02")
	runDemo("03")
	runDemo("04")
	runDemo("05")
}

func runDemo(path string) {
	basePath := "demo"

	inputFile := filepath.Join(basePath, path, "input.txt")
	tokenFile := filepath.Join(basePath, path, "token.json")
	astFile := filepath.Join(basePath, path, "ast.json")
	outputFile := filepath.Join(basePath, path, "output.txt")

	expression, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	tokens, _ := ParseTokens(string(expression), map[string]ExpressionFunction{
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

	code := ast.Generate()
	err = os.WriteFile(outputFile, []byte(code), 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

}
