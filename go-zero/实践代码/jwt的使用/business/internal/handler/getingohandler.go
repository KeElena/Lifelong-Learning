package handler

import (
	"net/http"

	"business/internal/logic"
	"business/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func getIngoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewGetIngoLogic(r.Context(), svcCtx)
		resp, err := l.GetIngo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
