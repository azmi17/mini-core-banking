package routingindukrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newRoutingIndukMysqlImpl(apexConn *sql.DB) RoutingIndukRepo {
	return &RoutingIndukMysqlImpl{
		apexDb: apexConn,
	}
}

type RoutingIndukMysqlImpl struct {
	apexDb *sql.DB
}

func (r *RoutingIndukMysqlImpl) GetRoutingRekInduk(kodeLkm string) (routing entities.RoutingRekIndukData, er error) {
	row := r.apexDb.QueryRow(`SELECT 
		norek,
		norek_induk
	FROM routing_rek_induk 
	WHERE norek = ? LIMIT 1`, kodeLkm)
	er = row.Scan(
		&routing.KodeLkm,
		&routing.NorekInduk,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return routing, err.NoRecord
		} else {
			return routing, errors.New(fmt.Sprint("error while get routing induk data: ", er.Error()))
		}
	}
	return
}

func (r *RoutingIndukMysqlImpl) GetListSysApexRoutingRekInduk(payload entities.GlobalFilter, limitOffset entities.LimitOffsetLkmUri) (list []entities.RoutingRekIndukData, er error) {
	var rows *sql.Rows

	args := []interface{}{}
	sqlCond := ""
	sqlStmt := `SELECT 
		norek, 
		norek_induk 
	FROM 
	routing_rek_induk `

	if payload.Filter == "" {
		if limitOffset.Limit > 0 {
			sqlCond = "LIMIT ? OFFSET ?"
			args = append(args, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = "LIMIT ? OFFSET ?"
			args = append(args, -1, limitOffset.Offset)
		}
		rows, er = r.apexDb.Query(sqlStmt+sqlCond+``, args...)
	} else {
		if limitOffset.Limit > 0 {
			sqlCond = `
			WHERE
			(norek LIKE "%` + payload.Filter + `%" OR norek_induk LIKE "%` + payload.Filter + `%") 
			LIMIT ? OFFSET ?`
			args = append(args, limitOffset.Limit, limitOffset.Offset)
		} else {
			sqlCond = `
			WHERE
			(norek LIKE "%` + payload.Filter + `%" OR norek_induk LIKE "%` + payload.Filter + `%") 
			LIMIT ? OFFSET ?`
			args = append(args, -1, limitOffset.Offset)
		}
		rows, er = r.apexDb.Query(sqlStmt+sqlCond+``, args...)
	}
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var routings entities.RoutingRekIndukData
		if er = rows.Scan(&routings.KodeLkm, &routings.NorekInduk); er != nil {
			return list, er
		}

		list = append(list, routings)
	}

	if len(list) == 0 {
		return list, err.NoRecord
	} else {
		return
	}
}

func (r *RoutingIndukMysqlImpl) CreateSysApexRoutingRekInduk(bankCode, norekInduk string) (routing entities.RoutingRekIndukData, er error) {

	stmt, er := r.apexDb.Prepare(`INSERT INTO routing_rek_induk(
		norek,
		norek_induk
	) VALUES(?,?)`)
	if er != nil {
		return routing, errors.New(fmt.Sprint("error while prepare add routing rek induk: ", er.Error()))
	}
	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(
		bankCode,
		norekInduk,
	); er != nil {
		return routing, errors.New(fmt.Sprint("error while add routing rek induk: ", er.Error()))
	} else {
		routing.KodeLkm = bankCode
		routing.NorekInduk = norekInduk
		return routing, nil
	}
}

func (r *RoutingIndukMysqlImpl) UpdateSysApexRoutingRekInduk(newBankCode, norekInduk, currentBankCode string) (routing entities.RoutingRekIndukData, er error) {

	thisRepo, _ := NewRoutingIndukRepo()
	_, er = thisRepo.GetRoutingRekInduk(currentBankCode)
	if er != nil {
		return routing, err.NoRecord
	}

	stmt, er := r.apexDb.Prepare("UPDATE routing_rek_induk SET norek = ?, norek_induk = ? WHERE norek = ?")
	if er != nil {
		return routing, errors.New(fmt.Sprint("error while prepare update routing rek induk: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(newBankCode, norekInduk, currentBankCode); er != nil {
		return routing, errors.New(fmt.Sprint("error while update routing rek induk: ", er.Error()))
	}

	routing.KodeLkm = newBankCode
	routing.NorekInduk = norekInduk
	return routing, nil
}

func (r *RoutingIndukMysqlImpl) DeleteSysApexRoutingRekInduk(kodeLkm ...string) (er error) {

	// thisRepo, _ := NewRoutingIndukRepo()
	// _, er = thisRepo.GetRoutingRekInduk(kodeLkm)
	// if er != nil {
	// 	return err.NoRecord
	// }

	stmt, er := r.apexDb.Prepare("DELETE FROM routing_rek_induk WHERE norek = ?")
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare delete routing rek induk: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	for _, v := range kodeLkm {
		if _, er := stmt.Exec(v); er != nil {
			return errors.New(fmt.Sprint("error while delete routing rek induk: ", er.Error()))
		}
	}

	return nil
}
