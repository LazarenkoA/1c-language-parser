package ast

import (
	"crypto/rand"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Next(t *testing.T) {

	t.Run("var & Identifier", func(t *testing.T) {
		tok := new(Token)
		token, err := tok.Next(` Перем     ввв; ввв = 3;`)
		assert.NoError(t, err)
		assert.Equal(t, "Перем", tok.literal)
		assert.Equal(t, token, Var)

		token, err = tok.Next(` Перем     ввв; ввв = 3;`)
		assert.NoError(t, err)
		assert.Equal(t, "ввв", tok.literal)
		assert.Equal(t, token, Identifier)
	})

	t.Run("error", func(t *testing.T) {
		tok := new(Token)
		_, err := tok.Next(`   2ввв`)
		assert.EqualError(t, err, "identifier immediately follow the number")
	})

	t.Run("Number", func(t *testing.T) {
		tok := new(Token)
		token, err := tok.Next(`32323 `)
		assert.NoError(t, err)
		assert.Equal(t, "32323", tok.literal)
		assert.Equal(t, token, Number)
	})

	t.Run("String", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			tok := new(Token)
			code := `тестПерем = "тест тест"`

			token, err := tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "тестПерем", tok.literal)
			assert.Equal(t, token, Identifier)

			token, err = tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "=", tok.literal)
			assert.Equal(t, token, int('='))

			token, err = tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "тест тест", tok.literal)
			assert.Equal(t, token, String)
		})
		t.Run("", func(t *testing.T) {
			tok := new(Token)
			code := `тестПерем = "тест ""тест"" fd"`

			token, err := tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "тестПерем", tok.literal)
			assert.Equal(t, token, Identifier)

			token, err = tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "=", tok.literal)
			assert.Equal(t, token, int('='))

			token, err = tok.Next(code)
			assert.NoError(t, err)
			assert.Equal(t, "тест \"\"тест\"\" fd", tok.literal)
			assert.Equal(t, token, String)
		})
		t.Run("date", func(t *testing.T) {
			t.Run("pass", func(t *testing.T) {
				tok := new(Token)
				code := `тестПерем = '00010101'`

				token, err := tok.Next(code)
				assert.NoError(t, err)
				assert.Equal(t, "тестПерем", tok.literal)
				assert.Equal(t, token, Identifier)

				token, err = tok.Next(code)
				assert.NoError(t, err)
				assert.Equal(t, "=", tok.literal)
				assert.Equal(t, token, int('='))

				token, err = tok.Next(code)
				assert.NoError(t, err)
				assert.Equal(t, "00010101", tok.literal)
				assert.Equal(t, token, Date)
			})
			t.Run("error", func(t *testing.T) {
				tok := new(Token)
				code := `тестПерем = '0001k101'`

				token, err := tok.Next(code)
				assert.NoError(t, err)
				assert.Equal(t, "тестПерем", tok.literal)
				assert.Equal(t, token, Identifier)

				token, err = tok.Next(code)
				assert.NoError(t, err)
				assert.Equal(t, "=", tok.literal)
				assert.Equal(t, token, int('='))

				token, err = tok.Next(code)
				assert.EqualError(t, err, "incorrect Date type constant")
			})
		})
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест тест`

		token, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест тест
`

		token, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOL")
	})

	t.Run("String", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww"`

		token, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, `тест 	тест
				|ааа fd
				| wqwq ww`, tok.literal)
		assert.Equal(t, token, String)
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		code := `тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww`

		token, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(code)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("operators", func(t *testing.T) {
		tok := new(Token)
		code := `Если РЗ <> Неопределено И ппп или ррр Тогда`

		token, err := tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Если", tok.literal)
		assert.Equal(t, token, If)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "РЗ", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "<>", tok.literal)
		assert.Equal(t, token, NeEq)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Неопределено", tok.literal)
		assert.Equal(t, token, Undefind)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "И", tok.literal)
		assert.Equal(t, token, And)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "ппп", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "или", tok.literal)
		assert.Equal(t, token, Or)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "ррр", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(code)
		assert.NoError(t, err)
		assert.Equal(t, "Тогда", tok.literal)
		assert.Equal(t, token, Then)
	})

	t.Run("comment", func(t *testing.T) {
		tok := new(Token)
		code := ` Перем     ввв;
					// ввв = 3;
 					d = 0;`

		result := []string{"Перем", "ввв", ";", "d", "=", "0", ";"}

		i := 0
		for token, err := tok.Next(code); err == nil && token > 0; token, err = tok.Next(code) {
			assert.Equal(t, result[i], tok.literal)
			i++
		}
	})

	t.Run("math", func(t *testing.T) {
		tok := new(Token)
		code := `тест = 2+2*2+(2-1);`

		result := []string{"тест", "=", "2", "+", "2", "*", "2", "+", "(", "2", "-", "1", ")", ";"}

		i := 0
		for token, err := tok.Next(code); err == nil && token > 0; token, err = tok.Next(code) {
			assert.Equal(t, result[i], tok.literal)
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

		for token, err := tok.Next(code); err == nil && token > 0; token, err = tok.Next(code) {
			assert.Equal(t, token, result[tok.literal])
		}
	})
}

func Benchmark(b *testing.B) {
	b.Run("isDigit", func(b *testing.B) {
		str := "12324567376566736kl"
		b.Run("tLoop", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsDigit(str)
			}
		})
		b.Run("RegExp1", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IsDigitRegExp(str)
			}
		})
		b.Run("RegExp2", func(b *testing.B) {
			re := regexp.MustCompile(`[0-9]+`)
			for i := 0; i < b.N; i++ {
				re.MatchString(str)
			}
		})
	})
	b.Run("toLower", func(b *testing.B) {
		str := "АывыввввввввввввввввввввввввввввввсавукавамвамваепмкеккеАа"
		b.Run("classic", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				strings.ToLower(str)
			}
		})
		b.Run("bicycle", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fastToLower(str)
			}
		})
	})
	b.Run("string concatenation", func(b *testing.B) {
		var test string
		b.Run("strings.Builder", func(b *testing.B) {
			builder := strings.Builder{}
			for i := 0; i < b.N; i++ {
				builder.WriteString(generateRandomString(50))
			}
			test = builder.String()
		})
		b.Run("fmt.Sprintf", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test = fmt.Sprintf("%s%s", test, generateRandomString(50))
			}
		})
		b.Run("concatenation", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				test += generateRandomString(50)
			}
		})
		_ = test
	})
}

func IsDigitRegExp(str string) bool {
	re := regexp.MustCompile(`[0-9]+`)
	if re.MatchString(str) {
		return true
	}
	return false
}

func generateRandomString(length int) string {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return strings.TrimRight(string(b), "\x00")
}
