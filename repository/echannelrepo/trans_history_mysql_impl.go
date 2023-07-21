package echannelrepo

import (
	"database/sql"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newEchannelMysqlImpl(apexConn *sql.DB) EchannelRepo {
	return &echannelMysqlImpl{
		apexDb: apexConn,
	}
}

type echannelMysqlImpl struct {
	apexDb *sql.DB
}

func (e *echannelMysqlImpl) GetEchannelTransHistories(payload entities.TransHistoryRequest, limitOffset entities.LimitOffsetLkmUri) (list []entities.TransHistoryResponse, er error) {
	var rows *sql.Rows

	/*
		Dibawah adalah slice of interface untuk kebutuhan custom query,
		di deklarasikan dengan scope Global Variable agar re-usable pada 2 kondisi
	*/
	args := []interface{}{}
	sqlCond := ""
	sqlStmt := `
	SELECT 
		th.trans_id,
		DATE_FORMAT(th.tgl_trans_str, "%d/%m/%Y") AS tgl_trans,
		th.bank_code,
		b.bank_name,
		th.subscriber_id,
		th.amount,
		th.dc,
		th.response_code,
		th.stan,
		th.ref,
		th.rek_id,
		th.biller_code,
		th.processing_code,
		th.product_code
	FROM trans_history AS th 
	LEFT JOIN bank b ON(th.bank_code=b.bank_code) 
	WHERE `

	if payload.Filter == "" {
		if limitOffset.Limit > 0 {
			sqlCond = "th.tgl_trans_str >= ? AND th.tgl_trans_str <= ? AND th.dc <> '' LIMIT ? OFFSET ?"
			args = append(args, payload.TglAwal, payload.TglAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "th.tgl_trans_str >= ? AND th.tgl_trans_str <= ? AND th.dc <> '' LIMIT ? OFFSET ?"
			args = append(args, payload.TglAwal, payload.TglAkhir, -1, limitOffset.Offset)
		}
		rows, er = e.apexDb.Query(sqlStmt+sqlCond+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			sqlCond = `
			th.tgl_trans_str >= ? 
			AND th.tgl_trans_str <= ? 
			AND (th.bank_code LIKE "%` + payload.Filter + `%" OR th.ref LIKE "%` + payload.Filter + `%" OR th.subscriber_id LIKE "%` + payload.Filter + `%") 
			AND th.dc <> '' 
			LIMIT ? OFFSET ?`
			args = append(args, payload.TglAwal, payload.TglAkhir, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = `
			th.tgl_trans_str >= ? 
			AND th.tgl_trans_str <= ? 
			AND (th.bank_code LIKE "%` + payload.Filter + `%" OR th.ref LIKE "%` + payload.Filter + `%" OR th.subscriber_id LIKE "%` + payload.Filter + `%") 
			AND th.dc <> '' 
			LIMIT ? OFFSET ?`
			args = append(args, payload.TglAwal, payload.TglAkhir, -1, limitOffset.Offset)
		}
		rows, er = e.apexDb.Query(sqlStmt+sqlCond+``, args...)
	}

	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var transHistories entities.TransHistoryResponse
		if er = rows.Scan(
			&transHistories.TransID,
			&transHistories.TglTrans,
			&transHistories.KodeLKM,
			&transHistories.NamaLembaga,
			&transHistories.SubscriberId,
			&transHistories.Amount,
			&transHistories.Dc,
			&transHistories.ResponseCode,
			&transHistories.Stan,
			&transHistories.Ref,
			&transHistories.RekeningID,
			&transHistories.BillerCode,
			&transHistories.ProcessingCode,
			&transHistories.ProductCode,
		); er != nil {
			return list, er
		}
		list = append(list, transHistories)
	}

	if len(list) == 0 {
		return list, err.NoRecord
	}

	return list, nil
}
