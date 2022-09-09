package constant

import (
	"apex-ems-integration-clean-arch/myutils"
)

// Convert Mysql Data Type
var (
	SQLVendor  myutils.FieldString
	SQLAlamat  myutils.FieldString
	SQLKontak  myutils.FieldString
	SQLPlafond myutils.FieldFloat
)

// func ConvertSQLDataType() {
// 	r := web.GetDetailLKMInfo{}
// 	r.Vendor = SQLVendor.String
// 	r.Alamat = SQLAlamat.String
// 	r.Kontak = SQLKontak.String
// 	r.Plafond = SQLPlafond.Float64
// }
