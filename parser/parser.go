package parser

import (
	"fmt"
)

func newASTNode(token *ExpressionToken) *ASTNode {
	return &ASTNode{Token: token, Children: []*ASTNode{}}
}

type Parser struct {
	tokens []ExpressionToken
	pos    int
}

func NewParser(tokens []ExpressionToken) *Parser {
	return &Parser{tokens: tokens, pos: 0}
}

func (p *Parser) Parse() (*ASTNode, error) {
	return p.parseExpression(0)
}

func (p *Parser) parseExpression(precedence int) (*ASTNode, error) {
	left, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for {
		token := p.peek()
		if token == nil {
			break
		}

		// log.Printf("parseExpression peek token: %s, start %d end %d\n", token.Raw, token.Start, token.End)

		tokenPrecedence := p.getPrecedence(token)
		if tokenPrecedence < precedence {
			break
		}

		if token.Kind == CLAUSE_CLOSE {
			break
		}

		left, err = p.parseBinaryExpression(left, tokenPrecedence)
		if err != nil {
			return nil, err
		}
	}

	return left, nil
}

func (p *Parser) parseBinaryExpression(left *ASTNode, precedence int) (*ASTNode, error) {
	token := p.peek()

	// log.Printf("parseBinaryExpression peek token: %s, start %d end %d\n", token.Raw, token.Start, token.End)

	switch token.Kind {
	case LOGICALOP:
		return p.parseLogicalOp(left, precedence)
	case COMPARATOR:
		return p.parseComparator(left, precedence)
	default:
		// node, err = p.parseExpression(precedence + 1)
		// log.Fatalf("parseBinaryExpression unexpected token: %v", token)
		return left, fmt.Errorf("parseBinaryExpression unexpected token: %v", token)
	}
}

func (p *Parser) parsePrimaryExpression() (*ASTNode, error) {
	token := p.peek()
	if token == nil {
		return nil, fmt.Errorf("unexpected end of tokens")
	}

	// log.Printf("parsePrimaryExpression peek token: %s, kind %s, start %d, end %d\n", token.Raw, token.Kind.String(), token.Start, token.End)

	switch token.Kind {
	case PREFIX:
		return p.parsePrefix()
	case NUMERIC:
		return p.parseNumeric()
	case BOOLEAN:
		return p.parseBoolean()
	case STRING:
		return p.parseString()
	case PATTERN:
		return p.parsePattern()
	case TIME:
		return p.parseTime()
	case VARIABLE:
		return p.parseVariable()
	case FUNCTION:
		return p.parseFunction()
	case ACCESSOR:
		return p.parseAccessor()
	case MODIFIER:
		return p.parseModifier()
	case CLAUSE:
		return p.parseClause()
	case TERNARY:
		return p.parseTernary()
	}

	return nil, fmt.Errorf("unexpected token: %v", token)
}

func (p *Parser) parsePrefix() (*ASTNode, error) {
	token := p.next()
	if token.Kind != PREFIX {
		return nil, fmt.Errorf("expected prefix token, got %v", token)
	}

	// log.Printf("parsePrefix peek token: %s, start %d end %d\n", token.Raw, token.Start, token.End)

	node := newASTNode(token)

	// Parse the following expression after the prefix
	expr, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}
	node.Children = append(node.Children, expr)

	return node, nil
}

func (p *Parser) parseNumeric() (*ASTNode, error) {
	return p.parseToken(NUMERIC)
}

func (p *Parser) parseBoolean() (*ASTNode, error) {
	return p.parseToken(BOOLEAN)
}

func (p *Parser) parseString() (*ASTNode, error) {
	return p.parseToken(STRING)
}

func (p *Parser) parsePattern() (*ASTNode, error) {
	return p.parseToken(PATTERN)
}

func (p *Parser) parseTime() (*ASTNode, error) {
	return p.parseToken(TIME)
}

func (p *Parser) parseVariable() (*ASTNode, error) {
	return p.parseToken(VARIABLE)
}

func (p *Parser) parseFunction() (*ASTNode, error) {
	node, err := p.parseToken(FUNCTION)
	if err != nil {
		return nil, err
	}

	if err := p.expectToken(CLAUSE); err != nil {
		return nil, err
	}

	// Parse function arguments
	args, err := p.parseArguments()
	if err != nil {
		return nil, err
	}

	node.Children = append(node.Children, args...)

	return node, nil
}

func (p *Parser) parseArguments() ([]*ASTNode, error) {
	var args []*ASTNode

	for {
		if p.peek() == nil {
			return nil, fmt.Errorf("unexpected end of tokens in function arguments")
		}

		// End of arguments list
		if p.peek().Kind == CLAUSE_CLOSE {
			p.next() // consume ')'
			break
		}

		// Parse individual argument
		arg, err := p.parsePrimaryExpression()
		if err != nil {
			return nil, err
		}
		args = append(args, arg)

		// Arguments are separated by commas
		if p.peek().Kind == SEPARATOR {
			p.next() // Consume ','
		}
	}

	return args, nil
}

func (p *Parser) parseAccessor() (*ASTNode, error) {
	token := p.next()
	if token.Kind != ACCESSOR {
		return nil, fmt.Errorf("expected accessor token, got %v", token)
	}

	node := newASTNode(token)

	ptoken := p.peek()
	if ptoken != nil && ptoken.Kind == CLAUSE {
		p.next()
		if p.expectToken(CLAUSE_CLOSE) != nil {
			return nil, fmt.Errorf("expected ')' to close accessor function")
		}
		node.Children = append(node.Children, newASTNode(ptoken))
	}

	return node, nil
}

func (p *Parser) parseComparator(left *ASTNode, precedence int) (*ASTNode, error) {
	node, err := p.parseToken(COMPARATOR)
	if err != nil {
		return nil, err
	}
	node.Children = append(node.Children, left)

	if p.peek().Kind == CLAUSE {
		right, err := p.parseClauseOrArray()
		if err != nil {
			return nil, err
		}
		node.Children = append(node.Children, right)
	} else {
		right, err := p.parseExpression(precedence + 1)
		if err != nil {
			return nil, err
		}
		node.Children = append(node.Children, right)
	}

	return node, nil
}

func (p *Parser) parseLogicalOp(left *ASTNode, precedence int) (*ASTNode, error) {
	node, err := p.parseToken(LOGICALOP)
	if err != nil {
		return nil, err
	}

	node.Children = append(node.Children, left)

	right, err := p.parseExpression(precedence + 1)
	if err != nil {
		return nil, err
	}
	node.Children = append(node.Children, right)

	return node, nil
}

func (p *Parser) parseModifier() (*ASTNode, error) {
	node, err := p.parseToken(MODIFIER)
	if err != nil {
		return nil, err
	}

	expr, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}
	node.Children = append(node.Children, expr)

	return node, nil
}

func (p *Parser) parseTernary() (*ASTNode, error) {
	condition, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}

	if p.peek() == nil || p.peek().Kind != TERNARY || p.peek().Raw != "?" {
		return nil, fmt.Errorf("expected '?' for ternary operator")
	}
	p.next() // consume '?'

	trueExpr, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}

	if p.peek() == nil || p.peek().Kind != TERNARY || p.peek().Raw != ":" {
		return nil, fmt.Errorf("expected ':' in ternary operator")
	}
	p.next() // consume ':'

	falseExpr, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}

	node := newASTNode(&ExpressionToken{Kind: TERNARY, Raw: "?:", Value: nil})
	node.Children = append(node.Children, condition, trueExpr, falseExpr)

	return node, nil
}

func (p *Parser) parseClause() (*ASTNode, error) {
	token := p.next()
	if token.Kind != CLAUSE {
		return nil, fmt.Errorf("expected %v token, got %v", CLAUSE, token)
	}

	// log.Printf("parseClause\n")

	expr, err := p.parseExpression(0)
	if err != nil {
		return nil, err
	}

	if err := p.expectToken(CLAUSE_CLOSE); err != nil {
		return nil, err
	}

	node := newASTNode(token)
	node.Children = append(node.Children, expr)

	return node, nil
}

func (p *Parser) parseClauseOrArray() (*ASTNode, error) {
	// 开括号
	if err := p.expectToken(CLAUSE); err != nil {
		return nil, err
	}

	token := p.next() // consume '('

	node := newASTNode(token)
	array := newASTNode(&ExpressionToken{Kind: ARRAY})
	isArray := false

	for {
		if p.peek() == nil {
			return nil, fmt.Errorf("unexpected end of tokens")
		}

		if p.peek().Kind == CLAUSE_CLOSE {
			p.next() // consume ')'
			break
		}

		// 判断是否为分隔符，如果是，则继续解析下一个元素
		if p.peek().Kind == SEPARATOR {
			p.next() // consume ','
			if !isArray {
				array.Children = append(array.Children, node)
			}
			isArray = true

			continue
		}

		element, err := p.parsePrimaryExpression()
		if err != nil {
			return nil, err
		}
		if isArray {
			array.Children = append(array.Children, element)
		} else {
			node.Children = append(node.Children, element)
		}
	}

	if isArray {
		return array, nil
	}

	return node, nil
}

func (p *Parser) parseToken(expected TokenKind) (*ASTNode, error) {
	token := p.next()
	if token.Kind != expected {
		return nil, fmt.Errorf("expected %v token, got %v", expected, token)
	}

	// log.Printf("parseToken expected %s token: %s, start %d end %d\n", expected, token.Raw, token.Start, token.End)

	return newASTNode(token), nil
}

func (p *Parser) expectToken(expected TokenKind) error {
	token := p.peek()
	if token == nil || token.Kind != expected {
		return fmt.Errorf("expected %v token, got %v", expected, token)
	}
	p.next() // Consume the token
	return nil
}

func (p *Parser) peek() *ExpressionToken {
	if p.pos >= len(p.tokens) {
		return nil
	}
	return &p.tokens[p.pos]
}

func (p *Parser) next() *ExpressionToken {
	if p.pos >= len(p.tokens) {
		return nil
	}
	token := &p.tokens[p.pos]
	p.pos++
	return token
}

func (p *Parser) getPrecedence(token *ExpressionToken) int {
	switch token.Kind {
	case MODIFIER:
		return 9000
	case COMPARATOR:
		return 8000
	case LOGICALOP:
		return 7000
	default:
		return 0
	}
}
