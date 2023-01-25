package getsnaps

import (
	"net/http"

	"gews_more/service/internal/logic/getsnaps"
	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetsnapsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Snapreque
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := getsnaps.NewGetsnapsLogic(r.Context(), svcCtx)
		resp, err := l.Getsnaps(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
