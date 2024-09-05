package electricity

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"yxy-go/internal/logic/electricity"
	"yxy-go/internal/svc"
	"yxy-go/internal/types"
	"yxy-go/pkg/response"
)

func GetElectricitySurplusHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetElectricitySurplusReq
		if err := httpx.Parse(r, &req); err != nil {
			response.ParamErrorResponse(r, w, err)
			return
		}

		l := electricity.NewGetElectricitySurplusLogic(r.Context(), svcCtx)
		resp, err := l.GetElectricitySurplus(&req)
		response.HttpResponse(r, w, resp, err)
	}
}