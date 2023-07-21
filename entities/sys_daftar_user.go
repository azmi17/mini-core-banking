package entities

import "time"

type SysDaftarUser struct {
	User_Id                int
	User_Name              string
	User_Password          string
	Nama_Lengkap           string
	Penerimaan             float32
	Pengeluaran            float32
	Unit_Kerja             string
	Jabatan                string
	User_Code              string
	Tgl_Expired            time.Time
	TglExpiredStr          string
	Flag                   int
	Status_Aktif           int
	User_Web_Password      string
	User_Web_Password_Hash string
}
