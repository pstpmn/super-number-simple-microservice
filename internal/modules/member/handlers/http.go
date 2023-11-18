package memberHandler

import (
	"super-number-simple-microservice/internal/modules/member"
	"super-number-simple-microservice/pkg"

	echo "github.com/labstack/echo/v4"
)

type (
	IHttpHandler interface {
		Registration(c echo.Context) error
	}
	httpHandler struct {
		Response   pkg.IResponse
		MemUseCase member.IUseCase
	}
)

// Registration implements IHttpHandler.
func (*httpHandler) Registration(c echo.Context) error {
	panic("unimplemented")
}

func NewHttpHandler(memUseCase member.IUseCase, response pkg.IResponse) IHttpHandler {
	return &httpHandler{}
}
