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

func DeleteRoutingRekIndukByKodeLKM(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := entities.MultipleKodeLKM{}
	httpio.Bind(&payload)

	usecase := usecase.NewRoutingIndukUsecase()
	er := usecase.DeleteSysApexRoutingRekInduk(payload.ListOfKodeLKM)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, nil)
	}

}
