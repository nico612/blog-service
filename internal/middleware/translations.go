package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	//多语言包，是从 CLDR 项目（Unicode 通用语言环境数据存储库）生成的一组多语言环境，主要在 i18n 软件包中使用，该库是与 universal-translator 配套使用的。
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	"github.com/go-playground/validator/v10"

	// validator翻译器
	en_translations "github.com/go-playground/validator/v10/translations/en"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"

	// 通用翻译器，是一个使用CLDR数据+复数规则的go语言i18n转换器
	ut "github.com/go-playground/universal-translator"
)

// Translations 国际化中间件
func Translations() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 创建支持语言
		uni := ut.New(en.New(), zh.New(), zh_Hant_TW.New())

		// 通过 GetHeader 方法去获取约定的 header 参数 locale，用于判别当前请求的语言类别是 en 又或是 zh
		locale := c.GetHeader("locale")
		trans, _ := uni.GetTranslator(locale)

		v, ok := binding.Validator.Engine().(*validator.Validate)
		if ok {
			switch locale {
			case "zh":
				//将验证器和对应语言类型的 Translator 注册进来,实现验证器的多语言支持
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			case "en":
				_ = en_translations.RegisterDefaultTranslations(v, trans)
				break
			default:
				_ = zh_translations.RegisterDefaultTranslations(v, trans)
				break
			}
			//将Translator 存储到全局上下文中，便于后续翻译时的使用。
			c.Set("trans", trans)
		}
		c.Next()
	}
}
