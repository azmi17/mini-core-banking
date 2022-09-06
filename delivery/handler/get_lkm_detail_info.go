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

func GetLkmInfo(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form (Randy's Framework implementations)
	payload := web.KodeLKMUri{}
	httpio.BindUri(&payload)

	// GIN Implementations..
	// er := ctx.ShouldBindUri(&payload)
	// if er != nil {
	// 	response := helper.ApiResponse("Failed to get detail of campaign", "error", nil)
	// 	ctx.JSON(http.StatusBadRequest, response)
	// 	return
	// }

	usecase := usecase.NewApexUsecase()
	getUser, er := usecase.GetLkmDetailInfo(payload.UserName)
	if er != nil {
		if er == err.DuplicateEntry {
			httpio.ResponseString(statuscode.StatusDuplicate, "institution data is available!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {
		response := helper.ApiResponse("Get institution detail succeeded", "success", getUser)
		httpio.Response(http.StatusOK, response)
	}

}
