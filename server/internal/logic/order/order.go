package order

import (
	"server/internal/service"
)

type sOrder struct{}

func init() {
	service.RegisterOrder(&sOrder{})
}
