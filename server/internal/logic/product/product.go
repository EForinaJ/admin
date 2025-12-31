package product

import "server/internal/service"

type sProduct struct{}

func init() {
	service.RegisterProduct(&sProduct{})
}
