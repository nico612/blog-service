package app

import (
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"strings"
)

// 接口校验

type ValidError struct {
	Key     string
	Message string
}

// 实现error接口
func (v *ValidError) Error() string {
	return v.Message
}

// ValidErrors 定义ValidErrors类型
type ValidErrors []*ValidError

// 实现error接口
func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string

	for _, err := range v {
		// 这里的err是指针类型
		errs = append(errs, err.Error())
	}
	return errs
}

// BindAndValid 对入参校验进行二次封装
func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	// 参数绑定和入参校验
	err := c.ShouldBind(v)

	// 如果发生错误，通过中间件设置的Translator 来对错误消息体进行具体的翻译行为
	if err != nil {
		// 获取中间件中设置的翻译器
		v := c.Value("trans")

		trans, _ := v.(ut.Translator)

		verrs, ok := err.(validator.ValidationErrors)
		if !ok {
			return false, errs
		}
		// 处理错误
		for key, value := range verrs.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return false, errs
	}

	return true, nil

}
