package aftersales

import (
	"server/internal/service"
)

type sAftersales struct{}

func init() {
	service.RegisterAftersales(&sAftersales{})
}
