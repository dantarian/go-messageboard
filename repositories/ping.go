package repositories

import "pencethren/go-messageboard/entities"

type IPingRepository interface {
	Add(entities.Ping)
	Count() int
}
