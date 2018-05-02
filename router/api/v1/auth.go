package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/redis"
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

	c.BindJSON(&signinCommand)

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

	if !model.ExistUserByMail(mail) {
		errors = append(errors, "账户不存在")

		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	user := model.CheckUserMd5Pwd(mail, utils.Md5(mail + pwd)) // return User model
	if user.ID == 0 {
		errors = append(errors, "用户名和密码不匹配")

		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	var err error
	token := redis.GetUserToken(user.ID)
	log.Println("token111")
	log.Println(token)
	if token == "" {
		token, err = utils.GenerateToken(user.ID)
		if err != nil {
			errors = append(errors, "token生成失败")

			c.JSON(200, gin.H{
				"code" : 500,
				"msg" : "failed",
				"data" : errors,
			})

			return
		}
	}

	redis.StoreUserPrivate(user.ID, user.ShowPrivate)
	err = redis.StoreUserToken(user.ID, token)
	if err != nil {
		log.Println(err)

		errors = append(errors, "token存储失败")

		c.JSON(200, gin.H{
			"code" : 500,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	c.JSON(200, gin.H{
		"code" : 200,
		"msg" : "success",
		"data" : token,
	})
}

func Signout(c *gin.Context) {
	var errors []string
	userid, exists := c.Get("userid")
	if !exists {
		errors = append(errors, "读取用户信息失败")
		c.JSON(200, gin.H{
			"code" : 500,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	redis.DelUserToken(userid.(int))
	c.JSON(200, gin.H{
		"code" : 200,
		"msg" : "success",
		"data" : userid,
	})
}
