package dashboard

import "server/internal/service"

type sDashboard struct{}

func init() {
	setData()
	service.RegisterDashboard(&sDashboard{})
}
