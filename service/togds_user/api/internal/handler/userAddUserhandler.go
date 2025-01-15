package handler

import (
	"net/http"

	"togds/service/togds_user/api/internal/logic"
	"togds/service/togds_user/api/internal/svc"
	"togds/service/togds_user/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func UserAddUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewAddUserLogic(r.Context(), svcCtx)
		resp, err := l.UserCheckToken(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
