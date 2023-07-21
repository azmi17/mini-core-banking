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

func GetDetailLKMInfo(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form (Randy's Framework implementations)
	payload := entities.KodeLKMUri{}
	httpio.BindUri(&payload)

	usecase := usecase.NewLkmUsecase()
	getUser, er := usecase.GetLKMDetailInfo(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, getUser)
	}

}
