package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"apex-ems-integration-clean-arch/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AppInfo(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	appInfo := map[string]interface{}{
		"App Name":        helper.AppAuthor,
		"App Description": helper.AppDescription,
		"App Version":     helper.AppVersion,
		"App Author":      helper.AppAuthor,
	}

	httpio.Response(http.StatusOK, appInfo)
}
