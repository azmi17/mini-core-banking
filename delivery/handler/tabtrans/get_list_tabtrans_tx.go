package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/statuscode"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTabungansLkmInfoList(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)

	uriPayload := web.LimitOffsetLkmUri{}
	httpio.BindUri(&uriPayload)

	formPayload := web.GetListTabtransByDate{}
	rerr := httpio.BindWithErr(&formPayload)
	if rerr != nil {
		errors := helper.FormatValidationError(rerr)
		errorMesage := gin.H{"errors": errors}
		response := helper.ApiResponse("get tabtrans transaction failed", http.StatusUnprocessableEntity, "failed", errorMesage)
		httpio.Response(http.StatusUnprocessableEntity, response)
		return
	}

	usecase := usecase.NewTabtransUsecase()
	tabtransTxList, total, er := usecase.GetListTabtransInfo(formPayload, uriPayload)

	resp := web.GetListTabtransInfoWithCountSumResp{}
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
		resp.TotalTrx = total.TotalTrx
		resp.TotalPokok = total.TotalPokok
		resp.Data = &tabtransTxList
	}
	httpio.Response(http.StatusOK, resp)
}
