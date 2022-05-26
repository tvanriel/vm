package preprocessor

import (
	"strings"

	lex "github.com/bbuck/go-lexer"
)

const uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const lowercase = "abcdefghijklmnopqrstuvwxyz"
const underscore = "_"
const whitespace = " \t"
const quote = '"'

const constants = uppercase + underscore
const directive = lowercase

const (
	TextToken lex.TokenType = iota
	DirectiveToken
	StringToken
	ConstantToken
	QuoteToken
)

func LexText(l *lex.L) lex.StateFunc {
	onStartOfLine := true
	for !(l.Peek() == '#' && onStartOfLine) {
		l.Next()
		onStartOfLine = l.Peek() == '\n'
		if l.Peek() == lex.EOFRune {
			l.Next()
			l.Emit(TextToken)
			return nil
		}
	}
	l.Emit(TextToken)
	return LexDirective
}

func LexDirective(l *lex.L) lex.StateFunc {
	l.Next()
	l.Take(directive)
	l.Emit(DirectiveToken)
	SkipWhitespace(l)
	if l.Peek() == quote {
		return LexString
	}
	if l.Peek() == '\n' {
		return LexText
	}
	if strings.ContainsRune(constants, l.Peek()) {
		return LexConstant
	}
	return LexText
}

func LexString(l *lex.L) lex.StateFunc {
	l.Next() // "
	l.Ignore()
	for l.Peek() != quote {
		l.Next()
	}
	l.Emit(StringToken)
	l.Next() // "
	l.Ignore()
	return LexText
}

func LexConstant(l *lex.L) lex.StateFunc {
	l.Take(constants)
	l.Emit(ConstantToken)
	return LexText
}

func SkipWhitespace(l *lex.L) {
	l.Take(whitespace)
	l.Ignore()
}

func Tokenize(source string) []*lex.Token {
	tokens := []*lex.Token{}
	l := lex.New(source, LexText)
	l.StartSync()
	for {
		token, done := l.NextToken()
		if done {
			return tokens
		}
		tokens = append(tokens, token)
	}
}
