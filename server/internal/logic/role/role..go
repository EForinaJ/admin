package role

import (
	"server/internal/service"
)

type sRole struct{}

func init() {
	service.RegisterRole(&sRole{})
}
