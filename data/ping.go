package data

import "pencethren/go-messageboard/entities"

type inMemoryPingRepository struct {
	pings []entities.Ping
}

func NewInMemoryPingRepository() entities.IPingRepository {
	return &inMemoryPingRepository{}
}

func (r *inMemoryPingRepository) Add(ping entities.Ping) {
	r.pings = append(r.pings, ping)
}

func (r *inMemoryPingRepository) Count() int {
	return len(r.pings)
}
