package category

import "server/internal/service"

type sCategory struct{}

func init() {
	service.RegisterCategory(&sCategory{})
}
