package token_analyzer

import (
	"regexp"

	"goodhumored/lr2_types_memory/token"
)

// Вспомогатеьлная структура для установки соответствия шаблонов лексем
// с их фабричными функциями
type TokenPattern struct {
	Pattern *regexp.Regexp
	Type    func(string, token.Position) token.Token
}

// Массив соответствий шаблонов лексем
var tokenPatterns = []TokenPattern{
	{regex("\\btype\\b"), token.Type},
	{regex("\\bvar\\b"), token.Var},
	{regex("\\brecord\\b"), token.Record},
	{regex(":"), token.TypeSeparator},
	{regex("\\bend\\b"), token.End},
	{regex("[a-zA-Z][a-zA-Z0-9]+"), token.Identifier},
	{regex("="), token.Assignment},
	{regex("#.*"), token.Comment},
	{regex("[(]"), token.Record},
	{regex("[)]"), token.End},
	{regex(";"), token.Delimiter},
}

// вспомогательная функция создающая объект регулярного выражения
// добавляющая в начале шаблона признак начала строки
func regex(pattern string) *regexp.Regexp {
	return regexp.MustCompile("^" + pattern)
}
