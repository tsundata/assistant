package action

import "fmt"

type ErrorType string

const (
	LexerErrorType    ErrorType = "LexerError"
	ParserErrorType   ErrorType = "ParserError"
	SemanticErrorType ErrorType = "SemanticError"
)

type ErrorCode string

const (
	UnexpectedToken ErrorCode = "Unexpected token"
	IdNotFound      ErrorCode = "Identifier not found"
	RepeatOpcode    ErrorCode = "Repeat opcode"
)

type Error struct {
	ErrorCode ErrorCode
	Token     *Token
	Message   string
	Type      ErrorType
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}
