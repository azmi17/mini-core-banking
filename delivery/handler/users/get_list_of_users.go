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

func GetListOfUsers(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	payload := entities.LimitOffsetLkmUri{}
	httpio.BindUri(&payload)

	formPayload := entities.GlobalFilter{}
	httpio.Bind(&formPayload)

	usecase := usecase.NewSysUserUsecase()
	data, er := usecase.GetListOfUsers(formPayload, payload)
	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "record not found.", nil)
		} else if er == err.BadRequest {
			httpio.ResponseString(http.StatusBadRequest, "invalid parameters", nil)
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		httpio.Response(http.StatusOK, data)
	}

}
