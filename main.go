package main

import (
	"fmt"
	"os"

	"goodhumored/lr2_types_memory/syntax_analyzer"
	"goodhumored/lr2_types_memory/token_analyzer"
	typeanalyzer "goodhumored/lr2_types_memory/type_analyzer"
)

func main() {
	source := getInput("./input-everything.txt") // читаем файл

	defer func() {
		println("\n=================Исходный код:=================\n")
		fmt.Println(source)
	}()

	// запускаем распознание лексем
	tokenTable := token_analyzer.RecogniseTokens(source)

	// выводим лексемы
	fmt.Println("Таблица лексем:")
	fmt.Println(tokenTable)

	// Проверяем на ошибки
	if errors := tokenTable.GetErrors(); len(errors) > 0 {
		fmt.Printf("Во время лексического анализа было обнаружено: %d ошибок:\n", len(errors))
		for _, error := range errors {
			fmt.Printf("Неожиданный символ '%s'\n", error.Value())
		}
		return
	}

	// запускаем синтаксический анализатор
	tree, err := syntax_analyzer.AnalyzeSyntax(rulesTable, *tokenTable, precedenceMatrix)
	if err != nil {
		fmt.Printf("Ошибка при синтаксическом анализе строки: %s\n", err)
		return
	} else {
		fmt.Println("Строка принята!!!")
		tree.Print()
	}
	typeAnalyzer := typeanalyzer.NewTypeAnalyzer(8, true)
	typeAnalyzerWithoutAlignment := typeanalyzer.NewTypeAnalyzer(1, false)
	err = typeAnalyzer.AnalyzeTypes(tree)
	_ = typeAnalyzerWithoutAlignment.AnalyzeTypes(tree)
	if err != nil {
		fmt.Printf("Ошибка при подсчёте выделяемой памяти: %s\n", err)
		return
	}
	println("\n\n===============================Выделение памяти с кратностью распределения:=====================")
	typeAnalyzer.PrintGatheredInfo()
	println("\n\n===============================Выделение памяти без кратности распределения:=====================")
	typeAnalyzerWithoutAlignment.PrintGatheredInfo()
}

// Читает файл с входными данными, вызывает панику в случае неудачи
func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
