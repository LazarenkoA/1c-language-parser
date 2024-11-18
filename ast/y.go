// Code generated by goyacc .\grammar.y. DO NOT EDIT.

//line .\grammar.y:1

package ast

import __yyfmt__ "fmt"

//line .\grammar.y:3

//line .\grammar.y:43
type yySymType struct {
	stmt                       Statement
	opt_directive              *Token
	goToLabel                  *GoToLabelStatement
	opt_export                 *Token
	opt_goToLabel              *GoToLabelStatement
	stmt_if                    *IfStatement
	stmt_loop                  *LoopStatement
	funcProc                   *FunctionOrProcedure
	opt_explicit_variables     map[string]VarStatement
	explicit_variables         map[string]VarStatement
	declarations_method_param  ParamStatement
	body                       []Statement
	exprs                      []Statement
	opt_elseif_list            []Statement
	declarations_method_params []ParamStatement
	global_variables           []GlobalVariables
	opt_body                   []Statement
	identifiers                []Token
	opt_else                   []Statement
	token                      Token
	yys                        int
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
const Execute = 57389
const UNARMinus = 57390
const UNARYPlus = 57391

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
	"Execute",
	"'='",
	"'>'",
	"'<'",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"'%'",
	"UNARMinus",
	"UNARYPlus",
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

//line .\grammar.y:379

//line yacctab:1
var yyExca = [...]int8{
	-1, 0,
	4, 15,
	5, 15,
	-2, 7,
	-1, 1,
	1, -1,
	-2, 0,
	-1, 2,
	1, 1,
	-2, 16,
	-1, 3,
	8, 7,
	9, 7,
	34, 7,
	-2, 15,
}

const yyPrivate = 57344

const yyLast = 605

var yyAct = [...]uint8{
	7, 5, 49, 193, 115, 45, 157, 145, 132, 17,
	20, 122, 133, 36, 67, 69, 156, 41, 51, 70,
	71, 84, 87, 73, 75, 76, 68, 77, 121, 182,
	81, 82, 63, 64, 65, 61, 62, 176, 172, 161,
	137, 135, 135, 135, 60, 58, 59, 53, 54, 55,
	56, 57, 136, 120, 96, 97, 98, 99, 100, 101,
	102, 103, 104, 105, 106, 107, 108, 63, 64, 65,
	61, 62, 143, 139, 134, 89, 135, 135, 135, 60,
	58, 59, 53, 54, 55, 56, 57, 69, 123, 124,
	69, 93, 126, 135, 135, 53, 54, 55, 56, 57,
	85, 119, 60, 58, 59, 53, 54, 55, 56, 57,
	88, 174, 127, 129, 65, 134, 92, 43, 125, 4,
	38, 69, 43, 40, 60, 58, 59, 53, 54, 55,
	56, 57, 152, 153, 140, 154, 35, 142, 140, 149,
	167, 69, 148, 163, 91, 52, 51, 166, 38, 138,
	128, 140, 162, 86, 160, 55, 56, 57, 72, 44,
	158, 168, 80, 78, 44, 52, 51, 179, 95, 173,
	114, 183, 178, 173, 165, 150, 177, 159, 188, 201,
	185, 184, 189, 190, 186, 187, 52, 51, 79, 195,
	116, 171, 175, 39, 43, 200, 199, 197, 22, 155,
	202, 203, 118, 23, 117, 206, 205, 111, 208, 12,
	19, 51, 24, 40, 11, 25, 131, 207, 28, 27,
	38, 52, 51, 14, 13, 52, 51, 52, 51, 6,
	31, 32, 34, 46, 33, 21, 44, 83, 198, 43,
	30, 29, 191, 22, 170, 52, 51, 18, 23, 42,
	130, 204, 52, 51, 12, 19, 169, 24, 40, 11,
	25, 146, 181, 28, 27, 38, 52, 51, 14, 13,
	47, 2, 43, 192, 3, 31, 32, 34, 1, 33,
	21, 44, 74, 15, 50, 30, 29, 48, 19, 10,
	194, 40, 18, 37, 42, 147, 28, 27, 38, 90,
	26, 66, 94, 164, 43, 8, 16, 9, 31, 32,
	34, 0, 33, 21, 44, 0, 0, 0, 30, 29,
	19, 64, 65, 40, 0, 18, 0, 42, 28, 27,
	38, 0, 60, 58, 59, 53, 54, 55, 56, 57,
	31, 32, 34, 0, 33, 21, 44, 0, 0, 0,
	30, 29, 63, 64, 65, 61, 62, 18, 0, 42,
	0, 0, 0, 0, 60, 58, 59, 53, 54, 55,
	56, 57, 63, 64, 65, 61, 62, 0, 141, 0,
	0, 0, 0, 0, 60, 58, 59, 53, 54, 55,
	56, 57, 0, 0, 0, 196, 63, 64, 65, 61,
	62, 0, 0, 0, 0, 0, 0, 0, 60, 58,
	59, 53, 54, 55, 56, 57, 0, 0, 0, 144,
	63, 64, 65, 61, 62, 43, 0, 0, 0, 0,
	0, 0, 60, 58, 59, 53, 54, 55, 56, 57,
	63, 64, 65, 109, 40, 0, 0, 180, 0, 28,
	27, 0, 60, 58, 59, 53, 54, 55, 56, 57,
	0, 31, 32, 34, 0, 33, 0, 44, 0, 0,
	0, 30, 29, 63, 64, 65, 61, 62, 151, 0,
	42, 0, 0, 0, 0, 60, 58, 59, 53, 54,
	55, 56, 57, 0, 0, 0, 63, 64, 65, 61,
	62, 113, 0, 0, 0, 0, 0, 0, 60, 58,
	59, 53, 54, 55, 56, 57, 0, 0, 0, 63,
	64, 65, 61, 62, 112, 0, 0, 0, 0, 0,
	0, 60, 58, 59, 53, 54, 55, 56, 57, 110,
	0, 0, 0, 63, 64, 65, 61, 62, 0, 0,
	0, 0, 0, 0, 0, 60, 58, 59, 53, 54,
	55, 56, 57, 0, 0, 63, 64, 65, 61, 62,
	0, 0, 0, 0, 0, 0, 0, 60, 58, 59,
	53, 54, 55, 56, 57, 63, 64, 65, 0, 62,
	0, 0, 0, 0, 0, 0, 0, 60, 58, 59,
	53, 54, 55, 56, 57,
}

var yyPact = [...]int16{
	187, -1000, -1000, 187, -1000, 262, -1000, -6, -1000, -1000,
	-1000, -1000, -1000, 297, 297, -1000, -1000, -1000, 297, 297,
	-1000, 97, 297, 265, 297, 232, 154, -1000, -1000, 297,
	297, -1000, -1000, -1000, -1000, -1000, -45, -1000, 93, -1000,
	-1000, -41, 50, 15, 84, 262, -1000, -1000, 232, -1000,
	-1000, -1000, -1000, 297, 297, 297, 297, 297, 297, 297,
	297, 297, 297, 297, 297, 297, -1000, -6, -1000, -6,
	382, 52, -1000, 527, 200, 505, 481, 141, 183, 197,
	195, -1000, -1000, 110, -1000, -7, 297, 297, 297, 297,
	-1000, 297, -1000, -1000, -1000, -1000, 100, 100, -1000, -1000,
	-1000, 42, 42, 52, 547, 402, 282, 74, 52, -1000,
	232, 132, 297, -1000, -1000, 28, -1000, -8, -20, -41,
	297, 12, -1000, 314, 29, 11, 358, 248, 115, 458,
	232, 232, 206, 192, -1000, -1000, 153, 153, -22, -1000,
	297, -1000, 297, -1000, -1000, 160, 297, 120, -45, -1000,
	87, -1000, 223, 161, -1000, -1000, -23, 61, -1000, 185,
	-24, -1000, -1000, 29, 157, 232, 435, -1000, -32, 232,
	-1000, -1000, 69, 153, 418, -1000, 69, 297, -1000, 262,
	232, 232, -1000, 221, -1000, 61, -1000, 180, 334, 248,
	217, -1000, 180, 232, 170, 183, -1000, -1000, -1000, 232,
	241, 183, 13, 182, -1000, 13, -1000, -1000, -1000,
}

var yyPgo = [...]int16{
	0, 270, 1, 119, 307, 306, 305, 7, 303, 302,
	301, 28, 0, 8, 300, 9, 16, 6, 11, 299,
	13, 295, 10, 293, 3, 290, 4, 289, 17, 136,
	287, 2, 284, 283, 278, 274, 229, 273, 12, 262,
	256, 250, 237, 216,
}

var yyR1 = [...]int8{
	0, 34, 34, 35, 35, 36, 36, 14, 14, 13,
	13, 33, 37, 5, 5, 2, 2, 1, 1, 9,
	9, 30, 30, 24, 24, 25, 25, 6, 7, 7,
	8, 8, 23, 39, 4, 40, 4, 41, 4, 21,
	21, 21, 3, 3, 3, 3, 3, 3, 3, 3,
	10, 10, 20, 20, 28, 28, 28, 28, 28, 19,
	19, 43, 27, 12, 12, 12, 12, 12, 12, 12,
	12, 12, 12, 12, 12, 12, 12, 12, 12, 12,
	12, 18, 18, 17, 17, 17, 16, 16, 16, 22,
	22, 22, 15, 15, 15, 15, 15, 15, 15, 15,
	15, 15, 15, 29, 11, 11, 26, 26, 31, 32,
	38, 42,
}

var yyR2 = [...]int8{
	0, 1, 2, 1, 2, 1, 1, 0, 1, 0,
	1, 5, 0, 11, 10, 0, 1, 1, 3, 0,
	1, 1, 1, 0, 1, 3, 4, 7, 0, 5,
	0, 2, 8, 0, 9, 0, 8, 0, 6, 1,
	1, 3, 1, 1, 1, 1, 1, 1, 2, 2,
	0, 1, 1, 3, 1, 4, 4, 2, 4, 1,
	1, 0, 6, 1, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 2, 1,
	2, 0, 1, 1, 2, 3, 0, 1, 3, 2,
	5, 4, 1, 1, 2, 2, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 3, 1, 3, 1, 1,
	1, 1,
}

var yyChk = [...]int16{
	-1000, -34, -1, -35, -3, -2, -36, -12, -6, -4,
	-27, 27, 22, 37, 36, -33, -5, -15, 60, 23,
	-22, 48, 11, 16, 25, 28, -14, 32, 31, 54,
	53, 43, 44, 47, 45, -29, -20, -23, 33, 6,
	26, -28, 62, 7, 49, -2, -36, -1, -30, -31,
	-32, 5, 4, 53, 54, 55, 56, 57, 51, 52,
	50, 41, 42, 38, 39, 40, -10, -12, -18, -12,
	-12, -12, -29, -12, 17, -12, -12, -2, 9, 34,
	8, -12, -12, -42, 66, 7, 60, 63, 60, 60,
	-19, 60, 32, 7, -9, -3, -12, -12, -12, -12,
	-12, -12, -12, -12, -12, -12, -12, -12, -12, 61,
	12, 7, 19, 20, 29, -26, 7, 7, 7, -28,
	60, -11, -18, -12, -12, -11, -12, -2, 18, -12,
	-41, -43, -13, -38, 46, 65, 60, 60, -11, 61,
	-38, 64, -38, 61, 61, -7, 13, -21, -20, -22,
	60, 20, -2, -2, -31, 7, -16, -17, 7, 24,
	-16, 61, -18, -12, -8, 14, -12, 20, -22, -40,
	21, 30, 61, -38, 50, 7, 61, -38, 15, -2,
	12, -39, 61, -2, -13, -17, -15, -13, -12, -2,
	-2, 21, -37, -24, -25, 9, 61, -7, 21, -24,
	-2, 9, -26, -2, 10, -26, -31, 35, -31,
}

var yyDef = [...]int8{
	-2, -2, -2, -2, 17, 0, 3, 42, 43, 44,
	45, 46, 47, 50, 81, 5, 6, 63, 0, 0,
	79, 0, 0, 0, 0, 15, 0, 92, 93, 0,
	0, 96, 97, 98, 99, 100, 101, 102, 0, 8,
	103, 52, 0, 54, 0, 2, 4, 16, 19, 21,
	22, 108, 109, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 48, 51, 49, 82,
	0, 78, 80, 0, 0, 0, 0, 0, 0, 0,
	0, 94, 95, 0, 111, 89, 81, 0, 0, 81,
	57, 0, 59, 60, 18, 20, 65, 66, 67, 68,
	69, 70, 71, 72, 73, 74, 75, 76, 77, 64,
	15, 0, 0, 37, 61, 9, 106, 0, 0, 53,
	81, 0, 104, 0, 0, 0, 0, 28, 0, 0,
	15, 15, 0, 0, 10, 110, 86, 86, 0, 91,
	81, 56, 0, 55, 58, 30, 0, 0, 39, 40,
	0, 35, 0, 0, 11, 107, 0, 87, 83, 0,
	0, 90, 105, 0, 0, 15, 0, 33, 0, 15,
	38, 62, 9, 0, 0, 84, 9, 0, 27, 31,
	15, 15, 41, 0, 12, 88, 85, 23, 0, 28,
	0, 36, 23, 15, 24, 0, 32, 29, 34, 15,
	0, 0, 0, 0, 14, 0, 25, 13, 26,
}

var yyTok1 = [...]int8{
	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 57, 3, 3,
	60, 61, 55, 53, 65, 54, 66, 56, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 4, 5,
	52, 50, 51, 62, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 63, 3, 64,
}

var yyTok2 = [...]int8{
	2, 3, 6, 7, 8, 9, 10, 11, 12, 13,
	14, 15, 16, 17, 18, 19, 20, 21, 22, 23,
	24, 25, 26, 27, 28, 29, 30, 31, 32, 33,
	34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
	44, 45, 46, 47, 48, 49, 58, 59,
}

var yyTok3 = [...]int8{
	0,
}

var yyErrorMessages = [...]struct {
	msg   string
	state int
	token int
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
//line .\grammar.y:92
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].body, yylex)
			}
		}
	case 2:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:97
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[2].opt_body, yylex)
			}
		}
	case 5:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:107
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].global_variables, yylex)
			}
		}
	case 6:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:112
		{
			if ast, ok := yylex.(*AstNode); ok {
				ast.ModuleStatement.Append(yyDollar[1].funcProc, yylex)
			}
		}
	case 7:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:119
		{
			yyVAL.opt_directive = nil
		}
	case 8:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:120
		{
			yyVAL.opt_directive = &yyDollar[1].token
		}
	case 9:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:123
		{
			yyVAL.opt_export = nil
		}
	case 10:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:124
		{
			yyVAL.opt_export = &yyDollar[1].token
		}
	case 11:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:127
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
//line .\grammar.y:140
		{
			isFunction(true, yylex)
		}
	case 13:
		yyDollar = yyS[yypt-11 : yypt+1]
//line .\grammar.y:141
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeFunction, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[9].opt_explicit_variables, yyDollar[10].opt_body)
			isFunction(false, yylex)
		}
	case 14:
		yyDollar = yyS[yypt-10 : yypt+1]
//line .\grammar.y:146
		{
			yyVAL.funcProc = createFunctionOrProcedure(PFTypeProcedure, yyDollar[1].opt_directive, yyDollar[3].token.literal, yyDollar[5].declarations_method_params, yyDollar[7].opt_export, yyDollar[8].opt_explicit_variables, yyDollar[9].opt_body)
		}
	case 15:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:151
		{
			yyVAL.opt_body = nil
		}
	case 16:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:152
		{
			yyVAL.opt_body = yyDollar[1].body
		}
	case 17:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:156
		{
			yyVAL.body = []Statement{yyDollar[1].stmt}
		}
	case 18:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:157
		{
			if yyDollar[2].token.literal == ":" && len(yyDollar[1].opt_body) > 0 {
				if _, ok := yyDollar[1].opt_body[len(yyDollar[1].opt_body)-1].(*GoToLabelStatement); !ok {
					yylex.Error("semicolon (;) is expected")
				}
			}
			if yyDollar[3].stmt != nil {
				yyVAL.body = append(yyVAL.body, yyDollar[3].stmt)
			}
		}
	case 19:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:170
		{
			yyVAL.stmt = nil
		}
	case 20:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:171
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 21:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:174
		{
			yyVAL.token = yyDollar[1].token
		}
	case 22:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:174
		{
			yyVAL.token = yyDollar[1].token
		}
	case 23:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:178
		{
			yyVAL.opt_explicit_variables = map[string]VarStatement{}
		}
	case 24:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:179
		{
			yyVAL.opt_explicit_variables = yyDollar[1].explicit_variables
		}
	case 25:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:182
		{
			if vars, err := appendVarStatements(map[string]VarStatement{}, yyDollar[2].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 26:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:189
		{
			if vars, err := appendVarStatements(yyDollar[1].explicit_variables, yyDollar[3].identifiers); err != nil {
				yylex.Error(err.Error())
			} else {
				yyVAL.explicit_variables = vars
			}
		}
	case 27:
		yyDollar = yyS[yypt-7 : yypt+1]
//line .\grammar.y:200
		{
			yyVAL.stmt_if = &IfStatement{
				Expression:  yyDollar[2].stmt,
				TrueBlock:   yyDollar[4].opt_body,
				IfElseBlock: yyDollar[5].opt_elseif_list,
				ElseBlock:   yyDollar[6].opt_else,
			}
		}
	case 28:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:210
		{
			yyVAL.opt_elseif_list = []Statement{}
		}
	case 29:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:211
		{
			yyVAL.opt_elseif_list = append(yyDollar[5].opt_elseif_list, &IfStatement{
				Expression: yyDollar[2].stmt,
				TrueBlock:  yyDollar[4].opt_body,
			})
		}
	case 30:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:219
		{
			yyVAL.opt_else = nil
		}
	case 31:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:220
		{
			yyVAL.opt_else = yyDollar[2].opt_body
		}
	case 32:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:223
		{
			yyVAL.stmt = TernaryStatement{
				Expression: yyDollar[3].stmt,
				TrueBlock:  yyDollar[5].stmt,
				ElseBlock:  yyDollar[7].stmt,
			}
		}
	case 33:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:232
		{
			setLoopFlag(true, yylex)
		}
	case 34:
		yyDollar = yyS[yypt-9 : yypt+1]
//line .\grammar.y:232
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[3].token.literal,
				In:   yyDollar[5].stmt,
				Body: yyDollar[8].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 35:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:240
		{
			setLoopFlag(true, yylex)
		}
	case 36:
		yyDollar = yyS[yypt-8 : yypt+1]
//line .\grammar.y:240
		{
			yyVAL.stmt_loop = &LoopStatement{
				For:  yyDollar[2].stmt,
				To:   yyDollar[4].stmt,
				Body: yyDollar[7].opt_body,
			}
			setLoopFlag(false, yylex)
		}
	case 37:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:248
		{
			setLoopFlag(true, yylex)
		}
	case 38:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:248
		{
			yyVAL.stmt_loop = &LoopStatement{
				WhileExpr: yyDollar[2].stmt,
				Body:      yyDollar[5].opt_body,
			}
		}
	case 39:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:257
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 40:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:258
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 41:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:259
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 42:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:262
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 43:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:263
		{
			yyVAL.stmt = yyDollar[1].stmt_if
		}
	case 44:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:264
		{
			yyVAL.stmt = yyDollar[1].stmt_loop
		}
	case 45:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:265
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 46:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:266
		{
			yyVAL.stmt = ContinueStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 47:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:267
		{
			yyVAL.stmt = BreakStatement{}
			checkLoopOperator(yyDollar[1].token, yylex)
		}
	case 48:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:268
		{
			yyVAL.stmt = ThrowStatement{Param: yyDollar[2].stmt}
			checkThrowParam(yyDollar[1].token, yyDollar[2].stmt, yylex)
		}
	case 49:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:269
		{
			yyVAL.stmt = &ReturnStatement{Param: yyDollar[2].stmt}
			checkReturnParam(yyDollar[2].stmt, yylex)
		}
	case 50:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:272
		{
			yyVAL.stmt = nil
		}
	case 51:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:273
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 52:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:278
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 53:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:279
		{
			yyVAL.stmt = CallChainStatement{Unit: yyDollar[3].stmt, Call: yyDollar[1].stmt}
		}
	case 54:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:285
		{
			yyVAL.stmt = VarStatement{Name: yyDollar[1].token.literal}
		}
	case 55:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:286
		{
			yyVAL.stmt = MethodStatement{Name: yyDollar[1].token.literal, Param: yyDollar[3].exprs}
		}
	case 56:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:287
		{
			yyVAL.stmt = ItemStatement{Object: yyDollar[1].stmt, Item: yyDollar[3].stmt}
		}
	case 57:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:288
		{
			yyVAL.stmt = MethodStatement{Name: yyDollar[1].token.literal, Param: []Statement{yyDollar[2].stmt}}
		}
	case 58:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:289
		{
			yyVAL.stmt = MethodStatement{Name: yyDollar[1].token.literal, Param: []Statement{yyDollar[3].stmt}}
		}
	case 59:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:292
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 60:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:293
		{
			yyVAL.stmt = VarStatement{Name: yyDollar[1].token.literal}
		}
	case 61:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:296
		{
			setTryFlag(true, yylex)
		}
	case 62:
		yyDollar = yyS[yypt-6 : yypt+1]
//line .\grammar.y:296
		{
			yyVAL.stmt = TryStatement{Body: yyDollar[2].opt_body, Catch: yyDollar[5].opt_body}
			setTryFlag(false, yylex)
		}
	case 63:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:302
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 64:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:303
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 65:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:304
		{
			yyVAL.stmt = &ExpStatement{Operation: OpPlus, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 66:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:305
		{
			yyVAL.stmt = &ExpStatement{Operation: OpMinus, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 67:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:306
		{
			yyVAL.stmt = &ExpStatement{Operation: OpMul, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 68:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:307
		{
			yyVAL.stmt = &ExpStatement{Operation: OpDiv, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 69:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:308
		{
			yyVAL.stmt = &ExpStatement{Operation: OpMod, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 70:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:309
		{
			yyVAL.stmt = &ExpStatement{Operation: OpGt, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 71:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:310
		{
			yyVAL.stmt = &ExpStatement{Operation: OpLt, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 72:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:311
		{
			yyVAL.stmt = &ExpStatement{Operation: OpEq, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 73:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:312
		{
			yyVAL.stmt = &ExpStatement{Operation: OpOr, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 74:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:313
		{
			yyVAL.stmt = &ExpStatement{Operation: OpAnd, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 75:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:314
		{
			yyVAL.stmt = &ExpStatement{Operation: OpNe, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 76:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:315
		{
			yyVAL.stmt = &ExpStatement{Operation: OpLe, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 77:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:316
		{
			yyVAL.stmt = &ExpStatement{Operation: OpGe, Left: yyDollar[1].stmt, Right: yyDollar[3].stmt}
		}
	case 78:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:317
		{
			yyVAL.stmt = not(yyDollar[2].stmt)
		}
	case 79:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:318
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 80:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:319
		{
			yyVAL.stmt = GoToStatement{Label: yyDollar[2].goToLabel}
		}
	case 81:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:322
		{
			yyVAL.stmt = nil
		}
	case 82:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:322
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 83:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:325
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(nil, yyDollar[1].token)
		}
	case 84:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:326
		{
			yyVAL.declarations_method_param = *(&ParamStatement{}).Fill(&yyDollar[1].token, yyDollar[2].token)
		}
	case 85:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:327
		{
			yyVAL.declarations_method_param = *(yyVAL.declarations_method_param.DefaultValue(yyDollar[3].stmt))
		}
	case 86:
		yyDollar = yyS[yypt-0 : yypt+1]
//line .\grammar.y:330
		{
			yyVAL.declarations_method_params = []ParamStatement{}
		}
	case 87:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:331
		{
			yyVAL.declarations_method_params = []ParamStatement{yyDollar[1].declarations_method_param}
		}
	case 88:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:332
		{
			yyVAL.declarations_method_params = append(yyDollar[1].declarations_method_params, yyDollar[3].declarations_method_param)
		}
	case 89:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:340
		{
			yyVAL.stmt = NewObjectStatement{Constructor: yyDollar[2].token.literal}
		}
	case 90:
		yyDollar = yyS[yypt-5 : yypt+1]
//line .\grammar.y:341
		{
			yyVAL.stmt = NewObjectStatement{Constructor: yyDollar[2].token.literal, Param: yyDollar[4].exprs}
		}
	case 91:
		yyDollar = yyS[yypt-4 : yypt+1]
//line .\grammar.y:342
		{
			yyVAL.stmt = NewObjectStatement{Param: yyDollar[3].exprs}
		}
	case 92:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:345
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 93:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:346
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 94:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:347
		{
			yyVAL.stmt = unaryMinus(yyDollar[2].stmt)
		}
	case 95:
		yyDollar = yyS[yypt-2 : yypt+1]
//line .\grammar.y:348
		{
			yyVAL.stmt = yyDollar[2].stmt
		}
	case 96:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:349
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 97:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:350
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 98:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:351
		{
			yyVAL.stmt = yyDollar[1].token.value
		}
	case 99:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:352
		{
			yyVAL.stmt = UndefinedStatement{}
		}
	case 100:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:353
		{
			yyVAL.stmt = yyDollar[1].goToLabel
		}
	case 101:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:354
		{
			if tok, ok := yyDollar[1].stmt.(Token); ok {
				yyVAL.stmt = tok.literal
			} else {
				yyVAL.stmt = yyDollar[1].stmt
			}
		}
	case 102:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:361
		{
			yyVAL.stmt = yyDollar[1].stmt
		}
	case 103:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:364
		{
			yyVAL.goToLabel = &GoToLabelStatement{Name: yyDollar[1].token.literal}
		}
	case 104:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:366
		{
			yyVAL.exprs = []Statement{yyDollar[1].stmt}
		}
	case 105:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:367
		{
			yyVAL.exprs = append(yyVAL.exprs, yyDollar[3].stmt)
		}
	case 106:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:370
		{
			yyVAL.identifiers = []Token{yyDollar[1].token}
		}
	case 107:
		yyDollar = yyS[yypt-3 : yypt+1]
//line .\grammar.y:371
		{
			yyVAL.identifiers = append(yyVAL.identifiers, yyDollar[3].token)
		}
	case 108:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:374
		{
			yyVAL.token = yyDollar[1].token
		}
	case 109:
		yyDollar = yyS[yypt-1 : yypt+1]
//line .\grammar.y:375
		{
			yyVAL.token = yyDollar[1].token
		}
	}
	goto yystack /* stack new state and value */
}
