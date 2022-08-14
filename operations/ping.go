package operations

import (
	"pencethren/go-messageboard/entities"
)

func RespondToPing(pings entities.IPingRepository) string {
	pings.Add(entities.NewPing())
	return "pong"
}

func CountPingsReceived(pings entities.IPingRepository) int {
	return pings.Count()
}
