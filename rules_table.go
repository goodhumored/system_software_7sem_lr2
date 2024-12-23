package main

import (
	"goodhumored/lr2_types_memory/syntax_analyzer/nonterminal"
	"goodhumored/lr2_types_memory/syntax_analyzer/rule"
	"goodhumored/lr2_types_memory/token"
)

func one(symbols ...rule.Symbol) rule.RuleItem {
	return rule.NewRuleItem(false, false, symbols...)
}

func optional(symbols ...rule.Symbol) rule.RuleItem {
	return rule.NewRuleItem(true, false, symbols...)
}

func optionalMany(symbols ...rule.Symbol) rule.RuleItem {
	return rule.NewRuleItem(true, true, symbols...)
}

func many(symbols ...rule.Symbol) rule.RuleItem {
	return rule.NewRuleItem(false, true, symbols...)
}

var (
	rootRule = rule.NewRule(nonterminal.Root, []rule.RuleItem{
		one(token.StartType),
		one(nonterminal.TypeBlock),
		one(nonterminal.VarBlock),
		one(token.EOFType),
	})
	typeBlockRule = rule.NewRule(nonterminal.TypeBlock, []rule.RuleItem{
		one(token.TypeType),
		optionalMany(nonterminal.TypeDeclaration),
	})
	typeDeclarationRule = rule.NewRule(nonterminal.TypeDeclaration, []rule.RuleItem{
		one(token.IdentifierType),
		one(token.AssignmentType),
		one(token.IdentifierType, nonterminal.Record),
		one(token.DelimiterType),
	})
	recordRule = rule.NewRule(nonterminal.Record, []rule.RuleItem{
		one(token.RecordStartType),
		optionalMany(nonterminal.VarDeclaration),
		one(token.RecordEndType),
	})
	varBlockRule = rule.NewRule(nonterminal.VarBlock, []rule.RuleItem{
		one(token.VarType),
		optionalMany(nonterminal.VarDeclaration),
	})
	varDeclarationRule = rule.NewRule(nonterminal.VarDeclaration, []rule.RuleItem{
		one(token.IdentifierType),
		one(token.TypeSeparatorType),
		one(token.IdentifierType, nonterminal.Record),
		one(token.DelimiterType),
	})
)

var rulesTable = rule.RuleTable{Rules: []rule.Rule{
	varDeclarationRule,
	varBlockRule,
	recordRule,
	typeDeclarationRule,
	typeBlockRule,
	rootRule,
}}
