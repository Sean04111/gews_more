package trending

import (
	"net/http"

	"gews_more/service/internal/logic/trending"
	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func TrendingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Trendingreque
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := trending.NewTrendingLogic(r.Context(), svcCtx)
		resp, err := l.Trending(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
