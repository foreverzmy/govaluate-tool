package main

import (
	"fmt"

	. "github.com/piex/govaluate-tool/parser"
)

func main() {
	tokens := []ExpressionToken{
		{Kind: FUNCTION, Value: map[string]interface{}{"Name": "getAbUidInt64", "Parameters": []interface{}{}, "ReturnType": nil}, Raw: "getAbUidInt64", Start: 0, End: 13},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 13, End: 14},
		{Kind: VARIABLE, Value: "ctx", Raw: "ctx", Start: 15, End: 20},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 20, End: 21},
		{Kind: STRING, Value: "app", Raw: "app", Start: 22, End: 27},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 27, End: 28},
		{Kind: STRING, Value: "gc", Raw: "gc", Start: 29, End: 33},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 33, End: 34},
		{Kind: NUMERIC, Value: 0, Raw: "0", Start: 35, End: 37},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 37, End: 38},
		{Kind: COMPARATOR, Value: "==", Raw: "==", Start: 39, End: 42},
		{Kind: NUMERIC, Value: 2, Raw: "2", Start: 42, End: 44},
		{Kind: LOGICALOP, Value: "&&", Raw: "&&", Start: 44, End: 47},
		{Kind: PREFIX, Value: "!", Raw: "!", Start: 47, End: 48},
		{Kind: FUNCTION, Value: map[string]interface{}{"Name": "mapGet", "Parameters": []interface{}{}, "ReturnType": nil}, Raw: "mapGet", Start: 48, End: 54},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 54, End: 55},
		{Kind: VARIABLE, Value: "bp", Raw: "bp", Start: 56, End: 60},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 60, End: 61},
		{Kind: STRING, Value: "h", Raw: "h", Start: 62, End: 65},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 65, End: 66},
		{Kind: BOOLEAN, Value: false, Raw: "false", Start: 67, End: 73},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 73, End: 74},
		{Kind: LOGICALOP, Value: "&&", Raw: "&&", Start: 74, End: 77},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 77, End: 78},
		{Kind: PREFIX, Value: "!", Raw: "!", Start: 78, End: 79},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 79, End: 80},
		{Kind: ACCESSOR, Value: "[ctx AppId]", Raw: "ctx.AppId", Start: 80, End: 89},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 89, End: 90},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 90, End: 91},
		{Kind: COMPARATOR, Value: "in", Raw: "in", Start: 92, End: 95},
		{Kind: CLAUSE, Value: 40, Raw: "(", Start: 95, End: 96},
		{Kind: NUMERIC, Value: 1, Raw: "1", Start: 96, End: 97},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 97, End: 98},
		{Kind: NUMERIC, Value: 2, Raw: "2", Start: 98, End: 99},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 99, End: 100},
		{Kind: NUMERIC, Value: 3, Raw: "3", Start: 100, End: 101},
		{Kind: SEPARATOR, Value: ",", Raw: ",", Start: 101, End: 102},
		{Kind: NUMERIC, Value: 4, Raw: "4", Start: 102, End: 103},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 103, End: 104},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 104, End: 105},
		{Kind: CLAUSE_CLOSE, Value: 41, Raw: ")", Start: 105, End: 106},
	}
	parser := NewParser(tokens)

	ast, err := parser.Parse()
	if err != nil {
		panic(err)
	}

	fmt.Println(ast)
}
