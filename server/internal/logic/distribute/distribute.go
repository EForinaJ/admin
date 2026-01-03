package distribute

import (
	"server/internal/service"
)

type sDistribute struct{}

func init() {
	service.RegisterDistribute(&sDistribute{})
}
