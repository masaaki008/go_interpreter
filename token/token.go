package token

// TokensType 文字トークンタイプ
// golintでパッケージ名含む名前はwarningになる
type TokensType string

const (

	// ILLEGAL イリーガル
	ILLEGAL = "ILELEGAL"
	// EOF end of file
	EOF = "EOF"

	// IDENT 識別子
	IDENT = "IDENT"
	// INT int
	INT = "INT"

	// ASSIGN =
	ASSIGN = "="
	// PLUS +
	PLUS = "+"
	// MINUS -
	MINUS = "-"
	// BANG 驚嘆府
	BANG = "!"
	// ASTERISK アスタリスク
	ASTERISK = "*"
	// SLASH スラッシュ
	SLASH = "/"

	// LT less than
	LT = "<"
	// GT greater than
	GT = ">"

	// COMMA コンマ
	COMMA = ","
	// SEMICOLON セミコロン
	SEMICOLON = ";"
	// LPAREN 左パレンティス
	LPAREN = "("
	// RPAREN 右パレンティス
	RPAREN = ")"
	// LBRACE 左ブラケット
	LBRACE = "{"
	// RBRACE 右ブラケット
	RBRACE = "}"

	// FUNCTION 関数
	FUNCTION = "FUNCTION"
	// LET 変数宣言
	LET = "LET"
	// TRUE 真
	TRUE = "TRUE"
	// FALSE 偽
	FALSE = "FALSE"
	// IF if
	IF = "IF"
	// ELSE else
	ELSE = "ELSE"
	// RETURN リターン
	RETURN = "RETURN"
	// EQ equal
	EQ = "=="
	// NOTEQ not equal
	NOTEQ = "!="
)

// Token トークン構造体
type Token struct {
	Type    TokensType
	Literal string
}

var keywords = map[string]TokensType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 特定のキーワードが入力された場合その識別子を返す
func LookupIdent(ident string) TokensType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
