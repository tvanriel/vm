package lexer

import (
	"fmt"
	"strings"

	lex "github.com/bbuck/go-lexer"
)

const (
	WordToken lex.TokenType = iota
	NumberToken
	Whitespace
	LabelToken
	InstructionToken
	DirectiveToken
	ConstantToken

	HexadecimalToken
	DecimalToken
	OctalToken
	BinaryToken

	CommentToken
	DollarSign
	ColonSign
	CommaSign
	HashSign

	PlusSign
	MinusSign
	AsteriskSign
	SlashSign
	EqualsSign

	OpGreaterThan
	OpLessThan

	PeriodSign
	ParenthesisOpen
	ParenthesisClose
)

// Helpers
const lowercase = "abcdefghijklmnopqrstuvwxyz"
const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const operators = "+-/*^|&"
const whitespace = "\t "

// Number helpers
const decimal = "0123456789"
const hexadecimal = "0123456789abcdefABCDEF"
const octal = "01234567"
const binary = "01"
const hexIndicator = 'x'
const binIndicator = 'b'
const octalIndicator = 'o'

// Punctuation
const colon = ':'
const semicolon = ';'
const newLine = '\n'
const space = ' '
const period = '.'
const comma = ','
const equals = '='
const plus = '+'
const minus = '-'
const underscore = '_'
const asterisk = '*'
const caret = '^'
const pipe = '|'
const ampersand = '&'
const lessthan = '<'
const greaterthan = '>'
const slash = '/'
const parenthesisOpen = '('
const pathentesisClose = ')'
const dollar = '$'
const zero = '0'

// Wordsets
const instruction = lowercase
const constant = uppercase + string(underscore)
const label = lowercase + string(underscore)
const directive = lowercase

// Lex any valid line
func LexLine(l *lex.L) lex.StateFunc {
	l.Take("\n")
	l.Ignore()
	if l.Peek() == lex.EOFRune {
		return nil
	}
	if l.Peek() == semicolon {
		return LexComment
	}
	if strings.ContainsRune(label, l.Peek()) {
		return LexLabel
	}

	if l.Peek() == period {
		return LexDirective
	}

	if strings.ContainsRune(constant, l.Peek()) {
		return LexConstAssignment
	}

	if strings.ContainsRune(whitespace, l.Peek()) {
		SkipWhitespace(l)
		return LexInstruction
	}
	if l.Peek() == lex.EOFRune {
		return nil
	}

	fmt.Errorf("unknown char while reading line: %s\n", string(l.Peek()))
	return nil
}

func LexConstAssignment(l *lex.L) lex.StateFunc {
	LexConstant(l)
	SkipWhitespace(l)
	if l.Peek() == equals {
		l.Next()
		l.Emit(EqualsSign)
		SkipWhitespace(l)
	}
	LexExpression(l)

	if l.Peek() == semicolon {
		return LexComment
	}
	return LexLine
}

// Lex an instruction, skipping over the indented whitespace.
func LexInstruction(l *lex.L) lex.StateFunc {
	if l.Peek() == semicolon {
		return LexComment
	}
	l.Take(instruction)
	if l.Current() == "" {
		l.Ignore()
		return LexLabel
	}
	l.Emit(InstructionToken)
	SkipWhitespace(l)

	// If we get a newline already, we didn't have any arguments.
	if l.Peek() == newLine {
		return LexLine
	}

	if l.Peek() == semicolon {
		return LexComment
	}

	LexArgumentList(l)
	l.Take("\n")
	l.Ignore()
	return LexLine
}

func LexExpression(l *lex.L) {
	peeked := l.Peek()
	if strings.ContainsRune(label, peeked) {
		l.Take(label)
		l.Emit(WordToken)
		SkipWhitespace(l)
		return
	}
	if strings.ContainsRune(decimal, peeked) {
		LexNumber(l)
		SkipWhitespace(l)
		return
	}
	if strings.ContainsRune(constant, peeked) {
		LexConstant(l)
		SkipWhitespace(l)
		return
	}
	if l.Peek() == parenthesisOpen {
		l.Next()
		l.Emit(ParenthesisOpen)
		SkipWhitespace(l)
	}
	if strings.ContainsRune(constant, peeked) {
		LexConstant(l)
		SkipWhitespace(l)
	}

	if strings.ContainsRune(decimal, peeked) {
		LexNumber(l)
		SkipWhitespace(l)
	}

	if strings.ContainsRune(operators, peeked) {
		LexOperator(l)

		SkipWhitespace(l)
		if strings.ContainsRune(constant, peeked) {
			LexConstant(l)
		}

		if strings.ContainsRune(decimal, peeked) {
			LexNumber(l)
		}
		SkipWhitespace(l)

	}

	if l.Peek() == pathentesisClose {
		l.Next()
		l.Emit(ParenthesisClose)
	}
}

func LexOperator(l *lex.L) {
	if l.Peek() == plus {
		l.Next()
		l.Emit(PlusSign)
	}
	if l.Peek() == minus {
		l.Next()
		l.Emit(MinusSign)
	}
	if l.Peek() == asterisk {
		l.Next()
		l.Emit(AsteriskSign)
	}
	if l.Peek() == lessthan {
		l.Next()
		l.Next()
		l.Emit(OpLessThan)
	}
	if l.Peek() == greaterthan {
		l.Next()
		l.Next()
		l.Emit(OpGreaterThan)
	}
}

func LexNumber(l *lex.L) {
	if l.Peek() == minus {
		l.Next()
		l.Emit(MinusSign)
	}

	if l.Peek() == '0' {
		l.Next()
		if l.Peek() == hexIndicator {
			l.Rewind()
			LexHexadecimalNumber(l)
			return
		}

		if l.Peek() == binIndicator {
			l.Rewind()
			LexBinaryNumber(l)
			return
		}

		if l.Peek() == octalIndicator {
			l.Rewind()
			LexOctalNumber(l)
			return
		}

		l.Rewind()
		LexOctalNumberWithoutPrefix(l)
		return
	}

	LexDecimalNumber(l)
}

func LexConstant(l *lex.L) {
	l.Take(constant)
	l.Emit(ConstantToken)
}

func LexHexadecimalNumber(l *lex.L) {
	l.Next() // 0
	l.Next() // x
	l.Take(hexadecimal)
	l.Emit(HexadecimalToken)
}

func LexBinaryNumber(l *lex.L) {
	l.Next() // 0
	l.Next() // b
	l.Take(binary)
	l.Emit(BinaryToken)
}
func LexOctalNumber(l *lex.L) {
	l.Next() // 0
	l.Next() // o
	l.Take(octal)
	l.Emit(OctalToken)
}
func LexOctalNumberWithoutPrefix(l *lex.L) {
	l.Next() // 0
	l.Take(octal)
	l.Emit(OctalToken)
}
func LexDecimalNumber(l *lex.L) {
	l.Take(decimal)
	l.Emit(DecimalToken)
}

func LexArgumentList(l *lex.L) {
	LexExpression(l)
	SkipWhitespace(l)
	for l.Peek() == comma {
		l.Next()
		l.Emit(CommaSign)
		SkipWhitespace(l)
		if l.Peek() == semicolon {
			LexComment(l)
		}
		LexExpression(l)
		SkipWhitespace(l)
		if l.Peek() == semicolon {
			LexComment(l)
		}
	}
}

func LexDirective(l *lex.L) lex.StateFunc {
	l.Next()
	l.Emit(PeriodSign)
	l.Take(directive)
	l.Emit(DirectiveToken)
	SkipWhitespace(l)

	if l.Peek() == semicolon {
		return LexComment
	}
	if l.Peek() == newLine {
		return LexLine
	}
	LexArgumentList(l)
	l.Take("\n")
	l.Ignore()
	return LexLine
}

func SkipWhitespace(l *lex.L) {
	l.Take(whitespace)
	l.Ignore()

}

func LexLabel(l *lex.L) lex.StateFunc {
	l.Take(label)
	l.Emit(LabelToken)
	SkipWhitespace(l)
	l.Next() // Read colon
	l.Emit(ColonSign)
	SkipWhitespace(l)

	l.Take("\n")
	l.Ignore()
	return LexLine
}

func LexComment(l *lex.L) lex.StateFunc {
	// Take to the end of the line and emit the comment
	for l.Peek() != newLine {
		l.Next()
	}
	l.Emit(CommentToken)
	l.Next() // Read over the \n
	l.Ignore()
	return LexLine
}

func Tokenize(source string) []*lex.Token {
	tokens := []*lex.Token{}
	l := lex.New(source, LexLine)
	l.StartSync()
	for {
		token, done := l.NextToken()
		if done {
			return tokens
		}
		fmt.Println(token.Value)
		tokens = append(tokens, token)
	}
}
