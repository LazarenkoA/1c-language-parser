package main

import (
	"errors"
	"fmt"
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
		"исключение":        Catch,
		"конецпопытки":      EndTry,
		"массив":            Array,
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
	}

	directives = map[string]int{
		"&наклиенте":                      Directive,
		"&насервере":                      Directive,
		"&насерверебезконтекста":          Directive,
		"&наклиентенасерверебезконтекста": Directive,
		"&наклиентенасервере":             Directive,
	}
)

func (t *Token) Next(srs string) (int, string, error) {
	t.one.Do(func() {
		t.srs = srs
	})

	t.skipSpace()

	switch let := t.currentLet(); {
	case isLetter(let):
		literal, err := t.scanIdentifier()
		if err != nil {
			return EOF, emptyLit, err
		}

		lowLit := strings.ToLower(literal)
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
		literal, err := t.scanString('"')
		if err != nil {
			return EOF, emptyLit, err
		}

		return String, literal, nil
	case let == '=' || let == '-' || let == '+' || let == '*' || let == '(' || let == ')' || let == '[' || let == ']' || let == ':' || let == ';' || let == ',' || let == '%':
		t.nextPos()
		return int(let), string(let), nil
	case let == '<':
		if t.nextLet() == '>' {
			t.nextPos()
			t.nextPos()
			return NeEq, "<>", nil
		} else {
			return int(let), string(let), nil
		}
	case let == '>':
		t.nextPos()
		return int(let), string(let), nil
	case let == '/':
		if t.nextLet() == '/' {
			for ch := t.currentLet(); ch != '\n' && ch != EOF; ch = t.currentLet() {
				t.nextPos()
			}
			return EOL, emptyLit, nil
		}

		t.nextPos()
		return int(let), string(let), nil
	case let == '&':
		t.nextPos()
		pos := t.offset

		literal, err := t.scanIdentifier()
		if err != nil {
			return EOF, emptyLit, err
		}

		lowLit := strings.ToLower("&" + literal)
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
		case '\n':
			t.nextPos()
			return int(let), string(let), nil
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

func (t *Token) scanString(l rune) (string, error) {
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
		case cl == l:
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
	return Position{
		Line:   strings.Count(t.srs[:t.offset], "\n") + 1,
		Column: len([]rune(t.srs[:t.offset])) + 1,
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
	return unicode.IsLetter(ch)
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isSpace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\r'
}
