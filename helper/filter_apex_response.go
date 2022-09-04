package helper

import (
	"apex-ems-integration-clean-arch/entities"
	"apex-ems-integration-clean-arch/entities/web"
)

func ApexFilterLKMResponse(nasabah entities.Nasabah, tabung entities.Tabung, sysDaftarUser entities.SysDaftarUser) web.LKMCreateResponse {
	return web.LKMCreateResponse{
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

func ApexUpdateLKMResponse(nasabah entities.Nasabah, sysDaftarUser entities.SysDaftarUser) web.LKMUpdateResponse {
	return web.LKMUpdateResponse{
		KodeLkm:      nasabah.Nasabah_Id,
		Nama_Lembaga: sysDaftarUser.Nama_Lengkap,
		Alamat:       nasabah.Alamat,
		Telepon:      nasabah.Telepon,
		User_Id:      nasabah.UserId,
	}
}
