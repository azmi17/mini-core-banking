package handler

import (
	"net/http"
	"new-apex-api/delivery/handler/httpio"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/statuscode"
	"new-apex-api/helper"
	"new-apex-api/usecase"

	"github.com/gin-gonic/gin"
)

func GetEchannelTransHistories(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	uriPayload := entities.LimitOffsetLkmUri{}
	httpio.BindUri(&uriPayload)

	formPayload := entities.TransHistoryRequest{}
	rerr := httpio.BindWithErr(&formPayload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("get trans_histories transaction failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewtransHisotryEchannelUsecase()
	listTrx, er := usecase.TransHistoriesLists(formPayload, uriPayload)

	if er != nil {
		if er == err.NoRecord {
			httpio.ResponseString(statuscode.StatusNoRecord, "record not found", nil)
			return
		} else if er == err.BadRequest {
			httpio.ResponseString(http.StatusBadRequest, "invalid parameters", nil)
			return
		} else {
			entities.PrintError(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
			return
		}
	} else {
		httpio.Response(http.StatusOK, listTrx)
	}
}
