package entities

type ScGroup struct {
	KodeGroup      string `json:"kode_group"`
	DeskripsiGroup string `json:"deskripsi_group"`
}

type JenisTransaksi struct {
	KodeTrans      string `json:"kode_trans"`
	DeskripsiTrans string `json:"deskripsi_transaksi"`
}

type BankGroup struct {
	Id             int    `json:"id"`
	NamaBank       string `json:"nama_bank"`
	NoRekeningBank string `json:"no_rekening_bank"`
	Deskripsi      string `json:"deskripsi"`
}

type ProdukTabungan struct {
	KodeProduk      string `json:"kode_produk_tabungan"`
	DeskripsiProduk string `json:"deskripsi_produk_tabungan"`
}

type Otorisators struct {
	UserID         int    `json:"user_id"`
	NamaOtorisator string `json:"nama_otorisator"`
	Jabatan        string `json:"jabatan"`
}

type JenisPembayaranSLA struct {
	KodePembayaran string `json:"kode_va"`
	NamaPembayaran string `json:"nama_va"`
}

type TabunganIntegrasi struct {
	KodeIntegrasi      string `json:"kode_integrasi"`
	DeskripsiIntegrasi string `json:"deskripsi_integrasi"`
}
