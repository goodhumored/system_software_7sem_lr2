package main

import (
	"goodhumored/lr2_types_memory/syntax_analyzer/precedence"
	"goodhumored/lr2_types_memory/token"
)

var (
	eq = precedence.Eq
	lt = precedence.Lt
	gt = precedence.Gt
)

// Матрица предшествования
var precedenceMatrix = precedence.Matrix{
	token.TypeType:          precedence.Row{token.IdentifierType: lt, token.VarType: gt},
	token.IdentifierType:    precedence.Row{token.AssignmentType: eq, token.TypeSeparatorType: lt, token.DelimiterType: lt},
	token.AssignmentType:    precedence.Row{token.IdentifierType: lt, token.RecordStartType: eq, token.DelimiterType: eq},
	token.VarType:           precedence.Row{token.IdentifierType: lt},
	token.TypeSeparatorType: precedence.Row{token.IdentifierType: lt, token.RecordStartType: eq, token.DelimiterType: eq},
	token.RecordStartType:   precedence.Row{token.IdentifierType: lt, token.RecordEndType: eq, token.DelimiterType: lt},
	token.RecordEndType:     precedence.Row{token.DelimiterType: gt},
	token.DelimiterType:     precedence.Row{token.IdentifierType: gt, token.RecordEndType: gt, token.VarType: gt},
}
