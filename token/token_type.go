package token

type TokenType struct {
	Name string
}

func (tokenType TokenType) GetName() string {
	return tokenType.Name
}

var (
	TypeType          = TokenType{"type_keyword"}   // Ключевое слово type
	IdentifierType    = TokenType{"identifier"}     // имя переменной
	AssignmentType    = TokenType{"assignment"}     // =
	VarType           = TokenType{"var_keyword"}    // ключевое слово var
	TypeSeparatorType = TokenType{"type_separator"} // Разделитель
	RecordStartType   = TokenType{"record_start"}   // ключевое слово record
	RecordEndType     = TokenType{"record_end"}     // ключевое слово end
	DelimiterType     = TokenType{"delimiter"}      // Разделитель
	ErrorType         = TokenType{"error"}          // Ошибка
	CommentType       = TokenType{"comment"}        // Комментарий
	StartType         = TokenType{"start"}          // Начало
	EOFType           = TokenType{"EOF"}            // Конец
)
