package ast

import (
	"crypto/rand"
	"fmt"
	mock_ast "github.com/LazarenkoA/1c-language-parser/ast/mock"
	"github.com/golang/mock/gomock"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Next(t *testing.T) {
	c := gomock.NewController(t)
	defer c.Finish()

	t.Run("var & Identifier", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(` Перем     ввв; ввв = 3;`).AnyTimes()
		tok := new(Token)
		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "Перем", tok.literal)
		assert.Equal(t, token, Var)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "ввв", tok.literal)
		assert.Equal(t, token, Identifier)
	})

	t.Run("error", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`   2ввв`).AnyTimes()
		tok := new(Token)
		_, err := tok.Next(ast)
		assert.EqualError(t, err, "identifier immediately follow the number")
	})

	t.Run("Number", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`32323 `).AnyTimes()
		tok := new(Token)
		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "32323", tok.literal)
		assert.Equal(t, token, Number)
	})

	t.Run("String", func(t *testing.T) {
		t.Run("", func(t *testing.T) {
			tok := new(Token)
			ast := mock_ast.NewMockIast(c)
			ast.EXPECT().SrsCode().Return(`тестПерем = "тест тест"`).AnyTimes()

			token, err := tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "тестПерем", tok.literal)
			assert.Equal(t, token, Identifier)

			token, err = tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "=", tok.literal)
			assert.Equal(t, token, int('='))

			token, err = tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "тест тест", tok.literal)
			assert.Equal(t, token, String)
		})
		t.Run("", func(t *testing.T) {
			ast := mock_ast.NewMockIast(c)
			ast.EXPECT().SrsCode().Return(`тестПерем = "тест ""тест"" fd"`).AnyTimes()
			tok := new(Token)

			token, err := tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "тестПерем", tok.literal)
			assert.Equal(t, token, Identifier)

			token, err = tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "=", tok.literal)
			assert.Equal(t, token, int('='))

			token, err = tok.Next(ast)
			assert.NoError(t, err)
			assert.Equal(t, "тест \"\"тест\"\" fd", tok.literal)
			assert.Equal(t, token, String)
		})
		t.Run("date", func(t *testing.T) {
			t.Run("pass", func(t *testing.T) {
				ast := mock_ast.NewMockIast(c)
				ast.EXPECT().SrsCode().Return(`тестПерем = '00010101'`).AnyTimes()
				tok := new(Token)

				token, err := tok.Next(ast)
				assert.NoError(t, err)
				assert.Equal(t, "тестПерем", tok.literal)
				assert.Equal(t, token, Identifier)

				token, err = tok.Next(ast)
				assert.NoError(t, err)
				assert.Equal(t, "=", tok.literal)
				assert.Equal(t, token, int('='))

				token, err = tok.Next(ast)
				assert.NoError(t, err)
				assert.Equal(t, "00010101", tok.literal)
				assert.Equal(t, token, Date)
			})
			t.Run("error", func(t *testing.T) {
				ast := mock_ast.NewMockIast(c)
				ast.EXPECT().SrsCode().Return(`тестПерем = 'gfdgfg'`).AnyTimes()
				tok := new(Token)

				token, err := tok.Next(ast)
				assert.NoError(t, err)
				assert.Equal(t, "тестПерем", tok.literal)
				assert.Equal(t, token, Identifier)

				token, err = tok.Next(ast)
				assert.NoError(t, err)
				assert.Equal(t, "=", tok.literal)
				assert.Equal(t, token, int('='))

				token, err = tok.Next(ast)
				assert.EqualError(t, err, "incorrect Date type constant")
			})
		})
	})

	t.Run("String error", func(t *testing.T) {
		tok := new(Token)
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`тестПерем = "тест тест`).AnyTimes()

		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(ast)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("String error", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`тестПерем = "тест тест
`).AnyTimes()
		tok := new(Token)
		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(ast)
		assert.EqualError(t, err, "unexpected EOL")
	})

	t.Run("String", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww"`).AnyTimes()
		tok := new(Token)

		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, `тест 	тест
				|ааа fd
				| wqwq ww`, tok.literal)
		assert.Equal(t, token, String)
	})

	t.Run("String error", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`тестПерем = "тест 	тест
				|ааа fd
				| wqwq ww`).AnyTimes()

		tok := new(Token)

		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "тестПерем", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "=", tok.literal)
		assert.Equal(t, token, int('='))

		token, err = tok.Next(ast)
		assert.EqualError(t, err, "unexpected EOF")
	})

	t.Run("operators", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`Если РЗ <> Неопределено И ппп или ррр Тогда`).AnyTimes()

		tok := new(Token)

		token, err := tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "Если", tok.literal)
		assert.Equal(t, token, If)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "РЗ", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "<>", tok.literal)
		assert.Equal(t, token, NeEq)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "Неопределено", tok.literal)
		assert.Equal(t, token, Undefind)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "И", tok.literal)
		assert.Equal(t, token, And)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "ппп", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "или", tok.literal)
		assert.Equal(t, token, Or)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "ррр", tok.literal)
		assert.Equal(t, token, Identifier)

		token, err = tok.Next(ast)
		assert.NoError(t, err)
		assert.Equal(t, "Тогда", tok.literal)
		assert.Equal(t, token, Then)
	})

	t.Run("comment", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(` Перем     ввв;
					// ввв = 3;
 					d = 0;`).AnyTimes()
		tok := new(Token)

		result := []string{"Перем", "ввв", ";", "d", "=", "0", ";"}

		i := 0
		for token, err := tok.Next(ast); err == nil && token > 0; token, err = tok.Next(ast) {
			assert.Equal(t, result[i], tok.literal)
			i++
		}
	})

	t.Run("math", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`тест = 2+2*2+(2-1);`).AnyTimes()

		tok := new(Token)

		result := []string{"тест", "=", "2", "+", "2", "*", "2", "+", "(", "2", "-", "1", ")", ";"}

		i := 0
		for token, err := tok.Next(ast); err == nil && token > 0; token, err = tok.Next(ast) {
			assert.Equal(t, result[i], tok.literal)
			i++
		}
	})

	t.Run("directive", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`&НаСервере
				Процедура ДобавитьРегистрНаСервере()
	
				КонецПроцедуры`).AnyTimes()

		tok := new(Token)

		result := map[string]int{
			"&НаСервере":               Directive,
			"Процедура":                Procedure,
			"ДобавитьРегистрНаСервере": Identifier,
			"КонецПроцедуры":           EndProcedure,
			"(":                        '(',
			")":                        ')',
			"\n":                       10,
		}

		for token, err := tok.Next(ast); err == nil && token > 0; token, err = tok.Next(ast) {
			assert.Equal(t, token, result[tok.literal])
		}
	})

	t.Run("goto", func(t *testing.T) {
		ast := mock_ast.NewMockIast(c)
		ast.EXPECT().SrsCode().Return(`Процедура ДобавитьРегистрНаСервере()
				перейти ~метка;
				~метка:
				КонецПроцедуры`).AnyTimes()

		tok := new(Token)

		result := map[string]int{
			"Процедура": Procedure,
			"ДобавитьРегистрНаСервере": Identifier,
			"КонецПроцедуры":           EndProcedure,
			"(":                        '(',
			")":                        ')',
			"перейти":                  GoTo,
			"метка":                    GoToLabel,
			";":                        ';',
			":":                        ':',
		}

		for token, err := tok.Next(ast); err == nil && token > 0; token, err = tok.Next(ast) {
			assert.Equal(t, token, result[tok.literal])
		}
	})
}

func Benchmark(b *testing.B) {
	c := gomock.NewController(b)
	defer c.Finish()

	ast := mock_ast.NewMockIast(c)

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
	b.Run("next", func(b *testing.B) {
		ast.EXPECT().SrsCode().Return(`

Процедура ОткрытьНавигационнуюСсылку(НавигационнаяСсылка, Знач Оповещение = Неопределено) Экспорт

	ПустаяДата = '00010101000000';
	ПустаяДата = '20131231235959';

	КлючЗаписиРегистра = Новый("РегистрСведенийКлючЗаписи.СостоянияОригиналовПервичныхДокументов", ПараметрыМассив);
	МассаДМ = ВыборкаЕдИзм.МассаДМ/Количество;
	
     стр = новый Структура("Цикл", 1);
     стр.Цикл = 0; 

Если КодСимвола < 1040 ИЛИ КодСимвола > 1103 И КодыДопустимыхСимволов.Найти(КодСимвола) = Неопределено И Не (Не УчитыватьРазделителиСлов И ЭтоРазделительСлов(КодСимвола)) Тогда
            Возврат ;
        КонецЕсли;

перейти ~метка;

МассивСтроки.Добавить(Новый ФорматированнаяСтрока(ЧастьСтроки.Значение, Новый Шрифт(,,Истина)));

	Позиция = Найти(Строка, Разделитель);
	Пока Позиция > 0 Цикл
		Подстрока = Лев(Строка, Позиция - 1);
		Если Не ПропускатьПустыеСтроки Или Не ПустаяСтрока(Подстрока) Тогда
			Если СокращатьНепечатаемыеСимволы Тогда
				Результат.Добавить(СокрЛП(Подстрока));
			Иначе
				Результат.Добавить(Подстрока);
			КонецЕсли;
		КонецЕсли;
		Строка = Сред(Строка, Позиция + СтрДлина(Разделитель));
		Позиция = Найти(Строка, Разделитель);
	КонецЦикла;

~метка:



	вы = ввывыв[0];
	СтрокаСпискаПП[ТекКолонка.Ключ].Вставить(ТекКолонкаЗначение.Ключ, УровеньГруппировки3[ПрефиксПоля + СтрЗаменить(ТекКолонкаЗначение.Значение, ".", "")]);

	Контекст = Новый Структура();
	Контекст.Вставить("НавигационнаяСсылка", НавигационнаяСсылка);
	Контекст.Вставить("Оповещение", Оповещение);
	
	ОписаниеОшибки = СтроковыеФункцииКлиентСервер.ПодставитьПараметрыВСтроку(
			НСтр("ru = 'Не удалось перейти по ссылке ""%1"" по причине: 
			           |Неверно задана навигационная ссылка.'"),
			НавигационнаяСсылка);
	
	Если Не ОбщегоНазначенияСлужебныйКлиент.ЭтоДопустимаяСсылка(НавигационнаяСсылка) Тогда 
		ОбщегоНазначенияСлужебныйКлиент.ОткрытьНавигационнуюСсылкуОповеститьОбОшибке(ОписаниеОшибки, Контекст);
		Возврат;
	КонецЕсли;
	
	Если ОбщегоНазначенияСлужебныйКлиент.ЭтоВебСсылка(НавигационнаяСсылка)
		Или ОбщегоНазначенияСлужебныйКлиент.ЭтоНавигационнаяСсылка(НавигационнаяСсылка) Тогда 
		
		Попытка
			а = а /0;
		Исключение
			ОбщегоНазначенияСлужебныйКлиент.ОткрытьНавигационнуюСсылкуОповеститьОбОшибке(ОписаниеОшибки, Контекст);
			Возврат;
		КонецПопытки;
		
		Если Оповещение <> Неопределено Тогда 
			ПриложениеЗапущено = Истина;
			ВыполнитьОбработкуОповещения(Оповещение, ПриложениеЗапущено);
		КонецЕсли;
		
		Возврат;
	КонецЕсли;
	
	Если ОбщегоНазначенияСлужебныйКлиент.ЭтоСсылкаНаСправку(НавигационнаяСсылка) Тогда 
		ОткрытьСправку(НавигационнаяСсылка);
		Возврат;
	КонецЕсли;
КонецПроцедуры`)

		tok := new(Token)
		tok.ast = ast
		for i := 0; i < b.N; i++ {
			for _, _, err := tok.next(); err != nil; _, _, err = tok.next() {

			}
		}
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

func Benchmark_fastToLower(b *testing.B) {
	str := "ЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHЫСВАМВАОЛОРЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJЫСВАМВАОЛОРИИОРОВИЫОРСЫВРООЫОРВЫРОJHDSDHSDSJDDSHDSJHDJSDERKRVEKVJDSJHDJSDERKRVEKVJ"

	b.Run("sdk", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.ToLower(str)
		}
	})
	b.Run("fastToLower-old", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fastToLower_old(str)
		}
	})
	b.Run("fastToLower", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			fastToLower(str)
		}
	})
}
