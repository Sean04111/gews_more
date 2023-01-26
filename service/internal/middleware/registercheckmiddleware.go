package middleware

import (
	"encoding/json"
	"fmt"
	"gews_more/service/internal/types"
	"net/http"

	"github.com/go-redis/redis"
)

type RegistercheckMiddleware struct {
}

func NewRegistercheckMiddleware() *RegistercheckMiddleware {
	return &RegistercheckMiddleware{}
}

type Reque struct {
	Email    string
	Name     string
	Code     string
	Password string
}

func (m *RegistercheckMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		//code的校验在middleware中实现
		// Passthrough to next handler if need
		var NewR Reque
		err := json.NewDecoder(r.Body).Decode(&NewR)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("Parse Error!")
			return
		}
		realcode, e := FromRedis(NewR.Name)
		if e != nil {
			fmt.Println("Redis 读取失败！！！")
			return
		} else {
			if realcode == NewR.Code {
				next(w, r)
			} else {
				fail := types.Registerrespo{
					Error_code: 1,
				}
				message, _ := json.Marshal(fail)
				w.Write(message)
				return
			}
		}
	}
}
func FromRedis(name string) (string, error) {
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
