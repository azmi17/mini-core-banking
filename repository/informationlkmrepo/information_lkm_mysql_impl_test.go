package informationlkmrepo

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

func TestRekeningKoranList(t *testing.T) {
	db := GetConnectionApx()
	lkmInfo := newInformationLKMMysqlImpl(db)

	data, err := lkmInfo.RekeningKoranLKMDetailHeader("0102")
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println("Result: ", data)
}
