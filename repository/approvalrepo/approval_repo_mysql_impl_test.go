package approvalrepo

import (
	"database/sql"
	"fmt"
	"new-apex-api/entities"
	"new-apex-api/helper"
	"testing"
	"time"

	"github.com/kpango/glg"
)

func GetApexConn() *sql.DB {
	dataSource := "root:azmic0ps@tcp(localhost:3317)/apex20230409?parseTime=true"
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

func TestCalculateReposingResult(t *testing.T) {
	db := GetApexConn()
	approvalRepo := newApprovalRepoImpl(db)

	appv := entities.Approval{}
	appv.UserID = 1024037
	appv.OtorisatorID = 2024
	appv.Token = helper.String(8)
	appv.Status = 1
	appv.Description = "Request: Delete tabtrans Data"
	// currentTime := time.Now()
	// appv.Time = currentTime
	// appv.Expired = currentTime.Add(time.Minute * 5)

	data, err := approvalRepo.CreateNewApproval(appv)
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println(data)

}

func TestTimeNow(t *testing.T) {
	var time1 = time.Now()
	fmt.Printf("time1 %v\n", time1)

	var time2 = time.Date(2011, 12, 24, 10, 20, 0, 0, time.UTC)
	fmt.Printf("time2 %v\n", time2)
	// time2 2011-12-24 10:20:00 +0000 UTC
}

func TestRandomString(t *testing.T) {
	for i := 1; i <= 100; i++ {
		fmt.Println(helper.String(6))
	}
}
