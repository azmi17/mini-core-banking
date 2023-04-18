package tabtransrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
	"new-apex-api/entities/web"
	"new-apex-api/helper"
	"sync"
)

func newTabtransMysqlImpl(apexConn *sql.DB) TabtransRepo {
	return &TabtransMysqlImpl{
		apexDb: apexConn,
	}
}

type TabtransMysqlImpl struct {
	apexDb *sql.DB
}

func (t *TabtransMysqlImpl) GetNextIDWithUserID() (transId int, er error) {

	userId := helper.GetUserIDApp()
	row := t.apexDb.QueryRow(`SELECT ibs_get_next_id_with_userid(?) AS trans_id`, userId)
	er = row.Scan(
		&transId,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return transId, err.NoRecord
		} else {
			return transId, errors.New(fmt.Sprint("error while get data: ", er.Error()))
		}
	}
	return transId, nil
}

func (t *TabtransMysqlImpl) GetSingleTabtransTrx(tabtransID int) (data web.GetListTabtransTrx, er error) {
	row := t.apexDb.QueryRow(`SELECT 
		tabtrans_id,
		DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS tgl_trans,
		no_rekening,
		pokok,
		CASE my_kode_trans WHEN '200' THEN 'D' ELSE 'K' END AS dk,
		COALESCE(pay_lkm_norek,'') AS lkm_norek,
		COALESCE(pay_idpel,'') AS idpel,
		kode_trans,
		kuitansi,
		keterangan,
		COALESCE(pay_biller_code,'') AS biller_code,
		COALESCE(pay_product_code, '') AS product_code,
		userid
	FROM tabtrans WHERE tabtrans_id = ?`, tabtransID)
	er = row.Scan(
		&data.TabtransID,
		&data.TglTrans,
		&data.KodeLKM,
		&data.Pokok,
		&data.Dk,
		&data.Lkm_Norek,
		&data.Idpel,
		&data.KodeTrans,
		&data.Kuitansi,
		&data.Keterangan,
		&data.BillerCode,
		&data.ProductCode,
		&data.UserID,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return data, err.NoRecord
		} else {
			return data, errors.New(fmt.Sprint("error while get tabtrans data: ", er.Error()))
		}
	}
	return
}

func (t *TabtransMysqlImpl) GetListsTabtransTrx(payload web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (list []web.GetListTabtransTrx, total web.GetCountWithSumTabtransTrx, er error) {
	var rows *sql.Rows

	/*
		Dibawah adalah slice of interface untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar re-usable pada 2 kondisi
	*/
	args := []interface{}{}
	sqlCond := ""
	sqlStmt := `
	SELECT 
			t.tabtrans_id,
			DATE_FORMAT(t.tgl_trans, "%d/%m/%Y") AS tgl_trans,
			t.no_rekening AS KodeLkm,
			COALESCE(n.nama_nasabah,'') AS nama_lembaga,
			t.pokok,
			CASE t.my_kode_trans WHEN '200' THEN 'D' ELSE 'K' END AS dk,
			COALESCE(t.pay_lkm_norek,'') AS lkm_norek,
			COALESCE(t.pay_idpel,'') AS idpel,
			t.kode_trans,
			t.kuitansi,
			t.keterangan,
			COALESCE(t.pay_biller_code,'') AS biller_code,
			COALESCE(t.pay_product_code, '') AS product_code,
			t.userid
		FROM tabtrans AS t 
		LEFT JOIN tabung AS tb ON(t.no_rekening=tb.no_rekening) 
		LEFT JOIN nasabah AS n ON (tb.nasabah_id=n.nasabah_id) WHERE `

	if payload.BankCode == "" {
		if limitOffset.Limit > 0 {
			sqlCond = "t.tgl_trans >= ? AND t.tgl_trans <= ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "t.tgl_trans >= ? AND t.tgl_trans <= ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, -1, limitOffset.Offset)
		}
		rows, er = t.apexDb.Query(sqlStmt+sqlCond+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			sqlCond = "t.tgl_trans >= ? AND t.tgl_trans <= ? AND t.no_rekening = ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, payload.BankCode, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "t.tgl_trans >= ? AND t.tgl_trans <= ? AND t.no_rekening = ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, payload.BankCode, -1, limitOffset.Offset)
		}
		rows, er = t.apexDb.Query(sqlStmt+sqlCond+``, args...)
	}

	if er != nil {
		return list, total, er
	}

	defer func() {
		_ = rows.Close()
	}()

	sum := 0.0
	for rows.Next() {
		var tabtransListTx web.GetListTabtransTrx
		if er = rows.Scan(
			&tabtransListTx.TabtransID,
			&tabtransListTx.TglTrans,
			&tabtransListTx.KodeLKM,
			&tabtransListTx.NamaLembaga,
			&tabtransListTx.Pokok,
			&tabtransListTx.Dk,
			&tabtransListTx.Lkm_Norek,
			&tabtransListTx.Idpel,
			&tabtransListTx.KodeTrans,
			&tabtransListTx.Kuitansi,
			&tabtransListTx.Keterangan,
			&tabtransListTx.BillerCode,
			&tabtransListTx.ProductCode,
			&tabtransListTx.UserID,
		); er != nil {
			return list, total, er
		}
		list = append(list, tabtransListTx)
		sum += tabtransListTx.Pokok
	}

	if len(list) == 0 {
		return list, total, err.NoRecord
	}

	total.TotalTrx = len(list)
	total.TotalPokok = sum

	return list, total, nil
}

func (t *TabtransMysqlImpl) GetListsTabtransTrxBySTAN(stan string) (trxList []web.GetListTabtransTrx, er error) {
	rows, er := t.apexDb.Query(`SELECT
		tabtrans_id,
		DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS tgl_trans,
		no_rekening,
		pokok,
		CASE my_kode_trans WHEN '200' THEN 'D' ELSE 'K' END AS dk,
		COALESCE(pay_lkm_norek,'') AS lkm_norek,
		COALESCE(pay_idpel,'') AS idpel,
		kode_trans,
		kuitansi,
		keterangan,
		COALESCE(pay_biller_code,'') AS biller_code,
		COALESCE(pay_product_code, '') AS product_code,
		userid
	FROM tabtrans WHERE kuitansi LIKE "%` + stan + `%"`)
	if er != nil {
		return trxList, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabtrans web.GetListTabtransTrx
		if er = rows.Scan(
			&tabtrans.TabtransID,
			&tabtrans.TglTrans,
			&tabtrans.KodeLKM,
			&tabtrans.Pokok,
			&tabtrans.Dk,
			&tabtrans.Lkm_Norek,
			&tabtrans.Idpel,
			&tabtrans.KodeTrans,
			&tabtrans.Kuitansi,
			&tabtrans.Keterangan,
			&tabtrans.BillerCode,
			&tabtrans.ProductCode,
			&tabtrans.UserID,
		); er != nil {
			return trxList, er
		}

		trxList = append(trxList, tabtrans)
	}

	if len(trxList) == 0 {
		return trxList, err.NoRecord
	} else {
		return
	}
}

func (t *TabtransMysqlImpl) DeleteTabtransTrx(tabtransID int) (er error) {

	thisRepo, _ := NewTabtransRepo()
	_, er = thisRepo.GetSingleTabtransTrx(tabtransID)
	if er != nil {
		return err.NoRecord
	}

	stmt, er := t.apexDb.Prepare(`DELETE FROM tabtrans WHERE tabtrans_id = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete tabtrans transaction: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(tabtransID); er != nil {
		return errors.New(fmt.Sprint("error while delete tabtrans transaction: ", er.Error()))
	}

	return nil
}

func (t *TabtransMysqlImpl) ChangeDateOnTabtransTrx(tabtransID int, tglTrans string) (data web.GetListTabtransTrx, er error) {

	thisRepo, _ := NewTabtransRepo()
	tx, er := thisRepo.GetSingleTabtransTrx(tabtransID)
	if er != nil {
		return data, err.NoRecord
	}

	stmt, er := t.apexDb.Prepare(`UPDATE tabtrans SET tgl_trans = ? WHERE tabtrans_id = ?`)
	if er != nil {
		return data, errors.New(fmt.Sprint("error while prepare update tgl_trans: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(tglTrans, tabtransID); er != nil {
		return data, errors.New(fmt.Sprint("error while update tgl_trans: ", er.Error()))
	}

	// pay attention here..
	y := tglTrans[0:4]
	m := tglTrans[4:6]
	d := tglTrans[6:8]
	tglFformat := fmt.Sprintf("%s/%s/%s", d, m, y)
	tx.TglTrans = tglFformat

	return tx, nil

}

func (t *TabtransMysqlImpl) CalculateSaldoOnRekeningLKM(kodeLKM string) (data web.CalculateSaldoResult, er error) {
	var tabtrans web.RepostingData
	rows, err := t.apexDb.Query(`SELECT 
	  tab.no_rekening,
	  SUM(CASE WHEN trans.my_kode_trans='100' THEN trans.pokok ELSE 0 END) AS total_kredit,
	  SUM(CASE WHEN trans.my_kode_trans='200' THEN trans.pokok ELSE 0 END) AS total_debet
	FROM tabung AS tab LEFT JOIN tabtrans AS trans ON (tab.no_rekening = trans.no_rekening)
	WHERE tab.no_rekening = ? GROUP BY tab.no_rekening`, kodeLKM)
	if err != nil {
		return data, er
	}
	for rows.Next() {
		rows.Scan(
			&tabtrans.KodeLKM,
			&tabtrans.TotalKredit,
			&tabtrans.TotalDebet,
		)
	}
	data.KodeLKM = tabtrans.KodeLKM
	data.SaldoAkhir = tabtrans.TotalKredit - tabtrans.TotalDebet

	return data, nil

}

func (t *TabtransMysqlImpl) RepostingSaldoOnRekeningLKM(listOfKodeLKM ...string) (er error) {
	var wg sync.WaitGroup

	for _, each := range listOfKodeLKM {
		wg.Add(1)
		go func(each string, w *sync.WaitGroup) {
			defer w.Done()
			t.doRepostingSaldoProcs(each)
		}(each, &wg)
	}
	wg.Wait()

	return
}

func (t *TabtransMysqlImpl) RepostingSaldoOnRekeningLKMByScheduler(PrintRepoResultChan chan entities.PrintRepo, listOfKodeLKM ...string) (er error) {
	var wg sync.WaitGroup

	for _, each := range listOfKodeLKM {

		wg.Add(1)

		go func(each string, w *sync.WaitGroup) {
			defer w.Done()

			var status = entities.PRINT_SUCCESS_STATUS_REPO_CHAN
			var msg = entities.PRINT_SUCCESS_STATUS_REPO_CHAN

			er := t.doRepostingSaldoProcs(each)
			if er != nil {
				status = entities.PRINT_FAILED_STATUS_REPO_CHAN
				msg = er.Error()
			}

			var printRepo = entities.PrintRepo{
				KodeLKM: each,
				Status:  status,
				Message: msg,
			}

			PrintRepoResultChan <- printRepo

		}(each, &wg)
	}

	wg.Wait()

	return
}

func (t *TabtransMysqlImpl) doRepostingSaldoProcs(data string) (er error) {

	lkm, er := t.CalculateSaldoOnRekeningLKM(data)
	if er != nil {
		return errors.New(fmt.Sprint("error while calculating saldo: ", er.Error()))
	}

	stmt, er := t.apexDb.Prepare(`UPDATE tabung SET saldo_akhir = ? WHERE no_rekening = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare reposting saldo: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er = stmt.Exec(
		lkm.SaldoAkhir,
		lkm.KodeLKM,
	); er != nil {
		return errors.New(fmt.Sprint("error while processing reposting saldo: ", er.Error()))
	}

	return
}

func (t *TabtransMysqlImpl) GetSaldo(kodeLKM, tglAwal string) (total float64, er error) {
	var tabtrans web.SaldoAwal

	row := t.apexDb.QueryRow(`SELECT 
	no_rekening,
	SUM(CASE WHEN my_kode_trans='100' THEN pokok ELSE 0 END) AS kredit,
	SUM(CASE WHEN my_kode_trans='200' THEN pokok ELSE 0 END) AS debit
  FROM tabtrans 
  WHERE no_rekening = ? AND tgl_trans <= ? GROUP BY no_rekening`, kodeLKM, tglAwal)
	if er != nil {
		return total, er
	}
	er = row.Scan(
		&tabtrans.KodeLKM,
		&tabtrans.Kredit,
		&tabtrans.Debit,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return total, err.NoRecord
		} else {
			return total, errors.New(fmt.Sprint("error while get saldo awal: ", er.Error()))
		}
	}
	total = tabtrans.Kredit - tabtrans.Debit

	return total, nil

}

func (t *TabtransMysqlImpl) GetRekeningKoranLKMDetail(kodeLKM, periodeAwal, periodeAkhir string) (data []web.RekeningKoran, er error) {
	rows, er := t.apexDb.Query(`SELECT 
	DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS tgl_trans,
	keterangan,
	kode_trans,
	pokok,
	my_kode_trans,
	kuitansi,
	pay_lkm_norek,
	pay_idpel,
	pay_biller_code,
	pay_product_code
  FROM tabtrans 
  WHERE no_rekening = ? AND tgl_trans >= ? AND tgl_trans <= ?`, kodeLKM, periodeAwal, periodeAkhir)
	if er != nil {
		return data, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var tabtrans web.RekeningKoran
		if er = rows.Scan(
			&tabtrans.TglTrans,
			&tabtrans.Keterangan,
			&tabtrans.KodeTrans,
			&tabtrans.Pokok,
			&tabtrans.MyKodeTrans,
			&tabtrans.Kuitansi,
			&tabtrans.PayLkmNorek,
			&tabtrans.PayIdpel,
			&tabtrans.PayBillerCode,
			&tabtrans.PayProductCode,
		); er != nil {
			return data, er
		}

		data = append(data, tabtrans)
	}

	if len(data) == 0 {
		return data, err.NoRecord
	} else {
		return
	}
}

func (t *TabtransMysqlImpl) GetNominatifDeposit(periode, beginLastMonthDt, EndLastmonthDt string, endLastMonthTg int, limitOffset web.LimitOffsetLkmUri) (laporan []web.NominatifDepositResponse, er error) {

	/*
		Dibawah adalah slice of interface untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar re-usable pada 2 kondisi
	*/
	args := []interface{}{}
	sqlStmt := `SELECT
	tabung.no_rekening,
	nasabah.nama_nasabah AS nama_lembaga,
	nasabah.alamat,
	SUM(IF(FLOOR(my_kode_trans/100)=1,pokok,0)) - SUM(IF(FLOOR(my_kode_trans/100)=2,pokok,0)) AS saldo_akhir,
	SUM(IF(FLOOR(my_kode_trans/100)=2 AND kode_trans<>"315" AND tgl_trans>=` + beginLastMonthDt + ` AND tgl_trans<=` + EndLastmonthDt + `,pokok,0)) as last_month_trans,
	SUM(IF(FLOOR(my_kode_trans/100)=1 AND keterangan1="Reversal" AND tgl_trans>=` + beginLastMonthDt + ` AND  tgl_trans<=` + EndLastmonthDt + `,pokok,0)) as last_month_reverse,
	DATE_FORMAT(MAX(tgl_trans), "%d/%m/%Y") AS last_activity
	FROM tabung, tabtrans, nasabah, tab_kode_group1 
	WHERE `
	sqlCond := `tabung.no_rekening=tabtrans.no_rekening 
		AND nasabah.nasabah_id=tabung.nasabah_id 
		AND tabung.kode_group1=tab_kode_group1.kode_group1
		AND tgl_trans <= ?
		GROUP BY tabung.no_rekening HAVING saldo_akhir > 0 ORDER BY saldo_akhir DESC LIMIT ? OFFSET ?`

	if limitOffset.Limit > 0 {
		args = append(args, periode, limitOffset.Limit, limitOffset.Offset)
	} else {
		args = append(args, periode, -1, limitOffset.Offset)
	}
	rows, er := t.apexDb.Query(sqlStmt+sqlCond+``, args...)
	if er != nil {
		return laporan, er
	}

	defer func() {
		_ = rows.Close()
	}()

	var nominatif web.NominatifDepositResponse
	for rows.Next() {
		var tabtrans web.RawQueryNominatifDeposit
		if er = rows.Scan(
			&tabtrans.NoRekening,
			&tabtrans.NamaLembaga,
			&tabtrans.Alamat,
			&tabtrans.SaldoAkhir,
			&tabtrans.LastMonthTrans,
			&tabtrans.LastMonthReverse,
			&tabtrans.LastActivity,
		); er != nil {
			return laporan, er
		}
		jmlTrans := tabtrans.LastMonthTrans - tabtrans.LastMonthReverse
		nominatif.KodeLKM = tabtrans.NoRekening
		nominatif.NamaLembaga = tabtrans.NamaLembaga
		nominatif.Alamat = tabtrans.Alamat
		nominatif.SaldoAkhir = tabtrans.SaldoAkhir
		nominatif.JmlTrans = jmlTrans
		nominatif.RataHarian = jmlTrans / float64(endLastMonthTg)
		nominatif.LastActivity = tabtrans.LastActivity
		laporan = append(laporan, nominatif)
	}

	if len(laporan) == 0 {
		return laporan, err.NoRecord
	} else {
		return
	}
}

func (t *TabtransMysqlImpl) GetLaporanTransaksi(payload web.DaftarTransaksiRequest, limitOffset web.LimitOffsetLkmUri) (laporan []web.DaftarTransaksiResponse, er error) {

	var tabtrans web.DaftarTransaksiResponse

	if payload.JenisTransaksi == "ALL" {
		args := []interface{}{}
		sqlStmt := `SELECT 
		tabung.no_rekening AS 'No.Rek',
		nama_nasabah AS 'Nama Nasabah',
		DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS 'Tgl Trans',
		kuitansi AS 'No Bukti',
		kode_trans AS 'Kode Trans',
		tabtrans.keterangan AS 'Deskripsi Transaksi',
		IF(my_kode_trans=100,pokok,0) AS Setoran,
		IF(my_kode_trans=200,pokok,0) AS Penarikan
		FROM tabung, tabtrans, nasabah
		WHERE `
		sqlCond := `tabung.no_rekening=tabtrans.no_rekening
		AND tabung.nasabah_id=nasabah.nasabah_id 
		AND tgl_trans>= ? 
		AND tgl_trans<= ? LIMIT ? OFFSET ?`

		if limitOffset.Limit > 0 {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, -1, limitOffset.Offset)
		}
		rows, er := t.apexDb.Query(sqlStmt+sqlCond+``, args...)
		if er != nil {
			return laporan, er
		}

		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {

			if er = rows.Scan(
				&tabtrans.NoRekening,
				&tabtrans.NamaLembaga,
				&tabtrans.TglTrans,
				&tabtrans.NoBukti,
				&tabtrans.KodeTrans,
				&tabtrans.DeskripsiTransaksi,
				&tabtrans.Setoran,
				&tabtrans.Penarikan,
			); er != nil {
				return laporan, er
			}

			laporan = append(laporan, tabtrans)
		}

	} else if payload.JenisTransaksi == "DEPOSIT" {
		args := []interface{}{}
		sqlStmt := `SELECT 
		tabung.no_rekening AS 'No.Rek',
		nama_nasabah AS 'Nama Nasabah',
		DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS 'Tgl Trans',
		kuitansi AS 'No Bukti',
		kode_trans AS 'Kode Trans',
		tabtrans.keterangan AS 'Deskripsi Transaksi',
		IF(my_kode_trans=100,pokok,0) AS Setoran,
		IF(my_kode_trans=200,pokok,0) AS Penarikan
		FROM tabung, tabtrans, nasabah
		WHERE `
		sqlCond := `tabung.no_rekening=tabtrans.no_rekening
		AND tabung.nasabah_id=nasabah.nasabah_id
		AND tgl_trans>= ?
		AND tgl_trans<= ?
		AND(
			kode_trans="100" OR 
			kode_trans="102" OR 
			kode_trans="111" OR 
			kode_trans="115" OR 
			kode_trans="200" OR 
			kode_trans="202" OR 
			kode_trans="211" OR 
			kode_trans="215" OR 
			kode_trans="250"
		) LIMIT ? OFFSET ?`

		if limitOffset.Limit > 0 {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, -1, limitOffset.Offset)
		}
		rows, er := t.apexDb.Query(sqlStmt+sqlCond+``, args...)
		if er != nil {
			return laporan, er
		}

		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {
			if er = rows.Scan(
				&tabtrans.NoRekening,
				&tabtrans.NamaLembaga,
				&tabtrans.TglTrans,
				&tabtrans.NoBukti,
				&tabtrans.KodeTrans,
				&tabtrans.DeskripsiTransaksi,
				&tabtrans.Setoran,
				&tabtrans.Penarikan,
			); er != nil {
				return laporan, er
			}
			laporan = append(laporan, tabtrans)
		}
	} else {
		args := []interface{}{}
		sqlStmt := `SELECT 
		tabung.no_rekening AS 'No.Rek',
		nama_nasabah AS 'Nama Nasabah',
		DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS 'Tgl Trans',
		kuitansi AS 'No Bukti',
		kode_trans AS 'Kode Trans',
		tabtrans.keterangan AS 'Deskripsi Transaksi',
		IF(my_kode_trans=100,pokok,0) AS Setoran,
		IF(my_kode_trans=200,pokok,0) AS Penarikan
		FROM tabung, tabtrans, nasabah
		WHERE `
		sqlCond := `tabung.no_rekening=tabtrans.no_rekening
		AND tabung.nasabah_id=nasabah.nasabah_id
		AND (kode_trans="150")
		AND tgl_trans>= ?
		AND tgl_trans<= ? LIMIT ? OFFSET ?`

		if limitOffset.Limit > 0 {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			args = append(args, payload.PeriodeAwal, payload.PeriodeAkhir, -1, limitOffset.Offset)
		}
		rows, er := t.apexDb.Query(sqlStmt+sqlCond+``, args...)
		if er != nil {
			return laporan, er
		}

		defer func() {
			_ = rows.Close()
		}()

		for rows.Next() {
			if er = rows.Scan(
				&tabtrans.NoRekening,
				&tabtrans.NamaLembaga,
				&tabtrans.TglTrans,
				&tabtrans.NoBukti,
				&tabtrans.KodeTrans,
				&tabtrans.DeskripsiTransaksi,
				&tabtrans.Setoran,
				&tabtrans.Penarikan,
			); er != nil {
				return laporan, er
			}

			laporan = append(laporan, tabtrans)
		}
	}

	if len(laporan) == 0 {
		return laporan, err.NoRecord
	} else {
		return
	}
}

func (t *TabtransMysqlImpl) GetListsTransaksiDeposit(payload web.GetListsDepositTrxReq) (list []web.GetListsDepositTrxRes, er error) {
	rows, er := t.apexDb.Query(`SELECT 
	tabtrans_id, 
	DATE_FORMAT(tgl_trans, "%d/%m/%Y") AS 'Tgl Trans', 
	tabung.no_rekening, 
	nasabah.nama_nasabah, 
	IF(kode_trans="100" OR kode_trans="200",pokok,0) AS tunai, 
	IF(kode_trans="115" OR kode_trans="215",pokok,0) AS transfer, 
	IF(kode_trans="111" OR kode_trans="211",pokok,0) AS piutang, 
	IF(kode_trans="150" OR kode_trans="250",pokok,0) AS branchless, 
	pokok, 
	kuitansi,
	tabtrans.keterangan,	
	kode_trans, 
	IF(my_kode_trans=100,"K","D") AS my_kode_trans, 
	tabtrans.userid
	   FROM tabtrans, tabung, nasabah
	 WHERE 
	   tabtrans.no_rekening=tabung.no_rekening 
	   AND tabung.nasabah_id=nasabah.nasabah_id 
	   AND 
	(
		kode_trans="100" 
		OR kode_trans="102" 
		OR kode_trans="111" 
		OR kode_trans="115" 
		OR kode_trans="150" 
	   	OR kode_trans="200" 
		OR kode_trans="202" 
		OR kode_trans="211" 
		OR kode_trans="215" 
		OR kode_trans="250"
	)
	AND tgl_trans >= ? 
	AND tgl_trans <= ? 
	`, payload.TanggalAwal, payload.TanggalAkhir)
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var laporanDeposit web.GetListsDepositTrxRes
		if er = rows.Scan(
			&laporanDeposit.TransID,
			&laporanDeposit.TglTrans,
			&laporanDeposit.KodeLKM,
			&laporanDeposit.NamaLembaga,
			&laporanDeposit.Tunai,
			&laporanDeposit.Transfer,
			&laporanDeposit.Piutang,
			&laporanDeposit.ViaBrhancless,
			&laporanDeposit.Total,
			&laporanDeposit.NoBukti,
			&laporanDeposit.Keterangan,
			&laporanDeposit.KodeTrans,
			&laporanDeposit.DK,
			&laporanDeposit.UserID,
		); er != nil {
			return list, er
		}

		list = append(list, laporanDeposit)
	}

	if len(list) == 0 {
		return list, err.NoRecord
	} else {
		return
	}
}
