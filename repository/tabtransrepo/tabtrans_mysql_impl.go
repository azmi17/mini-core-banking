package tabtransrepo

import (
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"apex-ems-integration-clean-arch/helper"
	"database/sql"
	"errors"
	"fmt"
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

func (t *TabtransMysqlImpl) GetListsTabtransTrx(payload web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
	list []web.GetListTabtransTrx,
	total web.GetCountWithSumTabtransTrx,
	er error,
) {
	var rows *sql.Rows

	/*
		Dibawah adalah slice of interface untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar re-usable pada 2 kondisi
	*/
	args := []interface{}{}
	q := ""

	if payload.BankCode == "" {
		if limitOffset.Limit > 0 {
			q = "t.tgl_trans >= ? AND t.tgl_trans <= ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			q = "t.tgl_trans >= ? AND t.tgl_trans <= ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, -1, limitOffset.Offset)
		}
		rows, er = t.apexDb.Query(`SELECT 
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
		LEFT JOIN nasabah AS n ON (tb.nasabah_id=n.nasabah_id) WHERE `+q+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			q = "t.tgl_trans >= ? AND t.tgl_trans <= ? AND t.no_rekening = ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, payload.BankCode, limitOffset.Limit, limitOffset.Offset)
		} else {
			q = "t.tgl_trans >= ? AND t.tgl_trans <= ? AND t.no_rekening = ? LIMIT ? OFFSET ?"
			args = append(args, payload.TanggalAwal, payload.TanggalAkhir, payload.BankCode, -1, limitOffset.Offset)
		}
		rows, er = t.apexDb.Query(`SELECT 
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
		LEFT JOIN nasabah AS n ON (tb.nasabah_id=n.nasabah_id) WHERE `+q+``, args...)
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

func (t *TabtransMysqlImpl) CountSaldoAkhirOnNoRekening(kodeLKM string) (data web.CalculateRepostingResult, er error) {
	var tabtrans web.RepostingData
	row := t.apexDb.QueryRow(`SELECT 
	  no_rekening,
	  SUM(CASE WHEN my_kode_trans='100' THEN pokok ELSE 0 END) AS total_kredit,
	  SUM(CASE WHEN my_kode_trans='200' THEN pokok ELSE 0 END) AS total_debet
	FROM tabtrans
	 WHERE
	no_rekening = ? GROUP BY no_rekening
	`, kodeLKM)
	er = row.Scan(
		&tabtrans.KodeLKM,
		&tabtrans.TotalKredit,
		&tabtrans.TotalDebet,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return data, err.NoRecord
		} else {
			return data, errors.New(fmt.Sprint("error while get reposting saldo akhir: ", er.Error()))
		}

	}
	data.KodeLKM = tabtrans.KodeLKM
	data.SaldoAkhir = tabtrans.TotalKredit - tabtrans.TotalDebet

	return data, nil
}

func (t *TabtransMysqlImpl) GetTotalTrxWithTotalPokok(TglTrans web.GetListTabtransByDate) (total web.GetCountWithSumTabtransTrx, er error) {

	rows, err := t.apexDb.Query(`SELECT 
		COUNT(tabtrans_id) AS total_trx,
		SUM(pokok) AS total_pokok 
	FROM tabtrans
	 WHERE
	  tgl_trans >= ?
	 AND tgl_trans <= ?
	`, TglTrans.TanggalAwal, TglTrans.TanggalAkhir)
	if err != nil {
		return total, er
	} else {
		for rows.Next() {
			rows.Scan(
				&total.TotalTrx,
				&total.TotalPokok,
			)
		}
		return total, nil
	}
}
