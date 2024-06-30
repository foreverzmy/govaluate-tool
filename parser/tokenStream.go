package parser

type tokenStream struct {
	tokens      []ExpressionToken
	index       int
	tokenLength int
}

func newTokenStream(tokens []ExpressionToken) *tokenStream {
	ret := new(tokenStream)
	ret.tokens = tokens
	ret.tokenLength = len(tokens)
	return ret
}

func (ts *tokenStream) next() ExpressionToken {

	token := ts.tokens[ts.index]

	ts.index += 1
	return token
}

func (ts tokenStream) hasNext() bool {
	return ts.index < ts.tokenLength
}
