package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL 未知のトークンや文字列
	ILLEGAL = "ILLEGAL"
	// EOF end of file ファイルの最後
	EOF = "EOF"

	ASSIGN = "="
	PLUS   = "+"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
