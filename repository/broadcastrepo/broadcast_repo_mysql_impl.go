package broadcastrepo

import (
	"database/sql"
	"errors"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/entities/err"
)

func newBroadcastRepoMysqlImpl(apexConn *sql.DB) BroadcastRepo {
	return &BroadCastRepoMysqlImpl{
		apexDb: apexConn,
	}
}

type BroadCastRepoMysqlImpl struct {
	apexDb *sql.DB
}

func (s *BroadCastRepoMysqlImpl) GetReceiverID(userID int) (data entities.Broadcast, er error) {
	row := s.apexDb.QueryRow(`SELECT
		id,
		user_id,
		name,
		receiver_id
	FROM broadcast_lists WHERE user_id = ?`, userID)
	er = row.Scan(
		&data.Id,
		&data.UserID,
		&data.Name,
		&data.ReceiverID,
	)
	if er != nil {
		if er == sql.ErrNoRows {
			return data, err.NoRecord
		} else {
			return data, errors.New(fmt.Sprint("error while get receiver id: ", er.Error()))
		}
	}
	return
}
