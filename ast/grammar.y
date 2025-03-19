%{
package ast
%}

%type<body> body
%type<opt_body> opt_body
%type<stmt> stmt
%type<stmt_loop> stmt_loop
%type<funcProc> funcProc
%type<stmt_if> stmt_if
%type<opt_elseif_list> opt_elseif_list
%type<opt_else> opt_else
%type<stmt> opt_stmt
%type<stmt> opt_param
%type<exprs> exprs 
%type<stmt> expr
%type<opt_export> opt_export
%type<opt_directive> opt_directive 
%type<stmt> simple_expr
%type<declarations_method_params> declarations_method_params
%type<declarations_method_param> declarations_method_param
%type<stmt> opt_expr
%type<stmt> execute_param
%type<stmt> through_dot
%type<stmt> loopExp
%type<stmt> new_object
%type<stmt> ternary
%type<opt_explicit_variables> opt_explicit_variables
%type<explicit_variables> explicit_variables
%type<identifiers> identifiers
%type<stmt> stmt_tryCatch
%type<stmt> identifier
%type<goToLabel> goToLabel
%type<token> separator
%type<token> semicolon
%type<token> colon
%type<token> ':'
%type<token> ';'
%type<global_variables> global_variables



%union {
    token Token
    stmt_if *IfStatement
    opt_elseif_list []Statement
    opt_else []Statement
    stmt    Statement
    stmt_loop *LoopStatement
    funcProc *FunctionOrProcedure
    body []Statement
    opt_body []Statement
    declarations_method_params []ParamStatement
    declarations_method_param ParamStatement
    exprs []Statement
    opt_export *Token
    opt_directive *Token
    explicit_variables map[string]VarStatement
    global_variables []GlobalVariables
    opt_explicit_variables map[string]VarStatement
    identifiers []Token
    goToLabel *GoToLabelStatement
    opt_goToLabel *GoToLabelStatement
}

%token<token> Directive Identifier Procedure Var EndProcedure If Then ElseIf Else EndIf For Each In To Loop EndLoop Break Not ValueParam While GoToLabel
%token<token> Continue Try Catch EndTry Number String New Function EndFunction Return Throw NeEq Le Ge Or And True False Undefind Export Date GoTo Execute


//%right '='
%left Or
%left And
%left NeEq
%left Le
%left Ge
%left Not

%left '='
%left Identifier

//%nonassoc NeEq '>' '<'
//%nonassoc NeEq
//%nonassoc Not
%left '>' '<'
%left '+' '-'
%left '*' '/' '%'
%right UNARMinus UNARYPlus


%%

module: body {
         if ast, ok := yylex.(*AstNode); ok {
            ast.ModuleStatement.Append($1, yylex)
        }
    }
    | main_items opt_body {
         if ast, ok := yylex.(*AstNode); ok {
            ast.ModuleStatement.Append($2, yylex)
        }
    };

main_items: main
    | main_items main
;

main: global_variables {  
        if ast, ok := yylex.(*AstNode); ok {
            ast.ModuleStatement.Append($1, yylex)
        }
    }
    | funcProc {
        if ast, ok := yylex.(*AstNode); ok {
            ast.ModuleStatement.Append($1, yylex)
        }
    }
;

opt_directive:  { $$ = nil}
        | Directive { $$ = &$1}
;

opt_export: { $$ = nil}
        | Export { $$ = &$1}
;

global_variables: opt_directive Var identifiers opt_export semicolon { 
        $$ = make([]GlobalVariables,  len($3), len($3))
        for i, v := range $3 {
            if $1 != nil {
                $$[i].Directive = $1.literal
            }

            $$[i].Export = $4 != nil 
            $$[i].Var = VarStatement { Name: v.literal }
        }
};


funcProc: opt_directive Function Identifier '(' declarations_method_params ')' opt_export { isFunction(true, yylex) } opt_explicit_variables opt_body EndFunction
        {  
            $$ = createFunctionOrProcedure(PFTypeFunction, $1, $3.literal, $5, $7, $9, $10)
            isFunction(false, yylex) 
        }
        | opt_directive Procedure Identifier '(' declarations_method_params ')' opt_export opt_explicit_variables opt_body EndProcedure
        { 
            $$ = createFunctionOrProcedure(PFTypeProcedure, $1, $3.literal, $5, $7, $8, $9)
        }
;

opt_body: { $$ = nil }
	| body { $$ = $1 }
;
    

body: stmt { $$ = []Statement{$1} }
    | opt_body separator opt_stmt { 
        if $2.literal == ":" && len($1) > 0 {
            if _, ok := $1[len($1)-1].(*GoToLabelStatement); !ok {
                yylex.Error("semicolon (;) is expected")
            }
        }
        if $3 != nil {
            $$ = append($$, $3) 
        }
    }
    
;

opt_stmt: { $$ = nil }
        | stmt { $$ = $1 }
;

separator: semicolon { $$ = $1} | colon { $$ = $1};


/* переменные */ 
opt_explicit_variables: { $$ = map[string]VarStatement{} }
            | explicit_variables { $$ = $1 }
;

explicit_variables: Var identifiers semicolon { 
                    if vars, err := appendVarStatements(map[string]VarStatement{}, $2); err != nil {
                        yylex.Error(err.Error()) 
                    } else {
                        $$ = vars
                    }
                }
            | explicit_variables Var identifiers semicolon {
                    if vars, err := appendVarStatements($1, $3); err != nil {
                        yylex.Error(err.Error()) 
                    } else {
                        $$ = vars
                    }
                }
;


/* Если Конецесли */
stmt_if : If expr Then opt_body opt_elseif_list opt_else EndIf {  
    $$ = &IfStatement {
        Expression: $2,
        TrueBlock:  $4,
        IfElseBlock: $5,
        ElseBlock: $6,
    }
};

/* ИначеЕсли */
opt_elseif_list : { $$ = []Statement{} }
        | ElseIf expr Then opt_body opt_elseif_list { 
             $$ = append($5, &IfStatement{
                Expression: $2,
                TrueBlock:  $4,
            })
        };

/* Иначе */
opt_else : { $$ = nil }
        | Else opt_body { $$ = $2 };

/* тернарный оператор */
ternary: '?' '(' expr comma expr comma expr ')' { 
    $$ = TernaryStatement{
            Expression: $3,
            TrueBlock: $5,
            ElseBlock: $7,
        } 
};

/* циклы */
stmt_loop: For Each Identifier In loopExp Loop { setLoopFlag(true, yylex) } opt_body EndLoop {
        $$ = &LoopStatement{
            For: $3.literal,
            In: $5,
            Body: $8,
        }
        setLoopFlag(false, yylex) 
    } 
    | For expr To expr Loop { setLoopFlag(true, yylex) } opt_body EndLoop {
        $$ = &LoopStatement{
            For: $2,
            To: $4,
            Body: $7,
        }
        setLoopFlag(false, yylex)
    }
    |While expr Loop { setLoopFlag(true, yylex) } opt_body EndLoop {
        $$ = &LoopStatement{
            WhileExpr: $2,
            Body: $5,
        }
};


/* описыввает выражения которые можно использовать в циккле Для Каждого */
loopExp: through_dot { $$ = $1 }
        | new_object { $$ = $1 }
        |'(' new_object ')' { $$ = $2 }
;

stmt : expr { $$ = $1 }
    | stmt_if { $$ = $1 }
    | stmt_loop {$$ = $1 }
    | stmt_tryCatch { $$ = $1 }
    | Continue { $$ = ContinueStatement{}; checkLoopOperator($1, yylex) }
    | Break { $$ = BreakStatement{}; checkLoopOperator($1, yylex) }
    | Throw opt_param { $$ = ThrowStatement{ Param: $2 }; checkThrowParam($1, $2, yylex) }
    | Return opt_expr { $$ = &ReturnStatement{ Param: $2 }; checkReturnParam($2, yylex) }
;

opt_param: { $$ = nil } 
            | expr { $$ = $1 }
;


/* вызовы через точку */
through_dot: identifier { $$ = $1 }
        | through_dot dot identifier { $$ = CallChainStatement{ Unit: $3, Call:  $1 } }
;

/* вызовы процедур, функций */
/* вызовы выполнить */
/* выполнить может вызываться так выполнить("что-то") или так выполнить "что-то" */
identifier: Identifier { $$ = VarStatement{ Name: $1.literal } }
        | Identifier '(' exprs ')' { $$ = MethodStatement{ Name: $1.literal, Param: $3 } }
        | identifier '[' expr ']' { $$ = ItemStatement{ Object: $1, Item: $3 } }
        | Execute execute_param { $$ = MethodStatement{ Name: $1.literal, Param: []Statement{$2} } }
        | Execute '(' expr ')' { $$ = MethodStatement{ Name: $1.literal, Param:  []Statement{$3} } }
;

execute_param: String { $$ = $1.value  }
             | Identifier { $$ = VarStatement{ Name: $1.literal }};

/* попытка */
stmt_tryCatch: Try opt_body Catch { setTryFlag(true, yylex) } opt_body EndTry { 
    $$ = TryStatement{ Body: $2, Catch: $5 }
    setTryFlag(false, yylex)
};

/* выражения */
expr : simple_expr { $$ = $1 }
    |'(' expr ')' { $$ = $2 }
    | expr '+' expr { $$ = &ExpStatement{Operation: OpPlus, Left: $1, Right: $3} }
    | expr '-' expr { $$ = &ExpStatement{Operation: OpMinus, Left: $1, Right: $3} }
    | expr '*' expr { $$ = &ExpStatement{Operation: OpMul, Left: $1, Right: $3} }
    | expr '/' expr { $$ = &ExpStatement{Operation: OpDiv, Left: $1, Right: $3} }
    | expr '%' expr { $$ = &ExpStatement{Operation: OpMod, Left: $1, Right: $3} }
    | expr '>' expr { $$ = &ExpStatement{Operation: OpGt, Left: $1, Right: $3} }
    | expr '<' expr { $$ = &ExpStatement{Operation: OpLt, Left: $1, Right: $3} }
	| expr '=' expr { $$ = &ExpStatement{Operation: OpEq, Left: $1, Right: $3 } }
    | expr Or expr {  $$ = &ExpStatement{Operation: OpOr, Left: $1, Right: $3 } } 
    | expr And expr { $$ = &ExpStatement{Operation: OpAnd, Left: $1, Right: $3 } } 
    | expr NeEq expr { $$ = &ExpStatement{Operation: OpNe, Left: $1, Right: $3 } }
    | expr Le expr { $$ = &ExpStatement{Operation: OpLe, Left: $1, Right: $3 } }
    | expr Ge expr { $$ = &ExpStatement{Operation: OpGe, Left: $1, Right: $3 } }
    | Not expr { $$ = not($2) }
    | new_object { $$ = $1 } 
    | GoTo goToLabel { $$ = GoToStatement{ Label: $2 } }
;

opt_expr: { $$ = nil } | expr { $$ = $1 };

// опиасываются правила по которым можно объявлять параметры в функции или процедуре
declarations_method_param: Identifier {  $$ = *(&ParamStatement{}).Fill(nil, $1) } // обычный параметр
            | ValueParam Identifier { $$ = *(&ParamStatement{}).Fill(&$1, $2) } // знач
            | declarations_method_param '=' simple_expr { $$ = *($$.DefaultValue($3)) } // необязательный параметр 
;

declarations_method_params : { $$ = []ParamStatement{} }
                | declarations_method_param  { $$ = []ParamStatement{$1} }
                | declarations_method_params comma declarations_method_param { $$ = append($1, $3) }
;

// для ключевого слова Новый
// 1С допускает такие конструкции
// новый Структура(), новый Массив() ...
// но так же и такие
// Новый("РегистрСведенийКлючЗаписи.СостоянияОригиналовПервичныхДокументов", ПараметрыМассив);
new_object:  New Identifier { $$ = NewObjectStatement{ Constructor: $2.literal } }
            | New Identifier '(' exprs ')' { $$ = NewObjectStatement{ Constructor: $2.literal, Param: $4 } }
            | New '(' exprs ')' { $$ = NewObjectStatement{ Param: $3 } }
;

simple_expr:  String { $$ = $1.value  }
            | Number { $$ =  $1.value }
            | '-' expr %prec UNARMinus { $$ = unaryMinus($2) }
            | '+' expr %prec UNARYPlus { $$ = $2 }
            | True { $$ =  $1.value  }
            | False { $$ =  $1.value  }
            | Date { $$ =  $1.value  }
            | Undefind { $$ = UndefinedStatement{} }
            | goToLabel { $$ = $1}
            | through_dot { 
                if tok, ok := $1.(Token); ok {
                    $$ = tok.literal
                } else {
                    $$ =  $1
                }
            }
            | ternary { $$ =  $1  } // тернарный оператор
;

goToLabel: GoToLabel { $$ = &GoToLabelStatement{ Name: $1.literal } }

exprs : opt_expr {$$ = []Statement{$1} }
	| exprs comma opt_expr { $$ = append($$, $3);  }
;    

identifiers: Identifier { $$ = []Token{$1} }
        | identifiers comma Identifier {$$ = append($$, $3) }
;

semicolon: ';' {$$ = $1};
colon: ':'{$$ = $1};
comma: ',';
dot: '.';

%%