package cvalidator

import (
	"log"
	"sync"

	zhongwen "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/go-playground/validator/v10/translations/zh"
)

type CV struct {
	Trans ut.Translator
}

type CustomTag struct {
	Tag         string
	Func        validator.FuncCtx
	Translation string
	Override    bool
}

var once sync.Once

func GetCV() *CV {
	//once.Do(func() {
	zhTrans := zhongwen.New()
	uni := ut.New(zhTrans, zhTrans)
	trans, _ := uni.GetTranslator("zh")
	cv := CV{}
	cv.Trans = trans
	//})
	return &cv
}

func (cv *CV) DecorateValidator(v *validator.Validate, tags []CustomTag) error {
	err := zh.RegisterDefaultTranslations(v, cv.Trans)
	if nil != err {
		return err
	}
	for _, tag := range tags {
		e := v.RegisterValidationCtx(tag.Tag, tag.Func)
		if nil != e {
			return e
		}
		e = v.RegisterTranslation(tag.Tag, cv.Trans, registrationFunc(tag.Tag, tag.Translation, tag.Override), translateFunc)
		if nil != e {
			return e
		}
	}

	return nil
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) (err error) {
		if err = ut.Add(tag, translation, override); err != nil {
			return
		}
		return
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	// ut.T 函数的 param 为翻译模板中的 {0}、{1} ...
	t, err := ut.T(fe.Tag(), fe.Field(), fe.Param())
	if err != nil {
		log.Printf("警告: 翻译字段错误: %#v", fe)
		return fe.(error).Error()
	}

	return t
}
