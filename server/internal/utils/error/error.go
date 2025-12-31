package utils_error

import (
	"server/internal/utils/response"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func Err(code response.ResponseCode, message string) error {
	return gerror.NewCode(gcode.New(int(code), "", nil), message)
}
