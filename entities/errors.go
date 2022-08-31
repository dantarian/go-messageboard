package entities

import (
	"fmt"
	"strings"
)

type ValidationError struct {
	reasons []string
	Msg     string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%v: %v", e.Msg, strings.Join(e.reasons, "; "))
}
