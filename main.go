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
		println("Исходный код:")
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
	typeAnalyzer := typeanalyzer.NewTypeAnalyzer()
	err = typeAnalyzer.AnalyzeTypes(tree)
	if err != nil {
		fmt.Printf("Ошибка при подсчёте выделяемой памяти: %s\n", err)
		return
	}
	varInfos := typeAnalyzer.GetVariablesMemory()
	totalMemo := 0
	fmt.Println("\nВыделение памяти под переменные: ")
	for i, varInfo := range varInfos {
		fmt.Printf("%d) %s: %d Байт\n", i, varInfo.Name, varInfo.Size)
		totalMemo += varInfo.Size
	}
	fmt.Printf("\nВсего памяти выделяется под переменные: %v Байт\n\n", totalMemo)
}

// Читает файл с входными данными, вызывает панику в случае неудачи
func getInput(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
