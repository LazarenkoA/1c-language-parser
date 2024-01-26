package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestParseIF(t *testing.T) {
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `если test = 1 Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `если (test = 1) Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если test2 = 1 И test1 = 1 ИЛИ Е = 4 Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если test2 = 1 И test1 = 1 ИЛИ  Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.EqualError(t, err, "syntax error. line: 1, column: 32 (unexpected literal: \"Тогда\")")
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если тест() Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если тест(1, в) Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если тест(1,
// 							в) Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если тест(1 в) Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.EqualError(t, err, "syntax error. line: 1, column: 12 (unexpected literal: \"в\")")
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `Если (test2 = 1) И (test1 = 1 ИЛИ f = 2) Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `если (test = 1 Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.EqualError(t, err, "syntax error. line: 1, column: 15 (unexpected literal: \"Тогда\")")
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `есл (test = 1 Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.EqualError(t, err, "syntax error. line: 1, column: 0 (unexpected literal: \"есл\")")
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `есл test = 1 Тогда КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.EqualError(t, err, "syntax error. line: 1, column: 0 (unexpected literal: \"есл\")")
// 	})
// 	t.Run("parse IF", func(t *testing.T) {
// 		code := `если test = 1 Тогда Иначе КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// 	t.Run("parse IFELSE", func(t *testing.T) {
// 		code := `если test = 1 Тогда
// 				ИначеЕсли test = 2 И кк = 9 Тогда
// 				КонецЕсли;`
//
// 		a := NewAST(code)
// 		err := a.Parse()
// 		assert.NoError(t, err)
// 	})
// }

func TestParseBase(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		code := ``

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("bosh", func(t *testing.T) {
		code := `fdfd`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 1, column: 0 (unexpected literal: \"fdfd\")")
	})
}

func TestParseIF(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если 1 = 1 Тогда
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если (1 = 1) Тогда 

					КонецЕсли; 
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И а = 1 или у = 3 Тогда

					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И а = 1 или у = 3 Тогда

					КонецЕсли

					;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда

					КонецЕсли

					;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда
						Иначе
						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

						ИначеЕсли ввв Тогда

						ИначеЕсли авыав Тогда

						Иначе

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

						ИначеЕсли ввв Тогда

						ИначеЕсли авыав Тогда

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда
							а = 1 + 3 *4;
						ИначеЕсли ввв Тогда

						ИначеЕсли авыав Тогда

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда
							а = 1 + 3 * 4
						ИначеЕсли ввв Тогда

						ИначеЕсли авыав Тогда

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 5, column: 142 (unexpected literal: \"ИначеЕсли\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

						ИначеЕсли ввв Тогда

						ИначеЕсли авыав 

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 8, column: 165 (unexpected literal: \"\\n\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 6, column: 136 (unexpected literal: \"КонецПроцедуры\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если  Тогда
	
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 51 (unexpected literal: \"Тогда\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если f Тогда
	
					КонецЕсли
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 5, column: 79 (unexpected literal: \"КонецПроцедуры\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если f Тогд
	
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 52 (unexpected literal: \"Тогд\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если (1 = 1) Тогда 
						f = 0
					КонецЕсли; 
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 4, column: 82 (unexpected literal: \"КонецЕсли\")")
	})
}

func TestParseFunctionProcedure(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		t.Run("with directive", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("bad directive", func(t *testing.T) {
			code := `&НасервереБез
					Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 1, column: 0 (unexpected literal: \"&\")")
		})
		t.Run("without directive", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка, вы, выыыыы) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("export", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка) Экспорт

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("error", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецФунки`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 3, column: 50 (unexpected literal: \"КонецФунки\")")
		})
		t.Run("error", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 3, column: 50 (unexpected literal: \"КонецПроцедуры\")")
		})
	})
	t.Run("Procedure", func(t *testing.T) {
		t.Run("with directive", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Процедура ПодключитьВнешнююОбработку() 

					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("bad directive", func(t *testing.T) {
			code := `&НасервереБез
					Процедура ПодключитьВнешнююОбработку(Ссылка) 

					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 1, column: 0 (unexpected literal: \"&\")")
		})
		t.Run("export", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) Экспорт

					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 3, column: 52 (unexpected literal: \"КонецФункции\")")
		})
	})
}

func TestParseBaseExpression(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) ds = 222; uu = 9; КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = 222; uu = 9;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = 222



					; uu = 9;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = 222; 
					uu = 9;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = 222 
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 65 (unexpected literal: \"КонецПроцедуры\")")
	})
}

func BenchmarkString(b *testing.B) {
	b.Run("string", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			test("rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd dfdf dsfd")
		}
	})

	b.Run("ptr string", func(b *testing.B) {
		str := "rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd dfdf dsfd"
		for i := 0; i < b.N; i++ {
			testPt(&str)
		}
	})
}

func test(str string) {

}

func testPt(str *string) {

}

func testGoodCode() string {
	return `  Процедура ТипЗамераПриИзменении(Элемент)
	Если Объект.ТипЗамера=ПредопределенноеЗначение("Перечисление.ТипыЗамеров.Произвольный") Тогда
		Элементы.ДополнительнаяОбработка.Видимость = Истина;
	Иначе
		Элементы.ДополнительнаяОбработка.Видимость = Ложь;
	КонецЕсли;

	
	Старт = ТекущаяУниверсальнаяДатаВМиллисекундах();
	Для а = 0 По Количество Цикл   
		Если а > 0 И (а%Гранулярность = 0 ИЛИ а = Количество) Тогда  
			Стоп = ТекущаяУниверсальнаяДатаВМиллисекундах();
			ДобавитьТочку(Серия, а, Стоп-Старт);	     
			Старт = ТекущаяУниверсальнаяДатаВМиллисекундах();    
		КонецЕсли;
		
		новСправочник = Справочники.Справочник1.СоздатьЭлемент();
		новСправочник.Наименование = "спр"+ Строка(а);
		новСправочник.Записать();
	КонецЦикла;    

КонецПроцедуры
`
}
