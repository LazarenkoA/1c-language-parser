package main

import (
	"fmt"

	"github.com/pkg/errors"
)

type ASTOneS struct {
	code       string
	err        error
	pos        Position
	currentLit string
}

type (
	statement interface {
	}
	statements []statement
)

const EOF = -1 // end of file

func NewAST(code string) *ASTOneS {
	return &ASTOneS{
		code: code,
	}
}

func (ast *ASTOneS) Parse() error {
	if len(ast.code) == 0 {
		return nil
	}

	yyParse(ast)
	if ast.err != nil {
		errors.Wrap(ast.err, "parse error")
	}
	return ast.err
}

func (ast *ASTOneS) Lex(lval *yySymType) int {
	if len(ast.code) == 0 {
		return EOF
	}

	token, lit, err := lval.token.Next(ast.code)
	if err != nil {
		ast.err = err
	}
	if token == EOF {
		return EOF
	}

	ast.currentLit = lit
	ast.pos = lval.token.GetPosition()
	ast.pos.Column -= len([]rune(lit)) + 1

	return token
}

func (ast *ASTOneS) Error(s string) {
	ast.err = fmt.Errorf("%s. line: %d, column: %d (unexpected literal: %q)", s, ast.pos.Line, ast.pos.Column, ast.currentLit)
}
