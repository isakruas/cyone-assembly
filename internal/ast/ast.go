package ast

import (
	"encoding/json"
	"fmt"
)

// Node defines the base interface for all AST nodes
type Node interface{}

// Program represents the entire parsed program
type Program struct {
	Variables []*VariableDeclaration //`json:"variables"`
	Start     *StartBlock            //`json:"start_block"`
	Blocks    []*Block               //`json:"blocks"`
}

func (p *Program) String() (string, error) {
	bytes, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to marshal program: %v", err)
	}
	return string(bytes), nil
}

// VariableDeclaration represents the declaration of a variable
type VariableDeclaration struct {
	Name    string //`json:"name"`
	Address string //`json:"address"`
}

// StartBlock represents the 'start' block of the program
type StartBlock struct {
	Address string //`json:"address"`
}

// Block represents a code block
type Block struct {
	Address    string      //`json:"address,omitempty"`
	Statements []Statement //`json:"statements"`
}

// Statement represents a statement within a block
type Statement interface{}

// Assignment represents an assignment statement (e.g., x = 0x0A)
type Assignment struct {
	VariableName string     //`json:"variable_name"`
	Expression   Expression //`json:"expression"`
}

// IfStatement represents an 'if' statement with optional 'else' block
type IfStatement struct {
	ConditionExpression Expression //`json:"condition_expression"`
	ThenBlock           *Block     //`json:"then_block,omitempty"`
	ElseBlock           *Block     //`json:"else_block,omitempty"`
}

// Call represents a call statement (e.g., call fn(0x0200, 0x0200, 0x0200); )
type Call struct {
	FunctionName string       //`json:"function_expression"`
	Parameters   []Expression //`json:"parameters"`
}

// Goto represents a goto statement (e.g., goto 0x0200)
type Goto struct {
	Address string //`json:"goto_address"`
}

// Expression represents an expression, which can be a constant or a variable
type Expression interface{}

// ParameterValue represents a value for a parameter
type ParameterValue interface{}

// ByteValue represents a byte literal value
type ByteValue struct {
	Value string //`json:"value"`
}

// MemoryLocation represents a location in memory (e.g., a variable or address)
type MemoryLocation struct {
	Address string //`json:"address"`
}

// MemoryAssignment represents a direct memory operation (e.g., mem[0x0004] = resultado)
type MemoryAssignment struct {
	MemoryAddress Expression //`json:"memory_address"`
	Value         Expression //`json:"value"`
}

// BinaryExpression represents a binary expression (e.g., x + y)
type BinaryExpression struct {
	LeftExpression  Expression //`json:"left_expression"`
	Operator        string     //`json:"operator"`
	RightExpression Expression //`json:"right_expression"`
}

// Variable represents a variable expression
type Variable struct {
	Name string //`json:"name"`
}

// Constant represents a constant value expression
type Constant struct {
	Value string //`json:"value"`
}
