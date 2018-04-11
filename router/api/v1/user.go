// TODO: 验证码
package v1

import (
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

type SingupCommand struct {
	Mail string `json:"mail"`
	Password string `json:"password"`
}

func AddUser(c *gin.Context) {
	model.AddUser("acd@mail.com", "123321")

	fmt.Println("add success")

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : make(map[string]string),
    })
}

func Signup(c *gin.Context) {
	var errors []string
	var signupCommand SingupCommand
	c.BindJSON(&signupCommand)

	mail := signupCommand.Mail
	pwd := signupCommand.Password

	valid := validation.Validation{}
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
