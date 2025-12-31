package site

import "server/internal/service"

type sSite struct{}

func init() {
	service.RegisterSite(&sSite{})
}
