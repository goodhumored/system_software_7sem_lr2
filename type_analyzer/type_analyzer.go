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
	types     map[string]Type
	variables map[string]VariableInfo
}

func (ta TypeAnalyzer) GetVariablesMemory() []VariableInfo {
	variableInfos := []VariableInfo{}
	for _, v := range ta.variables {
		variableInfos = append(variableInfos, v)
	}
	return variableInfos
}

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

func (ta *TypeAnalyzer) calculateTypeSize(typeNode *parse_tree.Node) (int, error) {
	if typeNode.Symbol == nonterminal.Record {
		return ta.calculateRecordSize(typeNode)
	}
	typeInfo, ok := ta.types[typeNode.Value]
	if !ok {
		return 0, fmt.Errorf("unknown type %s\n", typeNode.Value)
	}
	return typeInfo.Size, nil
}

func (ta TypeAnalyzer) calculateRecordSize(recordNode *parse_tree.Node) (int, error) {
	totalSize := 0
	varDeclarations := recordNode.Children[1 : len(recordNode.Children)-1]
	for _, varDeclaration := range varDeclarations {
		size, err := ta.calculateVarSize(varDeclaration)
		if err != nil {
			return 0, err
		}
		if size <= 4 {
			totalSize += calculateOffset(totalSize, size)
		} else {
			totalSize += calculateOffset(totalSize, 8)
		}
		totalSize += size
	}
	totalSize += calculateOffset(totalSize, 8)
	return totalSize, nil
}

func (ta TypeAnalyzer) calculateVarSize(varDeclarationNode *parse_tree.Node) (int, error) {
	typeNode := varDeclarationNode.Children[2]
	return ta.calculateTypeSize(typeNode)
}

func NewTypeAnalyzer() TypeAnalyzer {
	return TypeAnalyzer{
		types:     primitivesTypeMap(),
		variables: map[string]VariableInfo{},
	}
}

func primitivesTypeMap() map[string]Type {
	typeMap := map[string]Type{}
	typeMap["byte"] = Type{Name: "byte", Size: 1}
	typeMap["extended"] = Type{Name: "extended", Size: 10}
	return typeMap
}

func calculateOffset(totalSize int, allignment int) int {
	if offset := totalSize % allignment; offset > 0 {
		return allignment - offset
	}
	return 0
}
