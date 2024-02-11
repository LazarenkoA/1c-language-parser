// Code generated by goyacc .\grammar.y. DO NOT EDIT.

//line .\grammar.y:2
package ast

import __yyfmt__ "fmt"

//line .\grammar.y:2

//line .\grammar.y:41
type yySymType struct {
	yys                        int
	token                      Token
	stmt_if                    *IfStatement
	opt_elseif_list            []Statement
	opt_else                   []Statement
	stmt                       Statement
	opt_stmt                   Statement
	stmt_tryCatch              Statement
	stmt_loop                  *LoopStatement
	funcProc                   *FunctionOrProcedure
	body                       []Statement
	opt_body                   []Statement
	opt_param                  Statement
	through_dot                Statement
	manyfuncProc               []Statement
	declarations_method_params []ParamStatement
	declarations_method_param  ParamStatement
	expr                       Statement
	opt_expr                   Statement
	//opt_exprs []Statement
	exprs                  []Statement
	opt_export             *Token
	opt_directive          *Token
	simple_expr            Statement
	new_object             Statement
	ternary                Statement
	explicit_variables     map[string]VarStatement
	opt_explicit_variables map[string]VarStatement
	identifiers            []Token
	identifier             Statement
	goToLabel              *GoToLabelStatement
	opt_goToLabel          *GoToLabelStatement
}

const Directive = 57346
const Identifier = 57347
const Procedure = 57348
const Var = 57349
const EndProcedure = 57350
const If = 57351
const Then = 57352
const ElseIf = 57353
const Else = 57354
const EndIf = 57355
const For = 57356
const Each = 57357
const In = 57358
const To = 57359
const Loop = 57360
const EndLoop = 57361
const Break = 57362
const Not = 57363
const ValueParam = 57364
const While = 57365
const GoToLabel = 57366
const Continue = 57367
const Try = 57368
const Catch = 57369
const EndTry = 57370
const Number = 57371
const String = 57372
const New = 57373
const Function = 57374
const EndFunction = 57375
const Return = 57376
const Throw = 57377
const NeEq = 57378
const Le = 57379
const Ge = 57380
const Or = 57381
const And = 57382
const True = 57383
const False = 57384
const Undefind = 57385
const Export = 57386
const Date = 57387
const GoTo = 57388
const UNARY = 57389

var yyToknames = [...]string{
	"$end",
	"error",
	"$unk",
	"':'",
	"';'",
	"Directive",
	"Identifier",
	"Procedure",
	"Var",
	"EndProcedure",
	"If",
	"Then",
	"ElseIf",
	"Else",
	"EndIf",
	"For",
	"Each",
	"In",
	"To",
	"Loop",
	"EndLoop",
	"Break",
	"Not",
	"ValueParam",
	"While",
	"GoToLabel",
	"Continue",
	"Try",
	"Catch",
	"EndTry",
	"Number",
	"String",
	"New",
	"Function",
	"EndFunction",
	"Return",
	"Throw",
	"NeEq",
	"Le",
	"Ge",
	"Or",
	"And",
	"True",
	"False",
	"Undefind",
	"Export",
	"Date",
	"GoTo",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARY",
	"'('",
	"')'",
	"'?'",
	"'['",
	"']'",
	"','",
	"'.'",
}

var yyStatenames = [...]string{}

const yyEofCode = 1
const yyErrCode = 2
const yyInitialStackSize = 16

//line .\grammar.y:338

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 2,
}

const yyPrivate = 57344

const yyLast = 490

var yyAct = [...]uint8{
	68, 65, 159, 19, 67, 35, 110, 63, 62, 95,
	96, 97, 93, 94, 40, 169, 44, 108, 45, 81,
	92, 90, 91, 85, 86, 87, 88, 89, 157, 21,
	25, 104, 20, 37, 20, 20, 23, 18, 29, 28,
	20, 20, 34, 87, 88, 89, 53, 61, 64, 138,
	30, 31, 33, 83, 32, 47, 98, 99, 100, 44,
	46, 12, 11, 38, 20, 103, 105, 39, 85, 86,
	87, 88, 89, 123, 172, 113, 64, 115, 117, 118,
	48, 119, 8, 114, 147, 122, 124, 125, 126, 127,
	128, 129, 130, 131, 132, 133, 134, 135, 136, 121,
	120, 37, 101, 54, 139, 15, 64, 52, 7, 182,
	180, 24, 142, 140, 171, 153, 95, 96, 97, 93,
	94, 14, 16, 51, 122, 27, 173, 92, 90, 91,
	85, 86, 87, 88, 89, 41, 167, 160, 148, 64,
	84, 26, 13, 151, 106, 152, 154, 150, 80, 50,
	40, 149, 158, 144, 105, 17, 82, 163, 164, 161,
	102, 168, 22, 10, 40, 9, 111, 110, 76, 174,
	5, 3, 177, 77, 6, 156, 43, 178, 179, 73,
	57, 181, 78, 37, 72, 79, 155, 40, 29, 28,
	60, 170, 176, 75, 74, 42, 1, 116, 109, 107,
	30, 31, 33, 57, 32, 59, 37, 71, 49, 36,
	56, 29, 28, 60, 58, 55, 4, 39, 40, 2,
	112, 141, 166, 30, 31, 33, 69, 32, 59, 70,
	66, 0, 0, 56, 57, 96, 97, 37, 55, 0,
	39, 0, 29, 28, 60, 92, 90, 91, 85, 86,
	87, 88, 89, 0, 30, 31, 33, 0, 32, 59,
	0, 0, 0, 0, 56, 0, 0, 0, 0, 55,
	0, 39, 95, 96, 97, 93, 94, 0, 0, 0,
	0, 0, 0, 92, 90, 91, 85, 86, 87, 88,
	89, 175, 0, 165, 95, 96, 97, 93, 94, 0,
	0, 0, 0, 0, 0, 92, 90, 91, 85, 86,
	87, 88, 89, 0, 0, 137, 0, 95, 96, 97,
	93, 94, 162, 0, 0, 0, 0, 0, 92, 90,
	91, 85, 86, 87, 88, 89, 0, 0, 0, 0,
	95, 96, 97, 93, 94, 146, 0, 0, 0, 0,
	0, 92, 90, 91, 85, 86, 87, 88, 89, 0,
	0, 0, 0, 95, 96, 97, 93, 94, 145, 0,
	0, 0, 0, 0, 92, 90, 91, 85, 86, 87,
	88, 89, 143, 0, 0, 0, 0, 95, 96, 97,
	93, 94, 0, 0, 0, 0, 0, 0, 92, 90,
	91, 85, 86, 87, 88, 89, 0, 0, 95, 96,
	97, 93, 94, 0, 0, 0, 0, 0, 0, 92,
	90, 91, 85, 86, 87, 88, 89, 95, 96, 97,
	93, 94, 0, 0, 0, 0, 0, 0, 92, 90,
	91, 85, 86, 87, 88, 89, 95, 96, 97, 0,
	94, 0, 0, 0, 95, 96, 97, 92, 90, 91,
	85, 86, 87, 88, 89, 92, 90, 91, 85, 86,
	87, 88, 89, 97, 92, 90, 91, 85, 86, 87,
	88, 89, 92, 90, 91, 85, 86, 87, 88, 89,
}

var yyPact = [...]int16{
	164, -1000, 164, -1000, 74, -1000, -1000, 158, 156, 4,
	3, 98, 98, -22, -20, -1000, 155, -23, -16, 98,
	-1000, 7, -1000, -16, -1000, -1000, -20, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -48, -1000, -1000, -43, 2,
	-3, 140, 140, 143, -1000, 211, 211, 211, 157, 139,
	149, 157, -43, 78, -1000, 211, 211, 211, -1000, 75,
	153, -29, -28, -1000, 389, 134, 162, -1000, 389, -1000,
	-1000, -1000, -1000, -1000, 211, 211, 211, 180, 211, 157,
	149, 1, -1000, 38, -1000, 211, 211, 211, 211, 211,
	211, 211, 211, 211, 211, 211, 211, 211, 256, -1000,
	425, -1000, -9, 211, -1000, 211, -1000, 157, -1000, -1000,
	-1000, -1000, -1000, 389, -1000, 370, 146, 349, 325, 55,
	1, -1000, 144, -1000, -11, -11, -1000, -1000, -1000, 16,
	16, 425, 408, 416, 196, 433, 425, -1000, 211, -29,
	-1000, -1000, -1000, 157, 97, 211, -1000, -1000, -1000, -1000,
	-31, 211, 124, 143, 302, 157, 157, -1000, 234, 122,
	211, -5, -1000, 93, 44, -1000, 111, 157, 279, -1000,
	157, -1000, -1000, -1000, -1000, 157, 157, 89, 124, 88,
	-1000, -1000, -1000,
}

var yyPgo = [...]uint8{
	0, 230, 1, 4, 229, 171, 226, 2, 222, 221,
	220, 219, 8, 0, 111, 216, 103, 142, 121, 7,
	5, 214, 209, 80, 208, 19, 207, 63, 42, 199,
	17, 198, 196, 195, 3, 192, 191, 186, 176, 175,
}

var yyR1 = [...]int8{
	0, 32, 15, 15, 14, 14, 11, 11, 33, 5,
	5, 2, 2, 1, 1, 9, 9, 29, 29, 23,
	23, 24, 24, 6, 7, 7, 8, 8, 22, 35,
	4, 36, 4, 37, 4, 3, 3, 3, 3, 3,
	3, 3, 3, 10, 10, 20, 20, 27, 27, 27,
	39, 26, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 19, 19, 18, 18, 18, 17, 17, 17, 21,
	21, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	28, 12, 12, 25, 25, 30, 31, 34, 38,
}

var yyR2 = [...]int8{
	0, 1, 0, 1, 0, 1, 1, 2, 0, 11,
	10, 0, 1, 1, 3, 0, 1, 1, 1, 0,
	1, 3, 4, 7, 0, 5, 0, 2, 8, 0,
	9, 0, 8, 0, 6, 1, 1, 1, 1, 1,
	1, 2, 2, 0, 1, 1, 3, 1, 4, 4,
	0, 6, 1, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 2, 3, 3, 3, 3, 3, 2, 1,
	2, 0, 1, 1, 2, 3, 0, 1, 3, 2,
	5, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 3, 1, 3, 1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -32, -11, -5, -15, 6, -5, 34, 8, 7,
	7, 58, 58, -17, -18, 7, 24, -17, 59, -34,
	63, 49, 7, 59, -14, 46, -18, -16, 32, 31,
	43, 44, 47, 45, -28, -20, -22, 26, -27, 60,
	7, -14, -33, -38, 64, 61, 58, 58, -23, -24,
	9, -23, -27, -13, -16, 58, 53, 23, -21, 48,
	33, -13, -12, -19, -13, -2, -1, -3, -13, -6,
	-4, -26, 27, 22, 37, 36, 11, 16, 25, 28,
	9, -25, 7, -2, 62, 52, 53, 54, 55, 56,
	50, 51, 49, 41, 42, 38, 39, 40, -13, -13,
	-13, -28, 7, -34, 59, -34, 10, -29, -30, -31,
	5, 4, -10, -13, -19, -13, 17, -13, -13, -2,
	-25, -30, -34, 35, -13, -13, -13, -13, -13, -13,
	-13, -13, -13, -13, -13, -13, -13, 59, 58, -13,
	-19, -9, -3, 12, 7, 19, 20, 29, -30, 7,
	-12, -34, -2, 18, -13, -37, -39, 59, -13, -7,
	13, -20, 20, -2, -2, 59, -8, 14, -13, 20,
	-36, 21, 30, 15, -2, 12, -35, -2, -2, -2,
	21, -7, 21,
}

var yyDef = [...]int8{
	2, -2, -2, 6, 0, 3, 7, 0, 0, 0,
	0, 76, 76, 0, 77, 73, 0, 0, 4, 0,
	97, 0, 74, 4, 8, 5, 78, 75, 81, 82,
	83, 84, 85, 86, 87, 88, 89, 90, 45, 0,
	47, 19, 19, 0, 98, 0, 0, 71, 11, 20,
	0, 11, 46, 0, 52, 0, 0, 0, 69, 0,
	0, 0, 0, 91, 72, 0, 12, 13, 35, 36,
	37, 38, 39, 40, 43, 71, 0, 0, 0, 11,
	0, 0, 93, 0, 49, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 62,
	68, 70, 79, 0, 48, 71, 10, 15, 17, 18,
	95, 96, 41, 44, 42, 0, 0, 0, 0, 0,
	0, 21, 0, 9, 54, 55, 56, 57, 58, 59,
	60, 61, 63, 64, 65, 66, 67, 53, 71, 0,
	92, 14, 16, 11, 0, 0, 33, 50, 22, 94,
	0, 0, 24, 0, 0, 11, 11, 80, 0, 26,
	0, 0, 31, 0, 0, 28, 0, 11, 0, 29,
	11, 34, 51, 23, 27, 11, 11, 0, 24, 0,
	32, 25, 30,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 56, 3, 3,
	58, 59, 54, 52, 63, 53, 64, 55, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 5,
	51, 49, 50, 60, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 61, 3, 62,
}

var yyTok2 = [...]int8{
	2, 3, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 57,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	state int
	token int
	msg   string
}{}

//line yaccpar:1

/*	parser for yacc output	*/

var (
	yyDebug        = 0
	yyErrorVerbose = false
)

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

type yyParser interface {
	Parse(yyLexer) int
	Lookahead() int
}

type yyParserImpl struct {
	lval  yySymType
	stack [yyInitialStackSize]yySymType
	char  int
}

func (p *yyParserImpl) Lookahead() int {
	return p.char
}

func yyNewParser() yyParser {
	return &yyParserImpl{}
}

const yyFlag = -1000

func yyTokname(c int) string {
	if c >= 1 && c-1 < len(yyToknames) {
		if yyToknames[c-1] != "" {
			return yyToknames[c-1]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yyErrorMessage(state, lookAhead int) string {
	const TOKSTART = 4

	if !yyErrorVerbose {
		return "syntax error"
	}

	for _, e := range yyErrorMessages {
		if e.state == state && e.token == lookAhead {
			return "syntax error: " + e.msg
		}
	}

	res := "syntax error: unexpected " + yyTokname(lookAhead)

	// To match Bison, suggest at most four expected tokens.
	expected := make([]int, 0, 4)

	// Look for shiftable tokens.
	base := int(yyPact[state])
	for tok := TOKSTART; tok-1 < len(yyToknames); tok++ {
		if n := base + tok; n >= 0 && n < yyLast && int(yyChk[int(yyAct[n])]) == tok {
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}
	}

	if yyDef[state] == -2 {
		i := 0
		for yyExca[i] != -1 || int(yyExca[i+1]) != state {
			i += 2
		}

		// Look for tokens that we accept or reduce.
		for i += 2; yyExca[i] >= 0; i += 2 {
			tok := int(yyExca[i])
			if tok < TOKSTART || yyExca[i+1] == 0 {
				continue
			}
			if len(expected) == cap(expected) {
				return res
			}
			expected = append(expected, tok)
		}

		// If the default action is to accept or reduce, give up.
		if yyExca[i+1] != 0 {
			return res
		}
	}

	for i, tok := range expected {
		if i == 0 {
			res += ", expecting "
		} else {
			res += " or "
		}
		res += yyTokname(tok)
	}
	return res
}

func yylex1(lex yyLexer, lval *yySymType) (char, token int) {
	token = 0
	char = lex.Lex(lval)
	if char <= 0 {
		token = int(yyTok1[0])
		goto out
	}
	if char < len(yyTok1) {
		token = int(yyTok1[char])
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			token = int(yyTok2[char-yyPrivate])
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			token = int(yyTok3[i+1])
			goto out
		}
	}

out:
	if token == 0 {
		token = int(yyTok2[1]) /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(token), uint(char))
	}
	return char, token
}

func yyParse(yylex yyLexer) int {
	return yyNewParser().Parse(yylex)
}

func (yyrcvr *yyParserImpl) Parse(yylex yyLexer) int {
	var yyn int
	var yyVAL yySymType
	var yyDollar []yySymType
	_ = yyDollar // silence set and not used
	yyS := yyrcvr.stack[:]

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yyrcvr.char = -1
	yytoken := -1 // yyrcvr.char translated into internal numbering
	defer func() {
		// Make sure we report no lookahead when not parsing.
		yystate = -1
		yyrcvr.char = -1
		yytoken = -1
	}()
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yytoken), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = int(yyPact[yystate])
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yyrcvr.char < 0 {
		yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
	}
	yyn += yytoken
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = int(yyAct[yyn])
	if int(yyChk[yyn]) == yytoken { /* valid shift */
		yyrcvr.char = -1
		yytoken = -1
		yyVAL = yyrcvr.lval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = int(yyDef[yystate])
	if yyn == -2 {
		if yyrcvr.char < 0 {
			yyrcvr.char, yytoken = yylex1(yylex, &yyrcvr.lval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && int(yyExca[xi+1]) == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = int(yyExca[xi+0])
			if yyn < 0 || yyn == yytoken {
				break
			}
		}
		yyn = int(yyExca[xi+1])
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error(yyErrorMessage(yystate, yytoken))
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yytoken))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = int(yyPact[yyS[yyp].yys]) + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = int(yyAct[yyn]) /* simulate a shift of "error" */
					if int(yyChk[yystate]) == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yytoken))
			}
			if yytoken == yyEofCode {
				goto ret1
			}
			yyrcvr.char = -1
			yytoken = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= int(yyR2[yyn])
	// yyp is now the index of $0. Perform the default action. Iff the
	// reduced production is ε, $1 is possibly out of range.
	if yyp+1 >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = int(yyR1[yyn])
	yyg := int(yyPgo[yyn])
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = int(yyAct[yyg])
	} else {
		yystate = int(yyAct[yyj])
		if int(yyChk[yystate]) != -yyn {
			yystate = int(yyAct[yyg])
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:101
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement = ModuleStatement{
					Name: "",
					Body: yyDollar[1].manyfuncProc,
				}
			}
		}
	case 2:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:110
		{
			yyVAL.opt_directive = nil
		}
	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:111
		{
			yyVAL.opt_directive = &yyDollar[1].token
		}
	case 4:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:114
		{
			yyVAL.opt_export = nil
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:115
		{
			yyVAL.opt_export = &yyDollar[1].token
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:118
		{
			yyVAL.manyfuncProc = []Statement{yyDollar[1].funcProc}
		}
	case 7:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:119
		{
			yyVAL.manyfuncProc = append(yyDollar[1].manyfuncProc, yyDollar[2].funcProc)
		}
	case 8:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:121
		{
			isFunction(true, yylex)
		}
	case 9:
		yyDollar = yyS[yypt-11 : yypt+1]
//line .\grammar.y:122
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeFunction, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[9].opt_explicit_variables, yyDollar[10].opt_body)
			isFunction(false, yylex)
		}
	case 10:
		yyDollar = yyS[yypt-10 : yypt+1]
//line .\grammar.y:127
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeProcedure, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[8].opt_explicit_variables, yyDollar[9].opt_body)
		}
	case 11:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:132
		{
			yyVAL.opt_body = nil
		}
	case 12:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:133
		{
			yyVAL.opt_body = yyDollar[1].body
		}
	case 13:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:136
		{
			yyVAL.body = []Statement{yyDollar[1].stmt}
		}
	case 14:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:137
		{
			if yyDollar[2].token.literal == ":" && len(yyDollar[1].body) > 0 {
				if _, ok := yyDollar[1].body[len(yyDollar[1].body)-1].(*GoToLabelStatement); !ok {
					yylex.Error("semicolon (;) is expected")
				}
			}
			if yyDollar[3].opt_stmt != nil {
				yyVAL.body = append(yyVAL.body, yyDollar[3].opt_stmt)
			}
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:150
		{
			yyVAL.opt_stmt = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:151
		{
			yyVAL.opt_stmt = yyDollar[1].stmt
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:154
		{
			yyVAL.token = yyDollar[1].token
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:154
		{
			yyVAL.token = yyDollar[1].token
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:158
		{
			yyVAL.opt_explicit_variables = map[string]VarStatement{}
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:159
		{
			yyVAL.opt_explicit_variables = yyDollar[1].explicit_variables
		}
	case 21:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:162
		{
			if vars, err := appendVarStatements(map[string]VarStatement{}, yyDollar[2].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 22:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:169
		{
			if vars, err := appendVarStatements(yyDollar[1].explicit_variables, yyDollar[3].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 23:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:179
		{
			yyVAL.stmt_if = &IfStatement{
				Expression:  yyDollar[2].expr,
				TrueBlock:   yyDollar[4].opt_body,
				IfElseBlock: yyDollar[5].opt_elseif_list,
				ElseBlock:   yyDollar[6].opt_else,
			}
		}
	case 24:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:189
		{
			yyVAL.opt_elseif_list = []Statement{}
		}
	case 25:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:190
		{
			yyVAL.opt_elseif_list = append(yyDollar[5].opt_elseif_list, &IfStatement{
				Expression: yyDollar[2].expr,
				TrueBlock:  yyDollar[4].opt_body,
			})
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:198
		{
			yyVAL.opt_else = nil
		}
	case 27:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:199
		{
			yyVAL.opt_else = yyDollar[2].opt_body
		}
	case 28:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:202
		{
			yyVAL.ternary = TernaryStatement{
				Expression: yyDollar[3].expr,
				TrueBlock:  yyDollar[5].expr,
				ElseBlock:  yyDollar[7].expr,
			}
		}
	case 29:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:211
		{
			setLoopFlag(true, yylex)
		}
	case 30:
		yyDollar = yyS[yypt-9 : yypt+1]
//line .\grammar.y:211
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[3].token.literal,
				In:   yyDollar[5].through_dot,
				Body: yyDollar[8].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 31:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:219
		{
			setLoopFlag(true, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:219
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[2].expr,
				To:   yyDollar[4].expr,
				Body: yyDollar[7].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:227
		{
			setLoopFlag(true, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:227
		{
			yyVAL.stmt_loop = &LoopStatement{
				WhileExpr: yyDollar[2].expr,
				Body:      yyDollar[5].opt_body,
			}
		}
	case 35:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:234
		{
			yyVAL.stmt = yyDollar[1].expr
		}
	case 36:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:235
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:236
		{
			yyVAL.stmt = yyDollar[1].stmt_loop
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:237
		{
			yyVAL.stmt = yyDollar[1].stmt_tryCatch
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:238
		{
			yyVAL.stmt = ContinueStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:239
		{
			yyVAL.stmt = BreakStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 41:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:240
		{
			yyVAL.stmt = ThrowStatement{Param: yyDollar[2].opt_param}
			checkThrowParam(yyDollar[1].token, yyDollar[2].opt_param, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:241
		{
			yyVAL.stmt = ReturnStatement{Param: yyDollar[2].opt_expr}
			checkReturnParam(yyDollar[2].opt_expr, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:244
		{
			yyVAL.opt_param = nil
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:245
		{
			yyVAL.opt_param = yyDollar[1].expr
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:250
		{
			yyVAL.through_dot = yyDollar[1].identifier
		}
	case 46:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:251
		{
			yyVAL.through_dot = CallChainStatement{Unit: yyDollar[3].identifier, Call: yyDollar[1].through_dot}
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:254
		{
			yyVAL.identifier = VarStatement{Name: yyDollar[1].token.literal}
		}
	case 48:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:255
		{
			yyVAL.identifier = MethodStatement{Name: yyDollar[1].token.literal, Param: yyDollar[3].exprs}
		}
	case 49:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:256
		{
			yyVAL.identifier = ItemStatement{Object: yyDollar[1].identifier, Item: yyDollar[3].expr}
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:261
		{
			setTryFlag(true, yylex)
		}
	case 51:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:261
		{
			yyVAL.stmt_tryCatch = TryStatement{Body: yyDollar[2].opt_body, Catch: yyDollar[5].opt_body}
			setTryFlag(false, yylex)
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:267
		{
			yyVAL.expr = yyDollar[1].simple_expr
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:268
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:269
		{
			yyVAL.expr = &ExpStatement{Operation: OpPlus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:270
		{
			yyVAL.expr = &ExpStatement{Operation: OpMinus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:271
		{
			yyVAL.expr = &ExpStatement{Operation: OpMul, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:272
		{
			yyVAL.expr = &ExpStatement{Operation: OpDiv, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:273
		{
			yyVAL.expr = &ExpStatement{Operation: OpMod, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:274
		{
			yyVAL.expr = &ExpStatement{Operation: OpGt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:275
		{
			yyVAL.expr = &ExpStatement{Operation: OpLt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:276
		{
			yyVAL.expr = &ExpStatement{Operation: OpEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:277
		{
			yyVAL.expr = unary(yyDollar[2].expr)
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:278
		{
			yyVAL.expr = &ExpStatement{Operation: OpOr, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:279
		{
			yyVAL.expr = &ExpStatement{Operation: OpAnd, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:280
		{
			yyVAL.expr = &ExpStatement{Operation: OpNe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:281
		{
			yyVAL.expr = &ExpStatement{Operation: OpLe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:282
		{
			yyVAL.expr = &ExpStatement{Operation: OpGe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:283
		{
			yyVAL.expr = not(yyDollar[2].expr)
		}
	case 69:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:284
		{
			yyVAL.expr = yyDollar[1].new_object
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:285
		{
			yyVAL.expr = GoToStatement{Label: yyDollar[2].goToLabel}
		}
	case 71:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:288
		{
			yyVAL.opt_expr = nil
		}
	case 72:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:288
		{
			yyVAL.opt_expr = yyDollar[1].expr
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:291
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(nil, yyDollar[1].token)
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:292
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(&yyDollar[1].token, yyDollar[2].token)
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:293
		{
			yyVAL.declarations_method_param = *(yyVAL.declarations_method_param.DefaultValue(yyDollar[3].simple_expr))
		}
	case 76:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:296
		{
			yyVAL.declarations_method_params = []ParamStatement{}
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:297
		{
			yyVAL.declarations_method_params = []ParamStatement{yyDollar[1].declarations_method_param}
		}
	case 78:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:298
		{
			yyVAL.declarations_method_params = append(yyDollar[1].declarations_method_params, yyDollar[3].declarations_method_param)
		}
	case 79:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:302
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal}
		}
	case 80:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:303
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal, Param: yyDollar[4].exprs}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:306
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:307
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:308
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:309
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:310
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:311
		{
			yyVAL.simple_expr = UndefinedStatement{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:312
		{
			yyVAL.simple_expr = yyDollar[1].goToLabel
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:313
		{
			if tok, ok := yyDollar[1].through_dot.(Token); ok {
				yyVAL.simple_expr = tok.literal
			} else {
				yyVAL.simple_expr = yyDollar[1].through_dot
			}
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:320
		{
			yyVAL.simple_expr = yyDollar[1].ternary
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:323
		{
			yyVAL.goToLabel = &GoToLabelStatement{Name: yyDollar[1].token.literal}
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:325
		{
			yyVAL.exprs = []Statement{yyDollar[1].opt_expr}
		}
	case 92:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:326
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[3].opt_expr)
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:329
		{
			yyVAL.identifiers = []Token{yyDollar[1].token}
		}
	case 94:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:330
		{
			yyVAL.identifiers = append(yyVAL.identifiers, yyDollar[3].token)
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:333
		{
			yyVAL.token = yyDollar[1].token
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:334
		{
			yyVAL.token = yyDollar[1].token
		}
	}
	goto yystack /* stack new state and value */
}
