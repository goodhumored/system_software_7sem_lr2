package syntax_analyzer

import (
	"errors"
	"fmt"

	"goodhumored/lr2_types_memory/syntax_analyzer/nonterminal"
	"goodhumored/lr2_types_memory/syntax_analyzer/parse_tree"
	"goodhumored/lr2_types_memory/syntax_analyzer/precedence"
	"goodhumored/lr2_types_memory/syntax_analyzer/rule"
	"goodhumored/lr2_types_memory/token"
	"goodhumored/lr2_types_memory/token_table"
)

// Функция для анализа синтаксиса, принимает таблицу токенов, список правил и матрицу предшествования
func AnalyzeSyntax(ruleTable rule.RuleTable, tokenTable token_table.TokenTable, matrix precedence.Matrix) (parse_tree.ParseTree, error) {
	// Создаём дерево
	rootNode := parse_tree.CreateNode(nonterminal.Root)
	tree := parse_tree.ParseTree{Root: &rootNode}
	// Получаем лексемы из таблицы
	tokens := tokenTable.GetTokens()
	tokenIndex := 1
	// Создаём стек
	stack := symbolStack{tokens[0]}

	for {
		// Берём ближайший к вершине терминал
		stackTerminal := stack.PeekTopTerminal()
		// если токены кончились - смотрим если можем свернуть корень и вернуть результат
		if len(tokens) <= tokenIndex {
			ok, err := isInputAccepted(stack, ruleTable)
			if err != nil {
				return tree, err
			}
			if ok {
				return tree, nil
			}
			return tree, errors.New("Токены закончились, до конца свернуть не удалось")
		}
		// Берём текущий символ входной строки
		inputToken := tokens[tokenIndex]
		// Если комментарий - пропускаем
		if inputToken.Type == token.CommentType {
			tokenIndex += 1
			continue
		}

		fmt.Printf("Лексема: '%s' \n", tokens[tokenIndex].Value())
		fmt.Printf("Стек: %s \n", stack)

		// Получаем предшествование из матрицы
		prec := matrix.GetPrecedence(stackTerminal.Type, inputToken.Type)

		// Если предшествование или =, тогда сдвигаем
		if prec == precedence.Lt || prec == precedence.Eq {
			print("Сдвигаем\n")
			node := &parse_tree.Node{Value: inputToken.Value(), Symbol: inputToken, Children: []*parse_tree.Node{}}
			tree.AddNode(node) // Добавляем узел в дерево
			stack = stack.Push(inputToken)
			tokenIndex += 1
		} else if prec == precedence.Gt { // Иначе сворачиваем
			print("Сворачиваем\n")
			// сворачиваем стек
			newStack, rule, err := reduce(stack, ruleTable)
			if err != nil {
				return tree, err
			}
			stack = newStack
			// сворачиваем дерево
			tree.Reduce(*rule)
		} else {
			// Если предшествование не определено - выдаем ошибку
			return tree, fmt.Errorf("Ошибка в синтексе, неожиданное сочетание символов %s и %s (%d)", stackTerminal.GetName(), inputToken.GetName(), inputToken.Position.End)
		}
		println("==============")
	}
}

// Проверка на завершённость
func isInputAccepted(stack symbolStack, ruleTable rule.RuleTable) (bool, error) {
	// Пытаемся свернуть
	newStack, _, err := reduce(stack, ruleTable)
	if err != nil {
		return false, err
	}
	// Если в стеке верхушка теперь рут
	if newStack.Peek().GetName() == nonterminal.Root.Name {
		return true, nil
	}
	return false, nil
}

// Функция свёртки стека
func reduce(stack symbolStack, ruleTable rule.RuleTable) (symbolStack, *rule.Rule, error) {
	for {
		// Если есть применимое к стеку правило
		if rule, count := ruleTable.GetRuleByRightSide(stack); rule != nil {
			fmt.Printf("Нашлось правило: %v, пушим %s в стек\n", rule, rule.Left)
			// обновляем стек
			stack = append(stack[:len(stack)-count], rule.Left)
			return stack, rule, nil
		} else {
			// Если нет выдаем ошибку
			return stack, nil, fmt.Errorf("Не найдено правил для свёртки")
		}
	}
}
