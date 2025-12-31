package manage

import (
	"server/internal/service"
)

type sManage struct{}

func init() {
	service.RegisterManage(&sManage{})
}
