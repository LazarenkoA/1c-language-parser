%{
package main
%}

%type<end_stmt> end_stmt
%type<EOLs> EOLs

%union {
    token Token
    terminal                   Token
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

end_stmt: opt_EOL semi {};


opt_EOL :
        |EOLs;

EOLs: EOL {}
     | EOLs EOL;


EOL : '\n';
semi: ';';

%%