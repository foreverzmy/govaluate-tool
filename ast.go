package main

import (
	"fmt"
)

// ASTNode 表示 AST 的节点
type ASTNode struct {
	Token    *ExpressionToken
	Children []*ASTNode
}

// parseExpression 解析表达式标记数组生成 AST
func parseExpression(tokens []ExpressionToken) (*ASTNode, error) {
	if len(tokens) == 0 {
		return nil, fmt.Errorf("no tokens to parse")
	}

	stack := []*ASTNode{}

	for _, token := range tokens {
		node := &ASTNode{Token: &token}

		switch token.Kind {
		case VARIABLE, BOOLEAN, NUMERIC, STRING, TIME:
			stack = append(stack, node)
		case PATTERN:
		case FUNCTION:
			stack = append(stack, node)
		case SEPARATOR:
		case ACCESSOR:
		case COMPARATOR, LOGICALOP:
			left := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			node.Children = append(node.Children, left)
			stack = append(stack, node)
		case MODIFIER:
		case CLAUSE:
			stack = append(stack, node)
		case CLAUSE_CLOSE:
			subAstLength := extractSubAST(stack)
			var subAST []*ASTNode
			stack, subAST = splitLastN(stack, subAstLength)
			stack[len(stack)-1].Children = subAST[1:]
		}
	}

	for _, ast := range stack {
		fmt.Println(ast.Token.Kind.String(), ast.Token.Content)
	}

	if len(stack) != 1 {
		return nil, fmt.Errorf("invalid expression, could not resolve to a single root node")
	}

	return stack[0], nil
}

// extractSubExpression 提取子表达式标记
func extractSubAST(ast []*ASTNode) (count int) {
	for i := len(ast) - 1; i >= 0; i-- {
		count++
		if ast[i].Token.Kind == CLAUSE {
			return
		}
	}

	return
}

func splitLastN(nodes []*ASTNode, n int) ([]*ASTNode, []*ASTNode) {
	if n > len(nodes) {
		n = len(nodes)
	}
	lastN := nodes[len(nodes)-n:]
	remaining := nodes[:len(nodes)-n]
	return remaining, lastN
}
