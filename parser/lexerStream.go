package parser

type lexerStream struct {
	source   []rune
	position int
	length   int
}

func newLexerStream(source string) *lexerStream {
	var ret *lexerStream
	var runes []rune

	for _, character := range source {
		runes = append(runes, character)
	}

	ret = new(lexerStream)
	ret.source = runes
	ret.length = len(runes)
	return ret
}

func (s *lexerStream) readCharacter() rune {
	character := s.source[s.position]
	s.position += 1
	return character
}

func (s *lexerStream) rewind(amount int) {
	s.position -= amount
}

func (s lexerStream) canRead() bool {
	return s.position < s.length
}
