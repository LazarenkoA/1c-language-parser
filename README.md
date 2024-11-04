### Парсер языка 1С
Этот репозиторий содержит парсер для языка 1С, написанный на Go. Парсер использует yacc для эффективного синтаксического анализа и создает абстрактное синтаксическое дерево (AST) представления разобранного кода 1С.

### Использование
`go get  github.com/LazarenkoA/1c-language-parser@master`

```go
package main

import (
	"fmt"

	"github.com/LazarenkoA/1c-language-parser/ast"
)

func main() {
	code := `Процедура ПодключитьВнешнююОбработку() 
                Если в = 1 И а = 1 или у = 3 Тогда

                КонецЕсли
            КонецПроцедуры`

	a := ast.NewAST(code)
	if err := a.Parse(); err == nil {
		jdata, _ := a.JSON()
		fmt.Println(string(jdata))
	}
}

```
### Примеры использования
* [examples/pretty_code](examples/pretty_code)
* [Obfuscator-1C](https://github.com/LazarenkoA/Obfuscator-1C)
* [FuncGraphView](https://github.com/LazarenkoA/FuncGraphView)


### Примеры AST
Вот несколько примеров, демонстрирующих возможности парсера языка 1С:
```1C
Процедура Пример()
	Сообщить("Привет, мир!");
КонецПроцедуры
```
AST:
```
main.ModuleStatement{
  Name: "",
  Body: []main.Statement{
    main.FunctionOrProcedure{
      Type: 1,
      Name: "Пример",
      Body: []main.Statement{
        main.MethodStatement{
          Name:  "Сообщить",
          Param: []main.Statement{
            "Привет, мир!",
          },
        },
      },
      Export:            false,
      Params:            []main.ParamStatement{},
      Directive:         "",
      ExplicitVariables: map[string]main.VarStatement{},
    },
  },
}
```


```1C
Если a = b и c = 8 или истина Тогда
	Сообщить("Условие выполнено");
ИначеЕсли ВтороеУсловие Тогда
	Сообщить("Второе условие выполнено");
Иначе
	Сообщить("Ни одно из условий не выполнено");
КонецЕсли

```
AST:
```
main.ModuleStatement{
  Name: "",
  Body: []main.Statement{
    main.FunctionOrProcedure{
      Type: 1,
      Name: "Пример",
      Body: []main.Statement{
        main.IfStatement{
          Expression: main.ExpStatement{
            Operation: 11,
            Left:      main.ExpStatement{
              Operation: 12,
              Left:      main.ExpStatement{
                Operation: 4,
                Left:      main.VarStatement{
                  Name:  "a",
                  unary: false,
                  not:   false,
                },
                Right: main.VarStatement{
                  Name:  "b",
                  unary: false,
                  not:   false,
                },
                unary: false,
                not:   false,
              },
              Right: main.ExpStatement{
                Operation: 4,
                Left:      main.VarStatement{
                  Name:  "c",
                  unary: false,
                  not:   false,
                },
                Right: 8.000000,
                unary: false,
                not:   false,
              },
              unary: false,
              not:   false,
            },
            Right: true,
            unary: false,
            not:   false,
          },
          TrueBlock: []main.Statement{
            main.MethodStatement{
              Name:  "Сообщить",
              Param: []main.Statement{
                "Условие выполнено",
              },
            },
          },
          IfElseBlock: []*main.IfStatement{
            &main.IfStatement{
              Expression: main.VarStatement{
                Name:  "ВтороеУсловие",
                unary: false,
                not:   false,
              },
              TrueBlock: []main.Statement{
                main.MethodStatement{
                  Name:  "Сообщить",
                  Param: []main.Statement{
                    "Второе условие выполнено",
                  },
                },
              },
              IfElseBlock: []*main.IfStatement{},
              ElseBlock:   []main.Statement{},
            },
          },
          ElseBlock: []main.Statement{
            main.MethodStatement{
              Name:  "Сообщить",
              Param: []main.Statement{
                "Ни одно из условий не выполнено",
              },
            },
          },
        },
      },
      Export:            false,
      Params:            []main.ParamStatement{},
      Directive:         "",
      ExplicitVariables: map[string]main.VarStatement{},
    },
  },
}

```

<details>
<summary>Более сложный пример (код случайный)</summary>


```1C
Процедура ОткрытьНавигационнуюСсылку(НавигационнаяСсылка, Знач Оповещение = Неопределено) Экспорт
	
	Контекст = Новый Структура;
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
```

</details>

<details>
<summary>получаем такое AST</summary>

```
main.ModuleStatement{
  Name: "",
  Body: []main.Statement{
    main.FunctionOrProcedure{
      Type: 1,
      Name: "ОткрытьНавигационнуюСсылку",
      Body: []main.Statement{
        main.ExpStatement{
          Operation: 4,
          Left:      main.VarStatement{
            Name:  "Контекст",
            unary: false,
            not:   false,
          },
          Right: main.NewObjectStatement{
            Constructor: "Структура",
            Param:       []main.Statement{},
          },
          unary: false,
          not:   false,
        },
        main.CallChainStatement{
          Unit: main.MethodStatement{
            Name:  "Вставить",
            Param: []main.Statement{
              "НавигационнаяСсылка",
              main.VarStatement{
                Name:  "НавигационнаяСсылка",
                unary: false,
                not:   false,
              },
            },
          },
          Call: main.VarStatement{
            Name:  "Контекст",
            unary: false,
            not:   false,
          },
        },
        main.CallChainStatement{
          Unit: main.MethodStatement{
            Name:  "Вставить",
            Param: []main.Statement{
              "Оповещение",
              main.VarStatement{
                Name:  "Оповещение",
                unary: false,
                not:   false,
              },
            },
          },
          Call: main.VarStatement{
            Name:  "Контекст",
            unary: false,
            not:   false,
          },
        },
        main.ExpStatement{
          Operation: 4,
          Left:      main.VarStatement{
            Name:  "ОписаниеОшибки",
            unary: false,
            not:   false,
          },
          Right: main.CallChainStatement{
            Unit: main.MethodStatement{
              Name:  "ПодставитьПараметрыВСтроку",
              Param: []main.Statement{
                main.MethodStatement{
                  Name:  "НСтр",
                  Param: []main.Statement{
                    "ru = 'Не удалось перейти по ссылке \"\"%1\"\" по причине: \n\t\t\t           |Неверно задана навигационная ссылка.'",
                  },
                },
                main.VarStatement{
                  Name:  "НавигационнаяСсылка",

                  unary: false,
                  not:   false,
                },
              },
            },
            Call: main.VarStatement{
              Name:  "СтроковыеФункцииКлиентСервер",
              unary: false,
              not:   false,
            },
          },
          unary: false,
          not:   false,
        },
        main.IfStatement{
          Expression: main.CallChainStatement{
            Unit: main.MethodStatement{
              Name:  "ЭтоДопустимаяСсылка",
              Param: []main.Statement{
                main.VarStatement{
                  Name:  "НавигационнаяСсылка",
                  unary: false,
                  not:   false,
                },
              },
            },
            Call: main.VarStatement{
              Name:  "ОбщегоНазначенияСлужебныйКлиент",
              unary: false,
              not:   false,
            },
          },
          TrueBlock: []main.Statement{
            main.CallChainStatement{
              Unit: main.MethodStatement{
                Name:  "ОткрытьНавигационнуюСсылкуОповеститьОбОшибке",
                Param: []main.Statement{
                  main.VarStatement{
                    Name:  "ОписаниеОшибки",
                    unary: false,
                    not:   false,
                  },
                  main.VarStatement{
                    Name:  "Контекст",
                    unary: false,
                    not:   false,
                  },
                },
              },
              Call: main.VarStatement{
                Name:  "ОбщегоНазначенияСлужебныйКлиент",
                unary: false,
                not:   false,
              },
            },
            main.ReturnStatement{
              Param: nil,
            },
          },
          IfElseBlock: []*main.IfStatement{},
          ElseBlock:   []main.Statement{},
        },
        main.IfStatement{
          Expression: main.ExpStatement{
            Operation: 11,
            Left:      main.CallChainStatement{
              Unit: main.MethodStatement{
                Name:  "ЭтоВебСсылка",
                Param: []main.Statement{
                  main.VarStatement{
                    Name:  "НавигационнаяСсылка",
                    unary: false,
                    not:   false,
                  },
                },
              },
              Call: main.VarStatement{
                Name:  "ОбщегоНазначенияСлужебныйКлиент",
                unary: false,
                not:   false,
              },
            },
            Right: main.CallChainStatement{
              Unit: main.MethodStatement{
                Name:  "ЭтоНавигационнаяСсылка",
                Param: []main.Statement{
                  main.VarStatement{
                    Name:  "НавигационнаяСсылка",
                    unary: false,
                    not:   false,
                  },
                },
              },
              Call: main.VarStatement{
                Name:  "ОбщегоНазначенияСлужебныйКлиент",
                unary: false,
                not:   false,
              },
            },
            unary: false,
            not:   false,
          },
          TrueBlock: []main.Statement{
            main.TryStatement{
              Body: []main.Statement{
                main.ExpStatement{
                  Operation: 4,
                  Left:      main.VarStatement{
                    Name:  "а",
                    unary: false,
                    not:   false,
                  },
                  Right: main.ExpStatement{
                    Operation: 3,
                    Left:      main.VarStatement{
                      Name:  "а",
                      unary: false,
                      not:   false,
                    },
                    Right: 0.000000,
                    unary: false,
                    not:   false,
                  },
                  unary: false,
                  not:   false,
                },
              },
              Catch: []main.Statement{
                main.CallChainStatement{
                  Unit: main.MethodStatement{
                    Name:  "ОткрытьНавигационнуюСсылкуОповеститьОбОшибке",
                    Param: []main.Statement{
                      main.VarStatement{
                        Name:  "ОписаниеОшибки",
                        unary: false,
                        not:   false,
                      },
                      main.VarStatement{
                        Name:  "Контекст",
                        unary: false,
                        not:   false,
                      },
                    },
                  },
                  Call: main.VarStatement{
                    Name:  "ОбщегоНазначенияСлужебныйКлиент",
                    unary: false,
                    not:   false,
                  },
                },
                main.ReturnStatement{
                  Param: nil,
                },
              },
            },
            main.IfStatement{
              Expression: main.ExpStatement{
                Operation: 7,
                Left:      main.VarStatement{
                  Name:  "Оповещение",
                  unary: false,
                  not:   false,
                },
                Right: nil,
                unary: false,
                not:   false,
              },
              TrueBlock: []main.Statement{
                main.ExpStatement{
                  Operation: 4,
                  Left:      main.VarStatement{
                    Name:  "ПриложениеЗапущено",
                    unary: false,
                    not:   false,
                  },
                  Right: true,
                  unary: false,
                  not:   false,
                },
                main.MethodStatement{
                  Name:  "ВыполнитьОбработкуОповещения",
                  Param: []main.Statement{
                    main.VarStatement{
                      Name:  "Оповещение",
                      unary: false,
                      not:   false,
                    },
                    main.VarStatement{
                      Name:  "ПриложениеЗапущено",
                      unary: false,
                      not:   false,
                    },
                  },
                },
              },
              IfElseBlock: []*main.IfStatement{},
              ElseBlock:   []main.Statement{},
            },
            main.ReturnStatement{
              Param: nil,
            },
          },
          IfElseBlock: []*main.IfStatement{},
          ElseBlock:   []main.Statement{},
        },
        main.IfStatement{
          Expression: main.CallChainStatement{
            Unit: main.MethodStatement{
              Name:  "ЭтоСсылкаНаСправку",
              Param: []main.Statement{
                main.VarStatement{
                  Name:  "НавигационнаяСсылка",
                  unary: false,
                  not:   false,
                },
              },
            },
            Call: main.VarStatement{
              Name:  "ОбщегоНазначенияСлужебныйКлиент",
              unary: false,
              not:   false,
            },
          },
          TrueBlock: []main.Statement{
            main.MethodStatement{
              Name:  "ОткрытьСправку",
              Param: []main.Statement{
                main.VarStatement{
                  Name:  "НавигационнаяСсылка",
                  unary: false,
                  not:   false,
                },
              },
            },
            main.ReturnStatement{
              Param: nil,
            },
          },
          IfElseBlock: []*main.IfStatement{},
          ElseBlock:   []main.Statement{},
        },
      },
      Export: true,
      Params: []main.ParamStatement{
        main.ParamStatement{
          Name:    "НавигационнаяСсылка",
          IsValue: false,
          Default: nil,
        },
        main.ParamStatement{
          Name:    "Оповещение",
          IsValue: true,
          Default: main.UndefinedStatement{},
        },
      },
      Directive:         "",
      ExplicitVariables: map[string]main.VarStatement{},
    },
  },
}
```
</details>

Если у вас возникнут проблемы или у вас есть предложения по улучшению, не стесняйтесь создать [issue](https://github.com/LazarenkoA/ast_parser_1c/issues). Также приветствуются ваши вклады!
