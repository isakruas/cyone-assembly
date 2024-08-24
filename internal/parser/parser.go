package parser

import (
	"cyone/internal/ast"
	"cyone/internal/token"
	"fmt"
)

// Parser is responsible for parsing the list of tokens into an AST
type Parser struct {
	tokens  []token.Token
	current int
}

// NewParser creates a new Parser instance
func NewParser(tokens []token.Token) *Parser {
	return &Parser{
		tokens:  tokens,
		current: 0,
	}
}

// Parse initiates parsing and returns the constructed AST
func (p *Parser) Parse() (*ast.Program, error) {
	var program ast.Program
	for p.current < len(p.tokens) {
		currentToken, err := p.peek()
		if err != nil {
			return nil, err
		}

		switch currentToken.Type {
		case token.START:
			startBlock, err := p.parseStartBlock()
			if err != nil {
				return nil, fmt.Errorf("failed to parse start block: %v", err)
			}
			program.Start = startBlock
		case token.LOC:
			variableDeclaration, err := p.parseVariableDeclaration()
			if err != nil {
				return nil, fmt.Errorf("failed to parse variable declaration: %v", err)
			}
			program.Variables = append(program.Variables, variableDeclaration)
		case token.BLOCK:
			block, err := p.parseBlock()
			if err != nil {
				return nil, fmt.Errorf("failed to parse block: %v", err)
			}
			program.Blocks = append(program.Blocks, block)
		default:
			return nil, fmt.Errorf("unexpected token: %s", currentToken.Literal)
		}
	}

	return &program, nil
}

// parseVariableDeclaration parses a variable declaration
func (p *Parser) parseVariableDeclaration() (*ast.VariableDeclaration, error) {
	if _, err := p.expect(token.LOC); err != nil {
		return nil, err
	}
	nameToken, err := p.expect(token.IDENTIFIER)
	if err != nil {
		return nil, err
	}
	name := nameToken.Literal
	if _, err := p.expect(token.AT); err != nil {
		return nil, err
	}
	addressToken, err := p.expect(token.HEXNUMBER)
	if err != nil {
		return nil, err
	}
	address := addressToken.Literal
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, err
	}

	return &ast.VariableDeclaration{
		Name:    name,
		Address: address,
	}, nil
}

// parseStartBlock parses the 'start' block
func (p *Parser) parseStartBlock() (*ast.StartBlock, error) {
	if _, err := p.expect(token.START); err != nil {
		return nil, err
	}
	if _, err := p.expect(token.AT); err != nil {
		return nil, err
	}
	addressToken, err := p.expect(token.HEXNUMBER)
	if err != nil {
		return nil, err
	}
	address := addressToken.Literal
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, fmt.Errorf("expected semicolon after start block address")
	}

	return &ast.StartBlock{
		Address: address,
	}, nil
}

// parseBlock parses a 'block'
func (p *Parser) parseBlock() (*ast.Block, error) {
	if _, err := p.expect(token.BLOCK); err != nil {
		return nil, err
	}
	addressToken, err := p.expect(token.HEXNUMBER)
	if err != nil {
		return nil, err
	}
	address := addressToken.Literal
	if _, err := p.expect(token.LBRACE); err != nil {
		return nil, err
	}
	blockContent, err := p.parseBlockContent()
	if err != nil {
		return nil, fmt.Errorf("failed to parse block content: %v", err)
	}
	blockContent.Address = address

	return blockContent, nil
}

// parseBlockContent parses statements within a block until it encounters a closing brace
func (p *Parser) parseBlockContent() (*ast.Block, error) {
	var statements []ast.Statement
	for {
		currentToken, err := p.peek()
		if err != nil {
			return nil, err
		}
		if currentToken.Type == token.RBRACE {
			break
		}
		statement, err := p.parseStatement()
		if err != nil {
			return nil, fmt.Errorf("failed to parse statement: %v", err)
		}
		statements = append(statements, statement)
	}
	if _, err := p.expect(token.RBRACE); err != nil {
		return nil, fmt.Errorf("expected closing brace at the end of block")
	}

	return &ast.Block{
		Statements: statements,
	}, nil
}

// parseStatement parses individual statements inside a block
func (p *Parser) parseStatement() (ast.Statement, error) {
	currentToken, err := p.peek()
	if err != nil {
		return nil, err
	}
	switch currentToken.Type {
	case token.IDENTIFIER:
		return p.parseAssignment()
	case token.IF:
		return p.parseIfStatement()
	case token.CALL:
		return p.parseCall()
	case token.GOTO:
		return p.parseGoto()
	case token.MEM:
		return p.parseMemoryAssignment()
	default:
		return nil, fmt.Errorf("unexpected token: %s", currentToken.Literal)
	}
}

// parseAssignment parses an assignment statement
func (p *Parser) parseAssignment() (*ast.Assignment, error) {
	variableToken, err := p.expect(token.IDENTIFIER)
	if err != nil {
		return nil, err
	}
	variable := variableToken.Literal
	if _, err := p.expect(token.ASSIGN); err != nil {
		return nil, err
	}
	value, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse expression in assignment: %v", err)
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, fmt.Errorf("expected semicolon after assignment")
	}
	return &ast.Assignment{
		VariableName: variable,
		Expression:   value,
	}, nil
}

// parseIfStatement parses an if statement with optional else block
func (p *Parser) parseIfStatement() (*ast.IfStatement, error) {
	if _, err := p.expect(token.IF); err != nil {
		return nil, err
	}
	if _, err := p.expect(token.LPAREN); err != nil {
		return nil, err
	}
	condition, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse condition of if statement: %v", err)
	}
	if _, err := p.expect(token.RPAREN); err != nil {
		return nil, fmt.Errorf("expected closing parenthesis after if condition")
	}
	if _, err := p.expect(token.LBRACE); err != nil {
		return nil, fmt.Errorf("expected opening brace for if block")
	}
	thenBlock, err := p.parseBlockContent()
	if err != nil {
		return nil, fmt.Errorf("failed to parse then block: %v", err)
	}

	var elseBlock *ast.Block
	nextToken, err := p.peek()
	if err != nil {
		return nil, err
	}
	if nextToken.Type == token.ELSE {
		if _, err := p.expect(token.ELSE); err != nil {
			return nil, err
		}
		if _, err := p.expect(token.LBRACE); err != nil {
			return nil, fmt.Errorf("expected opening brace for else block")
		}
		elseBlock, err = p.parseBlockContent()
		if err != nil {
			return nil, fmt.Errorf("failed to parse else block: %v", err)
		}
	}

	return &ast.IfStatement{
		ConditionExpression: condition,
		ThenBlock:           thenBlock,
		ElseBlock:           elseBlock,
	}, nil
}

// parseCall parses a function call expression, including its parameters and semicolon.
func (p *Parser) parseCall() (*ast.Call, error) {
	if _, err := p.expect(token.CALL); err != nil {
		return nil, err
	}
	funcNameToken, err := p.expect(token.IDENTIFIER)
	if err != nil {
		return nil, err
	}
	funcName := funcNameToken.Literal
	if _, err := p.expect(token.LPAREN); err != nil {
		return nil, fmt.Errorf("expected opening parenthesis after function name")
	}

	var params []ast.Expression
	for {
		nextToken, err := p.peek()
		if err != nil {
			return nil, err
		}
		if nextToken.Type == token.RPAREN {
			break
		}

		param, err := p.parseParameter()
		if err != nil {
			return nil, fmt.Errorf("failed to parse parameter: %v", err)
		}
		params = append(params, param)

		nextToken, err = p.peek()
		if err != nil {
			return nil, err
		}
		if nextToken.Type == token.COMMA {
			p.current++
		} else if nextToken.Type == token.RPAREN {
			break
		} else {
			return nil, fmt.Errorf("expected comma or closing parenthesis, got %s", nextToken.Type)
		}
	}

	if _, err := p.expect(token.RPAREN); err != nil {
		return nil, fmt.Errorf("expected closing parenthesis after function parameters")
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, fmt.Errorf("expected semicolon after function call")
	}

	return &ast.Call{
		FunctionName: funcName,
		Parameters:   params,
	}, nil
}

// parseParameter
func (p *Parser) parseParameter() (ast.ParameterValue, error) {
	currentToken, err := p.peek()
	if err != nil {
		return nil, err
	}

	switch currentToken.Type {
	case token.IDENTIFIER:
		identifierToken, err := p.expect(token.IDENTIFIER)
		if err != nil {
			return nil, err
		}
		return &ast.MemoryLocation{Address: identifierToken.Literal}, nil
	case token.HEXNUMBER:
		hexToken, err := p.expect(token.HEXNUMBER)
		if err != nil {
			return nil, err
		}
		return &ast.ByteValue{Value: hexToken.Literal}, nil
	default:
		return nil, fmt.Errorf("expected identifier or hexadecimal number, but got %v", currentToken.Type)
	}
}

// parseGoto parses a goto statement
func (p *Parser) parseGoto() (*ast.Goto, error) {
	if _, err := p.expect(token.GOTO); err != nil {
		return nil, err
	}
	addressToken, err := p.expect(token.HEXNUMBER)
	if err != nil {
		return nil, err
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, fmt.Errorf("expected semicolon after goto statement")
	}

	return &ast.Goto{Address: addressToken.Literal}, nil
}

// parseMemoryAssignment parses a memory assignment statement (mem[addr] = value)
func (p *Parser) parseMemoryAssignment() (*ast.MemoryAssignment, error) {
	if _, err := p.expect(token.MEM); err != nil {
		return nil, err
	}
	if _, err := p.expect(token.LBRACKET); err != nil {
		return nil, err
	}
	address, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse memory address: %v", err)
	}
	if _, err := p.expect(token.RBRACKET); err != nil {
		return nil, fmt.Errorf("expected closing bracket after memory address")
	}
	if _, err := p.expect(token.ASSIGN); err != nil {
		return nil, fmt.Errorf("expected equals sign after memory address")
	}
	value, err := p.parseExpression()
	if err != nil {
		return nil, fmt.Errorf("failed to parse value for memory assignment: %v", err)
	}
	if _, err := p.expect(token.SEMICOLON); err != nil {
		return nil, fmt.Errorf("expected semicolon after memory assignment")
	}

	return &ast.MemoryAssignment{
		MemoryAddress: address,
		Value:         value,
	}, nil
}

// parseExpression parses a complex binary expression
func (p *Parser) parseExpression() (ast.Expression, error) {
	leftExpression, err := p.parsePrimaryExpression()
	if err != nil {
		return nil, err
	}

	for {
		nextToken, err := p.peek()
		if err != nil {
			return nil, err
		}

		if nextToken.Type != token.MOD && nextToken.Type != token.PLUS && nextToken.Type != token.MINUS && nextToken.Type != token.EQ &&
			nextToken.Type != token.NOT_EQ && nextToken.Type != token.GT && nextToken.Type != token.LT && nextToken.Type != token.ASTERISK {
			break
		}

		operatorToken, err := p.expectAny(token.MOD, token.PLUS, token.MINUS, token.EQ, token.NOT_EQ, token.GT, token.LT, token.ASTERISK)
		if err != nil {
			return nil, fmt.Errorf("failed to parse operator: %v", err)
		}

		rightExpression, err := p.parsePrimaryExpression()
		if err != nil {
			return nil, fmt.Errorf("failed to parse right-hand expression: %v", err)
		}

		leftExpression = &ast.BinaryExpression{
			LeftExpression:  leftExpression,
			Operator:        operatorToken.Literal,
			RightExpression: rightExpression,
		}
	}

	return leftExpression, nil
}

// parsePrimaryExpression parses the most basic expressions (constants, variables, memory access)
func (p *Parser) parsePrimaryExpression() (ast.Expression, error) {
	currentToken, err := p.peek()
	if err != nil {
		return nil, err
	}

	switch currentToken.Type {
	case token.IDENTIFIER:
		p.current++
		return &ast.Variable{Name: currentToken.Literal}, nil
	case token.HEXNUMBER:
		p.current++
		return &ast.Constant{Value: currentToken.Literal}, nil
	case token.MEM:
		expr, err := p.parseMemoryAccess()
		if err != nil {
			return nil, fmt.Errorf("failed to parse memory access: %v", err)
		}
		return expr, nil
	default:
		return nil, fmt.Errorf("unexpected token: %v", currentToken.Type)
	}
}

// parseMemoryAccess parses expressions like mem[0x0005]
func (p *Parser) parseMemoryAccess() (ast.Expression, error) {
	if _, err := p.expect(token.MEM); err != nil {
		return nil, err
	}
	if _, err := p.expect(token.LBRACKET); err != nil {
		return nil, err
	}
	addressToken, err := p.expect(token.HEXNUMBER)
	if err != nil {
		return nil, fmt.Errorf("failed to parse memory address: %v", err)
	}
	if _, err := p.expect(token.RBRACKET); err != nil {
		return nil, fmt.Errorf("expected closing bracket after memory address")
	}

	return &ast.MemoryLocation{Address: addressToken.Literal}, nil
}

// expect consumes the next token if it matches the expected type
func (p *Parser) expect(expectedType token.TokenType) (token.Token, error) {
	currentToken, err := p.peek()
	if err != nil {
		return token.Token{}, err
	}

	for currentToken.Type == token.COMMENT {
		p.current++
		currentToken, err = p.peek()
		if err != nil {
			return token.Token{}, err
		}
	}

	if currentToken.Type != expectedType {
		return token.Token{}, fmt.Errorf("expected token %v but got %v", expectedType, currentToken.Type)
	}

	p.current++
	return currentToken, nil
}

// expectAny consumes the next token if it matches any of the expected types
func (p *Parser) expectAny(expectedTypes ...token.TokenType) (token.Token, error) {
	currentToken, err := p.peek()
	if err != nil {
		return token.Token{}, err
	}

	for currentToken.Type == token.COMMENT {
		p.current++
		currentToken, err = p.peek()
		if err != nil {
			return token.Token{}, err
		}
	}

	for _, expectedType := range expectedTypes {
		if currentToken.Type == expectedType {
			p.current++
			return currentToken, nil
		}
	}

	return token.Token{}, fmt.Errorf("expected one of %v but got %v", expectedTypes, currentToken.Type)
}

// peek returns the current token without consuming it and skips comment tokens
func (p *Parser) peek() (token.Token, error) {
	if p.current >= len(p.tokens) {
		return token.Token{}, fmt.Errorf("reached end of input")
	}

	for p.current < len(p.tokens) && p.tokens[p.current].Type == token.COMMENT {
		p.current++
	}

	if p.current >= len(p.tokens) {
		return token.Token{}, fmt.Errorf("reached end of input")
	}

	return p.tokens[p.current], nil
}
