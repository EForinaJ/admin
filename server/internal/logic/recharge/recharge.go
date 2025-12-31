package recharge

import "server/internal/service"

type sRecharge struct{}

func init() {
	service.RegisterRecharge(&sRecharge{})
}
