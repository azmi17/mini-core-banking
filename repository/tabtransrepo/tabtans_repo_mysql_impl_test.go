package tabtransrepo

import (
	"database/sql"
	"fmt"
	"new-apex-api/entities/web"
	"new-apex-api/helper"
	"strings"
	"testing"
	"time"

	"github.com/kpango/glg"
)

func GetConnectionApx() *sql.DB {
	dataSource := "root:azmic0ps@tcp(localhost:3317)/apex_backup?parseTime=true"
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
	db := GetConnectionApx()
	tabtransRepo := newTabtransMysqlImpl(db)

	lkm, err := tabtransRepo.CalculateSaldoOnRekeningLKM("0001")
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println(lkm)

}

func TestCalculateGetSaldoAkhir(t *testing.T) {
	db := GetConnectionApx()
	tabtransRepo := newTabtransMysqlImpl(db)

	saldoAwal, err := tabtransRepo.GetSaldo("0102", "2023-01-01")
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println("Result: ", saldoAwal)
}

func TestRekeningKoranList(t *testing.T) {
	// db := GetConnectionApx()
	// tabtransRepo := newTabtransMysqlImpl(db)

	// data, err := tabtransRepo.GetRekeningKoranLKMDetail("0102", "2020-01-01", "2023-01-01")
	// if err != nil {
	// 	_ = glg.Log(err.Error())
	// }
	// fmt.Println("Result: ", data[2])

	input := "2023-03-20" // gimana caranya agar motong tanggal supaya (diabaikan) -> catatan, kalo akses pake index harus pastikan length-nya >= (6)

	date, er := time.Parse("2006-01", input[0:7]) // index mutlak 0-7 (kasih kondisi kalo kurang dari 7)
	if er != nil {
		_ = glg.Log(er.Error())
	}

	YYYYMM := date.AddDate(0, -1, 0)  // Get begin date in last month
	YYYYMM2 := date.AddDate(0, 0, -1) // Get end date in last month

	parseBeginDt := YYYYMM.Format(helper.YYYYMMDD)
	parseEndDt2 := YYYYMM2.Format(helper.YYYYMMDD)
	parseEndTg := YYYYMM2.Day()

	fmt.Println(parseBeginDt)
	fmt.Println(parseEndDt2)
	fmt.Println(parseEndTg)

	// kemarin, _ := helper.ParseTimeStrToDate(input) // dapetin end date
	// fmt.Println(kemarin)

	// timeFormatStr, _ := helper.FormatTimeStrDDMMYYY("2023-03-17")
	// fmt.Println(timeFormatStr)
}

func TestNominatifDeposit(t *testing.T) {
	db := GetConnectionApx()
	tabtransRepo := newTabtransMysqlImpl(db)

	limitoffset := web.LimitOffsetLkmUri{
		Limit:  4,
		Offset: 0,
	}

	data, err := tabtransRepo.GetNominatifDeposit("20230323", "20230201", "20230228", 28, limitoffset)
	if err != nil {
		_ = glg.Log(err.Error())
	}
	fmt.Println("Results: ", len(data))
}

func TestSplitJoinStr(t *testing.T) {
	item := strings.Split("100", ",") // daimbil dari konfigurasi

	beginText := "kode_trx=\""
	endText := " OR"

	for i, v := range item {
		value := v + "\""

		if i < len(item)-1 {
			fmt.Println(beginText + value + endText)
		} else {
			fmt.Println(beginText + value)
		}
	}
}

func TestSplitJoinStr2(t *testing.T) {
	item := strings.Split("100,102,111,115,200,202,211,215,250", ",") // konfgiurasi

	kodeTrx := kodeTrx(item)
	fmt.Println(kodeTrx)
}

func kodeTrx(item []string) string {
	lastIndex := len(item) - 1
	beginText := "kode_trans=\""
	endText := " OR "

	var kodeTrans string
	for i, v := range item {
		text := beginText + v + "\""
		if i < lastIndex {
			text += endText
		}
		kodeTrans += text
	}
	return kodeTrans
}
