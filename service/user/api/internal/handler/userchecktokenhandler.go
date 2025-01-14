package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"togds/service/user/api/internal/logic"
	"togds/service/user/api/internal/svc"
	"togds/service/user/api/internal/types"
)

func UserCheckTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUserCheckTokenLogic(r.Context(), svcCtx)
		resp, err := l.UserCheckToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
