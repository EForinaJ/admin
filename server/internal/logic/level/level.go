package level

import "server/internal/service"

type sLevel struct{}

func init() {
	service.RegisterLevel(&sLevel{})
}
