package parser

/*
Represents a single parsed token.
*/
type ExpressionToken struct {
	Kind  TokenKind
	Value interface{}
	Raw   string
	Start int
	End   int
}
