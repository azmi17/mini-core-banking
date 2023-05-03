package web

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
	KodeProduk      string `json:"kode_produk"`
	DeskripsiProduk string `json:"deskripsi_produk"`
}
