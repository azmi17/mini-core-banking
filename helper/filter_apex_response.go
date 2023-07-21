package helper

import (
	"new-apex-api/entities"
)

func ApexCreateLKMResponse(nasabah entities.Nasabah, tabung entities.Tabung, sysDaftarUser entities.SysDaftarUser) entities.LKMCreateResponse {
	return entities.LKMCreateResponse{
		KodeLkm:        nasabah.Nasabah_Id,
		NamaLembaga:    nasabah.Nama_Nasabah,
		Alamat:         nasabah.Alamat,
		Telepon:        nasabah.Telepon,
		KodeSC:         tabung.Kode_Group2,
		NoRekening:     tabung.No_Rekening,
		Status:         tabung.Status,
		Saldo:          int(tabung.Saldo_Akhir),
		Plafond:        int(tabung.Minimum),
		SetoranMinimum: int(tabung.Setoran_Minimum),
		IsSaldoChecked: tabung.IsSaldoChecked,
		UserNameSmec:   sysDaftarUser.User_Name,
		PwdSmec:        sysDaftarUser.User_Web_Password_Hash,
		User_Id:        tabung.UserId,
	}
}

func ApexUpdateLKMResponse(nasabah entities.Nasabah, tabungan entities.UpdateRekeningLKM, sysDaftarUser entities.SysDaftarUser) entities.LKMUpdateResponse {
	return entities.LKMUpdateResponse{
		KodeLkm:        nasabah.Nasabah_Id,
		NamaLembaga:    sysDaftarUser.Nama_Lengkap,
		KodeSC:         nasabah.Kode_Group2,
		Alamat:         nasabah.Alamat,
		Telepon:        nasabah.Telepon,
		Plafond:        tabungan.Plafond,
		SetoranMinimum: tabungan.SetoranMinimum,
		Status:         tabungan.Status,
		IsSaldoChecked: tabungan.IsSaldoChecked,
		UserID:         nasabah.UserId,
	}
}

func SysUserLoginResponseFilter(sysDaftarUser entities.SysDaftarUser) entities.LoginData {
	return entities.LoginData{
		User_Id:         sysDaftarUser.User_Id,
		User_Name:       sysDaftarUser.User_Name,
		Nama_lengkap:    sysDaftarUser.Nama_Lengkap,
		Tanggal_Expried: string(sysDaftarUser.Tgl_Expired.Format("02-01-2006")),
	}
}
