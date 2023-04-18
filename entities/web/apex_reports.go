package web

// LAPORAN REKENING KORAN
type SaldoAwal struct {
	KodeLKM string
	Debit   float64
	Kredit  float64
}

type RekeningKoran struct {
	TglTrans       string
	Keterangan     string
	KodeTrans      string
	Pokok          float64
	MyKodeTrans    string
	Kuitansi       string
	PayLkmNorek    string
	PayIdpel       string
	PayBillerCode  string
	PayProductCode string
}

type RekeningKoranHeader struct {
	Norek       string
	NamaLembaga string
	ProdukTab   string
	NamaSC      string
}

type RekeningKoranRequest struct {
	KodeLKM      string `form:"kode_lkm"`
	PeriodeAwal  string `form:"periode_awal"`
	PeriodeAkhir string `form:"periode_akhir"`
}

type RekeningKoranResponse struct {
	Norek        string                `json:"norek"`
	NamaLembaga  string                `json:"nama_lembaga"`
	ProdukTab    string                `json:"produk_tab"`
	NamaSC       string                `json:"nama_sc"`
	PeriodeAwal  string                `json:"periode_awal"`
	PeriodeAkhir string                `json:"periode_akhir"`
	SaldoAwal    float64               `json:"saldo_awal"`
	Detail       []RekeningKoranDetail `json:"detail"`
}

type RekeningKoranDetail struct {
	TglTrans  string  `json:"tgl_trans"`
	Uraian    string  `json:"uraian"`
	KodeTrans string  `json:"kode_trans"`
	Debet     float64 `json:"debet"`
	Kredit    float64 `json:"kredit"`
	Saldo     float64 `json:"saldo"`
	Kuitansi  string  `json:"kuitansi"`
	NorekLKM  string  `json:"norek_lkm"`
	Idpel     string  `json:"idpel"`
	Biller    string  `json:"biller"`
	Produk    string  `json:"produk"`
}

// LAPORAN NOMINTAIF DEPOSIT
type RawQueryNominatifDeposit struct {
	NoRekening       string
	NamaLembaga      string
	Alamat           string
	SaldoAkhir       float64
	LastMonthTrans   float64
	LastMonthReverse float64
	LastActivity     string
}

type NominatifDepositRequest struct {
	TanggalHitung string `form:"tgl_hitung"`
	// Order         string
}

type NominatifDepositResponse struct {
	KodeLKM      string  `json:"kode_lkm"`
	NamaLembaga  string  `json:"nama_lembaga"`
	Alamat       string  `json:"alamat"`
	SaldoAkhir   float64 `json:"saldo_akhir"`
	JmlTrans     float64 `json:"jml_trans"`
	RataHarian   float64 `json:"rata_rata_harian"`
	LastActivity string  `json:"last_activity"`
}

// LAPORAN DAFTAR TRANSAKSI
type DaftarTransaksiRequest struct {
	PeriodeAwal    string `form:"periode_awal"`
	PeriodeAkhir   string `form:"periode_akhir"`
	JenisTransaksi string `form:"jenis_transaksi"`
}

type DaftarTransaksiResponse struct {
	NoRekening         string  `json:"kode_lkm"`
	NamaLembaga        string  `json:"nama_lembaga"`
	TglTrans           string  `json:"tgl_trans"`
	NoBukti            string  `json:"no_bukti"`
	KodeTrans          string  `json:"kode_trans"`
	DeskripsiTransaksi string  `json:"deskripsi_transaksi"`
	Setoran            float64 `json:"setoran"`
	Penarikan          float64 `json:"penarikan"`
}
