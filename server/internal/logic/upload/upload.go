package upload

import (
	"server/internal/service"
)

type sUpload struct{}

func init() {
	service.RegisterUpload(&sUpload{})
}
