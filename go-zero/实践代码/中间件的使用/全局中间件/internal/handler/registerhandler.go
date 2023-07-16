package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"middlewares/internal/logic"
	"middlewares/internal/svc"
)

func registerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewRegisterLogic(r.Context(), svcCtx)
		resp, err := l.Register()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
