package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/statuscode"
	"new-apex-api/entities/web"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func DeleteLKM(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := web.KodeLKMUri{}
	httpio.BindUri(&payload)

	usecase := usecase.NewLkmUsecase()
	er := usecase.DeleteLkm(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, nil)
	}

}

func HardDeleteLKM(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := web.KodeLKMFilter{}
	httpio.Bind(&payload)

	usecase := usecase.NewLkmUsecase()
	er := usecase.HardDeleteLkm(payload.KodeLkm)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "Record not found!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, nil)
	}

}
