package middleware

import (
	"net/http"
)

type GetcodeMiddleware struct {
}
type Request struct {
	Email string
	Name  string
}

func NewGetcodeMiddleware() *GetcodeMiddleware {
	return &GetcodeMiddleware{}
}

func (m *GetcodeMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation
		// Passthrough to next handler if need
		//
		//var NewR Request
		//err := json.NewDecoder(r.Body).Decode(&NewR)
		//if err != nil {
			//http.Error(w, err.Error(), http.StatusBadRequest)
		//	fmt.Println("Parse Error!")
		//	return
		//}
		//
		//
		//
		//
		//
		next(w, r)
	}
}
