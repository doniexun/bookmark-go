package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

type SingupCommand struct {
	Mail string `json:"mail"`
	Password string `json:"password"`
	Tick string `json:"tick"`
	Captcha string `json:"captcha"`
}

// 获取验证码
func GetCaptcha(c *gin.Context) {
	capID, capImg := utils.GetCaptcha()

	data := map[string]string{"capImg": capImg, "capID": capID}

	c.JSON(200, gin.H{
		"code" : 200,
		"msg" : "success",
		"data" : data,
	})
}

func Signup(c *gin.Context) {
	var errors []string
	var signupCommand SingupCommand
	c.BindJSON(&signupCommand)

	mail := signupCommand.Mail
	pwd := signupCommand.Password
	tick := signupCommand.Tick
	captcha := signupCommand.Captcha

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

	if model.ExistUserByMail(mail) {
		errors = append(errors, "注册邮箱已存在")

		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	model.AddUser(mail, utils.Md5(mail + pwd))

	c.JSON(200, gin.H{
		"code" : 200,
		"msg" : "success",
		"data" : errors,
	})
}
