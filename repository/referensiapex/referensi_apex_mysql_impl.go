package referensiapex

import (
	"database/sql"
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"os"
	"strings"
)

func newReferensiApexMysqlImpl(apexConn *sql.DB) ReferensiApexRepo {
	return &referensiApexMysqlImpl{
		apexDb: apexConn,
	}
}

type referensiApexMysqlImpl struct {
	apexDb *sql.DB
}

func (r *referensiApexMysqlImpl) GetListsScGroup() (lists []web.ScGroup, er error) {
	rows, er := r.apexDb.Query("SELECT kode_group2, deskripsi_group2 FROM tab_kode_group2")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var scGroup web.ScGroup
		if er = rows.Scan(&scGroup.KodeGroup, &scGroup.DeskripsiGroup); er != nil {
			return lists, er
		}

		lists = append(lists, scGroup)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsJenisTransaksiTabungan() (lists []web.JenisTransaksi, er error) {
	rows, er := r.apexDb.Query("SELECT kode_trans, deskripsi_trans FROM tab_kode_trans")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabKodeTrans web.JenisTransaksi
		if er = rows.Scan(&tabKodeTrans.KodeTrans, &tabKodeTrans.DeskripsiTrans); er != nil {
			return lists, er
		}

		lists = append(lists, tabKodeTrans)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsJenisTransaksiDeposit() (lists []web.JenisTransaksi, er error) {

	item := strings.Split(os.Getenv("app.jenis_transaksi_deposit"), ",")
	kodeTrans := (strings.Join(item, ","))

	rows, er := r.apexDb.Query(`SELECT kode_trans, deskripsi_trans FROM tab_kode_trans WHERE kode_trans IN (` + kodeTrans + `)`)
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabKodeTrans web.JenisTransaksi
		if er = rows.Scan(&tabKodeTrans.KodeTrans, &tabKodeTrans.DeskripsiTrans); er != nil {
			return lists, er
		}

		lists = append(lists, tabKodeTrans)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetListsBankGroup() (lists []web.BankGroup, er error) {
	rows, er := r.apexDb.Query("SELECT id, nama_bank, no_rekening, deskripsi FROM deposit_bank_info")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var banks web.BankGroup
		if er = rows.Scan(&banks.Id, &banks.NamaBank, &banks.NoRekeningBank, &banks.Deskripsi); er != nil {
			return lists, er
		}

		lists = append(lists, banks)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}

func (r *referensiApexMysqlImpl) GetlistsProdukTabungan() (lists []web.ProdukTabungan, er error) {
	rows, er := r.apexDb.Query("SELECT kode_produk, deskripsi_produk FROM tab_produk")
	if er != nil {
		return lists, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabProduk web.ProdukTabungan
		if er = rows.Scan(&tabProduk.KodeProduk, &tabProduk.DeskripsiProduk); er != nil {
			return lists, er
		}

		lists = append(lists, tabProduk)
	}

	if len(lists) == 0 {
		return lists, err.NoRecord
	} else {
		return
	}
}
