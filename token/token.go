package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL 	= "ILLEGAL"
	EOF     	= "EOF"
	IDENT   	= "IDENT"
	INT     	= "INT"
	ASSIGN  	= "="
	MINUS		= "-"
	ASTARISK 	= "*"
	SLASH		= "/"
	PLUS    	= "+"
	LT			= "<"
	GT			= ">"
	EQ			= "=="
	NE 		    = "!="
	LET     	= "LET"
	GOTO		= "GOTO"
	GOSUB		= "GOSUB"
	IF      	= "IF" 
	THEN    	= "THEN" 
)

var keywords = map[string]TokenType {
	"LET":  LET, 
	"GOTO": GOTO,
	"GOSUB": GOSUB,
	"IF": IF,
    "THEN": THEN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok 
	}
	return IDENT 
}