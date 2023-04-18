package tabunganrepo

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/kpango/glg"
)

func GetConnectionApx() *sql.DB {
	dataSource := "root:azmic0ps@tcp(localhost:3317)/integrasi_apex_ems?parseTime=true"
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}

func TestGetRekeningLKMByStatusActive(t *testing.T) {
	db := GetConnectionApx()
	tabungRepo := newTabunganMysqlImpl(db)

	lkm, err := tabungRepo.GetRekeningLKMByStatusActive()
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println(len(lkm))

}
