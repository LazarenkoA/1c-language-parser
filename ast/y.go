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
	declarations_method_params []ParamStatement
	declarations_method_param  ParamStatement
	expr                       Statement
	opt_expr                   Statement
	exprs                      []Statement
	opt_export                 *Token
	opt_directive              *Token
	simple_expr                Statement
	new_object                 Statement
	ternary                    Statement
	explicit_variables         map[string]VarStatement
	global_variables           []GlobalVariables
	opt_explicit_variables     map[string]VarStatement
	identifiers                []Token
	identifier                 Statement
	goToLabel                  *GoToLabelStatement
	opt_goToLabel              *GoToLabelStatement
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

//line .\grammar.y:358

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 5,
}

const yyPrivate = 57344

const yyLast = 496

var yyAct = [...]uint8{
	77, 74, 167, 16, 76, 44, 71, 43, 72, 103,
	104, 105, 101, 102, 177, 165, 21, 53, 11, 18,
	100, 98, 99, 93, 94, 95, 96, 97, 30, 158,
	54, 22, 30, 18, 18, 103, 104, 105, 101, 102,
	113, 33, 145, 56, 18, 18, 100, 98, 99, 93,
	94, 95, 96, 97, 17, 62, 70, 73, 53, 92,
	29, 55, 91, 110, 18, 106, 107, 108, 95, 96,
	97, 18, 47, 20, 112, 114, 109, 19, 90, 93,
	94, 95, 96, 97, 121, 73, 123, 125, 126, 18,
	127, 31, 117, 122, 131, 132, 133, 134, 135, 136,
	137, 138, 139, 140, 141, 142, 143, 129, 128, 17,
	57, 130, 73, 147, 111, 73, 63, 180, 146, 155,
	46, 150, 190, 148, 25, 61, 103, 104, 105, 101,
	102, 10, 8, 188, 15, 179, 161, 100, 98, 99,
	93, 94, 95, 96, 97, 156, 73, 173, 36, 24,
	114, 159, 157, 160, 162, 35, 26, 9, 181, 175,
	166, 114, 60, 168, 34, 171, 172, 169, 50, 176,
	28, 115, 49, 27, 49, 89, 85, 182, 59, 152,
	185, 86, 12, 32, 23, 186, 187, 82, 66, 189,
	87, 46, 81, 88, 14, 49, 38, 37, 69, 13,
	6, 84, 83, 119, 22, 124, 22, 164, 39, 40,
	42, 66, 41, 68, 46, 2, 52, 7, 65, 38,
	37, 69, 163, 64, 178, 48, 49, 184, 51, 1,
	3, 39, 40, 42, 118, 41, 68, 116, 80, 58,
	45, 65, 66, 104, 105, 46, 64, 67, 48, 5,
	38, 37, 69, 100, 98, 99, 93, 94, 95, 96,
	97, 120, 39, 40, 42, 149, 41, 68, 174, 78,
	4, 79, 65, 75, 0, 0, 183, 64, 0, 48,
	103, 104, 105, 101, 102, 0, 0, 0, 0, 0,
	0, 100, 98, 99, 93, 94, 95, 96, 97, 0,
	0, 144, 103, 104, 105, 101, 102, 170, 0, 0,
	0, 0, 0, 100, 98, 99, 93, 94, 95, 96,
	97, 0, 0, 0, 0, 103, 104, 105, 101, 102,
	154, 0, 0, 0, 0, 0, 100, 98, 99, 93,
	94, 95, 96, 97, 0, 0, 0, 0, 103, 104,
	105, 101, 102, 153, 0, 0, 0, 0, 0, 100,
	98, 99, 93, 94, 95, 96, 97, 151, 0, 0,
	0, 0, 103, 104, 105, 101, 102, 0, 0, 0,
	0, 0, 0, 100, 98, 99, 93, 94, 95, 96,
	97, 0, 0, 103, 104, 105, 101, 102, 0, 0,
	0, 0, 0, 0, 100, 98, 99, 93, 94, 95,
	96, 97, 103, 104, 105, 101, 102, 49, 0, 0,
	0, 0, 0, 100, 98, 99, 93, 94, 95, 96,
	97, 0, 0, 103, 104, 105, 46, 102, 0, 0,
	0, 38, 37, 0, 100, 98, 99, 93, 94, 95,
	96, 97, 0, 39, 40, 42, 0, 41, 0, 0,
	103, 104, 105, 0, 0, 0, 0, 0, 0, 0,
	48, 100, 98, 99, 93, 94, 95, 96, 97, 105,
	100, 98, 99, 93, 94, 95, 96, 97, 100, 98,
	99, 93, 94, 95, 96, 97,
}

var yyPact = [...]int16{
	194, 194, -1000, -1000, -1000, 123, -1000, -1000, 175, 192,
	187, 8, -1000, 19, 15, 201, 177, -1000, -1000, 149,
	149, -1000, -1000, -1000, 1, 42, -1000, 176, -18, 63,
	149, 410, -1000, 63, -1000, 42, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -47, -1000, -1000, -31, 3, -15,
	169, 169, 167, -1000, 219, 219, 219, 165, 166, 175,
	165, -31, -3, -1000, 219, 219, 219, -1000, 94, 56,
	-29, -19, -1000, 374, 161, 199, -1000, 374, -1000, -1000,
	-1000, -1000, -1000, 219, 219, 219, 188, 219, 165, 175,
	26, 76, -1000, 219, 219, 219, 219, 219, 219, 219,
	219, 219, 219, 219, 219, 219, 242, -1000, 431, -1000,
	-16, 219, 219, -1000, 219, -1000, 165, -1000, -1000, -1000,
	-1000, 374, -1000, 355, 172, 334, 310, 90, 26, -1000,
	-1000, 14, 14, -1000, -1000, -1000, 27, 27, 431, 395,
	422, 204, 439, 431, -1000, 219, -30, -29, -1000, -1000,
	-1000, 165, 118, 219, -1000, -1000, -1000, -44, -1000, 219,
	150, 167, 287, 165, 165, -1000, 88, 145, 219, -6,
	-1000, 114, 87, -1000, 143, 165, 264, -1000, 165, -1000,
	-1000, -1000, -1000, 165, 165, 112, 150, 101, -1000, -1000,
	-1000,
}

var yyPgo = [...]int16{
	0, 273, 1, 4, 271, 270, 269, 2, 268, 265,
	261, 6, 0, 134, 249, 116, 149, 124, 8, 5,
	247, 240, 110, 239, 18, 238, 72, 7, 237, 16,
	234, 230, 229, 215, 228, 3, 227, 224, 222, 216,
	207,
}

var yyR1 = [...]int8{
	0, 32, 32, 33, 33, 14, 14, 13, 13, 31,
	34, 5, 5, 2, 2, 1, 1, 9, 9, 28,
	28, 22, 22, 23, 23, 6, 7, 7, 8, 8,
	21, 36, 4, 37, 4, 38, 4, 3, 3, 3,
	3, 3, 3, 3, 3, 10, 10, 19, 19, 26,
	26, 26, 40, 25, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 18, 18, 17, 17, 17, 16, 16,
	16, 20, 20, 20, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 27, 11, 11, 24, 24, 29, 30,
	35, 39,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 1, 0, 1, 0, 1, 5,
	0, 11, 10, 0, 1, 1, 3, 0, 1, 1,
	1, 0, 1, 3, 4, 7, 0, 5, 0, 2,
	8, 0, 9, 0, 8, 0, 6, 1, 1, 1,
	1, 1, 1, 2, 2, 0, 1, 1, 3, 1,
	4, 4, 0, 6, 1, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 2, 3, 3, 3, 3, 3,
	2, 1, 2, 0, 1, 1, 2, 3, 0, 1,
	3, 2, 5, 4, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 1,
}

var yyChk = [...]int16{
	-1000, -32, -33, -31, -5, -14, 6, -33, 9, 34,
	8, -24, 7, 7, 7, -13, -35, 46, 63, 58,
	58, -29, 5, 7, -16, -17, 7, 24, -16, 59,
	-35, 49, 7, 59, -13, -17, -15, 32, 31, 43,
	44, 47, 45, -27, -19, -21, 26, -26, 60, 7,
	-13, -34, -39, 64, 61, 58, 58, -22, -23, 9,
	-22, -26, -12, -15, 58, 53, 23, -20, 48, 33,
	-12, -11, -18, -12, -2, -1, -3, -12, -6, -4,
	-25, 27, 22, 37, 36, 11, 16, 25, 28, 9,
	-24, -2, 62, 52, 53, 54, 55, 56, 50, 51,
	49, 41, 42, 38, 39, 40, -12, -12, -12, -27,
	7, 58, -35, 59, -35, 10, -28, -29, -30, 4,
	-10, -12, -18, -12, 17, -12, -12, -2, -24, -29,
	35, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, -12, -12, -12, 59, 58, -11, -12, -18, -9,
	-3, 12, 7, 19, 20, 29, -29, -11, 59, -35,
	-2, 18, -12, -38, -40, 59, -12, -7, 13, -19,
	20, -2, -2, 59, -8, 14, -12, 20, -37, 21,
	30, 15, -2, 12, -36, -2, -2, -2, 21, -7,
	21,
}

var yyDef = [...]int8{
	5, -2, 1, 3, 4, 0, 6, 2, 0, 0,
	0, 7, 96, 0, 0, 0, 0, 8, 100, 78,
	78, 9, 98, 97, 0, 79, 75, 0, 0, 7,
	0, 0, 76, 7, 10, 80, 77, 84, 85, 86,
	87, 88, 89, 90, 91, 92, 93, 47, 0, 49,
	21, 21, 0, 101, 0, 0, 73, 13, 22, 0,
	13, 48, 0, 54, 0, 0, 0, 71, 0, 0,
	0, 0, 94, 74, 0, 14, 15, 37, 38, 39,
	40, 41, 42, 45, 73, 0, 0, 0, 13, 0,
	0, 0, 51, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 64, 70, 72,
	81, 73, 0, 50, 73, 12, 17, 19, 20, 99,
	43, 46, 44, 0, 0, 0, 0, 0, 0, 23,
	11, 56, 57, 58, 59, 60, 61, 62, 63, 65,
	66, 67, 68, 69, 55, 73, 0, 0, 95, 16,
	18, 13, 0, 0, 35, 52, 24, 0, 83, 0,
	26, 0, 0, 13, 13, 82, 0, 28, 0, 0,
	33, 0, 0, 30, 0, 13, 0, 31, 13, 36,
	53, 25, 29, 13, 13, 0, 26, 0, 34, 27,
	32,
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

	case 3:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:103
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].global_variables, yylex)
			}
		}
	case 4:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:108
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].funcProc, yylex)
			}
		}
	case 5:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:114
		{
			yyVAL.opt_directive = nil
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:115
		{
			yyVAL.opt_directive = &yyDollar[1].token
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:118
		{
			yyVAL.opt_export = nil
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:119
		{
			yyVAL.opt_export = &yyDollar[1].token
		}
	case 9:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:122
		{
			yyVAL.global_variables = make([]GlobalVariables, len(yyDollar[3].identifiers), len(yyDollar[3].identifiers))
			for i, v := range yyDollar[3].identifiers {
				if yyDollar[1].opt_directive != nil {
					yyVAL.global_variables[i].Directive = yyDollar[1].opt_directive.literal
				}

				yyVAL.global_variables[i].Export = yyDollar[4].opt_export != nil
				yyVAL.global_variables[i].Var = VarStatement{Name: v.literal}
			}
		}
	case 10:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:135
		{
			isFunction(true, yylex)
		}
	case 11:
		yyDollar = yyS[yypt-11 : yypt+1]
//line .\grammar.y:136
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeFunction, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[9].opt_explicit_variables, yyDollar[10].opt_body)
			isFunction(false, yylex)
		}
	case 12:
		yyDollar = yyS[yypt-10 : yypt+1]
//line .\grammar.y:141
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeProcedure, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[8].opt_explicit_variables, yyDollar[9].opt_body)
		}
	case 13:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:146
		{
			yyVAL.opt_body = nil
		}
	case 14:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:147
		{
			yyVAL.opt_body = yyDollar[1].body
		}
	case 15:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:150
		{
			yyVAL.body = []Statement{yyDollar[1].stmt}
		}
	case 16:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:151
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
	case 17:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:164
		{
			yyVAL.opt_stmt = nil
		}
	case 18:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:165
		{
			yyVAL.opt_stmt = yyDollar[1].stmt
		}
	case 19:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:168
		{
			yyVAL.token = yyDollar[1].token
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:168
		{
			yyVAL.token = yyDollar[1].token
		}
	case 21:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:172
		{
			yyVAL.opt_explicit_variables = map[string]VarStatement{}
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:173
		{
			yyVAL.opt_explicit_variables = yyDollar[1].explicit_variables
		}
	case 23:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:176
		{
			if vars, err := appendVarStatements(map[string]VarStatement{}, yyDollar[2].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 24:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:183
		{
			if vars, err := appendVarStatements(yyDollar[1].explicit_variables, yyDollar[3].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 25:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:194
		{
			yyVAL.stmt_if = &IfStatement{
				Expression:  yyDollar[2].expr,
				TrueBlock:   yyDollar[4].opt_body,
				IfElseBlock: yyDollar[5].opt_elseif_list,
				ElseBlock:   yyDollar[6].opt_else,
			}
		}
	case 26:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:204
		{
			yyVAL.opt_elseif_list = []Statement{}
		}
	case 27:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:205
		{
			yyVAL.opt_elseif_list = append(yyDollar[5].opt_elseif_list, &IfStatement{
				Expression: yyDollar[2].expr,
				TrueBlock:  yyDollar[4].opt_body,
			})
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:213
		{
			yyVAL.opt_else = nil
		}
	case 29:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:214
		{
			yyVAL.opt_else = yyDollar[2].opt_body
		}
	case 30:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:217
		{
			yyVAL.ternary = TernaryStatement{
				Expression: yyDollar[3].expr,
				TrueBlock:  yyDollar[5].expr,
				ElseBlock:  yyDollar[7].expr,
			}
		}
	case 31:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:226
		{
			setLoopFlag(true, yylex)
		}
	case 32:
		yyDollar = yyS[yypt-9 : yypt+1]
//line .\grammar.y:226
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[3].token.literal,
				In:   yyDollar[5].through_dot,
				Body: yyDollar[8].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 33:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:234
		{
			setLoopFlag(true, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:234
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[2].expr,
				To:   yyDollar[4].expr,
				Body: yyDollar[7].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:242
		{
			setLoopFlag(true, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:242
		{
			yyVAL.stmt_loop = &LoopStatement{
				WhileExpr: yyDollar[2].expr,
				Body:      yyDollar[5].opt_body,
			}
		}
	case 37:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:249
		{
			yyVAL.stmt = yyDollar[1].expr
		}
	case 38:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:250
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:251
		{
			yyVAL.stmt = yyDollar[1].stmt_loop
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:252
		{
			yyVAL.stmt = yyDollar[1].stmt_tryCatch
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:253
		{
			yyVAL.stmt = ContinueStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:254
		{
			yyVAL.stmt = BreakStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 43:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:255
		{
			yyVAL.stmt = ThrowStatement{Param: yyDollar[2].opt_param}
			checkThrowParam(yyDollar[1].token, yyDollar[2].opt_param, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:256
		{
			yyVAL.stmt = &ReturnStatement{Param: yyDollar[2].opt_expr}
			checkReturnParam(yyDollar[2].opt_expr, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:259
		{
			yyVAL.opt_param = nil
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:260
		{
			yyVAL.opt_param = yyDollar[1].expr
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:265
		{
			yyVAL.through_dot = yyDollar[1].identifier
		}
	case 48:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:266
		{
			yyVAL.through_dot = CallChainStatement{Unit: yyDollar[3].identifier, Call: yyDollar[1].through_dot}
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:269
		{
			yyVAL.identifier = VarStatement{Name: yyDollar[1].token.literal}
		}
	case 50:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:270
		{
			yyVAL.identifier = MethodStatement{Name: yyDollar[1].token.literal, Param: yyDollar[3].exprs}
		}
	case 51:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:271
		{
			yyVAL.identifier = ItemStatement{Object: yyDollar[1].identifier, Item: yyDollar[3].expr}
		}
	case 52:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:276
		{
			setTryFlag(true, yylex)
		}
	case 53:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:276
		{
			yyVAL.stmt_tryCatch = TryStatement{Body: yyDollar[2].opt_body, Catch: yyDollar[5].opt_body}
			setTryFlag(false, yylex)
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:282
		{
			yyVAL.expr = yyDollar[1].simple_expr
		}
	case 55:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:283
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 56:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:284
		{
			yyVAL.expr = &ExpStatement{Operation: OpPlus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:285
		{
			yyVAL.expr = &ExpStatement{Operation: OpMinus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:286
		{
			yyVAL.expr = &ExpStatement{Operation: OpMul, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:287
		{
			yyVAL.expr = &ExpStatement{Operation: OpDiv, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:288
		{
			yyVAL.expr = &ExpStatement{Operation: OpMod, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:289
		{
			yyVAL.expr = &ExpStatement{Operation: OpGt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:290
		{
			yyVAL.expr = &ExpStatement{Operation: OpLt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:291
		{
			yyVAL.expr = &ExpStatement{Operation: OpEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:292
		{
			yyVAL.expr = unary(yyDollar[2].expr)
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:293
		{
			yyVAL.expr = &ExpStatement{Operation: OpOr, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:294
		{
			yyVAL.expr = &ExpStatement{Operation: OpAnd, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:295
		{
			yyVAL.expr = &ExpStatement{Operation: OpNe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:296
		{
			yyVAL.expr = &ExpStatement{Operation: OpLe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:297
		{
			yyVAL.expr = &ExpStatement{Operation: OpGe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:298
		{
			yyVAL.expr = not(yyDollar[2].expr)
		}
	case 71:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:299
		{
			yyVAL.expr = yyDollar[1].new_object
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:300
		{
			yyVAL.expr = GoToStatement{Label: yyDollar[2].goToLabel}
		}
	case 73:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:303
		{
			yyVAL.opt_expr = nil
		}
	case 74:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:303
		{
			yyVAL.opt_expr = yyDollar[1].expr
		}
	case 75:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:306
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(nil, yyDollar[1].token)
		}
	case 76:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:307
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(&yyDollar[1].token, yyDollar[2].token)
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:308
		{
			yyVAL.declarations_method_param = *(yyVAL.declarations_method_param.DefaultValue(yyDollar[3].simple_expr))
		}
	case 78:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:311
		{
			yyVAL.declarations_method_params = []ParamStatement{}
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:312
		{
			yyVAL.declarations_method_params = []ParamStatement{yyDollar[1].declarations_method_param}
		}
	case 80:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:313
		{
			yyVAL.declarations_method_params = append(yyDollar[1].declarations_method_params, yyDollar[3].declarations_method_param)
		}
	case 81:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:321
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal}
		}
	case 82:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:322
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal, Param: yyDollar[4].exprs}
		}
	case 83:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:323
		{
			yyVAL.new_object = NewObjectStatement{Param: yyDollar[3].exprs}
		}
	case 84:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:326
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 85:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:327
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:328
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:329
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:330
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:331
		{
			yyVAL.simple_expr = UndefinedStatement{}
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:332
		{
			yyVAL.simple_expr = yyDollar[1].goToLabel
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:333
		{
			if tok, ok := yyDollar[1].through_dot.(Token); ok {
				yyVAL.simple_expr = tok.literal
			} else {
				yyVAL.simple_expr = yyDollar[1].through_dot
			}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:340
		{
			yyVAL.simple_expr = yyDollar[1].ternary
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:343
		{
			yyVAL.goToLabel = &GoToLabelStatement{Name: yyDollar[1].token.literal}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:345
		{
			yyVAL.exprs = []Statement{yyDollar[1].opt_expr}
		}
	case 95:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:346
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[3].opt_expr)
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:349
		{
			yyVAL.identifiers = []Token{yyDollar[1].token}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:350
		{
			yyVAL.identifiers = append(yyVAL.identifiers, yyDollar[3].token)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:353
		{
			yyVAL.token = yyDollar[1].token
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:354
		{
			yyVAL.token = yyDollar[1].token
		}
	}
	goto yystack /* stack new state and value */
}
