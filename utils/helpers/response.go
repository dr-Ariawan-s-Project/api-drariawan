package helpers

import "github.com/labstack/echo/v4"

type Response struct {
	Code    int    `json:"code,omitempty"`
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

func FailedResponse(message string) map[string]any {
	return map[string]any{
		"status":  "failed",
		"message": message,
	}
}

func SuccessResponse(message string) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
	}
}

func SuccessWithDataResponse(message string, data any) map[string]any {
	return map[string]any{
		"status":  "success",
		"message": message,
		"data":    data,
	}
}

/*
code: http response code.
message: message response.
data: if data response is available, put here. otherwise put nil.
*/
func WebResponseError(c echo.Context, err error) error {
	errCode, errMessage := CheckHandlerError(err)
	response := map[string]any{
		"meta": Response{
			Code:    errCode,
			Status:  "failed",
			Message: errMessage.Error(),
		},
	}
	return c.JSON(errCode, response)
}

func WebResponseSuccess(c echo.Context, msg string, data any) error {
	responseCode := CheckHandlerSuccess(msg)
	response := map[string]any{
		"meta": Response{
			Code:    responseCode,
			Status:  "success",
			Message: msg,
		},
		"data": data,
	}
	return c.JSON(responseCode, response)
}
