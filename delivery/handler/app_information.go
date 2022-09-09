package handler

import (
	"apex-ems-integration-clean-arch/delivery/handler/httpio"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AppInfo(ctx *gin.Context) {

	httpio := httpio.NewRequestIO(ctx)
	httpio.Recv()

	appInfo := map[string]interface{}{
		"App Name":        os.Getenv("application.name"),
		"App Description": os.Getenv("application.desc"),
		"App Version":     os.Getenv("application.version"),
		"App Author":      os.Getenv("application.author"),
		"Port Listener":   os.Getenv("app.listener_port"),
	}

	httpio.Response(http.StatusOK, appInfo)
}
