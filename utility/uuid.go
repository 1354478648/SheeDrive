package utility

import (
	"github.com/google/uuid"
)

func GenStrUUID() (id string) {
	id = uuid.New().String()
	return
}

func GenIntUUID() (id int64) {
	id = int64(uuid.New().ID())
	return
}
