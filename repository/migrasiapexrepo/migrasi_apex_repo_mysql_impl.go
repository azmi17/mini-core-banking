package migrasiapexrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newMigrasiApexMysqlImpl(apexConn *sql.DB) MigrasiApexRepo {
	return &MigrasiApexMysqlImpl{
		apexDb: apexConn,
	}
}

type MigrasiApexMysqlImpl struct {
	apexDb *sql.DB
}

func (t *MigrasiApexMysqlImpl) NorekLengthEqual4() (data []entities.NorekWithNID, er error) {
	rows, er := t.apexDb.Query("SELECT no_rekening, nasabah_id FROM tabung WHERE LENGTH(no_rekening)=4")
	if er != nil {
		return data, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var list entities.NorekWithNID
		if er = rows.Scan(&list.NoRekening, &list.NasabahID); er != nil {
			return data, er
		}

		data = append(data, list)
	}

	if len(data) == 0 {
		return data, err.NoRecord
	} else {
		return
	}
}

func (t *MigrasiApexMysqlImpl) UpdateNasabahIDWithNorekOnNasabah() (er error) {

	nID, er := t.NorekLengthEqual4()
	if er != nil {
		return er
	}

	for _, v := range nID {
		stmt1, er := t.apexDb.Prepare(`UPDATE nasabah SET nasabah_id = ? WHERE nasabah_id = ?`)
		if er != nil {
			return errors.New(fmt.Sprint("error while prepare update nasabah_id on nasabah: ", er.Error()))
		}

		defer func() {
			_ = stmt1.Close()
		}()

		if _, er = stmt1.Exec(
			v.NoRekening,
			v.NasabahID,
		); er != nil {
			return errors.New(fmt.Sprint("error while update nasabah_id on nasabah: ", er.Error()))
		}

		stmt2, er := t.apexDb.Prepare(`UPDATE tabung SET nasabah_id = ? WHERE no_rekening = ?`)
		if er != nil {
			return errors.New(fmt.Sprint("error while prepare update nasabah_id on tabung: ", er.Error()))
		}

		defer func() {
			_ = stmt2.Close()
		}()

		if _, er = stmt2.Exec(
			v.NoRekening,
			v.NoRekening,
		); er != nil {
			return errors.New(fmt.Sprint("error while update nasabah_id on tabung: ", er.Error()))
		}
	}

	return nil
}
