package main

import (
	"bufio"
	"cyone/internal/bytecode"
	"cyone/internal/lexer"
	"cyone/internal/parser"
	"cyone/internal/token"
	"flag"
	"fmt"
	"os"
	"strings"
)

// GOOS represents the operating system on which the program is running.
var GOOS string

// GOARCH represents the architecture of the operating system on which the program is running.
var GOARCH string

// CODEVERSION represents the version of the code.
var CODEVERSION string

// CODEBUILDDATE represents the date when the program was built.
var CODEBUILDDATE string

// CODEBUILDREVISION represents the revision of the code build.
var CODEBUILDREVISION string

func main() {

	// Defer a function to recover from a panic and handle errors gracefully.
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("an error occurred:", r)
		}
	}()

	// Define and parse command-line flags
	var info, license bool
	flag.BoolVar(&info, "info", false, "Display program compilation and version information")
	flag.BoolVar(&license, "license", false, "Display program license information")
	filename := flag.String("file", "", "Path to the file to be parsed")
	flag.Parse()

	if info {
		fmt.Printf("Version: %s\n", CODEVERSION)
		fmt.Printf("Operating System: %s\n", GOOS)
		fmt.Printf("System Architecture: %s\n", GOARCH)
		fmt.Printf("Build Date: %s\n", CODEBUILDDATE)
		fmt.Printf("Build Revision: %s\n", CODEBUILDREVISION)
		return
	}

	if license {
		fmt.Println("Copyright 2024 Isak Ruas")
		fmt.Println("Licensed under the Apache License, Version 2.0 (the 'License');")
		fmt.Println("you may not use this file except in compliance with the License.")
		fmt.Println("You may obtain a copy of the License at")
		fmt.Println("    http://www.apache.org/licenses/LICENSE-2.0")
		fmt.Println("Unless required by applicable law or agreed to in writing, software")
		fmt.Println("distributed under the License is distributed on an 'AS IS' BASIS,")
		fmt.Println("WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.")
		fmt.Println("See the License for the specific language governing permissions and")
		fmt.Println("limitations under the License.")
		return
	}

	// Ensure the filename argument is provided
	if *filename == "" {
		fmt.Println("Usage: cyone -file <filename>")
		return
	}

	// Open the file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error opening the file:", err)
		return
	}
	defer file.Close()

	// Read the file content
	var code strings.Builder
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		code.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading the file:", err)
		return
	}

	// Initialize the lexer
	lex := lexer.NewLexer(code.String())
	var tokens []token.Token
	for {
		tok, err := lex.NextToken()
		if err != nil {
			fmt.Println("Error tokenizing input:", err)
			return
		}
		if tok.Type == token.EOF {
			break
		}
		tokens = append(tokens, tok)
		// fmt.Printf("{Type: %s, Literal: \"%s\"}\n", tok.Type, tok.Literal)
	}

	// Initialize the parser
	par := parser.NewParser(tokens)
	program, err := par.Parse()
	if err != nil {
		fmt.Println("Error parsing tokens:", err)
		return
	}

	// Print the AST
	// fmt.Println(program.String())

	bytecodes, err := bytecode.GenerateBytecode(program)
	if err != nil {
		fmt.Println("Error generating bytecode:", err)
		return
	}

	for _, line := range bytecode.GenerateIntelHex(bytecodes) {
		fmt.Println(line)
	}

}
