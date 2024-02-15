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

//line .\grammar.y:371

//line yacctab:1
var yyExca = [...]int8{
	-1, 1,
	1, -1,
	-2, 0,
	-1, 3,
	1, 15,
	-2, 7,
}

const yyPrivate = 57344

const yyLast = 553

var yyAct = [...]uint8{
	6, 47, 43, 178, 107, 16, 143, 135, 123, 113,
	144, 124, 114, 64, 66, 45, 34, 39, 67, 68,
	69, 80, 83, 71, 73, 74, 65, 75, 60, 61,
	62, 58, 59, 162, 158, 128, 127, 126, 126, 57,
	55, 56, 50, 51, 52, 53, 54, 154, 160, 181,
	81, 88, 89, 90, 91, 92, 93, 94, 95, 96,
	97, 98, 99, 100, 60, 61, 62, 58, 59, 112,
	125, 85, 84, 126, 41, 57, 55, 56, 50, 51,
	52, 53, 54, 66, 115, 116, 66, 126, 148, 126,
	134, 80, 126, 38, 126, 117, 130, 111, 28, 27,
	126, 82, 125, 192, 118, 120, 52, 53, 54, 157,
	29, 30, 32, 66, 31, 50, 51, 52, 53, 54,
	4, 106, 129, 139, 140, 131, 141, 40, 133, 131,
	78, 76, 66, 38, 150, 147, 137, 153, 183, 33,
	176, 131, 156, 145, 149, 57, 55, 56, 50, 51,
	52, 53, 54, 119, 165, 159, 77, 168, 164, 159,
	146, 70, 163, 87, 173, 152, 171, 169, 174, 175,
	170, 172, 136, 189, 186, 180, 108, 161, 37, 41,
	185, 184, 182, 22, 142, 187, 188, 41, 23, 110,
	191, 190, 109, 193, 11, 19, 103, 24, 38, 10,
	25, 46, 45, 28, 27, 36, 45, 5, 13, 12,
	122, 48, 49, 2, 79, 29, 30, 32, 41, 31,
	21, 121, 22, 155, 167, 18, 177, 23, 3, 1,
	17, 14, 40, 11, 19, 44, 24, 38, 10, 25,
	42, 9, 28, 27, 36, 179, 35, 13, 12, 20,
	26, 63, 86, 151, 29, 30, 32, 7, 31, 21,
	15, 8, 0, 0, 18, 0, 0, 0, 0, 17,
	0, 40, 60, 61, 62, 58, 59, 0, 0, 0,
	0, 41, 0, 57, 55, 56, 50, 51, 52, 53,
	54, 72, 0, 0, 0, 0, 132, 19, 61, 62,
	38, 0, 0, 0, 41, 28, 27, 36, 57, 55,
	56, 50, 51, 52, 53, 54, 0, 29, 30, 32,
	19, 31, 21, 38, 0, 0, 0, 18, 28, 27,
	36, 0, 17, 0, 40, 0, 0, 0, 0, 0,
	29, 30, 32, 0, 31, 21, 0, 0, 0, 0,
	18, 0, 0, 0, 166, 17, 0, 40, 60, 61,
	62, 58, 59, 0, 0, 0, 0, 0, 0, 57,
	55, 56, 50, 51, 52, 53, 54, 0, 0, 101,
	60, 61, 62, 58, 59, 138, 0, 0, 0, 0,
	0, 57, 55, 56, 50, 51, 52, 53, 54, 0,
	0, 0, 0, 60, 61, 62, 58, 59, 105, 0,
	0, 0, 0, 0, 57, 55, 56, 50, 51, 52,
	53, 54, 0, 0, 0, 0, 60, 61, 62, 58,
	59, 104, 0, 0, 0, 0, 0, 57, 55, 56,
	50, 51, 52, 53, 54, 102, 0, 0, 0, 0,
	60, 61, 62, 58, 59, 0, 0, 0, 0, 0,
	0, 57, 55, 56, 50, 51, 52, 53, 54, 0,
	0, 60, 61, 62, 58, 59, 0, 0, 0, 0,
	0, 0, 57, 55, 56, 50, 51, 52, 53, 54,
	60, 61, 62, 58, 59, 0, 0, 0, 0, 0,
	0, 57, 55, 56, 50, 51, 52, 53, 54, 60,
	61, 62, 0, 59, 0, 0, 0, 60, 61, 62,
	57, 55, 56, 50, 51, 52, 53, 54, 57, 55,
	56, 50, 51, 52, 53, 54, 62, 0, 0, 0,
	0, 0, 0, 0, 0, 57, 55, 56, 50, 51,
	52, 53, 54,
}

var yyPact = [...]int16{
	172, -1000, 197, 172, -1000, -1000, 452, -1000, -1000, -1000,
	-1000, -1000, 297, 297, -1000, -1000, -1000, 297, 297, 297,
	-1000, 107, 297, 274, 297, 211, 122, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -43, -1000, 43, -1000, -1000, -39,
	14, 13, 211, -1000, -1000, -1000, -1000, -1000, -1000, 197,
	297, 297, 297, 297, 297, 297, 297, 297, 297, 297,
	297, 297, 297, -1000, 452, -1000, 452, 320, -1000, 96,
	-1000, 433, 189, 412, 388, 92, 169, 185, 182, 180,
	-1000, 11, 297, 297, 297, 297, -1000, -1000, 52, 52,
	-1000, -1000, -1000, 63, 63, 96, 471, 479, 259, 496,
	96, -1000, 211, 135, 297, -1000, -1000, 24, -1000, -22,
	-23, -39, 297, 37, -1000, 234, 26, 31, 159, 180,
	365, 211, 211, 201, 177, -1000, -1000, 136, 136, 29,
	-1000, 297, -1000, 297, -1000, 151, 297, 27, -1000, 121,
	79, -1000, -1000, -25, -1, -1000, 170, -26, -1000, -1000,
	26, 143, 211, 342, -1000, 211, -1000, -1000, 56, 136,
	67, -1000, 56, 297, -1000, -1000, 211, 211, 119, -1000,
	-1, -1000, 166, -10, 159, 117, -1000, 166, 211, 165,
	169, -1000, -1000, -1000, 211, 163, 169, 10, 68, -1000,
	10, -1000, -1000, -1000,
}

var yyPgo = [...]int16{
	0, 212, 1, 120, 261, 260, 257, 7, 253, 252,
	251, 9, 0, 8, 250, 5, 6, 10, 12, 16,
	249, 246, 3, 245, 4, 241, 17, 139, 240, 2,
	235, 231, 229, 228, 207, 226, 11, 224, 223, 221,
	214, 210,
}

var yyR1 = [...]int8{
	0, 32, 32, 33, 33, 34, 34, 14, 14, 13,
	13, 31, 35, 5, 5, 2, 2, 1, 1, 9,
	9, 28, 28, 22, 22, 23, 23, 6, 7, 7,
	8, 8, 21, 37, 4, 38, 4, 39, 4, 3,
	3, 3, 3, 3, 3, 3, 3, 10, 10, 19,
	19, 26, 26, 26, 41, 25, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 18, 18, 17, 17, 17,
	16, 16, 16, 20, 20, 20, 15, 15, 15, 15,
	15, 15, 15, 15, 15, 27, 11, 11, 24, 24,
	29, 30, 36, 40,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 2, 1, 1, 0, 1, 0,
	1, 5, 0, 11, 10, 0, 1, 1, 3, 0,
	1, 1, 1, 0, 1, 3, 4, 7, 0, 5,
	0, 2, 8, 0, 9, 0, 8, 0, 6, 1,
	1, 1, 1, 1, 1, 2, 2, 0, 1, 1,
	3, 1, 4, 4, 0, 6, 1, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 2, 3, 3, 3,
	3, 3, 2, 1, 2, 0, 1, 1, 2, 3,
	0, 1, 3, 2, 5, 4, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 3, 1, 3,
	1, 1, 1, 1,
}

var yyChk = [...]int16{
	-1000, -32, -1, -33, -3, -34, -12, -6, -4, -25,
	27, 22, 37, 36, -31, -5, -15, 58, 53, 23,
	-20, 48, 11, 16, 25, 28, -14, 32, 31, 43,
	44, 47, 45, -27, -19, -21, 33, 6, 26, -26,
	60, 7, -28, -29, -30, 5, 4, -2, -34, -1,
	52, 53, 54, 55, 56, 50, 51, 49, 41, 42,
	38, 39, 40, -10, -12, -18, -12, -12, -12, -12,
	-27, -12, 17, -12, -12, -2, 9, 34, 8, -40,
	64, 7, 58, 61, 58, 58, -9, -3, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, -12,
	-12, 59, 12, 7, 19, 20, 29, -24, 7, 7,
	7, -26, 58, -11, -18, -12, -12, -11, -2, 18,
	-12, -39, -41, -13, -36, 46, 63, 58, 58, -11,
	59, -36, 62, -36, 59, -7, 13, -19, 20, -2,
	-2, -29, 7, -16, -17, 7, 24, -16, 59, -18,
	-12, -8, 14, -12, 20, -38, 21, 30, 59, -36,
	49, 7, 59, -36, 15, -2, 12, -37, -2, -13,
	-17, -15, -13, -12, -2, -2, 21, -35, -22, -23,
	9, 59, -7, 21, -22, -2, 9, -24, -2, 10,
	-24, -29, 35, -29,
}

var yyDef = [...]int8{
	7, -2, 1, -2, 17, 3, 39, 40, 41, 42,
	43, 44, 47, 75, 5, 6, 56, 0, 0, 0,
	73, 0, 0, 0, 0, 15, 0, 86, 87, 88,
	89, 90, 91, 92, 93, 94, 0, 8, 95, 49,
	0, 51, 19, 21, 22, 100, 101, 2, 4, 16,
	0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 45, 48, 46, 76, 0, 66, 72,
	74, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	103, 83, 75, 0, 0, 75, 18, 20, 58, 59,
	60, 61, 62, 63, 64, 65, 67, 68, 69, 70,
	71, 57, 15, 0, 0, 37, 54, 9, 98, 0,
	0, 50, 75, 0, 96, 0, 0, 0, 28, 0,
	0, 15, 15, 0, 0, 10, 102, 80, 80, 0,
	85, 75, 53, 0, 52, 30, 0, 0, 35, 0,
	0, 11, 99, 0, 81, 77, 0, 0, 84, 97,
	0, 0, 15, 0, 33, 15, 38, 55, 9, 0,
	0, 78, 9, 0, 27, 31, 15, 15, 0, 12,
	82, 79, 23, 0, 28, 0, 36, 23, 15, 24,
	0, 32, 29, 34, 15, 0, 0, 0, 0, 14,
	0, 25, 13, 26,
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
//line .\grammar.y:99
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].body, yylex)
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:104
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[2].opt_body, yylex)
			}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:114
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].global_variables, yylex)
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:119
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].funcProc, yylex)
			}
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:126
		{
			yyVAL.opt_directive = nil
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:127
		{
			yyVAL.opt_directive = &yyDollar[1].token
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:130
		{
			yyVAL.opt_export = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:131
		{
			yyVAL.opt_export = &yyDollar[1].token
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:134
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
	case 12:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:147
		{
			isFunction(true, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-11 : yypt+1]
//line .\grammar.y:148
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeFunction, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[9].opt_explicit_variables, yyDollar[10].opt_body)
			isFunction(false, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-10 : yypt+1]
//line .\grammar.y:153
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeProcedure, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[8].opt_explicit_variables, yyDollar[9].opt_body)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:158
		{
			yyVAL.opt_body = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:159
		{
			yyVAL.opt_body = yyDollar[1].body
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:163
		{
			yyVAL.body = []Statement{yyDollar[1].stmt}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:164
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
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:177
		{
			yyVAL.opt_stmt = nil
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:178
		{
			yyVAL.opt_stmt = yyDollar[1].stmt
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:181
		{
			yyVAL.token = yyDollar[1].token
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:181
		{
			yyVAL.token = yyDollar[1].token
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:185
		{
			yyVAL.opt_explicit_variables = map[string]VarStatement{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:186
		{
			yyVAL.opt_explicit_variables = yyDollar[1].explicit_variables
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:189
		{
			if vars, err := appendVarStatements(map[string]VarStatement{}, yyDollar[2].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:196
		{
			if vars, err := appendVarStatements(yyDollar[1].explicit_variables, yyDollar[3].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:207
		{
			yyVAL.stmt_if = &IfStatement{
				Expression:  yyDollar[2].expr,
				TrueBlock:   yyDollar[4].opt_body,
				IfElseBlock: yyDollar[5].opt_elseif_list,
				ElseBlock:   yyDollar[6].opt_else,
			}
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:217
		{
			yyVAL.opt_elseif_list = []Statement{}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:218
		{
			yyVAL.opt_elseif_list = append(yyDollar[5].opt_elseif_list, &IfStatement{
				Expression: yyDollar[2].expr,
				TrueBlock:  yyDollar[4].opt_body,
			})
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:226
		{
			yyVAL.opt_else = nil
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:227
		{
			yyVAL.opt_else = yyDollar[2].opt_body
		}
	case 32:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:230
		{
			yyVAL.ternary = TernaryStatement{
				Expression: yyDollar[3].expr,
				TrueBlock:  yyDollar[5].expr,
				ElseBlock:  yyDollar[7].expr,
			}
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:239
		{
			setLoopFlag(true, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-9 : yypt+1]
//line .\grammar.y:239
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[3].token.literal,
				In:   yyDollar[5].through_dot,
				Body: yyDollar[8].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:247
		{
			setLoopFlag(true, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:247
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[2].expr,
				To:   yyDollar[4].expr,
				Body: yyDollar[7].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:255
		{
			setLoopFlag(true, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:255
		{
			yyVAL.stmt_loop = &LoopStatement{
				WhileExpr: yyDollar[2].expr,
				Body:      yyDollar[5].opt_body,
			}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:262
		{
			yyVAL.stmt = yyDollar[1].expr
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:263
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 41:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:264
		{
			yyVAL.stmt = yyDollar[1].stmt_loop
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:265
		{
			yyVAL.stmt = yyDollar[1].stmt_tryCatch
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:266
		{
			yyVAL.stmt = ContinueStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:267
		{
			yyVAL.stmt = BreakStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 45:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:268
		{
			yyVAL.stmt = ThrowStatement{Param: yyDollar[2].opt_param}
			checkThrowParam(yyDollar[1].token, yyDollar[2].opt_param, yylex)
		}
	case 46:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:269
		{
			yyVAL.stmt = &ReturnStatement{Param: yyDollar[2].opt_expr}
			checkReturnParam(yyDollar[2].opt_expr, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:272
		{
			yyVAL.opt_param = nil
		}
	case 48:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:273
		{
			yyVAL.opt_param = yyDollar[1].expr
		}
	case 49:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:278
		{
			yyVAL.through_dot = yyDollar[1].identifier
		}
	case 50:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:279
		{
			yyVAL.through_dot = CallChainStatement{Unit: yyDollar[3].identifier, Call: yyDollar[1].through_dot}
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:282
		{
			yyVAL.identifier = VarStatement{Name: yyDollar[1].token.literal}
		}
	case 52:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:283
		{
			yyVAL.identifier = MethodStatement{Name: yyDollar[1].token.literal, Param: yyDollar[3].exprs}
		}
	case 53:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:284
		{
			yyVAL.identifier = ItemStatement{Object: yyDollar[1].identifier, Item: yyDollar[3].expr}
		}
	case 54:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:289
		{
			setTryFlag(true, yylex)
		}
	case 55:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:289
		{
			yyVAL.stmt_tryCatch = TryStatement{Body: yyDollar[2].opt_body, Catch: yyDollar[5].opt_body}
			setTryFlag(false, yylex)
		}
	case 56:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:295
		{
			yyVAL.expr = yyDollar[1].simple_expr
		}
	case 57:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:296
		{
			yyVAL.expr = yyDollar[2].expr
		}
	case 58:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:297
		{
			yyVAL.expr = &ExpStatement{Operation: OpPlus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 59:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:298
		{
			yyVAL.expr = &ExpStatement{Operation: OpMinus, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 60:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:299
		{
			yyVAL.expr = &ExpStatement{Operation: OpMul, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:300
		{
			yyVAL.expr = &ExpStatement{Operation: OpDiv, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 62:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:301
		{
			yyVAL.expr = &ExpStatement{Operation: OpMod, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 63:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:302
		{
			yyVAL.expr = &ExpStatement{Operation: OpGt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:303
		{
			yyVAL.expr = &ExpStatement{Operation: OpLt, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:304
		{
			yyVAL.expr = &ExpStatement{Operation: OpEq, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 66:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:305
		{
			yyVAL.expr = unary(yyDollar[2].expr)
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:306
		{
			yyVAL.expr = &ExpStatement{Operation: OpOr, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:307
		{
			yyVAL.expr = &ExpStatement{Operation: OpAnd, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:308
		{
			yyVAL.expr = &ExpStatement{Operation: OpNe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:309
		{
			yyVAL.expr = &ExpStatement{Operation: OpLe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:310
		{
			yyVAL.expr = &ExpStatement{Operation: OpGe, Left: yyDollar[1].expr, Right: yyDollar[3].expr}
		}
	case 72:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:311
		{
			yyVAL.expr = not(yyDollar[2].expr)
		}
	case 73:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:312
		{
			yyVAL.expr = yyDollar[1].new_object
		}
	case 74:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:313
		{
			yyVAL.expr = GoToStatement{Label: yyDollar[2].goToLabel}
		}
	case 75:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:316
		{
			yyVAL.opt_expr = nil
		}
	case 76:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:316
		{
			yyVAL.opt_expr = yyDollar[1].expr
		}
	case 77:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:319
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(nil, yyDollar[1].token)
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:320
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(&yyDollar[1].token, yyDollar[2].token)
		}
	case 79:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:321
		{
			yyVAL.declarations_method_param = *(yyVAL.declarations_method_param.DefaultValue(yyDollar[3].simple_expr))
		}
	case 80:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:324
		{
			yyVAL.declarations_method_params = []ParamStatement{}
		}
	case 81:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:325
		{
			yyVAL.declarations_method_params = []ParamStatement{yyDollar[1].declarations_method_param}
		}
	case 82:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:326
		{
			yyVAL.declarations_method_params = append(yyDollar[1].declarations_method_params, yyDollar[3].declarations_method_param)
		}
	case 83:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:334
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal}
		}
	case 84:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:335
		{
			yyVAL.new_object = NewObjectStatement{Constructor: yyDollar[2].token.literal, Param: yyDollar[4].exprs}
		}
	case 85:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:336
		{
			yyVAL.new_object = NewObjectStatement{Param: yyDollar[3].exprs}
		}
	case 86:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:339
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:340
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 88:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:341
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 89:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:342
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 90:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:343
		{
			yyVAL.simple_expr = yyDollar[1].token.value
		}
	case 91:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:344
		{
			yyVAL.simple_expr = UndefinedStatement{}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:345
		{
			yyVAL.simple_expr = yyDollar[1].goToLabel
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:346
		{
			if tok, ok := yyDollar[1].through_dot.(Token); ok {
				yyVAL.simple_expr = tok.literal
			} else {
				yyVAL.simple_expr = yyDollar[1].through_dot
			}
		}
	case 94:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:353
		{
			yyVAL.simple_expr = yyDollar[1].ternary
		}
	case 95:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:356
		{
			yyVAL.goToLabel = &GoToLabelStatement{Name: yyDollar[1].token.literal}
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:358
		{
			yyVAL.exprs = []Statement{yyDollar[1].opt_expr}
		}
	case 97:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:359
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[3].opt_expr)
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:362
		{
			yyVAL.identifiers = []Token{yyDollar[1].token}
		}
	case 99:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:363
		{
			yyVAL.identifiers = append(yyVAL.identifiers, yyDollar[3].token)
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:366
		{
			yyVAL.token = yyDollar[1].token
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:367
		{
			yyVAL.token = yyDollar[1].token
		}
	}
	goto yystack /* stack new state and value */
}
