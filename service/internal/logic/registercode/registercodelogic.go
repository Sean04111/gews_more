package registercode

import (
	"context"
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"time"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"

	"github.com/go-redis/redis"
	"github.com/jordan-wright/email"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegistercodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegistercodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegistercodeLogic {
	return &RegistercodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegistercodeLogic) Registercode(req *types.Registercodereque) (resp *types.Registercoderespo, err error) {
	// todo: add your logic here and delete this line
	//计划使用redis来储存code
	//code发送在logic中实现
	realcode, err := l.SendCode(req.Email)
	l.ToRedis(req.Name, strconv.Itoa(realcode))
	if err != nil {
		return &types.Registercoderespo{
			Error_code: 1,
		}, nil
	} else {
		return &types.Registercoderespo{
			Error_code: 0,
		}, nil
	}
}
func (l *RegistercodeLogic) SendCode(receiver string) (int, error) {
	//随机生成验证码
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(10000)
	//Input the code
	text := "欢迎加入gews!!你的验证码是 :" + strconv.Itoa(code)
	em := email.NewEmail()
	em.From = "3408935702@qq.com"
	em.To = []string{receiver}
	em.Subject = "欢迎加入gews!"
	em.Text = []byte(text)
	em.HTML=[]byte(
		`<html></html>
		`)
	//注意!!!
	//QQ邮箱验证码需要定期更换
	//同时注意QQ盗号他用！！
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "3408935702@qq.com", "cesdwcniiuhtcjja", "smtp.qq.com"))
	if err != nil {
		fmt.Println("send error :", err)
	}
	//input code into the cache
	return code, err
}
func (l *RegistercodeLogic) ToRedis(name, code string) {
	client := redis.NewClient(&redis.Options{
		Addr: "121.36.131.50:6379",
		DB:   0,
	})
	pong,err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
	}

	if pong != "PONG" {
		fmt.Println("客户端连接redis服务端失败")
	}else {
		err:=client.Set(name,code,120).Err()
		if err!=nil{
			fmt.Println("插入redis值失败!!!")
		}
	}	
}
