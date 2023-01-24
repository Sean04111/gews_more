package registercode

import (
	"net/http"

	"gews_more/service/internal/logic/registercode"
	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func RegistercodeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Registercodereque
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := registercode.NewRegistercodeLogic(r.Context(), svcCtx)
		resp, err := l.Registercode(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
