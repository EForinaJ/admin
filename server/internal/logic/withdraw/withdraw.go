package withdraw

import (
	"server/internal/service"
)

type sWithdraw struct{}

func init() {
	service.RegisterWithdraw(&sWithdraw{})
}
