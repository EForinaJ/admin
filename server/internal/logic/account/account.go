package account

import "server/internal/service"

type sAccount struct{}

func init() {
	service.RegisterAccount(&sAccount{})
}
