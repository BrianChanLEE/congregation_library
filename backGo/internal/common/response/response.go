package response

import (
	"github.com/gin-gonic/gin"
)

// ErrorResponse 표준 에러 응답 포맷
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Detail  string `json:"detail,omitempty"`
}

// SendError 표준 에러 응답 전송
func SendError(c *gin.Context, status int, code, message, detail string) {
	c.JSON(status, ErrorResponse{
		Code:    code,
		Message: message,
		Detail:  detail,
	})
}
