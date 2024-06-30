package parser

import (
	"fmt"
	"strings"
)

// ASTNode 表示 AST 的节点
type ASTNode struct {
	Token    *ExpressionToken
	Children []*ASTNode
}

func (ast *ASTNode) Generate() string {
	return ast.generateWithIndent(0)
}

// GenerateWithIndent 生成带有缩进和换行的代码
func (ast *ASTNode) generateWithIndent(indent int) string {
	if ast.Token == nil {
		return ""
	}

	var sb strings.Builder
	indentation := strings.Repeat("  ", indent)

	switch ast.Token.Kind {
	case PREFIX:
		children := ast.Children[0].generateWithIndent(indent + 1)
		sb.WriteString(ast.Token.Raw)
		sb.WriteString("( ")
		if len(children) > 100 {
			sb.WriteString("\n")
			sb.WriteString(indentation)
		}
		sb.WriteString(children)
		if len(children) > 100 {
			sb.WriteString("\n")
		}
		sb.WriteString(")")
	case NUMERIC, BOOLEAN:
		sb.WriteString(ast.Token.Raw)
	case STRING:
		sb.WriteString(fmt.Sprintf("'%s'", ast.Token.Raw))
	case PATTERN:
	case TIME:
	case VARIABLE:
		sb.WriteString(fmt.Sprintf("[%s]", ast.Token.Raw))
	case FUNCTION:
		sb.WriteString(ast.Token.Raw)
		sb.WriteString("( ")
		for i, child := range ast.Children {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(child.generateWithIndent(0))
		}
		sb.WriteString(" )")
	case SEPARATOR:
	case ACCESSOR:
	case COMPARATOR:
		sb.WriteString(ast.Children[0].generateWithIndent(indent))
		sb.WriteString(" ")
		sb.WriteString(ast.Token.Value.(string))
		sb.WriteString(" ")
		sb.WriteString(ast.Children[1].generateWithIndent(indent))
	case LOGICALOP:
		isLeftLogical := ast.Children[0].Token.Kind == LOGICALOP
		isRightLogical := ast.Children[1].Token.Kind == LOGICALOP
		leftIndent := indent
		rightIndent := indent
		if isLeftLogical {
			leftIndent = leftIndent + 1
		}
		if isRightLogical {
			rightIndent = rightIndent + 1
		}
		left := ast.Children[0].generateWithIndent(leftIndent)
		right := ast.Children[1].generateWithIndent(rightIndent)
		sb.WriteString(indentation)
		if isLeftLogical {
			sb.WriteString("(\n")
		}
		sb.WriteString(left)
		sb.WriteString("\n")
		if isLeftLogical {
			sb.WriteString(indentation)
			sb.WriteString(")\n")
		}
		sb.WriteString(indentation)
		sb.WriteString(ast.Token.Value.(string))
		sb.WriteString("\n")
		sb.WriteString(indentation)
		if isRightLogical {
			sb.WriteString("(\n")
		}
		sb.WriteString(right)
		if isRightLogical {
			sb.WriteString("\n")
			sb.WriteString(indentation)
			sb.WriteString(")")
		}
	case MODIFIER:
	case CLAUSE:
		sb.WriteString("(")
	case CLAUSE_CLOSE:
		sb.WriteString(")")
	case TERNARY:
	case ARRAY:
		sb.WriteString("( ")
		for i, child := range ast.Children {
			if i > 0 {
				sb.WriteString(", ")
			}
			sb.WriteString(child.generateWithIndent(0))
		}
		sb.WriteString(" )")
	default:
		return ""
	}

	return sb.String()
}
