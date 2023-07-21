package constant

import (
	"new-apex-api/myutils"
)

// Convert Mysql Data Type
var (
	SQLVendor         myutils.FieldString
	SQLTglExpired     myutils.FieldString
	SQLJabatan        myutils.FieldString
	SQLUnitKerja      myutils.FieldString
	SQLUSerName       myutils.FieldString
	SQLNamaLengkap    myutils.FieldString
	SQLAlamat         myutils.FieldString
	SQLKontak         myutils.FieldString
	SQLPlafond        myutils.FieldFloat
	SQLSetoranMinimum myutils.FieldFloat
)

// func ConvertSQLDataType() {
// 	r := entities.GetDetailLKMInfo{}
// 	r.Vendor = SQLVendor.String
// 	r.Alamat = SQLAlamat.String
// 	r.Kontak = SQLKontak.String
// 	r.Plafond = SQLPlafond.Float64
// }
