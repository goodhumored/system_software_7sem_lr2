package nonterminal

// Структура представляющая нетерминалы
type NonTerminal struct {
	Name string
}

// Метод для соответствия нетерменалов интерфейсу символ
func (nt NonTerminal) GetName() string {
	return nt.Name
}

func (nt NonTerminal) Value() string {
	return ""
}

var (
	Record          = NonTerminal{"Record"}
	TypeBlock       = NonTerminal{"TypeBlock"}
	VarBlock        = NonTerminal{"VarBlock"}
	TypeDeclaration = NonTerminal{"TypeDeclaration"}
	VarDeclaration  = NonTerminal{"VarDeclaration"}
	Root            = NonTerminal{"/"} // Корневой нетерминал
)
