package rule

import "fmt"

// Таблица правил
type RuleTable struct {
	Rules []Rule
}

// Метод поиска правила по правой части
func (ruleTable RuleTable) GetRuleByRightSide(tokenTypes []Symbol) (*Rule, int) {
	for _, rule := range ruleTable.Rules {
		fmt.Printf("trying rule for %s:\n", rule.Left.GetName())
		if count, ok := IsApplyable(rule.Right, tokenTypes); ok {
			return &rule, count
		}
	}
	return nil, 0
}

// Проверка на применимость правила к целевым символам
func IsApplyable(ruleItems []RuleItem, targetSymbols []Symbol) (int, bool) {
	ruleItemI := len(ruleItems) - 1
	symbolI := len(targetSymbols) - 1
	repeating := false
	for ruleItemI >= 0 && symbolI >= 0 {
		fmt.Printf("ruleSymbol: %s; targetSymbol: %s\n", ruleItems[ruleItemI], targetSymbols[symbolI])
		if ruleItems[ruleItemI].Fits(targetSymbols[symbolI]) {
			symbolI--
			if ruleItems[ruleItemI].AllowRepeat() {
				repeating = true
			} else {
				ruleItemI--
			}
		} else {
			if ruleItems[ruleItemI].Optional() || repeating {
				ruleItemI--
				repeating = false
			} else {
				return 0, false
			}
		}
	}
	if ruleItemI > 0 && symbolI == 0 {
		return 0, false
	}
	return len(targetSymbols) - 1 - symbolI, true
}

func ContainsRule(arr []Symbol, el Symbol) bool {
	elName := el.GetName()
	for _, e := range arr {
		if e.GetName() == elName {
			return true
		}
	}
	return false
}
