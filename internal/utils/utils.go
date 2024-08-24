package utils

import (
	pkg_token "cyone/internal/token"
	"unicode"
)

// IsLetter check if a character is a letter
func IsLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// IsHexDigit check if a character is a hexadecimal digit
func IsHexDigit(ch byte) bool {
	return IsDigit(ch) || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

// IsDigit check if a character is a digit
func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// LookupIdent checks if an identifier is a keyword
func LookupIdent(ident string) pkg_token.TokenType {
	if tok, ok := pkg_token.Keywords[ident]; ok {
		return tok
	}
	return pkg_token.IDENTIFIER
}

// NewToken creates a new token of a given type from a character
func NewToken(tokenType pkg_token.TokenType, ch byte) pkg_token.Token {
	return pkg_token.Token{Type: tokenType, Literal: string(ch)}
}
