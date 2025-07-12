package ast

import (
	"fmt"
	"strings"
	"time"
)

type PrintConf struct {
	// Margin отступы (количество пробелов)
	Margin  int
	OneLine bool

	// автоматически расставить скобки в выражениях
	LispStyle bool
}

type astPrint struct {
	ast  *AstNode
	conf PrintConf
}

func (ast *AstNode) Print(conf PrintConf) string {
	if ast == nil {
		return ""
	}

	p := &astPrint{conf: conf, ast: ast}
	return p.print()
}

func (ast *AstNode) PrintStatement(stat Statement) string {
	if stat == nil {
		return ""
	}

	return ast.PrintStatementWithConf(stat, PrintConf{Margin: 4})
}

func (ast *AstNode) PrintStatementWithConf(stat Statement, conf PrintConf) string {
	if stat == nil {
		return ""
	}

	p := &astPrint{conf: conf}
	return p.printBodyItem(stat, 0)
}

func (p *astPrint) print() string {
	if len(p.ast.ModuleStatement.Body) == 0 {
		return ""
	}

	builder := &strings.Builder{}

	for _, v := range p.ast.ModuleStatement.GlobalVariables {
		builder.WriteString(p.printGlobalVariables(v))
		builder.WriteString(p.newLine(1))
	}

	for _, node := range p.ast.ModuleStatement.Body {
		if pf, ok := node.(*FunctionOrProcedure); ok {
			builder.WriteString(p.printFunctionOrProcedure(pf))
			builder.WriteString(p.newLine(3))
		} else {
			builder.WriteString(p.newLine(1))
			builder.WriteString(p.printBodyItem(node, 0))
		}
	}

	return builder.String()
}

func (p *astPrint) printGlobalVariables(variables GlobalVariables) string {
	builder := strings.Builder{}

	export := ""
	if variables.Export {
		export = " Экспорт "
	}

	directive := ""
	if variables.Directive != "" {
		directive = "\n" + variables.Directive + "\n"
	}

	builder.WriteString(directive)
	builder.WriteString("Перем ")
	builder.WriteString(variables.Var.Name)
	builder.WriteString(export)
	builder.WriteString(";")

	return builder.String()
}

func (p *astPrint) printFunctionOrProcedure(pf *FunctionOrProcedure) (result string) {
	builder := &strings.Builder{}
	defer func() { result = builder.String() }()

	declaration := ""
	if pf.Type == PFTypeFunction {
		declaration = "Функция"
		defer func() { builder.WriteString(p.newLine(2)); builder.WriteString("КонецФункции ") }()
	} else if pf.Type == PFTypeProcedure {
		declaration = "Процедура"
		defer func() { builder.WriteString(p.newLine(2)); builder.WriteString("КонецПроцедуры ") }()
	}

	var params []string
	// buildParam := strings.Builder{}
	for _, param := range pf.Params {
		val, def := "", ""
		if param.IsValue {
			val = "Знач "
		}

		if asText := p.printVarStatement(param.Default); asText != "" {
			def = " = " + asText
		}

		params = append(params, val+param.Name+def)
	}

	export := ""
	if pf.Export {
		export = "Экспорт "
	}

	directive := ""
	if pf.Directive != "" {
		directive = "\n" + pf.Directive + "\n"
	}

	depth := 1

	builder.WriteString(directive)
	builder.WriteString(declaration)
	builder.WriteString(" ")
	builder.WriteString(pf.Name)
	builder.WriteString("(")
	builder.WriteString(strings.Join(params, ", "))
	builder.WriteString(")")
	builder.WriteString(" ")
	builder.WriteString(export)
	builder.WriteString(p.newLine(1))
	builder.WriteString(p.printBody(pf.Body, depth))

	return
}

func (p *astPrint) printVarStatement(v Statement) string {
	switch val := v.(type) {
	case float64, float32:
		return fmt.Sprintf("%.4f", val)
	case int, int64, int32:
		return fmt.Sprintf("%d", val)
	case string:
		return fmt.Sprintf("\"%s\"", val)
	case bool:
		return IF[string](val, "Истина", "Ложь")
	case time.Time:
		return fmt.Sprintf(`'%s'`, val.Format("20060102150405"))
	case CallChainStatement:
		not := IF[string](val.not, "Не ", "")
		return not + p.printCallChainStatement(val)
	case UndefinedStatement:
		return "Неопределено"
	case MethodStatement:
		not := IF[string](val.not, "Не ", "")
		return not + val.Name + "(" + p.printParams(val.Param.Statements) + ")"
	case VarStatement:
		return val.Name
	case ItemStatement:
		return p.printVarStatement(val.Object) + "[" + p.printExpression(val.Item) + "]"
	case TernaryStatement:
		return fmt.Sprintf("?(%s, %s, %s)", p.printExpression(val.Expression), p.printExpression(val.TrueBlock), p.printExpression(val.ElseBlock))
	case NewObjectStatement:
		return fmt.Sprintf("Новый %s(%s)", val.Constructor, p.printParams(val.Param.Statements))
	case AssignmentStatement:
		return fmt.Sprintf("%s = %s", p.printVarStatement(val.Var), p.printExpression(val.Expr))
	case ExpStatement, ExprStatements, *ExpStatement, *ExprStatements:
		return p.printExpression(val)
	default:
		return ""
	}
}

func (p *astPrint) printParams(Params Statements) string {
	params := make([]string, len(Params), len(Params))
	for i, parm := range Params {
		params[i] = p.printVarStatement(parm)
	}

	return strings.Join(params, ", ")
}

func (p *astPrint) printBody(items Statements, depth int) string {
	builder := &strings.Builder{}

	for _, item := range items {
		builder.WriteString(p.newLine(1))
		builder.WriteString(p.printBodyItem(item, depth))
	}

	builder.WriteString(p.newLine(1))

	return builder.String()
}

func (p *astPrint) printBodyItem(item Statement, depth int) string {
	builder := &strings.Builder{}

	spaces := strings.Repeat(" ", p.conf.Margin*depth)
	builder.WriteString(spaces)

	switch v := item.(type) {
	case *IfStatement:
		builder.WriteString(p.printIfStatement(v, depth))
		builder.WriteString(";")
		builder.WriteString(p.newLine(1))
	case *ExpStatement:
		builder.WriteString(p.printExpression(v))
		builder.WriteString(";")
	case *LoopStatement:
		builder.WriteString(p.printLoopStatement(v, depth))
		builder.WriteString(";")
		builder.WriteString(p.newLine(1))
	case BreakStatement:
		builder.WriteString("Прервать;")
	case ContinueStatement:
		builder.WriteString("Продолжить;")
	case CallChainStatement:
		builder.WriteString(p.printCallChainStatement(v))
		builder.WriteString(";")
	case TryStatement:
		builder.WriteString(p.printTryStatement(v, depth))
		builder.WriteString(";")
		builder.WriteString(p.newLine(1))
	case ThrowStatement:
		builder.WriteString("ВызватьИсключение")
		if v.Param != nil {
			if param, ok := v.Param.(ExprStatements); ok {
				builder.WriteString("(" + p.printParams(param.Statements) + ")")
			} else {
				builder.WriteString("(" + p.printParams(Statements{v.Param}) + ")")
			}
		}
		builder.WriteString(";")
	case *ReturnStatement:
		builder.WriteString("Возврат")
		if v.Param != nil {
			builder.WriteString(" ")
			builder.WriteString(p.printExpression(v.Param))
		}
		builder.WriteString(";")
	case GoToStatement, *GoToLabelStatement:
		builder.WriteString(p.printGoTo(v, depth))
		builder.WriteString(p.newLine(1))
	default:
		builder.WriteString(p.printVarStatement(v))
		builder.WriteString(";")
	}

	return builder.String()
}

func (p *astPrint) printIfStatement(expr *IfStatement, depth int) string {
	builder := &strings.Builder{}

	spaces := strings.Repeat(" ", p.conf.Margin*depth)
	builder.WriteString("Если ")
	builder.WriteString(p.printExpression(expr.Expression))
	builder.WriteString(" Тогда ")

	builder.WriteString(p.printBody(expr.TrueBlock, depth+1))
	for _, item := range expr.IfElseBlock {
		builder.WriteString(spaces)
		builder.WriteString("ИначеЕсли ")
		builder.WriteString(p.printExpression(item.(*IfStatement).Expression))
		builder.WriteString(" Тогда ")
		builder.WriteString(p.printBody(item.(*IfStatement).TrueBlock, depth+1))
	}

	if expr.ElseBlock != nil {
		builder.WriteString(spaces)
		builder.WriteString("Иначе ")
		builder.WriteString(p.printBody(expr.ElseBlock, depth+1))
	}

	builder.WriteString(spaces)
	builder.WriteString("КонецЕсли")
	return builder.String()
}

func (p *astPrint) printLoopStatement(loop *LoopStatement, depth int) string {
	builder := &strings.Builder{}

	spaces := strings.Repeat(" ", p.conf.Margin*depth)
	if loop.WhileExpr != nil {
		builder.WriteString("Пока ")
		builder.WriteString(p.printExpression(loop.WhileExpr))
		builder.WriteString(" Цикл ")
	} else {
		builder.WriteString("Для ")
	}

	if loop.In != nil {
		builder.WriteString("Каждого ")
		builder.WriteString(loop.For.(string))
		builder.WriteString(" Из ")
		builder.WriteString(p.printExpression(loop.In))
		builder.WriteString(" Цикл ")
	}
	if loop.To != nil {
		builder.WriteString(p.printExpression(loop.For))
		builder.WriteString(" По ")
		builder.WriteString(p.printExpression(loop.To))
		builder.WriteString(" Цикл ")
	}

	builder.WriteString(p.printBody(loop.Body, depth+1))
	builder.WriteString(spaces)
	builder.WriteString("КонецЦикла")

	return builder.String()
}

func (p *astPrint) printExpression(expr Statement) string {
	builder := &strings.Builder{}

	switch v := expr.(type) {
	case ExprStatements:
		if v.not {
			builder.WriteString("Не ")
			builder.WriteString("(")
		}

		for _, s := range v.Statements {
			builder.WriteString(p.printExpression(s))
		}

		if v.not {
			builder.WriteString(")")
		}
	case *ExpStatement:
		if p.conf.LispStyle {
			builder.WriteString("(")
		}

		if v.not {
			builder.WriteString("Не ")
		}
		if v.unaryMinus {
			builder.WriteString("-")
		}

		if v.unaryMinus || v.not {
			builder.WriteString("(")
		}

		builder.WriteString(p.printExpression(v.Left))
		builder.WriteString(" ")
		builder.WriteString(v.Operation.String())
		builder.WriteString(" ")
		builder.WriteString(p.printExpression(v.Right))

		if v.unaryMinus || v.not {
			builder.WriteString(")")
		}

		if p.conf.LispStyle {
			builder.WriteString(")")
		}
	case VarStatement:
		if v.not {
			builder.WriteString("Не ")
		}
		if v.unaryMinus {
			builder.WriteString("-")
		}
		builder.WriteString(p.printVarStatement(v))
	default:
		builder.WriteString(p.printVarStatement(v))
	}

	return builder.String()
}

func (p *astPrint) printCallChainStatement(call Statement) string {
	switch v := call.(type) {
	case CallChainStatement:
		if v.Call != nil {
			return p.printCallChainStatement(v.Call) + "." + p.printVarStatement(v.Unit)
		}
	case VarStatement, ItemStatement, MethodStatement:
		return p.printVarStatement(call)
	}

	return ""
}

func (p *astPrint) printTryStatement(try TryStatement, depth int) string {
	builder := &strings.Builder{}

	spaces := strings.Repeat(" ", p.conf.Margin*depth)
	builder.WriteString("Попытка")

	if try.Body != nil {
		builder.WriteString(p.printBody(try.Body, depth+1))
	} else {
		builder.WriteString(p.newLine(1))
	}

	builder.WriteString(spaces)
	builder.WriteString("Исключение")
	if try.Catch != nil {
		builder.WriteString(p.printBody(try.Catch, depth+1))
	} else {
		builder.WriteString(p.newLine(1))
	}

	builder.WriteString(spaces)
	builder.WriteString("КонецПопытки")
	return builder.String()
}

func (p *astPrint) printGoTo(gotoStat Statement, depth int) string {
	builder := strings.Builder{}

	// spaces := strings.Repeat(" ", p.conf.Margin*depth)

	switch v := gotoStat.(type) {
	case *GoToLabelStatement:
		// builder.WriteString(spaces)
		builder.WriteString("~")
		builder.WriteString(v.Name)
		builder.WriteString(":")
	case GoToStatement:
		// builder.WriteString(spaces)
		builder.WriteString("Перейти ")
		builder.WriteString("~")
		builder.WriteString(v.Label.Name)
		builder.WriteString(";")
	}

	return builder.String()
}

func (p *astPrint) newLine(count int) string {
	if p.conf.OneLine {
		return ""
	}

	return strings.Repeat("\n", count)
}
