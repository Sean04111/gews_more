package middleware

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/smtp"
	"strconv"
	"time"

	"github.com/beego/beego/v2/adapter/cache"
	"github.com/jordan-wright/email"
)

type RegistercheckMiddleware struct {
}

func NewRegistercheckMiddleware() *RegistercheckMiddleware {
	return &RegistercheckMiddleware{}
}

func (m *RegistercheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		//code的校验在middleware中实现
		// Passthrough to next handler if need
		
		next(w, r)
	}
}

//邮箱验证码绑定方法
func Send(ca cache.Cache, receiver string, fromqq string, rightcode string) error {
	//get the code
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(10000)
	//Input the code
	InputCode(ca, code)
	text := "your verification code is :" + strconv.Itoa(code)
	em := email.NewEmail()
	em.From = fromqq
	em.To = []string{receiver}
	em.Subject = "Hello"
	em.Text = []byte(text)
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", fromqq, rightcode, "smtp.qq.com"))
	if err != nil {
		fmt.Println("send error :", err)
	}
	//input code into the cache
	return err
}

func Check(ca cache.Cache, input int) bool {
	real := GetCode(ca)
	if input == real {
		return true
	} else {
		return false
	}
}
func InputCode(ca cache.Cache, code int) error {
	err := ca.Put("code", code, 60*time.Second)
	return err
}
func GetCode(ca cache.Cache) int {
	out := ca.Get("code")
	return out.(int)
}
