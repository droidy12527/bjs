package constants

const (
	KEYWORD   = "KEYWORD"
	OPERATOR  = "OPERATOR"
	SPACE     = "SPACE"
	IDENTFIER = "IDENTIFIER"
	SEPERATOR = "SEPERATOR"
	PROMPT    = ">>"
)

const (
	INTEGER_OBJECT      = "INTEGER"
	BOOLEAN_OBJECT      = "BOOLEAN"
	NULL_OBJECT         = "NULL"
	RETURN_VALUE_OBJECT = "RETURN_VALUE"
	ERROR_OBJECT        = "ERROR"
	FUNCTION_OBJ        = "FUNCTION"
)

const (
	_ int = iota
	LOWEST
	EQUALS
	LESSGREATER
	SUM
	PRODUCT
	PREFIX
	CALL
)

const (
	LOGO = `
 /\_/\  
( o.o )  Pookie
 > ^ <
`
)
