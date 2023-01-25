package snappost

import (
	"net/http"

	"gews_more/service/internal/logic/snappost"
	"gews_more/service/internal/svc"
	"gews_more/service/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PostsnapHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Snappostreque
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := snappost.NewPostsnapLogic(r.Context(), svcCtx)
		resp, err := l.Postsnap(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
