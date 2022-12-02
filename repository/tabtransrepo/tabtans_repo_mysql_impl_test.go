package tabtransrepo

import (
	"database/sql"
	"time"
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

// func TestCalculateReposingResult(t *testing.T) {
// 	db := GetConnectionApx()
// 	tabtransRepo := newTabtransMysqlImpl(db)

// 	lkm1, err := tabtransRepo.CountSaldoAkhirOnNoRekening("0102")
// 	if err != nil {
// 		_ = glg.Log(err.Error())
// 	}
// 	fmt.Println(lkm1)

// 	lkm2, err := tabtransRepo.CountSaldoAkhirOnNoRekening("0095")
// 	if err != nil {
// 		_ = glg.Log(err.Error())
// 	}
// 	fmt.Println(lkm2)

// 	err = tabtransRepo.RepostingOnLkmAccount(lkm1, lkm2)
// 	if err != nil {
// 		_ = glg.Log(err.Error())
// 	}
// 	fmt.Println("Reposting saldo succeeded..")

// }
