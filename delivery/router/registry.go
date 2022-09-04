/*
 * Copyright (c) 2022 Randy Ardiansyah https://github.com/randyardiansyah25/<repo>
 *
 * Created Date: Wednesday, 16/03/2022, 10:32:08
 * Author: Randy Ardiansyah
 *
 * Filename: /home/Documents/workspace/go/src/router-template/delivery/router/registry.go
 * Project : /home/Documents/workspace/go/src/router-template/delivery/router
 *
 * HISTORY:
 * Date                  	By                 	Comments
 * ----------------------	-------------------	--------------------------------------------------------------------------------------------------------------------
 */

package router

import (
	"apex-ems-integration-clean-arch/delivery/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.Engine) {

	// API Versioning:
	apiv1 := router.Group("api/v1")

	// API Endpoint:
	apiv1.GET("/version", func(c *gin.Context) {
		c.String(http.StatusOK, "Router Template V1.0.0\n")
		c.String(http.StatusOK, "Apex API Endpoint")
	})
	apiv1.POST("/institutions", handler.CreateLKM)
	apiv1.PUT("/institutions", handler.UpdateLKM)
	apiv1.DELETE("/institutions", handler.DeleteLKM)

}
