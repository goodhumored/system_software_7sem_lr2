package rule

type RuleItemType struct {
	Name string
}

type RuleItem struct {
	symbols     []Symbol
	optional    bool
	allowRepeat bool
}

func (ri RuleItem) Fits(s Symbol) bool {
	for _, acceptableSymbol := range ri.symbols {
		if acceptableSymbol.GetName() == s.GetName() {
			return true
		}
	}
	return false
}

func (ri RuleItem) AllowRepeat() bool {
	return ri.allowRepeat
}

func (ri RuleItem) Optional() bool {
	return ri.optional
}

func (ri RuleItem) String() string {
	ruleItemString := ri.symbols[0].GetName()
	for _, symbol := range ri.symbols[1:] {
		ruleItemString += "|" + symbol.GetName()
	}
	if ri.allowRepeat && ri.optional {
		ruleItemString += "*"
	}
	if ri.allowRepeat && !ri.optional {
		ruleItemString += "+"
	}
	if !ri.allowRepeat && ri.optional {
		ruleItemString += "?"
	}
	return ruleItemString
}

func NewRuleItem(optional bool, allowRepeat bool, symbols ...Symbol) RuleItem {
	return RuleItem{
		symbols:     symbols,
		optional:    optional,
		allowRepeat: allowRepeat,
	}
}
