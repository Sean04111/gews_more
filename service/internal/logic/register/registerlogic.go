package register

import (
	"context"
	"fmt"
	"strconv"

	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"gews_more/service/model"

	"github.com/go-redis/redis"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.Registerreque) (resp *types.Registerrespo, err error) {
	// todo: add your logic here and delete this line
	var NewUser model.User
	NewUser.Email = req.Emial
	NewUser.Name = req.Name
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), 10)
	if err != nil {
		return
	}
	NewUser.Password = string(hashedpassword)

	olderid, e := l.FromRedis()
	if e != nil {
		fmt.Println("redis 读取失败")
		return &types.Registerrespo{
			Error_code: 1,
		}, nil
	}
	NewUser.Uid = int64(olderid) + 1
	if l.ToRedis(int(NewUser.Uid)) != nil {
		fmt.Println("redis 写入失败")
		return &types.Registerrespo{
			Error_code: 1,
		}, nil
	}
	realcode, _ := CodeFromRedis(req.Name)
	if req.Code == realcode {
		_, error := l.svcCtx.UserModel.Insert(l.ctx, &NewUser)
		if error != nil {
			return &types.Registerrespo{
				Error_code: 1,
			}, nil
		} else {
			return &types.Registerrespo{
				Error_code: 0,
				Data: types.Registerdata{
					Uid:   int(NewUser.Uid),
					Email: NewUser.Email,
					Name:  NewUser.Name,
				},
			}, nil
		}
	}else{
		return &types.Registerrespo{
			Error_code: 1,
		},nil
	}
}
func (l *RegisterLogic) ToRedis(id int) error {
	client := redis.NewClient(&redis.Options{
		Addr: "121.36.131.50:6379",
		DB:   0,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return err
	}
	if pong != "PONG" {
		fmt.Println("客户端连接redis服务端失败")
		return err
	}
	return client.Set("user", id, 0).Err()
}
func (l *RegisterLogic) FromRedis() (int, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "121.36.131.50:6379",
		DB:   0,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	if pong != "PONG" {
		fmt.Println("客户端连接redis服务端失败")
		return 0, err
	}
	num, e := client.Get("user").Result()
	intnum, _ := strconv.Atoi(num)
	return intnum, e
}
func CodeFromRedis(name string) (string, error) {
	client := redis.NewClient(&redis.Options{
		Addr: "121.36.131.50:6379",
		DB:   0,
	})
	pong, err := client.Ping().Result()

	if err != nil {
		fmt.Println(err)
		return "", nil
	}
	if pong != "PONG" {
		fmt.Println("客户端连接redis服务端失败")
		return "", nil
	}
	code, er := client.Get(name).Result()
	if er != nil {
		fmt.Println("redis 读取失败!!")
		return "", er
	} else {
		return code, nil
	}
}
