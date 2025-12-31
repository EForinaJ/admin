package user

import "server/internal/service"

type sUser struct{}

func init() {
	service.RegisterUser(&sUser{})
}
