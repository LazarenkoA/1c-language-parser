package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Next(t *testing.T) {

	t.Run("var & Identifier", func(t *testing.T) {
		tok := new(Token)
		token, lit, err := tok.Next(` Перем     ввв; ввв = 3;`)
		assert.NoError(t, err)
		assert.Equal(t, "Перем", lit)
		assert.Equal(t, token, Var)

		token, lit, err = tok.Next(` Перем     ввв; ввв = 3;`)
		assert.NoError(t, err)
		assert.Equal(t, "ввв", lit)
		assert.Equal(t, token, Identifier)
	})

	t.Run("error", func(t *testing.T) {
		tok := new(Token)
		_, _, err := tok.Next(`   2ввв`)
		assert.EqualError(t, err, "identifier immediately follow the number")
	})

	t.Run("Number", func(t *testing.T) {
		tok := new(Token)
		token, lit, err := tok.Next(`32323 `)
		assert.NoError(t, err)
		assert.Equal(t, "32323", lit)
		assert.Equal(t, token, Number)
	})

	t.Run("String", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест тест"`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", lit)
		assert.Equal(t, token, int('='))

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тест тест", lit)
		assert.Equal(t, token, String)
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест тест`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", lit)
		assert.Equal(t, token, int('='))

		token, lit, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест тест
`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", lit)
		assert.Equal(t, token, int('='))

		token, lit, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOL")
	})

	t.Run("String", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww"`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", lit)
		assert.Equal(t, token, int('='))

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, `тест 	тест
				|ааа fd
				| wqwq ww`, lit)
		assert.Equal(t, token, String)
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", lit)
		assert.Equal(t, token, int('='))

		token, lit, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("operators", func(t *testing.T) {
		tok := new(Token)
		code := `Если РЗ <> Неопределено И ппп или ррр Тогда`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Если", lit)
		assert.Equal(t, token, If)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "РЗ", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "<>", lit)
		assert.Equal(t, token, NeEq)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Неопределено", lit)
		assert.Equal(t, token, Undefind)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "И", lit)
		assert.Equal(t, token, And)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "ппп", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "или", lit)
		assert.Equal(t, token, Or)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "ррр", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Тогда", lit)
		assert.Equal(t, token, Then)
	})

	t.Run("comment", func(t *testing.T) {
		tok := new(Token)
		code := ` Перем     ввв;
					// ввв = 3;`

		token, lit, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Перем", lit)
		assert.Equal(t, token, Var)

		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "ввв", lit)
		assert.Equal(t, token, Identifier)

		token, lit, err = tok.Next(code) // ;
		token, lit, err = tok.Next(code) // \n
		token, lit, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "", lit)
		assert.Equal(t, int32(token), EOL)
	})

	t.Run("math", func(t *testing.T) {
		tok := new(Token)
		code := `тест = 2+2*2+(2-1);`

		result := []string{"тест", "=", "2", "+", "2", "*", "2", "+", "(", "2", "-", "1", ")", ";"}

		i := 0
		for token, lit, err := tok.Next(code); err == nil && token > 0; token, lit, err = tok.Next(code) {
			assert.Equal(t, result[i], lit)
			i++
		}
	})

	t.Run("directive", func(t *testing.T) {
		tok := new(Token)
		code := `&НаСервере
				Процедура ДобавитьРегистрНаСервере()
	
				КонецПроцедуры`

		result := map[string]int{
			"&НаСервере":               Directive,
			"Процедура":                Procedure,
			"ДобавитьРегистрНаСервере": Identifier,
			"КонецПроцедуры":           EndProcedure,
			"(":                        '(',
			")":                        ')',
			"\n":                       10,
		}

		for token, lit, err := tok.Next(code); err == nil && token > 0; token, lit, err = tok.Next(code) {
			assert.Equal(t, token, result[lit])
		}
	})
}
