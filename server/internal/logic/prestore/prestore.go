package prestore

import (
	"server/internal/service"
)

type sPrestore struct{}

func init() {
	service.RegisterPrestore(&sPrestore{})
}
