package token

// TokenType represents the type of token
type TokenType string

// Token represents a single token in the source code
type Token struct {
	Type    TokenType
	Literal string
}

// Define all different types of tokens
const (
	ILLEGAL TokenType = "ILLEGAL" // Value for illegal tokens
	EOF     TokenType = "EOF"     // Value for end-of-file token

	// Identifiers + literals
	IDENTIFIER TokenType = "IDENTIFIER" // Value for identifiers
	HEXNUMBER  TokenType = "HEXNUMBER"  // Value for hexadecimal numbers

	// Keywords
	LOC   TokenType = "LOC"   // Value for the 'LOC' keyword
	AT    TokenType = "AT"    // Value for the 'AT' keyword
	START TokenType = "START" // Value for the 'START' keyword
	BLOCK TokenType = "BLOCK" // Value for the 'BLOCK' keyword
	MEM   TokenType = "MEM"   // Value for the 'MEM' keyword
	IF    TokenType = "IF"    // Value for the 'IF' keyword
	ELSE  TokenType = "ELSE"  // Value for the 'ELSE' keyword
	GOTO  TokenType = "GOTO"  // Value for the 'GOTO' keyword
	CALL  TokenType = "CALL"  // Value for the 'CALL' keyword
	TO    TokenType = "TO"    // Value for the 'TO' keyword

	// Operators
	ASSIGN   TokenType = "ASSIGN"   // Value for the '=' operator
	PLUS     TokenType = "PLUS"     // Value for the '+' operator
	MINUS    TokenType = "MINUS"    // Value for the '-' operator
	ASTERISK TokenType = "ASTERISK" // Value for the '*' operator
	SLASH    TokenType = "SLASH"    // Value for the '/' operator
	EQ       TokenType = "EQ"       // Value for the '==' operator
	NOT_EQ   TokenType = "NOT_EQ"   // Value for the '!=' operator
	GT       TokenType = "GT"       // Value for the '>' operator
	LT       TokenType = "LT"       // Value for the '<' operator
	MOD      TokenType = "MOD"      // Value for the '<' operator

	// Delimiters
	COMMA     TokenType = "COMMA"     // Value for the ',' delimiter
	SEMICOLON TokenType = "SEMICOLON" // Value for the ';' delimiter
	COLON     TokenType = "COLON"     // Value for the ':' delimiter
	LPAREN    TokenType = "LPAREN"    // Value for the '(' delimiter
	RPAREN    TokenType = "RPAREN"    // Value for the ')' delimiter
	LBRACE    TokenType = "LBRACE"    // Value for the '{' delimiter
	RBRACE    TokenType = "RBRACE"    // Value for the '}' delimiter
	LBRACKET  TokenType = "LBRACKET"  // Value for the '[' delimiter
	RBRACKET  TokenType = "RBRACKET"  // Value for the ']' delimiter

	// Comments
	COMMENT TokenType = "COMMENT" // Value for comments
)

// Map of reserved keywords
var Keywords = map[string]TokenType{
	"loc":   LOC,
	"at":    AT,
	"start": START,
	"block": BLOCK,
	"mem":   MEM,
	"if":    IF,
	"else":  ELSE,
	"goto":  GOTO,
	"call":  CALL,
	"to":    TO,
}

// Map of single character tokens for quick lookup
var SingleCharTokens = map[byte]TokenType{
	'+': PLUS,
	'-': MINUS,
	'*': ASTERISK,
	'/': SLASH,
	'=': ASSIGN,
	';': SEMICOLON,
	':': COLON,
	',': COMMA,
	'(': LPAREN,
	')': RPAREN,
	'{': LBRACE,
	'}': RBRACE,
	'[': LBRACKET,
	']': RBRACKET,
	'>': GT,
	'<': LT,
	'%': MOD,
}

var MultiCharTokens = map[string]TokenType{
	"==": EQ,
	"!=": NOT_EQ,
}

// Define opcodes for each TokenType
const (
	OP_ILLEGAL    byte = 0x00
	OP_EOF        byte = 0x01
	OP_IDENTIFIER byte = 0x02
	OP_HEXNUMBER  byte = 0x03
	OP_LOC        byte = 0x04
	OP_AT         byte = 0x05
	OP_START      byte = 0x06
	OP_BLOCK      byte = 0x07
	OP_MEM        byte = 0x08
	OP_IF         byte = 0x09
	OP_ELSE       byte = 0x0A
	OP_GOTO       byte = 0x0B
	OP_CALL       byte = 0x0C
	OP_TO         byte = 0x0D
	OP_ASSIGN     byte = 0x0E
	OP_PLUS       byte = 0x0F
	OP_MINUS      byte = 0x10
	OP_ASTERISK   byte = 0x11
	OP_SLASH      byte = 0x12
	OP_EQ         byte = 0x13
	OP_NOT_EQ     byte = 0x14
	OP_GT         byte = 0x15
	OP_LT         byte = 0x16
	OP_MOD        byte = 0x17
	OP_COMMA      byte = 0x18
	OP_SEMICOLON  byte = 0x19
	OP_COLON      byte = 0x1A
	OP_LPAREN     byte = 0x1B
	OP_RPAREN     byte = 0x1C
	OP_LBRACE     byte = 0x1D
	OP_RBRACE     byte = 0x1E
	OP_LBRACKET   byte = 0x1F
	OP_RBRACKET   byte = 0x20
	OP_COMMENT    byte = 0x21
)

// Map of function names to their respective opcodes
var FunctionOpCodes = map[string]byte{
	"DRAW_LINE":      0x00,
	"DRAW_CIRCLE":    0x01,
	"SET_COLOR":      0x02,
	"DRAW_RECTANGLE": 0x03,
}

// Map of TokenType to opcode
var TokenOpcodes = map[TokenType]byte{
	ILLEGAL:    OP_ILLEGAL,
	EOF:        OP_EOF,
	IDENTIFIER: OP_IDENTIFIER,
	HEXNUMBER:  OP_HEXNUMBER,
	LOC:        OP_LOC,
	AT:         OP_AT,
	START:      OP_START,
	BLOCK:      OP_BLOCK,
	MEM:        OP_MEM,
	IF:         OP_IF,
	ELSE:       OP_ELSE,
	GOTO:       OP_GOTO,
	CALL:       OP_CALL,
	TO:         OP_TO,
	ASSIGN:     OP_ASSIGN,
	PLUS:       OP_PLUS,
	MINUS:      OP_MINUS,
	ASTERISK:   OP_ASTERISK,
	SLASH:      OP_SLASH,
	EQ:         OP_EQ,
	NOT_EQ:     OP_NOT_EQ,
	GT:         OP_GT,
	LT:         OP_LT,
	MOD:        OP_MOD,
	COMMA:      OP_COMMA,
	SEMICOLON:  OP_SEMICOLON,
	COLON:      OP_COLON,
	LPAREN:     OP_LPAREN,
	RPAREN:     OP_RPAREN,
	LBRACE:     OP_LBRACE,
	RBRACE:     OP_RBRACE,
	LBRACKET:   OP_LBRACKET,
	RBRACKET:   OP_RBRACKET,
	COMMENT:    OP_COMMENT,
}
