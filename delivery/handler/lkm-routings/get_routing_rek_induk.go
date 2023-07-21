package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/statuscode"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func GetRoutingRekInduk(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := entities.KodeLKMUri{}
	httpio.BindUri(&payload)

	usecase := usecase.NewRoutingIndukUsecase()
	data, er := usecase.GetRoutingRekInduk(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "record not found", nil)
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, data)
	}

}
