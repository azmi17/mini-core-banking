package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/statuscode"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateLKM(ctx *gin.Context) {

	// Init HTTP Request..
	httpio := httpio.NewRequestIO(ctx)

	// Call Payload and binding form
	payload := web.SaveLKMApex{}
	httpio.Bind(&payload)

	usecae := usecase.NewApexUsecase()
	lkm, er := usecae.CreateLkm(payload)
	if er != nil {
		if er == err.DuplicateEntry {
			httpio.ResponseString(statuscode.StatusDuplicate, "Institution data is available!", nil)
		} else {
			entities.PrintError(er.Error())
			entities.PrintLog(er.Error())
			httpio.ResponseString(http.StatusInternalServerError, "internal service error", nil)
		}
	} else {

		httpio.Response(http.StatusOK, lkm)
	}

}
