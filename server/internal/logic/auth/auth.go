package auth

import "server/internal/service"

type sAtuh struct{}

func init() {
	service.RegisterAuth(&sAtuh{})
}
