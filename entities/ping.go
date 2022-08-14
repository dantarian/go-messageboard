package entities

import (
	"time"

	"github.com/google/uuid"
)

type Ping struct {
	id        string
	timestamp int64
}

func NewPing() Ping {
	return Ping{id: uuid.NewString(), timestamp: time.Now().UnixMicro()}
}

type IPingRepository interface {
	Add(Ping)
	Count() int
}
