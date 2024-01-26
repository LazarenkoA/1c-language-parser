%{
package main
%}

//%type<stmts> stmts
%type<stmt> stmt
//%type<stmt_if> stmt_if
//%type<exprs> exprs
//%type<compstmt> compstmt

%union {
    token Token
    //terminal                   Token
    //terms                  Token
    stmt_if statement // оператор Если
    stmts   statements
    stmt    statement
    compstmt    statements
}

%token<token> Directive Identifier Procedure Var EndProcedure If Then ElseIf Else EndIf For Each In To Loop EndLoop Break Not
%token<token> Continue Try Catch EndTry Number String Array Function EndFunction Return Throw NeEq Or And True False Undefind Export

%right '='
%left Or
%left And
%left Identifier
//%nonassoc NeEq '>' '<'
%nonassoc NeEq
%left '>' '<'
%left '+' '-'
%left '*' '/' '%'



%%

main: opt_directive invoke {}
;

opt_directive: /* none */
| Directive EOL {}
;

opt_export: /* none */
| Export {}
;


invoke: Function Identifier '(' exprs ')' opt_export opt_stmt EndFunction opt_EOLs {}
        | Procedure Identifier '(' exprs ')' opt_export opt_stmt EndProcedure opt_EOLs {}
;

opt_stmt : /* none */
	| opt_EOLs stmts  {}
    | EOLs {}
;
    
stmts : stmt end_stmt opt_EOLs {}
        | stmts stmt end_stmt opt_EOLs {}
;


stmt : Throw throw_terms {}
    | Break {}
    | Continue {}
    | expr {}
    | stmt_if {}
;

throw_terms : end_stmt {}
	| String end_stmt {}
;


/* Если ИначеЕсли Иначе Конецесли */
stmt_if : If expr Then opt_stmt opt_elseif_list opt_else EndIf 
        {

        }
;

opt_elseif_list : /* none */
               | ElseIf expr Then opt_stmt opt_elseif_list
               {
                   // Обработка блока ElseIf
               };

opt_else : /* none */
         | Else opt_stmt
         {
             // Обработка блока Else
         };

/* выражения */
expr : Identifier{}
    |'(' expr ')' { }
    | expr '+' expr { }
    | expr '-' expr { }
    | expr '*' expr { }
    | expr '/' expr { }
    | expr '%' expr {}
    | String {}
    | Number { }
    | True { }
    | False { }
    | Undefind {}
    | expr '>' expr { }
    | expr '<' expr{}
	| expr '=' expr {}
    | Identifier  '=' expr {}
    | expr NeEq expr {}
    | expr And expr { }
    | expr Or expr {}
	| Identifier '(' exprs ')' {}
;

exprs : { }
	| expr {}
	| exprs ',' opt_EOLs expr {}
;    


// [\n]*;
end_stmt: opt_EOLs semi {} 

// опциональные переносы строк [\n]*
opt_EOLs : /* none */ 
        |EOLs;

// несколько переносов строк [\n]+
EOLs: EOL {} 
     | EOLs EOL;

EOL : '\n';
semi: ';';

%%