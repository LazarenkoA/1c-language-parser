package ast

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	code := `Процедура dsds() d = 864/63+607-177*906*27>737*429+84-270 КонецПроцедуры`

	a := NewAST(code)
	err := a.Parse()
	if assert.NoError(t, err) {
		p := a.Print(PrintConf{OneLine: true, LispStyle: true})
		assert.Equal(t, "Процедура dsds() d = ((((864 / 63) + 607) - ((177 * 906) * 27)) > (((737 * 429) + 84) - 270));КонецПроцедуры", strings.TrimSpace(p))
	}
}

func TestParse2(t *testing.T) {
	code := `
		// @strict-types
		
		
		#Если Сервер Или ТолстыйКлиентОбычноеПриложение Или ВнешнееСоединение Тогда
		
		#Область ОписаниеПеременных
		
		#КонецОбласти
		
		#Область ПрограммныйИнтерфейс
		
		// Код процедур и функций
		
		#КонецОбласти
		
		#Область ОбработчикиСобытий
		
		// Код процедур и функций
		
		#КонецОбласти
		
		#Область СлужебныйПрограммныйИнтерфейс
		
		// Код процедур и функций
		
		#КонецОбласти
		
		#Область СлужебныеПроцедурыИФункции
		
		// Код процедур и функций
		
		#КонецОбласти
		
		#Область Инициализация
		
		#КонецОбласти
		
		#КонецЕсли
		
		`

	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)
	assert.Nil(t, a.ModuleStatement.Body)
}

func TestParse3(t *testing.T) {
	code := `
		// См. УправлениеДоступомПереопределяемый.ПриЗаполненииСписковСОграничениемДоступа.
		Процедура ПриЗаполненииОграниченияДоступа(Ограничение) Экспорт
		//{{MRG[ <-> ]
		//
		//}}MRG[ <-> ]
			Ограничение.Текст =
			"РазрешитьЧтениеИзменение
			|ГДЕ
		//{{MRG[ <-> ]
			|	ЗначениеРазрешено(Организация)
			|	И ЗначениеРазрешено(ФизическоеЛицо)";
		//}}MRG[ <-> ]
		//{{MRG[ <-> ]
		//	|	ЗначениеРазрешено(Организация)";
		//
		//}}MRG[ <-> ]
		КонецПроцедуры
		
		`

	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)
	assert.NotNil(t, a.ModuleStatement.Body)
}

func TestParseModule(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		code := ``

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("global var", func(t *testing.T) {
		code := `&НаСервере
				Перем в, e;

				&НаКлиенте 
				Перем а Экспорт; Перем с;

				Процедура вв1()
				
				Конецпроцедуры
				
				&НаКлиенте
				Процедура вв2()

				Конецпроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("global var error", func(t *testing.T) {
		code := `Перем а;
				Перем а;
				
				Процедура вв()
				
				Конецпроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.ErrorContains(t, err, "variable has already been defined")
	})
	t.Run("global var error", func(t *testing.T) {
		code := `Перем в; 

				Процедура вв()
					
				Конецпроцедуры
				Перем а;`

		a := NewAST(code)
		err := a.Parse()
		assert.ErrorContains(t, err, "variable declarations must be placed at the beginning of the module")
	})
	t.Run("without FunctionProcedure pass", func(t *testing.T) {
		code := `
					Пока Истина Цикл
						
					КонецЦикла;
					
					
					ВызватьИсключение "";
					
					Если Истина Тогда
						а = 0;
					КонецЕсли`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("without FunctionProcedure pass", func(t *testing.T) {
		code := `Перем в; 
					Функция test1() 
					КонецФункции

Функция test1() 
					КонецФункции

					Пока Истина Цикл
						
					КонецЦикла;
					
					
					ВызватьИсключение "";
					
					Если Истина Тогда
						а = 0;
					КонецЕсли;`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)

		// fmt.Println(a.Print(PrintConf{Margin: 4}))
	})
	t.Run("without FunctionProcedure error", func(t *testing.T) {
		code := `
					Пока Истина Цикл
						
					КонецЦикла;
					
					
					ВызватьИсключение "";
					
					Если Истина Тогда
						а = 0;
					КонецЕсли;

					Процедура test()
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.Error(t, err)
		// assert.ErrorContains(t, err, "procedure and function definitions should be placed before the module body statements")
	})
}

func TestExecute(t *testing.T) {
	t.Run("Execute-1", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить Алгоритм;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		if assert.NoError(t, err) {
			jdata, _ := a.JSON()
			assert.Equal(t, `{"Name":"","Body":[{"ExplicitVariables":{},"Name":"ВыполнитьВБезопасномРежиме","Directive":"\u0026НаСервере","Body":[{"Name":"Выполнить","Param":{"Statements":[{"Name":"Алгоритм"}]}}],"Params":[{"Name":"Алгоритм","IsValue":true},{"Default":{},"Name":"Параметры","IsValue":true}],"Type":1,"Export":false}]}`, string(jdata))
		}
	})
	t.Run("Execute-2", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить(Алгоритм);
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)

		p := a.Print(PrintConf{Margin: 4})
		assert.Equal(t, true, compareHashes(code, p))
	})
	t.Run("Execute-3", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить "Алгоритм";
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("Execute-error", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить 32;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 16 (unexpected literal: \"32\")")
	})
	t.Run("Execute-error", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить "Алгоритм", "";
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 26 (unexpected literal: \",\")")
	})
	t.Run("Execute-error-2", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Выполнить ("Алгоритм", "");
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 27 (unexpected literal: \",\")")
	})
	t.Run("Eval-1", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						в = Вычислить(Алгоритм);
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("Eval-2", func(t *testing.T) {
		code := `&НаСервере
					Процедура ВыполнитьВБезопасномРежиме(Знач Алгоритм, Знач Параметры = Неопределено)
						Вычислить Алгоритм;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 16 (unexpected literal: \"Алгоритм\")")
	})
}

func TestParseIF(t *testing.T) {
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
						test = 2+2*2;
						а = 7;
						а = 7.2;
					ИначеЕсли Не 4 = 3 И Не 8 = 2 И 1 <> 3 Тогда;
						а = 5;
					ИначеЕсли Ложь Тогда;
					Иначе
						а = -(1+1);
						а = -s;
						а = -1;
						а = -7.42;
						а = Не истина;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)

		data, err := a.JSON()
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(data))
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
					Если Истина Тогда
	
					КонецЕсли // запятой может и не быть
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если (Истина или Ложь) Тогда
						а = 0;
					КонецЕсли
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если (1 = 1) Тогда 
						f = 0 // запятой может не быть
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
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Если ВерсияПлатформыЧислом < 80303641 Тогда
							ВызватьИсключение НСтр("ru = 'Для работы обработки требуется ""1С:Предприятие"" версии 8.3.3.641 или страше.';sys= ''", "ru");
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
							b = 1
						ИначеЕсли ввв Тогда

						ИначеЕсли авыав Тогда

						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 5, column: 7 (unexpected literal: \"b\")")
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
		assert.EqualError(t, err, "syntax error. line: 9, column: 6 (unexpected literal: \"КонецЕсли\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если в = 1 И (а = 1 или у = 3) Тогда
						Если в = 1 или у = 3 Тогда

					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 6, column: 4 (unexpected literal: \"КонецПроцедуры\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если  Тогда
	
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 11 (unexpected literal: \"Тогда\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если f Тогд
	
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 12 (unexpected literal: \"Тогд\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если ав f Тогда
	
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 13 (unexpected literal: \"f\")")
	})
	t.Run("\"not\" pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Если Не f Тогда

					КонецЕсли;

					Если Не f Тогда
						d = 0;
					ИначеЕсли 3 = 9 Тогда
						Если тогоСего Тогда
						
						КонецЕсли;
					Иначе
						Если Не f И не 1 = 1 ИЛИ не (а = 2 ИЛИ Истина) Тогда
						
						КонецЕсли;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
}

func TestParseLoop(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для Каждого ИзмененныйОбъект Из ОбъектыНазначения Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							ТипыИзмененныхОбъектов = 0;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)

		data, err := a.JSON()
		assert.NoError(t, err)
		assert.NotEqual(t, 0, len(data))
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для а = 0 По 100 Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							ТипыИзмененныхОбъектов = 0;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для а = 0 По 100 Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
   						Продолжить;

						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							Продолжить;
						Иначе
							Прервать;
						КонецЕсли;
					КонецЦикла;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для а = 0 По 100 Цикл            
						Для а = 0 По 100 Цикл
							Если Истина Тогда
								Прервать;
							КонецЕсли;
						КонецЦикла;
						
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							Продолжить;
						Иначе
							Прервать;
						КонецЕсли;
					КонецЦикла; 
					
					Если ТипыИзмененныхОбъектов  = Неопределено Тогда       
						Для а = 0 По 100 Цикл
							Если Истина Тогда
								Прервать;
							КонецЕсли;
						КонецЦикла;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку()
					Для Каждого КлючЗначение Из Новый Структура(СписокКолонок) Цикл
						
					КонецЦикла;
					Для Каждого КлючЗначение Из (Новый Структура(СписокКолонок2)) Цикл
						
					КонецЦикла;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
		if !t.Failed() {
			p := a.Print(PrintConf{Margin: 0})
			assert.Equal(t, "Процедура ПодключитьВнешнююОбработку() \nДля Каждого КлючЗначение Из Новый Структура(СписокКолонок) Цикл \nКонецЦикла;\nДля Каждого КлючЗначение Из Новый Структура(СписокКолонок2) Цикл \nКонецЦикла;\nКонецПроцедуры", deleteEmptyLine(p))
		}
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для а = 0 По 100 Цикл            
						Для а = 0 По 100 Цикл
							Если Истина Тогда
								Прервать;
							КонецЕсли;
						КонецЦикла;
						
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							Продолжить;
						Иначе
							Прервать;
						КонецЕсли;
					КонецЦикла; 
					
					Если ТипыИзмененныхОбъектов  = Неопределено Тогда       
						Для а = 0 По 100 Цикл
							Если Истина Тогда
								Прервать;
							КонецЕсли;
						КонецЦикла;

						Продолжить; // вне цикла нельзя
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Продолжить\" can only be used inside a loop. line: 23, column: 6 (unexpected literal: \"Продолжить\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Функция ПодключитьВнешнююОбработку() 
					Если 1 = 1 Тогда
							f = 1+1;
							Прервать; // вне цикла нельзя
					КонецЕсли;

					Для Каждого ИзмененныйОбъект Из ОбъектыНазначения Цикл
						Если 1 = 1 Тогда
							f = 1+1;
							Прервать;
						КонецЕсли;
					КонецЦикла;
				КонецФункции`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Прервать\" can only be used inside a loop. line: 4, column: 7 (unexpected literal: \"Прервать\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Функция ПодключитьВнешнююОбработку() 
					Если 1 = 1 Тогда
							f = 1+1;
							Прервать;
					КонецЕсли;
				КонецФункции`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Прервать\" can only be used inside a loop. line: 4, column: 7 (unexpected literal: \"Прервать\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Функция ПодключитьВнешнююОбработку() 
					Для Каждого ИзмененныйОбъект Из ОбъектыНазначения Цикл
						Если 1 = 1 Тогда
							f = 1+1;
							Прервать;
						КонецЕсли;
						продолжить;
					КонецЦикла;

					Если 1 = 1 Тогда
							f = 1+1;
							Прервать;
					КонецЕсли;
				КонецФункции`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Прервать\" can only be used inside a loop. line: 12, column: 7 (unexpected literal: \"Прервать\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Продолжить; // вне цикла нельзя
					Для а = 0 По 100 Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							Продолжить;
						Иначе
							Прервать;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Продолжить\" can only be used inside a loop. line: 2, column: 5 (unexpected literal: \"Продолжить\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Прервать; // вне цикла нельзя
					Для а = 0 По 100 Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							Продолжить;
						Иначе
							Прервать;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"Прервать\" can only be used inside a loop. line: 2, column: 5 (unexpected literal: \"Прервать\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для а = 0 По  Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							ТипыИзмененныхОбъектов = 0;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 19 (unexpected literal: \"Цикл\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					Для ИзмененныйОбъект Из ОбъектыНазначения Цикл
						Тип = ТипЗнч(ИзмененныйОбъект);
						Если ТипыИзмененныхОбъектов  = Неопределено Тогда
							ТипыИзмененныхОбъектов = 0;
						КонецЕсли;
					КонецЦикла;

				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 26 (unexpected literal: \"Из\")")
	})
}

func TestTryCatch(t *testing.T) {
	t.Run("throw", func(t *testing.T) {
		t.Run("pass", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку() 
						Если в = 1 И (а = 1 или у = 3) Тогда
							f = 0;
							ВызватьИсключение "dsdsd dsds";							
							f = 0;
							f = 0;
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку() 
						Если в = 1 И (а = 1 или у = 3) Тогда
							f = 0;
							ВызватьИсключение; // без параметров нельзя
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "operator \"ВызватьИсключение\" without arguments can only be used when handling an exception. line: 4, column: 24 (unexpected literal: \";\")")
		})
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
						Попытка 
							а = 1+1;					
						Исключение
							ВызватьИсключение "fff";
						КонецПопытки;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		if assert.NoError(t, err) {
			json, _ := a.JSON()
			assert.Contains(t, string(json), `{"Param":"fff"}`)
		}
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
						Попытка 
							Попытка 
								а = 1+1;					
							Исключение
								ВызватьИсключение;
							КонецПопытки;				
						Исключение
							ВызватьИсключение
						КонецПопытки;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку()
						Попытка 
							а = 1+1;
							ВызватьИсключение("dsdsd dsds");	  
							f = 0;
							f = 0		
						Исключение
							а = 1+1;
							а = 1+1;
							ВызватьИсключение;  // в блоке Исключение можно вызывать без параметров
							а = 1+1;
							
							Если истина Тогда
								ВызватьИсключение;  // в блоке Исключение можно вызывать без параметров
							КонецЕсли
						КонецПопытки;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
						Попытка 
							а = 1+1;
							Если в = 1 И (а = 1 или у = 3) Тогда
								ВызватьИсключение "SDSDSD";
							КонецЕсли;
						Исключение
							а = 1+1;
							ВызватьИсключение "dsd";
							а = 1+1;
						КонецПопытки;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
						Попытка 
							Попытка 
								а = 1+1;					
							Исключение
								ВызватьИсключение;
							КонецПопытки;		

							ВызватьИсключение ;
						Исключение
							ВызватьИсключение
						КонецПопытки;
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "operator \"ВызватьИсключение\" without arguments can only be used when handling an exception. line: 9, column: 25 (unexpected literal: \";\")")
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
						Попытка 
							Попытка 
								а = 1+1;					
							Исключение
								ВызватьИсключение;
							КонецПопытки;
						Исключение
							ВызватьИсключение
						КонецПопытки;

						ВызватьИсключение 
					КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.ErrorContains(t, err, "without arguments can only be used when handling an exception")
	})
	t.Run("pass", func(t *testing.T) {
		code := `Функция Команда1НаСервере()

					ВызватьИсключение(НСтр("ru = 'Недостаточно прав на использование сертификата.'"),
						КатегорияОшибки.НарушениеПравДоступа);

			 КонецФункции`

		a := NewAST(code)
		err := a.Parse()
		if assert.NoError(t, err) {
			json, _ := a.JSON()
			assert.Contains(t, string(json), `{"Name":"","Body":[{"ExplicitVariables":{},"Name":"Команда1НаСервере","Directive":"","Body":[{"Param":{"Statements":[{"Name":"НСтр","Param":{"Statements":["ru = 'Недостаточно прав на использование сертификата.'"]}},{"Unit":{"Name":"НарушениеПравДоступа"},"Call":{"Name":"КатегорияОшибки"}}]}}],"Params":[],"Type":2,"Export":false}]}`)
		}
	})
}

func TestParseMethod(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					а = ТипыИзмененныхОбъектов.Найти(Тип)
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					а = ТипыИзмененныхОбъектов.Test.Найти(Тип)
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					а = ТипыИзмененныхОбъектов(Тип);
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку() 
					а = ТипыИзмененныхОбъектов..Найти(Тип)
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 32 (unexpected literal: \".\")")
	})
}

func TestParseFunctionProcedure(t *testing.T) {
	t.Run("Function", func(t *testing.T) {
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = 1 + gggg - (fd +1 / 3);
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)

			data, err := a.JSON()
			assert.NoError(t, err)
			assert.NotEqual(t, 0, len(data))
		})
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = парапапапам; 
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			if assert.NoError(t, err) {
				p := a.Print(PrintConf{OneLine: true})
				assert.Equal(t, "&НасервереБезКонтекста\nФункция ПодключитьВнешнююОбработку(Ссылка) f = парапапапам;КонецФункции", strings.TrimSpace(p))
			}
		})
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = 221;
						возврат 2+2;
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			fmt.Println(a.Print(PrintConf{Margin: 2}))
			if assert.NoError(t, err) {
				p := a.Print(PrintConf{OneLine: true})
				assert.Equal(t, "&НасервереБезКонтекста\nФункция ПодключитьВнешнююОбработку(Ссылка) f = 221;Возврат 2 + 2;КонецФункции", strings.TrimSpace(p))
			}
		})
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = "вававава авава";
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			if assert.NoError(t, err) {
				p := a.Print(PrintConf{OneLine: true})
				assert.Equal(t, "&НасервереБезКонтекста\nФункция ПодключитьВнешнююОбработку(Ссылка) f = \"вававава авава\";КонецФункции", strings.TrimSpace(p))
			}
		})
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = Истина;
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			if assert.NoError(t, err) {
				p := a.Print(PrintConf{OneLine: true})
				assert.Equal(t, "&НасервереБезКонтекста\nФункция ПодключитьВнешнююОбработку(Ссылка) f = Истина;КонецФункции", strings.TrimSpace(p))
			}
		})
		t.Run("ast", func(t *testing.T) {
			code := `&НасервереБезКонтекста
					Функция ПодключитьВнешнююОбработку(Ссылка) 
						f = Ложь;
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			if assert.NoError(t, err) {
				p := a.Print(PrintConf{OneLine: true})
				assert.Equal(t, "&НасервереБезКонтекста\nФункция ПодключитьВнешнююОбработку(Ссылка) f = Ложь;КонецФункции", strings.TrimSpace(p))
			}
		})
		t.Run("bad directive", func(t *testing.T) {
			code := `&НасервереБез
					Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 1, column: 1 (unexpected literal: \"НасервереБез\")")
		})
		t.Run("without directive", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка, вы, выыыыы) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("return", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка, вы, выыыыы) 
						Перем а;
						Перем вы, в;

						Если истина Тогда
							ВызватьИсключение вызов("");
						ИначеЕсли 1 = 1 Тогда
						ИначеЕсли 2 = 2 Тогда
						Иначе
							б = а;
						КонецЕсли;
					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)

			// fmt.Println(a.Print(&PrintConf{Margin: 2}))
		})
		t.Run("return", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка, вы, выыыыы) 
						Перем а;
						Перем вы, в;

						Если истина Тогда
							Возврат;
						КонецЕсли;
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
			assert.EqualError(t, err, "syntax error. line: 3, column: 5 (unexpected literal: \"КонецФунки\")")
		})
		t.Run("error", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Ссылка) 

					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 3, column: 5 (unexpected literal: \"КонецПроцедуры\")")
		})
		t.Run("params def value", func(t *testing.T) {
			code := `Функция ПодключитьВнешнююОбработку(Парам1, Парам2 = Неопределено, Знач Парам3 = "вывыв", парам4 = 4) 

					КонецФункции`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
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
			assert.EqualError(t, err, "syntax error. line: 1, column: 1 (unexpected literal: \"НасервереБез\")")
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
			assert.EqualError(t, err, "syntax error. line: 3, column: 5 (unexpected literal: \"КонецФункции\")")
		})
		t.Run("with var pass", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Перем а;
						Перем вы, в;

						Если истина Тогда
							ВызватьИсключение "";
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
			assert.Equal(t, 3, len(a.ModuleStatement.Body[0].(*FunctionOrProcedure).ExplicitVariables))
		})
		t.Run("with var error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Перем а;
						Перем а, вы, в;

						Если истина Тогда
							ВызватьИсключение "";
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.ErrorContains(t, err, "variable has already been defined")
		})
		t.Run("return", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Перем а;
						Перем вы, в;

						Если истина Тогда
							возврат;
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("return error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Перем а;
						Перем а, вы, в;

						Если истина Тогда
							возврат "";
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.ErrorContains(t, err, "procedure cannot return a value")
		})
		t.Run("with var error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
						Перем а;
						Перем а, вы, в

						Если истина Тогда
							ВызватьИсключение "";
						КонецЕсли;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 5, column: 6 (unexpected literal: \"Если\")")
		})
		t.Run("with var error", func(t *testing.T) {
			code := `Процедура ПодключитьВнешнююОбработку(Ссылка)
						Если истина Тогда
							ВызватьИсключение "";
						КонецЕсли;

						Перем а, вы, в;
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.EqualError(t, err, "syntax error. line: 6, column: 6 (unexpected literal: \"Перем\")")
		})
		t.Run("with region", func(t *testing.T) {
			code := `#Область ПрограммныйИнтерфейс
// hg
#Область ПрограммныйИнтерфейс
					&НасервереБезКонтекста
					Процедура ПодключитьВнешнююОбработку()
						ТипЗначенияСтрокой = XMLТипЗнч(КлючДанных).ИмяТипа;

					КонецПроцедуры
					#КонецОбласти
#КонецОбласти

					#Область СлужебныеПроцедурыИФункции
					&НасервереБезКонтекста
						Функция ПодключитьВнешнююОбработку() 
							ВызватьИсключение "Нет соответствия шаблону! " + СтрокаТекста;

						КонецФункции
					#КонецОбласти`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)
		})
		t.Run("through_dot pass", func(t *testing.T) {
			code := `Процедура ЗагрузитьОбъекты(Задание, Отказ = Ложь) Экспорт
						Перем СоответствиеРеквизитовШапки;
					
						Организация  = Задание.Организация.ВыполнитьМетодСПараметрами(1, "ав", авава);
						Организация  = Задание.Организация.ВыполнитьМетодБезПараметров();
						Организация  = Задание.Организация.Код;
 
					КонецПроцедуры`

			a := NewAST(code)
			err := a.Parse()
			assert.NoError(t, err)

			data, err := a.JSON()
			assert.NoError(t, err)
			assert.NotEqual(t, 0, len(data))
		})
	})
	t.Run("many", func(t *testing.T) {
		code := `&Насервере
				Процедура ПодключитьВнешнююОбработку() 
					Возврат
				КонецПроцедуры

				&НаКлиенте
				Функция ОчиститьПараметрыТЖ(парам1 = 1, парам2 = Неопределено, парам3 = -1) Экспорт
					Возврат 100;
				КонецФункции

				Функция ПарамТарам(Знач парам1)
					возврат +1;
				КонецФункции`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
		if !t.Failed() {
			p := a.Print(PrintConf{Margin: 0})
			assert.Equal(t, "&Насервере\nПроцедура ПодключитьВнешнююОбработку() \nВозврат;\nКонецПроцедуры \n&НаКлиенте\nФункция ОчиститьПараметрыТЖ(парам1 = 1, парам2 = Неопределено, парам3 = -1) Экспорт \nВозврат 100;\nКонецФункции \nФункция ПарамТарам(Знач парам1) \nВозврат 1;\nКонецФункции", deleteEmptyLine(p))
		}
	})
}

func TestParseBaseExpression(t *testing.T) {
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) ds; КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) ds = 222; uu = 9; КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = 222; ds = 222; uu = 9
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
					uu = 9;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 3, column: 5 (unexpected literal: \"uu\")")
	})
	t.Run("pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = ИспользуемыеНастройки[0].Структура[0].Структура;
					fdfd = СтруктураКонтрагент();
					fdfd = f.СтруктураКонтрагент(gf, ghf);
					СтруктураКонтрагент.Наименование = СтрокаВывода[РезультатВывода.Колонки.Найти("СтруктураКонтрагентНаименование").Имя];
					СтрокаСпискаПП[ТекКолонка.Ключ]["РасшифровкаПлатежа"].Добавить(ВременнаяСтруктура);
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = ИспользуемыеНастройки[0].Структура[0].Структура;
					fdfd = СтруктураКонтрагент();
					fdfd = f.СтруктураКонтрагент(gf, ghf);
					СтруктураКонтрагент.Наименование = СтрокаВывода[РезультатВывода.Колонки.Найти("СтруктураКонтрагентНаименование.Имя];
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.Error(t, err)
	})
	t.Run("new pass", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					Контекст = Новый Структура;
					Контекст = Новый Структура();
					Контекст = Новый Структура("выыыы");
					Контекст = Новый Структура(какойтофункшин());
					Контекст = Новый Структура("какойтоимя", чето);
					Запрос = Новый Запрос(ТекстЗапросаЗадание());
					Оповещение = Новый ОписаниеОповещения(,, Контекст,
													"ОткрытьНавигационнуюСсылкуПриОбработкеОшибки", ОбщегоНазначенияСлужебныйКлиент);
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.NoError(t, err)
	})
	t.Run("new error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					Контекст = Новый Структура(;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.EqualError(t, err, "syntax error. line: 2, column: 32 (unexpected literal: \";\")")
	})
	t.Run("new error", func(t *testing.T) {
		code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					Контекст = Новый Структура("выыыы);
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()
		assert.Error(t, err)
	})
}

func TestParseAST(t *testing.T) {
	code := `

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
КонецПроцедуры

Если Оповещение <> Неопределено Тогда 
			ПриложениеЗапущено = Истина;
			ВыполнитьОбработкуОповещения(Оповещение, ПриложениеЗапущено);
		КонецЕсли;`

	a := NewAST(code)
	err := a.Parse()

	if assert.NoError(t, err) {
		p := a.Print(PrintConf{Margin: 4})
		assert.Equal(t, true, compareHashes(code, p))
	}
}

func TestParseEmpty(t *testing.T) {
	code := `

`

	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)
}

func TestBigProcedure(t *testing.T) {
	if _, err := os.Stat("testdata"); errors.Is(err, os.ErrNotExist) {
		t.Fatal("testdata file not found")
	}

	fileData, err := os.ReadFile("testdata")
	assert.NoError(t, err)

	a := NewAST(string(fileData))
	s := time.Now()
	err = a.Parse()
	fmt.Println("milliseconds -", time.Since(s).Milliseconds())
	assert.NoError(t, err)

	// p := a.Print(&PrintConf{Margin: 4})
	// fmt.Println(p)
}

func TestTernaryOperator(t *testing.T) {
	code := `Процедура ПодключитьВнешнююОбработку(Ссылка) 
					ds = ?(Истина, ?(dd = 3, а = 1, Наименование), СтруктураКонтрагент.Наименование);
				КонецПроцедуры`
	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)

	data, err := a.JSON()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(data))
}

func TestArrayStruct(t *testing.T) {
	code := `Процедура ПодключитьВнешнююОбработку()        
				м = Новый Массив();
				в = м[4];
				
				м = Новый Структура("ав", уцуцу);
				в = м["вывыв"];
			КонецПроцедуры`

	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)

	data, err := a.JSON()
	assert.NoError(t, err)
	assert.NotEqual(t, 0, len(data))
}

func TestPrint(t *testing.T) {
	code := `&НаКлиенте
					Процедура Проба() 
						Если в = 1 или у = 3 И 0 <> 3 и не гоого и Истина и ав = неопределено Тогда
							а=1 + 3 *4;
						а=1 + 3 *4;
							fgd = 1
						ИначеЕсли ввв Тогда Если в = 1 Тогда
							а = -(1 + 3 *4);
Если в = 1 Тогда
							а = 1 + 3 *4;
							КонецЕсли
							КонецЕсли
						ИначеЕсли авыав Тогда

						Иначе						
							ваывы = 1 + 3 *4;
ваывы = 1 + 3 *4
						КонецЕсли;

						а = 1 + 3 *4
					КонецПроцедуры

					Функция авава(пар1, пар2) экспорт
						Для Каждого ИзмененныйОбъект Из ОбъектыНазначения Цикл
							Тип = ТипЗнч(ИзмененныйОбъект);
							Если ТипыИзмененныхОбъектов = Неопределено Тогда
								ТипыИзмененныхОбъектов.Добавить(Тип);
							КонецЕсли;

Для Каждого ИзмененныйОбъект Из ОбъектыНазначения Цикл
        Тип = ТипЗнч(ИзмененныйОбъект);
        Если ТипыИзмененныхОбъектов = Неопределено Тогда
            ТипыИзмененныхОбъектов.Добавить(Тип);
        Иначе
        КонецЕсли;
    КонецЦикла;

						КонецЦикла;

						Для а = 0 По 100 Цикл
							Тип = ТипЗнч(ИзмененныйОбъект);
							Если ТипыИзмененныхОбъектов  = Неопределено Тогда
								Продолжить; Иначе 	Прервать;
							КонецЕсли;
						КонецЦикла;
					Конецфункции
					
					Процедура Опрпп(пар1, Знач пар2 = 2.2, пар1 = Неопределено, Пар3 = "авава") 

						Попытка 
							а = 1+1;
ВызватьИсключение ававава();
						Исключение
							ВызватьИсключение "";
ВызватьИсключение ;
						КонецПопытки;
					Конецпроцедуры`

	a := NewAST(code)
	err := a.Parse()
	assert.NoError(t, err)

	// p := a.Print(&PrintConf{Margin: 4})
	// fmt.Println(p)
}

func TestExpPriority(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		code := `А = d = 2 = d ИЛИ в = 3;
			Если 1 = 1 = 2 = 3 Тогда
			   ПриКомпоновкеРезультата();
			КонецЕсли`
		a := NewAST(code)
		err := a.Parse()
		if assert.NoError(t, err) {
			p := a.Print(PrintConf{Margin: 4, LispStyle: true})
			//fmt.Println(p)
			assert.Equal(t, "А=(((d=2)=d)ИЛИ(в=3));Если(((1=1)=2)=3)ТогдаПриКомпоновкеРезультата();КонецЕсли;", normalize(p))
		}
	})
	t.Run("test2", func(t *testing.T) {
		code := `Процедура ОткрытьНавигационнуюСсылку(НавигационнаяСсылка, Знач Оповещение = Неопределено) Экспорт
					Если в = 1 = 5 и не авав ИЛИ ааа Тогда
						в = 1 = 5 = 1 и не авав ИЛИ ааа;
					КонецЕсли;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()

		if assert.NoError(t, err) {
			p := a.Print(PrintConf{Margin: 4, LispStyle: true})
			assert.Equal(t, "ПроцедураОткрытьНавигационнуюСсылку(НавигационнаяСсылка,ЗначОповещение=Неопределено)ЭкспортЕсли((((в=1)=5)ИНеавав)ИЛИааа)Тогдав=((((1=5)=1)ИНеавав)ИЛИааа);КонецЕсли;КонецПроцедуры", normalize(p))
		}
	})
	t.Run("test3", func(t *testing.T) {
		code := `Процедура f()
					тест.куку.ууу = 1 = 5 = 1 и не авав ИЛИ ааа;
					тест[333] = 1 = 5 = 1 = 4 = fd;
				КонецПроцедуры`

		a := NewAST(code)
		err := a.Parse()

		if assert.NoError(t, err) {
			p := a.Print(PrintConf{OneLine: true, LispStyle: true})
			assert.Equal(t, "Процедура f() тест.куку.ууу = ((((1 = 5) = 1) И Не авав) ИЛИ ааа);тест[333] = ((((1 = 5) = 1) = 4) = fd);КонецПроцедуры", strings.TrimSpace(p))
		}
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
	b.Run("string count - 1", func(b *testing.B) {
		str := "rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd dfdf dsfd"
		for i := 0; i < b.N; i++ {
			strings.Count(str, "df")
		}
	})
	//b.Run("string count - 2", func(b *testing.B) {
	//	str := "rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs rdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd rdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfdrdedfs dfdf dsfd dfdf dsfd"
	//	for i := 0; i < b.N; i++ {
	//		stringCount(str, "df")
	//	}
	//})
}

func test(_ string)    {}
func testPt(_ *string) {}

func compareHashes(str1, str2 string) bool {
	str1 = normalize(str1)
	str2 = normalize(str2)

	hash1 := sha256.Sum256([]byte(fastToLower(str1)))
	hash2 := sha256.Sum256([]byte(fastToLower(str2)))

	return hash1 == hash2
}

func normalize(str string) string {
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "\t", "")
	str = strings.ReplaceAll(str, "\n", "")

	return str
}

func deleteEmptyLine(str string) string {
	result := strings.Builder{}
	for _, line := range strings.Split(str, "\n") {
		if strings.TrimSpace(line) != "" {
			result.WriteString(line + "\n")
		}
	}

	return strings.TrimSpace(result.String())
}
