package settlement

import "server/internal/service"

type sSettlement struct{}

func init() {
	service.RegisterSettlement(&sSettlement{})
}
