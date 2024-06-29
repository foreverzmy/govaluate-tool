package parser

import "encoding/json"

/*
Represents all valid types of tokens that a token can be.
*/
type TokenKind int

const (
	UNKNOWN TokenKind = iota

	PREFIX
	NUMERIC
	BOOLEAN
	STRING
	PATTERN
	TIME
	VARIABLE
	FUNCTION
	SEPARATOR
	ACCESSOR // 访问器，包括点操作符（.）用于访问对象的属性，或中括号操作符（[]）用于访问数组或列表的元素

	COMPARATOR
	LOGICALOP
	MODIFIER // 修饰符，如 递增（++）、递减（--）

	CLAUSE
	CLAUSE_CLOSE

	TERNARY // 三元运算符
	ARRAY   // 新增数组类型
)

/*
GetTokenKindString returns a string that describes the given TokenKind.
e.g., when passed the NUMERIC TokenKind, this returns the string "NUMERIC".
*/
func (kind TokenKind) String() string {

	switch kind {

	case PREFIX:
		return "PREFIX"
	case NUMERIC:
		return "NUMERIC"
	case BOOLEAN:
		return "BOOLEAN"
	case STRING:
		return "STRING"
	case PATTERN:
		return "PATTERN"
	case TIME:
		return "TIME"
	case VARIABLE:
		return "VARIABLE"
	case FUNCTION:
		return "FUNCTION"
	case SEPARATOR:
		return "SEPARATOR"
	case COMPARATOR:
		return "COMPARATOR"
	case LOGICALOP:
		return "LOGICALOP"
	case MODIFIER:
		return "MODIFIER"
	case CLAUSE:
		return "CLAUSE"
	case CLAUSE_CLOSE:
		return "CLAUSE_CLOSE"
	case TERNARY:
		return "TERNARY"
	case ACCESSOR:
		return "ACCESSOR"
	case ARRAY:
		return "ARRAY"
	}

	return "UNKNOWN"
}

func (k TokenKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(k.String())
}
