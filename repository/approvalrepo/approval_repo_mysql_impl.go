package approvalrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newApprovalRepoImpl(apexConn *sql.DB) ApprovalRepo {
	return &ApprovalMysqlImpl{
		apexDb: apexConn,
	}
}

type ApprovalMysqlImpl struct {
	apexDb *sql.DB
}

func (a *ApprovalMysqlImpl) CreateNewApproval(payload entities.Approval) (approval entities.Approval, er error) {
	stmt, er := a.apexDb.Prepare(`INSERT INTO approval_lists(
		user_id, 
		otorisator_id, 
		token, 
		status, 
		description, 
		time, 
		expired 
	) VALUES(?,?,?,?,?,?,?)`)
	if er != nil {
		return approval, errors.New(fmt.Sprint("error while prepare create approval request : ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if stmt, er := stmt.Exec(
		payload.UserID,
		payload.OtorisatorID,
		payload.Token,
		payload.Status,
		payload.Description,
		payload.Time,
		payload.Expired,
	); er != nil {
		return approval, errors.New(fmt.Sprint("error while create approval request : ", er.Error()))
	} else {

		// Get Last insert ID
		lastId, txErr := stmt.LastInsertId()
		if txErr != nil {
			return approval, errors.New(fmt.Sprint("error while get last insert id : ", txErr.Error()))
		}

		approval = payload
		approval.Id = int(lastId)

		return approval, nil
	}

}

func (a *ApprovalMysqlImpl) GetApproval(token string) (data entities.ApprovalResponse, er error) {
	row := a.apexDb.QueryRow(`SELECT
		al.id,
		al.user_id,
		app.name AS 'user',
		al.otorisator_id,
		al.token,
		al.status,
		al.description,
		DATE_FORMAT(al.time, "%Y-%m-%d %H:%i:%S"),
		DATE_FORMAT(al.expired, "%Y-%m-%d %H:%i:%S")
	FROM approval_lists AS al 
	INNER JOIN broadcast_lists AS app ON (al.user_id=app.user_id) WHERE al.token = ?`, token)
	er = row.Scan(
		&data.Id,
		&data.UserID,
		&data.UserName,
		&data.OtorisatorID,
		&data.Token,
		&data.Status,
		&data.Description,
		&data.Time,
		&data.Expired,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return data, err.NoRecord
		} else {
			return data, errors.New(fmt.Sprint("error while get approval data: ", er.Error()))
		}
	}
	return
}

func (a *ApprovalMysqlImpl) GetListsApproval(limitOffset entities.LimitOffsetLkmUri) (list []entities.ApprovalResponse, er error) {

	args := []interface{}{}
	limit := ""
	if limitOffset.Limit > 0 {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, limitOffset.Limit, limitOffset.Offset)
	} else {
		limit = "LIMIT ? OFFSET ?"
		args = append(args, -1, limitOffset.Offset)
	}
	rows, er := a.apexDb.Query(`SELECT
		al.id,
		al.user_id,
		app.name AS 'user',
		al.otorisator_id,
		al.token,
		al.status,
		al.description,
		DATE_FORMAT(al.time, "%d-%m-%Y %H:%i:%S"),
		DATE_FORMAT(al.expired, "%d-%m-%Y %H:%i:%S")
	FROM approval_lists AS al 
	INNER JOIN broadcast_lists AS app ON (al.user_id=app.user_id) `+limit+``, args...)
	if er != nil {
		return list, er
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		var data entities.ApprovalResponse
		if er = rows.Scan(
			&data.Id,
			&data.UserID,
			&data.UserName,
			&data.OtorisatorID,
			&data.Token,
			&data.Status,
			&data.Description,
			&data.Time,
			&data.Expired,
		); er != nil {
			return list, er
		}
		list = append(list, data)
	}

	if len(list) == 0 {
		return list, nil
	} else {
		return
	}
}

func (a *ApprovalMysqlImpl) UpdateStatusApproval(status int, token string) (er error) {

	stmt, er := a.apexDb.Prepare(`UPDATE approval_lists SET status = ? WHERE token = ?`)
	if er != nil {
		return errors.New(fmt.Sprint("error while prepare update approval status: ", er.Error()))
	}

	defer func() {
		_ = stmt.Close()
	}()

	if _, er := stmt.Exec(status, token); er != nil {
		return errors.New(fmt.Sprint("error while update approval status: ", er.Error()))
	}

	return nil
}
