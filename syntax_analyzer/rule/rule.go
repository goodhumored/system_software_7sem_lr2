package rule

import (
	"fmt"

	"goodhumored/lr2_types_memory/syntax_analyzer/nonterminal"
)

// Интерфейс для представления символа
type Symbol interface {
	GetName() string
}

// Правило
type Rule struct {
	Left  nonterminal.NonTerminal
	Right []RuleItem
}

// Метод получения строки из правила
func (r Rule) String() string {
	return fmt.Sprintf("%s -> %s", r.Left.GetName(), r.Right)
}

func NewRule(left nonterminal.NonTerminal, right []RuleItem) Rule {
	return Rule{
		Left:  left,
		Right: right,
	}
}
