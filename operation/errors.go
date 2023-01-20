package operation

import (
	"fmt"
)

type BusinessRuleError struct {
	reason string
	Msg    string
}

func NewBusinessRuleError(message string, reason string) *BusinessRuleError {
	return &BusinessRuleError{Msg: message, reason: reason}
}

func (e *BusinessRuleError) Error() string {
	return fmt.Sprintf("%v: %v", e.Msg, e.reason)
}
