package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	. "github.com/piex/govaluate-tool/parser"
)

func main() {
	// 创建一个日志文件
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		log.Fatalf("无法打开日志文件: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)

	runDemo("01")
	runDemo("02")
	runDemo("03")
	runDemo("04")
	runDemo("05")
	runDemo("06")
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
		"getAbUidStr": {
			Name:       "getAbUidStr",
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
