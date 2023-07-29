// Package errcode 编写常用的依稀错误处理公共方法，标准化错误输出
package errcode

import (
	"fmt"
	"net/http"
)

type Error struct {
	code    int      `json:"code"`
	msg     string   `json:"msg"`
	details []string `json:"details"`
}

// 储存错误码
var codes = map[int]string{}

func NewError(code int, msg string) *Error {

	// 判断错误码是否存在
	if _, ok := codes[code]; ok {
		// 抛出错误
		panic(fmt.Sprintf("错误码 %d 已存在，请更换一个", code))
	}

	// 储存错误
	codes[code] = msg

	return &Error{
		code: code,
		msg:  msg,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("错误码: %d, 错误信息 :%s", e.Code(), e.Msg())
}

func (e *Error) Code() int {
	return e.code
}
func (e *Error) Msg() string {
	return e.msg
}

func (e *Error) MsgF(args []interface{}) string {
	return fmt.Sprintf(e.msg, args...)
}

func (e *Error) Details() []string {
	return e.details
}

func (e *Error) WithDetails(details ...string) *Error {
	newError := e
	newError.details = []string{}
	for _, detail := range details {
		newError.details = append(newError.details, detail)
	}
	return newError
}

// StatusCode 处理返回错误码，
// 针对特定的错误进行状态码的转换
// 因为不同的内部错误码在 HTTP 状态码中都代表着不同的意义，我们需要将其区分开来，便于客户端以及监控/报警等系统的识别和监听。
func (e *Error) StatusCode() int {
	switch e.code {
	case Success.code:
		return http.StatusOK
	case ServerError.code:
		return http.StatusInternalServerError
	case InvalidParams.code:
		return http.StatusBadRequest
	case UnauthorizedAuthNotExist.code:
		fallthrough
	case UnauthorizedTokenError.code:
		fallthrough
	case UnauthorizedTokenGenerate.code:
		fallthrough
	case UnauthorizedTokenTimeout.code:
		return http.StatusUnauthorized
	case TooManyRequests.code:
		return http.StatusTooManyRequests
	}
	return http.StatusInternalServerError
}
