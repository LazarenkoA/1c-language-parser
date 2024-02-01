package ast

//go:generate goyacc  .\grammar.y

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sync/atomic"

	"github.com/pkg/errors"
)

type AstNode struct {
	ModuleStatement

	code       string
	err        error
	pos        Position
	currentLit string
	isLoop     atomic.Int32
	isTry      atomic.Int32
	isFunction bool
}

const EOF = -1 // end of file
var (
	errVariableAlreadyDefined = errors.New("variable has already been defined")
)

func NewAST(code string) *AstNode {
	return &AstNode{
		code: code,
	}
}

func (ast *AstNode) Parse() error {
	if len(ast.code) == 0 {
		return nil
	}

	yyParse(ast)
	if ast.err != nil {
		errors.Wrap(ast.err, "parse error")
	}
	return ast.err
}

func (ast *AstNode) JSON() ([]byte, error) {
	return json.Marshal(&ast.ModuleStatement)
}

func (ast *AstNode) Lex(lval *yySymType) int {
	if len(ast.code) == 0 {
		return EOF
	}

	token, err := lval.token.Next(ast.code)
	if err != nil {
		ast.err = errors.Wrap(err, "get token error")
	}
	if token == EOF {
		return EOF
	}

	ast.currentLit = lval.token.literal
	ast.pos = lval.token.GetPosition()
	ast.pos.Column -= len([]rune(lval.token.literal)) + 1

	return token
}

func (ast *AstNode) Error(s string) {
	ast.err = fmt.Errorf("%s. line: %d, column: %d (unexpected literal: %q)", s, ast.pos.Line, ast.pos.Column, ast.currentLit)
}

func checkLoopOperator(token Token, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		if ast.isLoop.Load() == 0 {
			yylex.Error(fmt.Sprintf("operator %q can only be used inside a loop", token.literal))
		}
	}
}

func checkThrowParam(token Token, param Statement, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		if ast.isTry.Load() == 0 && param == nil {
			yylex.Error(fmt.Sprintf("operator %q without arguments can only be used when handling an exception", token.literal))
		}
	}
}

func isFunction(isFunc bool, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		ast.isFunction = isFunc
	}
}

func checkReturnParam(param Statement, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		if !ast.isFunction && param != nil {
			yylex.Error("procedure cannot return a value")
		}
	}
}

func setLoopFlag(flag bool, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		if flag {
			ast.isLoop.Add(1)
		} else {
			ast.isLoop.Add(-1)
		}
	}
}

func setTryFlag(flag bool, yylex yyLexer) {
	if ast, ok := yylex.(*AstNode); ok {
		if flag {
			ast.isTry.Add(1)
		} else {
			ast.isTry.Add(-1)
		}
	}
}

func createFunctionOrProcedure(Type StatementType, directive Statement, name string, params []ParamStatement, export Statement, variables map[string]VarStatement, body []Statement) FunctionOrProcedure {
	result := FunctionOrProcedure{
		Type:              Type,
		Name:              name,
		Body:              body,
		Export:            export != nil && !reflect.ValueOf(export).IsNil(),
		Params:            params,
		ExplicitVariables: variables,
	}

	if tok, ok := directive.(*Token); ok && tok != nil {
		result.Directive = tok.literal
	}

	return result
}

func appendVarStatements(existingVariables map[string]VarStatement, newVariables []Token) (map[string]VarStatement, error) {
	for _, v := range newVariables {
		if _, ok := existingVariables[v.literal]; ok {
			return map[string]VarStatement{}, fmt.Errorf("%w: with the specified name %q", errVariableAlreadyDefined, v.literal)
		} else {
			existingVariables[v.literal] = VarStatement{Name: v.literal}
		}
	}
	return existingVariables, nil
}

func unary(iv interface{}) interface{} {
	switch v := iv.(type) {
	case int:
		return -v
	case int32:
		return -v
	case int64:
		return -v
	case float32:
		return -v
	case float64:
		return -v
	case IUnary:
		return v.Unary()
	case INot:
		return v.Not()
	default:
		return v
	}
}

func not(iv interface{}) interface{} {
	switch v := iv.(type) {
	case bool:
		return !v
	case ExpStatement:
		v.Not()
		return v
	case VarStatement:
		v.Not()
		return v
	default:
		return v
	}
}
