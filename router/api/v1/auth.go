package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
	// "github.com/GallenHu/bookmarkgo/model"
)

type SinginCommand struct {
	Mail string `json:"mail"`
	Password string `json:"password"`
	Tick string `json:"tick"`
	Captcha string `json:"captcha"`
}

func Signin(c *gin.Context) {
	var errors []string
	var signinCommand SinginCommand

	mail := signinCommand.Mail
	pwd := signinCommand.Password
	tick := signinCommand.Tick
	captcha := signinCommand.Captcha

	valid := validation.Validation{}
	valid.Required(tick, "tick").Message("tick不能为空")
	valid.Required(captcha, "captcha").Message("验证码不能为空")
	valid.Required(mail, "mail").Message("邮箱不能为空")
	valid.Required(pwd, "password").Message("密码不能为空")
	valid.Email(mail, "mailValidity").Message("邮箱输入有误")
	valid.MinSize(pwd, 6, "passwordLength").Message("密码至少为6位")

	if valid.HasErrors() {
        for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			errors = append(errors, err.Message)
		}

		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	if !utils.VerfiyCaptcha(tick, captcha) {
		errors = append(errors, "验证码错误")
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})
		return
	}
}

func Signout(c *gin.Context) {

}