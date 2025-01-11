package typeanalyzer

import (
	"fmt"

	"goodhumored/lr2_types_memory/syntax_analyzer/nonterminal"
	"goodhumored/lr2_types_memory/syntax_analyzer/parse_tree"
)

type VariableInfo struct {
	Name string
	Size int
}

// Тип анализатор принимает на входе дерево
// считает какой тип сколько занимает
// считает какая переменная сколько требует памяти
type TypeAnalyzer struct {
	types                     map[string]Type
	variables                 map[string]VariableInfo
	memoryAllocationAlignment int
	alignStructureElements    bool
}

// Функция возвращающая собранную в результате анализа типов информацию о переменных и их размерах
func (ta TypeAnalyzer) GetVariablesMemory() []VariableInfo {
	variableInfos := []VariableInfo{}
	for _, v := range ta.variables {
		variableInfos = append(variableInfos, v)
	}
	return variableInfos
}

func (ta TypeAnalyzer) PrintGatheredInfo() {
	varInfos := ta.GetVariablesMemory()
	totalMemo := 0
	fmt.Println("\nВыделение памяти под переменные: ")
	for i, varInfo := range varInfos {
		fmt.Printf("%d) %s: %d Байт\n", i, varInfo.Name, varInfo.Size)
		totalMemo += varInfo.Size
	}
	fmt.Printf("\nВсего памяти выделяется под переменные: %v Байт\n", totalMemo)
}

// функция для анализа типов входящего дерева вывода
func (ta TypeAnalyzer) AnalyzeTypes(tree parse_tree.ParseTree) error {
	typeDeclarationNodes := tree.BFS(nonterminal.TypeDeclaration)
	varDeclarationNodes := tree.BFS(nonterminal.VarBlock)[0].Children[1:]
	err := ta.fillTypeTable(typeDeclarationNodes)
	if err != nil {
		return err
	}
	err = ta.fillVarTable(varDeclarationNodes)
	if err != nil {
		return err
	}
	return nil
}

// функция для поиска объявления типов в дереве и заполнения таблицы типов
func (ta *TypeAnalyzer) fillTypeTable(nodes []*parse_tree.Node) error {
	for _, typeDeclarationNode := range nodes {
		name := typeDeclarationNode.Children[0].Value
		typeNode := typeDeclarationNode.Children[2]
		size, err := ta.calculateTypeSize(typeNode)
		if err != nil {
			return err
		}
		ta.types[name] = Type{Name: name, Size: size}
	}
	return nil
}

// Функция для поиска объявления переменных в дереве и заполнения таблицы переменных
func (ta *TypeAnalyzer) fillVarTable(nodes []*parse_tree.Node) error {
	for _, varNode := range nodes {
		name := varNode.Children[0].Value
		typeNode := varNode.Children[2]
		size, err := ta.calculateTypeSize(typeNode)
		if err != nil {
			return err
		}
		ta.variables[name] = VariableInfo{Name: name, Size: size}
	}
	return nil
}

// Функция для подсчёта размера типа
func (ta *TypeAnalyzer) calculateTypeSize(typeNode *parse_tree.Node) (int, error) {
	if typeNode.Symbol == nonterminal.Record {
		return ta.calculateRecordSize(typeNode)
	}
	typeInfo, ok := ta.types[typeNode.Value]
	if !ok {
		return 0, fmt.Errorf("Неизвестный тип %s\n", typeNode.Value)
	}
	return typeInfo.Size, nil
}

// Функция для подсчёта размера структуры
func (ta TypeAnalyzer) calculateRecordSize(recordNode *parse_tree.Node) (int, error) {
	totalSize := 0
	varDeclarations := recordNode.Children[1 : len(recordNode.Children)-1]
	for _, varDeclaration := range varDeclarations {
		size, err := ta.calculateVarSize(varDeclaration)
		if err != nil {
			return 0, err
		}
		if ta.alignStructureElements {
			totalSize += calculateOffset(totalSize, ta.memoryAllocationAlignment)
		}
		totalSize += size
	}
	totalSize += calculateOffset(totalSize, ta.memoryAllocationAlignment)
	return totalSize, nil
}

// Функция для подсчёта размера переменной
func (ta TypeAnalyzer) calculateVarSize(varDeclarationNode *parse_tree.Node) (int, error) {
	typeNode := varDeclarationNode.Children[2]
	return ta.calculateTypeSize(typeNode)
}

// функция для создания инстанса анализатора типов
func NewTypeAnalyzer(memoryAllocationAlignment int, alignStructureElements bool) TypeAnalyzer {
	return TypeAnalyzer{
		types:                     primitivesTypeMap(),
		variables:                 map[string]VariableInfo{},
		memoryAllocationAlignment: memoryAllocationAlignment,
		alignStructureElements:    alignStructureElements,
	}
}

// Функция возвращающая соответствие размеров скалярных встроенных типов их размерам
func primitivesTypeMap() map[string]Type {
	typeMap := map[string]Type{}
	typeMap["byte"] = Type{Name: "byte", Size: 1}
	typeMap["extended"] = Type{Name: "extended", Size: 10}
	return typeMap
}

// функция для рассчёта необходимого сдвига в памяти для выравнивания
func calculateOffset(totalSize int, allignment int) int {
	if offset := totalSize % allignment; offset > 0 {
		return allignment - offset
	}
	return 0
}
