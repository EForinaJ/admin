package capital

import "server/internal/service"

type sCapital struct{}

func init() {
	service.RegisterCapital(&sCapital{})
}
