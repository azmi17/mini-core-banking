package entities

import "time"

type Tabung struct {
	No_Rekening               string
	Nasabah_Id                string
	Kode_Bi_Pemilik           string
	Suku_Bunga                float32
	Persen_Pph                float32
	Tgl_Register              time.Time
	Saldo_Akhir               float64
	Kode_Group1               string
	Kode_Group2               string
	Kode_Group3               string
	Verifikasi                string
	Status                    int
	Kode_Kantor               string
	Kode_Integrasi            string
	Kode_Produk               string
	UserId                    int
	Minimum                   float32
	Setoran_Minimum           float32
	Jkw                       int
	Abp                       int
	Setoran_Wajib             float32
	Adm_Per_Bln               float32
	Target_Nominal            float32
	Saldo_Akhir_Titipan_bunga float32
	Kode_Bi_Lokasi            string
	Saldo_Titipan_Pokok       float32
	Saldo_Titipan_Bunga_Ks    float32
	Saldo_Blokir              float32
	Premi                     float32
	Kode_Keterkaitan          string
	Kode_Kantor_Kas           string
	No_Rekening_Virtual       string
}
