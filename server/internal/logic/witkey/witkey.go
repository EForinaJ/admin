package witkey

import "server/internal/service"

type sWitkey struct{}

func init() {
	service.RegisterWitkey(&sWitkey{})
}
