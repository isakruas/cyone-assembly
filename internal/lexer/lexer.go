package lexer

import (
	"cyone/internal/token"
	"cyone/internal/utils"
	"fmt"
	"unicode"
)

// Lexer represents the lexical analyzer
type Lexer struct {
	input       string
	currentPos  int
	nextPos     int
	currentChar byte
}

// NewLexer initializes a new Lexer
func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.advanceChar()
	return lexer
}

// advanceChar reads the next character and advances the lexer's position
func (l *Lexer) advanceChar() {
	if l.nextPos >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPos]
	}
	l.currentPos = l.nextPos
	l.nextPos++
}

// NextToken returns the next token from the input
func (l *Lexer) NextToken() (token.Token, error) {
	var tok token.Token
	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.peekChar() == '=' {
			tok = l.createTwoCharToken(token.EQ)
		} else {
			tok = utils.NewToken(token.ASSIGN, l.currentChar)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.createTwoCharToken(token.NOT_EQ)
		} else {
			return utils.NewToken(token.ILLEGAL, l.currentChar), fmt.Errorf("unexpected character: '%c' at position %d", l.currentChar, l.currentPos)
		}
	case '/':
		if l.peekChar() == '/' {
			tok.Type = token.COMMENT
			tok.Literal = l.readComment()
		} else {
			tok = utils.NewToken(token.SLASH, l.currentChar)
		}
	default:
		if tokenType, ok := token.SingleCharTokens[l.currentChar]; ok {
			tok = utils.NewToken(tokenType, l.currentChar)
		} else if utils.IsLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = utils.LookupIdent(tok.Literal)
			return tok, nil
		} else if l.currentChar == '0' && l.peekChar() == 'x' {
			tok.Type = token.HEXNUMBER
			tok.Literal = l.readHexNumber()
			return tok, nil
		} else if l.currentChar == 0 {
			tok.Literal = ""
			tok.Type = token.EOF
			return tok, nil
		} else {
			return utils.NewToken(token.ILLEGAL, l.currentChar), fmt.Errorf("unexpected character: '%c' at position %d", l.currentChar, l.currentPos)
		}
	}

	l.advanceChar()
	return tok, nil
}

// createTwoCharToken returns a token made up of two characters
func (l *Lexer) createTwoCharToken(tokenType token.TokenType) token.Token {
	previousChar := l.currentChar
	l.advanceChar()
	return token.Token{Type: tokenType, Literal: string(previousChar) + string(l.currentChar)}
}

// readIdentifier reads an identifier and advances lexer's position
func (l *Lexer) readIdentifier() string {
	startPos := l.currentPos
	for utils.IsLetter(l.currentChar) {
		l.advanceChar()
	}
	return l.input[startPos:l.currentPos]
}

// readHexNumber reads a hexadecimal number and advances lexer's position
func (l *Lexer) readHexNumber() string {
	startPos := l.currentPos
	l.advanceChar() // '0'
	l.advanceChar() // 'x'
	for utils.IsHexDigit(l.currentChar) {
		l.advanceChar()
	}
	return l.input[startPos:l.currentPos]
}

// readComment reads a comment until the end of line
func (l *Lexer) readComment() string {
	startPos := l.currentPos
	for l.currentChar != '\n' && l.currentChar != 0 {
		l.advanceChar()
	}
	return l.input[startPos:l.currentPos]
}

// skipWhitespace skips whitespace characters
func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(rune(l.currentChar)) {
		l.advanceChar()
	}
}

// peekChar returns the next character without advancing position
func (l *Lexer) peekChar() byte {
	if l.nextPos >= len(l.input) {
		return 0
	}
	return l.input[l.nextPos]
}
