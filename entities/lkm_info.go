package entities

type LkmInfo struct {
	KodeLKM         string  `json:"kode_lkm"`
	NasabahId       string  `json:"nasabah_id"`
	NamaLembaga     string  `json:"nama_lembaga"`
	Alamat          string  `json:"alamat"`
	TanggalRegister string  `json:"tanggal_register"`
	Plafond         float64 `json:"plafond"`
	SetoranMinimal  float64 `json:"setoran_minimal"`
	SaldoAkhir      float64 `json:"saldo_akhir"`
	StatusRekening  int     `json:"status_rekening"`
}
