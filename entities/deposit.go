package entities

import "time"

type ReversalDepositRequest struct {
	Kuitansi string `form:"kuitansi" binding:"required"`
	UserID   int    `form:"user_id" binding:"required"`
}

type Deposit struct {
	TabtransID    int
	Tgl_trans     time.Time
	NoRekening    string
	MyKodeTrans   string
	Pokok         float64
	Keterangan    string
	Verifikasi    string
	UserID        int
	ModulIDSource string
	TransIDSource int
	KodeTrans     string
	Tob           string
	PostedToGl    int
	KodePerkOB    string
	KodeKantor    string
	SandiTrans    string
	Kuitansi      string
	CounterSign   int
	NoRekeningABA string
}

type DepositRequest struct {
	KodeLKM         string  `form:"kode_lkm" binding:"required"`
	JenisTransaksi  string  `form:"jenis_transaksi" binding:"required"`
	JumlahTransaksi float64 `form:"jumlah_transaksi" binding:"required"`
	Keterangan      string  `form:"keterangan" binding:"required"`
	UserID          int     `form:"user_id" binding:"required"`
}
