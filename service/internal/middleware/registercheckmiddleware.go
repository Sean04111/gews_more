package middleware

import (
	"net/http"
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
		next(w, r)
	}
}
