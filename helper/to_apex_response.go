package helper

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

func ToApexResponse(nasabah entities.Nasabah, tabung entities.Tabung, sysDaftarUser entities.SysDaftarUser) web.ApexResponse {
	return web.ApexResponse{
		KodeLkm:        nasabah.Nasabah_Id,
		Nama_Lembaga:   nasabah.Nama_Nasabah,
		Alamat:         nasabah.Alamat,
		Telepon:        nasabah.Telepon,
		No_rekening:    tabung.No_Rekening,
		Saldo_Akhir:    int(tabung.Saldo_Akhir),
		User_Name_Smec: sysDaftarUser.User_Name,
		Password_Smec:  sysDaftarUser.User_Web_Password_Hash,
		User_Id:        tabung.UserId,
	}
}
