package parse_tree

import (
	"fmt"

	"goodhumored/lr2_types_memory/syntax_analyzer/rule"
)

// Узел дерева вывода
type Node struct {
	Symbol   rule.Symbol
	Value    string
	Children []*Node
}

// Вспомогательная функция для создания пустого узла
func CreateNode(s rule.Symbol) Node {
	return Node{Symbol: s, Children: []*Node{}, Value: ""}
}

// Метод, добавляющий дочерний узел
func (n *Node) AddChild(child *Node) {
	n.Children = append(n.Children, child)
}

// Метод свёртки узла дерева
func (node *Node) Reduce(rule rule.Rule) bool {
	node.Print("", true)
	// Если не можем применить правило к текущему узлу - уходим
	symbolsCount, ok := node.CanApplyRule(rule)
	if !ok {
		return false
	}

	lenDiff := len(node.Children) - symbolsCount
	fmt.Printf("lenDiff: %v, symbolsCount: %v, childrenCount: %v\n", lenDiff, symbolsCount, len(node.Children))
	// копируем слайс с нужными нам узлами, которые собираемся заменять
	nodes := make([]*Node, symbolsCount)
	copy(nodes, node.Children[lenDiff:])

	// перезаписываем дочерние узлы узла
	node.Children = append(node.Children[:lenDiff], &Node{Symbol: rule.Left, Children: nodes, Value: ""})
	node.Print("", true)
	return true
}

// Функция проверки возможности применения правила к дочерним узлам узла
func (node Node) CanApplyRule(ruleToCheck rule.Rule) (int, bool) {
	childrenSymbols := []rule.Symbol{}
	for _, child := range node.Children {
		childrenSymbols = append(childrenSymbols, child.Symbol)
	}
	return rule.IsApplyable(ruleToCheck.Right, childrenSymbols)
}

func (node Node) String() string {
	return node.Symbol.GetName()
}

// Метод для рекурсивного вывода узлов дерева в консоль
func (node *Node) Print(prefix string, isTail bool) {
	// Выводим символ узла с отступом
	var branch, prefixSuffix string
	if isTail {
		prefixSuffix = "    "
		branch = "└── "
	} else {
		branch = "├── "
		prefixSuffix = "│   "
	}
	fmt.Println(prefix + branch + node.Symbol.GetName() + " (" + node.Value + ")")

	// Рекурсивно выводим дочерние узлы
	for i := 0; i < len(node.Children)-1; i++ {
		node.Children[i].Print(prefix+prefixSuffix, false)
	}
	if len(node.Children) > 0 {
		node.Children[len(node.Children)-1].Print(prefix+prefixSuffix, true)
	}
}
