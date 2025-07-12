package ast

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"
	"unsafe"
)

//go:generate mockgen -source=$GOFILE -destination=./mock/mock.go
type Iast interface {
	SrsCode() string
}

type Position struct {
	Line   int
	Column int
}

type Token struct {
	ast      Iast
	value    interface{}
	literal  string
	position Position
	offset   int
	prevDot  bool
}

const (
	EOL      = '\n' // end of line.
	emptyLit = ""
)

var (
	tokens = map[string]int{
		"процедура":         Procedure,
		"перем":             Var,
		"перейти":           GoTo,
		"конецпроцедуры":    EndProcedure,
		"знач":              ValueParam,
		"если":              If,
		"тогда":             Then,
		"иначеесли":         ElseIf,
		"иначе":             Else,
		"конецесли":         EndIf,
		"для":               For,
		"каждого":           Each,
		"из":                In,
		"по":                To,
		"цикл":              Loop,
		"конеццикла":        EndLoop,
		"прервать":          Break,
		"продолжить":        Continue,
		"попытка":           Try,
		"новый":             New,
		"исключение":        Catch,
		"пока":              While,
		"конецпопытки":      EndTry,
		"функция":           Function,
		"конецфункции":      EndFunction,
		"возврат":           Return,
		"вызватьисключение": Throw,
		"и":                 And,
		"или":               OR,
		"истина":            True,
		"ложь":              False,
		"неопределено":      Undefind,
		"не":                Not,
		"экспорт":           Export,
		"выполнить":         Execute,
		//"вычислить":         Eval,
		// "массив":            Array,
		// "структура":         Struct,
		// "соответствие":      Dictionary,
	}

	directives = map[string]int{
		"&наклиенте":                      Directive,
		"&насервере":                      Directive,
		"&насерверебезконтекста":          Directive,
		"&наклиентенасерверебезконтекста": Directive,
		"&наклиентенасервере":             Directive,
	}
)

func (t *Token) Next(ast Iast) (token int, err error) {
	t.ast = ast
	token, t.literal, err = t.next()

	switch token {
	case Number:
		t.value, err = strconv.ParseFloat(t.literal, 64)
	case String:
		t.value = t.literal
	case Date:
		formats := []string{"20060102", "200601021504", "20060102150405"} // Допускается опускать либо время целиком, либо только секунды.
		for _, f := range formats {
			// если все 0 это равносильно пустой дате
			if strings.Count(t.literal, "0") == len(t.literal) {
				t.value = time.Time{}
				return
			}

			if t.value, err = time.Parse(f, t.literal); err == nil {
				break
			}
		}
	case Undefind:
		t.value = nil
	case True:
		t.value = true
	case False:
		t.value = false
	}

	return
}

func (t *Token) next() (int, string, error) {
	t.skipSpace()
	t.skipComment()
	t.skipRegions()

	if t.prevDot {
		defer func() { t.prevDot = false }()
	}

	switch let := t.currentLet(); {
	case isLetter(let):
		literal := t.scanIdentifier()
		lowLit := fastToLower(literal)

		if tName, ok := tokens[lowLit]; ok && !t.prevDot {
			return tName, literal, nil
		} else {
			return token_identifier, literal, nil
		}
	case let == '.':
		// если после точки у нас следует идентификатор то нам нужно читать его обычным идентификатором
		// Могут быть таие случаи стр.Истина = 1 или стр.Функция = 2 (стр в данном случае какой-то объект, например структура)
		// нам нужно что бы то что следует после точки считалось Identifier, а не определенным зарезервированным токеном
		t.prevDot = true

		t.nextPos()
		return int(let), string(let), nil
	case isDigit(let):
		if literal, err := t.scanNumber(); err != nil {
			return EOF, emptyLit, err
		} else {
			return Number, literal, nil
		}

	case let == 0x27:
		literal, err := t.scanString(let)
		if err != nil {
			return EOF, emptyLit, err
		}

		// В литерале даты игнорируются все значения, отличные от цифр.
		if literal = extractDigits(literal); literal == "" {
			return EOF, emptyLit, fmt.Errorf("incorrect Date type constant")
		}

		return Date, literal, nil
	case let == '/' || let == ';' || let == '(' || let == ')' || let == ',' || let == '-' || let == '+' || let == '*' || let == '?' || let == '[' || let == ']' || let == ':' || let == '%':
		t.nextPos()
		return int(let), string(let), nil
	case let == '=':
		t.nextPos()
		return EQUAL, string(let), nil
	case let == '"':
		literal, err := t.scanString(let)
		if err != nil {
			return EOF, emptyLit, err
		}

		return String, literal, nil
	case let == '<':
		if t.nextLet() == '>' {
			t.nextPos()
			t.nextPos()
			return NeEQ, "<>", nil
		} else if t.nextLet() == '=' {
			t.nextPos()
			t.nextPos()
			return LE, "<=", nil
		} else {
			t.nextPos()
			return int(let), string(let), nil
		}
	case let == '>':
		if t.nextLet() == '=' {
			t.nextPos()
			t.nextPos()
			return GE, ">=", nil
		} else {
			t.nextPos()
			return int(let), string(let), nil
		}

	// case let == '#':
	// literal, err := t.scanIdentifier()
	// if err != nil {
	// 	return EOF, emptyLit, err
	// }

	case let == '&':
		t.nextPos()
		pos := t.offset

		literal := t.scanIdentifier()
		lowLit := fastToLower("&" + literal)

		if tName, ok := directives[lowLit]; ok {
			return tName, "&" + literal, nil
		} else {
			t.offset = pos
			return int(let), string(let), fmt.Errorf(`syntax error %q`, string(let))
		}
	case let == '~':
		t.nextPos()
		return GoToLabel, t.scanIdentifier(), nil
	default:
		switch let {
		case EOF:
			t.nextPos()
			return EOF, emptyLit, nil
		// case '\n':
		// 	t.nextPos()
		// 	return int(let), string(let), nil
		default:
			t.nextPos()
			return int(let), string(let), fmt.Errorf(`syntax error %q`, string(let))
		}
	}
}

func (t *Token) scanIdentifier() string {
	ret := make([]rune, 0, 10) // как правило встречаются короткие идентификаторы и лучше предаллоцировать, это сильный буст дает

	for {
		let := t.currentLet()
		if !isLetter(let) && !isDigit(let) {
			break
		}

		ret = append(ret, let)
		t.nextPos()
	}

	return string(ret)
}

func (t *Token) scanString(end rune) (string, error) {
	var ret []rune

eos:
	for {
		t.nextPos()

		switch cl := t.currentLet(); {
		case cl == EOL:
			t.nextPos()

			t.skipSpace()
			t.skipComment() // комментарии могут быть в тексте, в тексте запроса например
			if cl = t.currentLet(); cl != '|' && !isSpace(cl) {
				return "", fmt.Errorf("unexpected EOL")
			}

			ret = append(append(ret, EOL), cl)
		case cl == EOF:
			return "", fmt.Errorf("unexpected EOF")
		case cl == end:
			// пропускаем двойные "
			if t.nextLet() == '"' {
				t.nextPos()
				ret = append(ret, '"', '"')
				continue
			}
			t.nextPos()
			t.skipSpace()

			// для случаев (в 1с такое допускается, считается одна строка)
			// d = "ввава"
			// "fffff"
			if t.currentLet() == '"' {
				continue
			}

			break eos
		default:
			ret = append(ret, cl)
		}
	}

	return string(ret), nil
}

func (t *Token) skipSpace() {
	for isSpace(t.currentLet()) {
		t.nextPos()
	}
}

func (t *Token) skipComment() {
	if t.currentLet() == '/' && t.nextLet() == '/' {
		for ch := t.currentLet(); ch != EOL && ch != EOF; ch = t.currentLet() {
			t.nextPos()
		}
		t.skipSpace()
	} else {
		return
	}

	// проверяем что на новой строке нет комментария или новой области, если есть, рекурсия
	if cl := t.currentLet(); cl == '/' {
		t.skipComment()
	} else if cl := t.currentLet(); cl == '#' {
		t.skipRegions()
	}
}

func (t *Token) skipRegions() {
	// todo пока будут пропускаться и условия типа #Если Не ВебКлиент Тогда, потом надо будет доработать
	if t.currentLet() == '#' {
		for ch := t.currentLet(); ch != EOL && ch != EOF; ch = t.currentLet() {
			t.nextPos()
		}
		t.skipSpace()
	}

	// проверяем что на новой строке нет комментария или новой области, если есть, рекурсия
	if cl := t.currentLet(); cl == '/' {
		t.skipComment()
	} else if cl := t.currentLet(); cl == '#' {
		t.skipRegions()
	}
}

func (t *Token) nextLet() rune {
	srsCode := t.ast.SrsCode()
	_, size := utf8.DecodeRuneInString(srsCode[t.offset:])
	t.offset += size
	defer func() { t.offset -= size }()

	return t.currentLet()
}

func (t *Token) currentLet() rune {
	srsCode := t.ast.SrsCode()

	if t.offset >= len(srsCode) {
		return EOF
	}

	char, _ := utf8.DecodeRuneInString(srsCode[t.offset:])
	if char == utf8.RuneError {
		fmt.Println(fmt.Errorf("error decoding the character"))
		return char
	}

	return char
}

func (t *Token) GetPosition() Position {
	srsCode := t.ast.SrsCode()
	eol := strings.LastIndex(srsCode[:t.offset], "\n") + 1
	lineBegin := IF[int](eol < 0, 0, eol)

	return Position{
		Line:   strings.Count(srsCode[:t.offset], "\n") + 1,
		Column: len([]rune(srsCode[lineBegin:t.offset])) + 1,
	}
}

func (t *Token) nextPos() {
	srsCode := t.ast.SrsCode()
	_, size := utf8.DecodeRuneInString(srsCode[t.offset:])
	t.offset += size
}

func (t *Token) scanNumber() (string, error) {
	var ret []rune

	let := t.currentLet()
	for ; isDigit(let) || let == '.'; let = t.currentLet() {
		ret = append(ret, let)
		t.nextPos()
	}

	if isLetter(let) {
		return "", fmt.Errorf("identifier immediately follow the number")
	}

	return string(ret), nil
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_'
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func IF[T any](condition bool, a, b T) T {
	if condition {
		return a
	} else {
		return b
	}
}

func IsDigit(str string) bool {
	for _, c := range str {
		if c < '0' || c > '9' {
			return false
		}
	}
	return true
}

func extractDigits(str string) string {
	result := make([]rune, 0, len(str))
	for _, c := range str {
		if c >= '0' && c <= '9' {
			result = append(result, c)
		}
	}
	return string(result)
}

func fastToLower_old(s string) string {
	rs := bytes.NewBuffer(make([]byte, 0, len(s)))
	for _, rn := range s {
		switch {
		case (rn >= 'А' && rn <= 'Я') || (rn >= 'A' && rn <= 'Z'):
			rs.WriteRune(rn + 0x20)
		case rn == 'Ё':
			rs.WriteRune('ё')
		default:
			rs.WriteRune(rn)
		}
	}
	return rs.String()
}

func fastToLower(s string) string {
	b := []byte(s)
	for i, r := range s {
		switch {
		case r >= 'A' && r <= 'Z':
			b[i] = s[i] + ('a' - 'A')
		case r >= 'А' && r <= 'Я':
			if s[i] == 208 && r > 'П' { // от "П" и дальше
				b[i], b[i+1] = b[i]+1, s[i+1]-('а'-'А')
			} else {
				b[i+1] = s[i+1] + ('а' - 'А')
			}
		case r == 'Ё':
			b[i], b[i+1] = 209, 145
		}
	}

	return unsafe.String(unsafe.SliceData(b), len(b))
}
