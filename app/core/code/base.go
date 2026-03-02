package code

import "fmt"

type Code struct {
	Code    int    `json:"code"`    // 响应码
	Message string `json:"message"` // 响应消息
}

func (c *Code) Error() string {
	return fmt.Sprintf("code: %d, message: %s", c.Code, c.Message)
}

func New(code int, message string) error {
	return &Code{Code: code, Message: message}
}

var (
	ErrUnknown            = New(99999, "unknown error")
	ErrNotFound           = New(10000, "not found")
	ErrInvalidParam       = New(10001, "invalid param")
	ErrInvalidToken       = New(10002, "invalid token")
	ErrUnauthorized       = New(10003, "unauthorized")
	ErrForbidden          = New(10004, "forbidden")
	ErrInternal           = New(10005, "internal error")
	ErrNotImplemented     = New(10006, "not implemented")
	ErrBadGateway         = New(10007, "bad gateway")
	ErrServiceUnavailable = New(10008, "service unavailable")
	ErrGatewayTimeout     = New(10009, "gateway timeout")
	ErrBadRequest         = New(10010, "bad request")
	ErrConflict           = New(10011, "conflict")
	ErrJsonDecode         = New(10012, "json decode error")
	ErrJsonEncode         = New(10013, "json encode error")
	ErrInvalidRequest     = New(10014, "invalid request")
)
