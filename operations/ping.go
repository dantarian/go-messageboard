package operations

import (
	"pencethren/go-messageboard/entities"
	"pencethren/go-messageboard/repositories"
)

func RespondToPing(pings repositories.IPingRepository) string {
	pings.Add(entities.NewPing())
	return "pong"
}

func CountPingsReceived(pings repositories.IPingRepository) int {
	return pings.Count()
}
