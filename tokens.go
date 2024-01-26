package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

//go:generate goyacc  .\grammar.y

type Position struct {
	Line   int
	Column int
}

type Token struct {
	offset   int
	one      sync.Once
	srs      string
	charSize int
	position Position
	literal  string
	value    interface{}
}

const (
	EOL      = '\n' // end of line.
	emptyLit = ""
)

var (
	tokens = map[string]int{
		"процедура":         Procedure,
		"перем":             Var,
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
		"конецпопытки":      EndTry,
		"функция":           Function,
		"конецфункции":      EndFunction,
		"возврат":           Return,
		"вызватьисключение": Throw,
		"и":                 And,
		"или":               Or,
		"истина":            True,
		"ложь":              False,
		"неопределено":      Undefind,
		"не":                Not,
		"экспорт":           Export,
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

func (t *Token) Next(srs string) (token int, err error) {
	t.one.Do(func() {
		t.srs = srs
	})

	token, t.literal, err = t.next()
	switch token {
	case Number:
		t.value, err = strconv.ParseFloat(t.literal, 64)
	//	t.value, _ = strconv.ParseInt(t.literal, 10, 64)
	case String:
		t.value = t.literal
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

	switch let := t.currentLet(); {
	case isLetter(let):
		literal, err := t.scanIdentifier()
		if err != nil {
			return EOF, emptyLit, err
		}

		lowLit := fastToLower(literal)
		if tName, ok := tokens[lowLit]; ok {
			return tName, literal, nil
		} else {
			return Identifier, literal, nil
		}
	case isDigit(let):
		if literal, err := t.scanNumber(); err != nil {
			return EOF, emptyLit, err
		} else {
			return Number, literal, nil
		}
	case let == '"':
		literal, err := t.scanString(let)
		if err != nil {
			return EOF, emptyLit, err
		}

		return String, literal, nil
	case let == 0x27:
		literal, err := t.scanString(let)
		if err != nil {
			return EOF, emptyLit, err
		}

		if !IsDigit(literal) {
			return EOF, emptyLit, errors.New("Incorrect Date type constant")
		}

		return String, literal, nil
	case let == '=' || let == '-' || let == '+' || let == '*' || let == '/' || let == '(' || let == '?' || let == ')' || let == '[' || let == ']' || let == ':' || let == ';' || let == '.' || let == ',' || let == '%':
		t.nextPos()
		return int(let), string(let), nil
	case let == '<':
		if t.nextLet() == '>' {
			t.nextPos()
			t.nextPos()
			return NeEq, "<>", nil
		} else if t.nextLet() == '=' {
			t.nextPos()
			t.nextPos()
			return Le, "<=", nil
		} else {
			t.nextPos()
			return int(let), string(let), nil
		}
	case let == '>':
		if t.nextLet() == '=' {
			t.nextPos()
			t.nextPos()
			return Ge, ">=", nil
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

		literal, err := t.scanIdentifier()
		if err != nil {
			return EOF, emptyLit, err
		}

		lowLit := fastToLower("&" + literal)
		if tName, ok := directives[lowLit]; ok {
			return tName, "&" + literal, nil
		} else {
			t.offset = pos
			return int(let), string(let), fmt.Errorf(`syntax error %q`, string(let))
		}
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

func (t *Token) scanIdentifier() (string, error) {
	var ret []rune

	for {
		let := t.currentLet()
		if !isLetter(let) && !isDigit(let) {
			break
		}

		ret = append(ret, let)
		t.nextPos()
	}

	return string(ret), nil
}

func (t *Token) scanString(end rune) (string, error) {
	var ret []rune

eos:
	for {
		t.nextPos()

		switch cl := t.currentLet(); {
		case cl == EOL:
			t.nextPos()
			if cl = t.currentLet(); cl != '|' && !isSpace(cl) {
				return "", errors.New("unexpected EOL")
			}

			ret = append(append(ret, EOL), cl)
		case cl == EOF:
			return "", errors.New("unexpected EOF")
		case cl == end:
			// пропускаем двойные "
			if t.nextLet() == '"' {
				t.nextPos()
				ret = append(ret, '"', '"')
				continue
			}

			t.nextPos()
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

	// проверяем что на новой строке нет комментария, если есть, рекурсия
	if t.currentLet() == '/' {
		t.skipComment()
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
	if cl := t.currentLet(); cl == '/' || cl == '#' {
		t.skipComment()
	}
}

func (t *Token) nextLet() rune {
	t.offset += t.charSize
	defer func() { t.offset -= t.charSize }()

	return t.currentLet()
}

func (t *Token) currentLet() rune {
	if t.offset >= len(t.srs) {
		return EOF
	}

	char, size := utf8.DecodeRuneInString(t.srs[t.offset:])
	if char == utf8.RuneError {
		fmt.Println(errors.New("error decoding the character"))
		return char
	}

	t.charSize = size
	return char
}

func (t *Token) GetPosition() Position {
	eol := strings.LastIndex(t.srs[:t.offset], "\n") + 1
	lineBegin := IF[int](eol < 0, 0, eol)

	return Position{
		Line:   strings.Count(t.srs[:t.offset], "\n") + 1,
		Column: len([]rune(t.srs[lineBegin:t.offset])) + 1,
	}
}

func (t *Token) nextPos() {
	t.offset += t.charSize
}

func (t *Token) scanNumber() (string, error) {
	var ret []rune

	let := t.currentLet()
	for ; isDigit(let) || let == '.'; let = t.currentLet() {
		ret = append(ret, let)
		t.nextPos()
	}

	if isLetter(let) {
		return "", errors.New("identifier immediately follow the number")
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

func fastToLower(s string) string {
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
