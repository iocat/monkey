package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 123

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	// EQ and NOT_EQ: These two operators are hard to parsed given that
	// we have overlapping cases which are the ASSIGN and
	// BANG operators
	EQ     = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// TRUE and FALSE are both literals and keywords
	// 	Literal: language's lexical strings that represent
	// fixed value or are used to create fixed value
	//
	// 	Keyword: Language's lexical strings that are used
	// as part of the code construct. It has fixed meaning
	// no identifiers can be created with the same string to name
	// (unless compiler is rather complicated
	// and is smart about the compilation context)
	TRUE  = "TRUE"
	FALSE = "FALSE"

	IF     = "IF"
	ELSE   = "ELSE"
	RETURN = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent returns if the type if identifier is a language keyword
// otherwise returns IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
