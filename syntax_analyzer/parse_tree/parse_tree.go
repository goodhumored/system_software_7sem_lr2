package parse_tree

import (
	"fmt"

	"goodhumored/lr2_types_memory/syntax_analyzer/rule"
)

// Дерево вывода
type ParseTree struct {
	Root *Node
}

// Метод добавления узлов в дерево
func (tree *ParseTree) AddNode(node *Node) {
	tree.Root.AddChild(node)
}

// Метод для свёртки дерева по правилу
func (tree *ParseTree) Reduce(rule rule.Rule) {
	fmt.Printf("Применяем правило %s к дереву\n", rule)
	if tree.Root.Reduce(rule) {
		fmt.Printf("Успешно применено\n")
	} else {
		fmt.Printf("Правило %s применить не удалось\n", rule)
	}
}

func (tree ParseTree) BFS(symbol rule.Symbol) []*Node {
	nodes := []*Node{}
	queue := []*Node{tree.Root}
	for len(queue) > 0 {
		curNode := queue[0]
		if curNode.Symbol.GetName() == symbol.GetName() {
			nodes = append(nodes, curNode)
		}
		queue = append(queue[1:], curNode.Children...)
	}
	return nodes
}

// Метод для вывода дерева
func (tree ParseTree) Print() {
	tree.Root.Print("", true)
}
