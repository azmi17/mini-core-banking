package entities

import "time"

type Nasabah struct {
	Nasabah_Id            string
	Nama_Nasabah          string
	Nama_Alias            string
	Nama_Nasabah_Sid      string
	Alamat                string
	Alamat2               string
	Telepon               string
	Jenis_Kelamin         string
	TempatLahir           string
	TglLahir              time.Time
	Jenis_Id              string
	No_Id                 string
	Status_Gelar          string
	Jenis_Debitur         string
	Kode_Area             string
	Negara_Domisili       string
	Gol_Debitur           string
	Langgar_Bmpk          string
	Lampaui_Bmpk          string
	Flag_Masa_Berlaku     string
	Status_Marital        string
	Kode_Group1           string
	Kode_Group2           string
	Kode_Group3           string
	Kode_Agama            string
	Desa                  string
	Kecamatan             string
	Kota_Kab              string
	Provinsi              string
	Verifikasi            string
	Hp                    string
	Hp1                   string
	Hp2                   string
	Tgl_Register          time.Time
	Nama_Ibu_Kandung      string
	Kodepos               string
	Kode_Kantor           string
	Status_Tempat_Tinggal string
	UserId                int
	Masa_Berlaku_Ktp      time.Time
}
