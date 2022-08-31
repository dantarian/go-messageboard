package data

import (
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/repositories"
)

type inMemoryPingRepository struct {
	pings []entities.Ping
}

func NewInMemoryPingRepository() repositories.IPingRepository {
	return &inMemoryPingRepository{}
}

func (r *inMemoryPingRepository) Add(ping entities.Ping) {
	r.pings = append(r.pings, ping)
}

func (r *inMemoryPingRepository) Count() int {
	return len(r.pings)
}
