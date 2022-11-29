package tabtransrepo

import (
	"apex-ems-integration-clean-arch/entities/err"
	"apex-ems-integration-clean-arch/entities/web"
	"database/sql"
)

func newTabtransMysqlImpl(apexConn *sql.DB) TabtransRepo {
	return &TabtransMysqlImpl{
		apexDb: apexConn,
	}
}

type TabtransMysqlImpl struct {
	apexDb *sql.DB
}

func (t *TabtransMysqlImpl) GetListTabtransInfo(payload web.GetListTabtransByDate, limitOffset web.LimitOffsetLkmUri) (
	list []web.GetListTabtransInfo,
	total web.GetCountWithSumTabtransTrx,
	er error,
) {
	var rows *sql.Rows

	/*
		Dibawah adalah interface dengan variadic untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar bisa re-usable pada 2 kondisi
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
		var tabtransListTx web.GetListTabtransInfo
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
