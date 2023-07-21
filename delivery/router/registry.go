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
	alterApexHandler "new-apex-api/delivery/handler/alter-apex"
	vaApexHandler "new-apex-api/delivery/handler/apex-virtualaccount"
	appInformationsHandler "new-apex-api/delivery/handler/app-informations"
	approvalHandler "new-apex-api/delivery/handler/approvals"
	authHandler "new-apex-api/delivery/handler/auth"
	echHandler "new-apex-api/delivery/handler/ech-transhistories"
	institutionsHandler "new-apex-api/delivery/handler/institutions"
	institutionRoutingsHandler "new-apex-api/delivery/handler/lkm-routings"
	referensiHandler "new-apex-api/delivery/handler/referensi-apex"
	repostingHandler "new-apex-api/delivery/handler/repostings"
	tabtransHandler "new-apex-api/delivery/handler/tabtrans"
	usersHandler "new-apex-api/delivery/handler/users"

	"github.com/gin-gonic/gin"
)

func RegisterHandler(router *gin.Engine) {
	apiv1 := router.Group("api/v2")
	{
		// DEVELOPMENT ENDPOINT
		apiv1.GET("/version", appInformationsHandler.AppInfo)
		apiv1.DELETE("/flush", institutionsHandler.HardDeleteLKM)

		// INSTITUTIONS ENDPOINT
		apiv1.POST("/institutions/:limit/:offset", institutionsHandler.GetLkmInfoLists)
		apiv1.POST("/institution", institutionsHandler.CreateLKM)

		// REPOSTING ENDPOINT
		apiv1.POST("/repostings", repostingHandler.RepostingSaldoByBulk)
		apiv1.POST("/repostings/all", repostingHandler.RepostingAllLKM)

		// USERS ENDPOINT
		apiv1.POST("/users/:limit/:offset", usersHandler.GetListOfUsers)
		apiv1.POST("/user", usersHandler.CreateSysUser)
		apiv1.POST("/user/search", usersHandler.FindSingleUserByUserName)
		apiv1.POST("/user/login", usersHandler.LoginUser)
		apiv1.PUT("/user", usersHandler.UpdateSysUser)
		apiv1.PUT("/user/reset-password", usersHandler.ResetApexPassword)

		// ROUTING-LKMS ENDPOINT
		apiv1.POST("/routings/:limit/:offset", institutionRoutingsHandler.GetListRoutingRekInduk)
		apiv1.POST("/routing", institutionRoutingsHandler.CreateRoutingRekInduk)
		apiv1.PUT("/routing", institutionRoutingsHandler.UpdateRoutingRekInduk)
		apiv1.DELETE("/routing", institutionRoutingsHandler.DeleteRoutingRekIndukByKodeLKM)

		// TABTRANS ENDPOINT
		apiv1.POST("/tabtrans/:limit/:offset", tabtransHandler.GetListsTabtransTransaction)

		// LAPORAN TABTRANS ENDPOINT
		apiv1.POST("/laporan/rekening_koran", tabtransHandler.GetRekeningKoranLKMDetail)
		apiv1.POST("/laporan/nominatif_deposit/:limit/:offset", tabtransHandler.GetNominatifDeposit)
		apiv1.POST("/laporan/daftar_transaksi/:limit/:offset", tabtransHandler.GetLaporanTransaksi)
		apiv1.POST("/laporan/deposits", tabtransHandler.GetListsDepositHisotry)

		// TRANSAKSI DEPOSIT ENDPOINT
		apiv1.POST("/deposit", tabtransHandler.TransaksiDeposit)

		// REFERENCES APEX DATA ENDPOINT
		apiv1.GET("/referensi/vendors", referensiHandler.GetListsScGroup)
		apiv1.GET("/referensi/banks", referensiHandler.GetListsBankGroup)
		apiv1.GET("/referensi/jenis_deposit", referensiHandler.GetListsJenisTransaksiDeposit)
		apiv1.GET("/referensi/jenis_tabungan", referensiHandler.GetListsJenisTransaksiTabungan)
		apiv1.GET("/referensi/tabproducts", referensiHandler.GetlistsProdukTabungan)
		apiv1.GET("/referensi/otorisators", referensiHandler.GetListsOtorisator)
		apiv1.GET("/referensi/jenis_pembayaran_sla", referensiHandler.GetListsJenisPembayaranVaSLA)
		apiv1.GET("/referensi/tabungan_integrasi", referensiHandler.GetListsTabunganIntegrasi)

		// ECHANNEL ENDPOINT
		apiv1.POST("/echannel/trans_histories/:limit/:offset", echHandler.GetEchannelTransHistories)

		// APPROVAL ENDPOINT
		apiv1.POST("/approval/token", approvalHandler.RequestNewTokenCode)
		apiv1.GET("/approvals/:limit/:offset", approvalHandler.GetApprovalLists)

		// SLA VIRTUAL ACCOUNT ENDPOINT
		apiv1.POST("/sla/virtual_account", vaApexHandler.CreateApexSLAVirtualAccount)
		apiv1.GET("/sla/virtual_accounts/:limit/:offset", vaApexHandler.GetApexSLAVirtualAccount)
		apiv1.PUT("/sla/virtual_account", vaApexHandler.UpdateApexSLAVirtualAccount)
		apiv1.POST("/sla/virtual_accounts/transaksi/:limit/:offset", vaApexHandler.GetListsSLATransactionVirtualAccount)

		// MIGRATE & TOOLS
		apiv1.GET("/migrate/nasabahid", alterApexHandler.AlterNasabahIDOnTabungAndNasabah)
	}

	// AUTHORIZE (PROTECT BY MIDDLEWARE)
	authorized := router.Group("api/v2", authHandler.AuthRequired)
	{
		authorized.POST("/deposit/reversal", tabtransHandler.PembatalanTransaksiDeposit)
		authorized.PUT("/institution", institutionsHandler.UpdateLKM)
		// authorized.PUT("/tabtrans/change_date", tabtransHandler.ChangeDateOnTabtransTrx)
		authorized.PUT("/tabtrans/change_date", tabtransHandler.MultipleChangeDateApexTransactions)
		authorized.DELETE("/tabtrans", tabtransHandler.PermanentlyDeleteApexTransaction)
		authorized.DELETE("/institution", institutionsHandler.DeleteLKM)
		authorized.DELETE("/sla/virtual_account", vaApexHandler.DeleteVAAccounts)
		authorized.DELETE("/user", usersHandler.PermanentlyDeleteUser)

	}

}
