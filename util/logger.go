package util

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ApplicationLogger struct {
	*zerolog.Logger
}

func NewLogger() *ApplicationLogger {
	baseLogger := log.Logger
	applicationLogger := &ApplicationLogger{&baseLogger}
	return applicationLogger
}

// Message Board Operation Messages

func (al *ApplicationLogger) FailedToInstantiateBoard(context string, err error) {
	al.Warn().Str("context", context).Err(err).Msg("failed to instantiate board")
}

func (al *ApplicationLogger) FailedToCheckForBoardNameReuse(context string, err error) {
	al.Error().Str("context", context).Err(err).Msg("failed to check for board name reuse")
}

func (al *ApplicationLogger) FailedToPersistBoard(context string, err error) {
	al.Error().Str("context", context).Err(err).Msg("failed to persist board")
}

func (al *ApplicationLogger) DuplicateBoardName(context string, name string) {
	al.Info().Str("context", context).Str("name", name).Msg("duplicate board name")
}
