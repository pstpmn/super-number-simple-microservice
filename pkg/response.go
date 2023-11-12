package pkg

import echo "github.com/labstack/echo/v4"

type (
	IResponse interface {
		ErrResponse(c echo.Context, statusCode int, message string) error
		SuccessResponse(c echo.Context, statusCode int, message string, result any) error
	}
	response struct{}

	resultResponse struct {
		Message    string `json:"message"`
		Result     any    `json:"result"`
		StatusCode int    `json:"statusCode"`
	}
)

// ErrResponse implements IResponse.
func (*response) ErrResponse(c echo.Context, statusCode int, message string) error {
	return c.JSON(statusCode, &resultResponse{
		Message:    message,
		StatusCode: statusCode,
	})
}

// SuccessResponse implements IResponse.
func (*response) SuccessResponse(c echo.Context, statusCode int, message string, result any) error {
	return c.JSON(statusCode, &resultResponse{
		Message:    message,
		StatusCode: statusCode,
		Result:     result,
	})
}

func NewResponse() IResponse {
	return &response{}
}
