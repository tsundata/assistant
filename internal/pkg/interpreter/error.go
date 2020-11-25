package interpreter

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
	DuplicateId     ErrorCode = "Duplicate id found"
)

type Error struct {
	ErrorCode ErrorCode
	Token     *Token
	Message   string
	Type      ErrorType
}

func (e Error) Error() string {
	return fmt.Sprintf("%s: %s", e.ErrorCode, e.Message)
}
