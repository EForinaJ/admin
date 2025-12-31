package title

import "server/internal/service"

type sTitle struct {
}

func init() {
	service.RegisterTitle(&sTitle{})
}
