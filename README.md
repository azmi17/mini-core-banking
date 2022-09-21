# API Specification
Document Version : 1.0.0
Last Update : Wednesday, 21-09-2022 

##  LIST API
- APEX
    - [Create](#create-LKM)
    - [Update](#update-LKM)
    - [Delete](#delete-LKM)
    - [List](#list-all-LKM)
    - [Get LKM Info By lkmCode](#get-lkm-by-lkmCode)
    - [List](#list-all-SCGroup)
    - [Update](#reset-password-LKM)

## Create LKM
Request :
- Method           : POST
- Endpoint         : `/v1/api/institutions`
- Header           : 
    - Content-Type : multipart/form-data
    - Accept : application/json
- Body             :
key:value
[kode_lkm]      = [string]
[nama_lembaga]  = [string]
[alamat]        = [string]
[telepon]       = [string]
[user_id]       = [integer]

Response :
```json
{
  "code" : "number",
  "data" : {
                "kode_lkm" : "string, unique",
                "nama_lembaga" : "string",
                "alamat" : "string",
                "Telepon" : "string",
                "no_rekening" : "string",
                "saldo_akhir" : "int",
                "user_name_smec" : "string",
                "password_smec" : "string",
                "user_id" : "int"
           }
}
```

## Update LKM
Request :
- Method : PUT
- Endpoint : `/v1/api/institutions`
- Header : 
    - Content-Type : multipart/form-data
    - Accept : application/json
- Body :
key:value
[kode_lkm]      = [string]
[nama_lembaga]  = [string]
[alamat]        = [string]
[telepon]       = [string]
[user_id]       = [integer]

Response :
```json
{
  "code" : "number",
  "data" : {
                "kode_lkm" : "string, unique",
                "nama_lembaga" : "string",
                "alamat" : "string",
                "Telepon" : "string",
                "user_id" : "int"
           }
}
```

## Delete LKM
Request :
- Method : DELETE
- Endpoint : `/v1/api/institutions`
- Header : 
    - Content-Type : multipart/form-data
    - Accept : application/json  
- Body :
key:value
[kode_lkm]      = [string]

Response :
```json
{
  "code" : "number",
  "data" : null
}
```

## List All LKM
Request :
- Method : GET
- Endpoint : `/v1/api/institutions/all/{limit}/{offset}`
- Header : 
    - Accept : application/json
Response :
```json
{
  "code" : "number",
  "data" : [
              {
                "kode_lkm" : "integer, unique",
                "nama_lembaga" : "string",
                "vendor" : "string",
                "alamat" : "integer, unique",
                "kontak" : "string",
                "apex_norek" : "string",
                "saldo_akhir" : "integer",
                "plafond" : "integer",
                "status_tab" : "string"
              },
              {
                "kode_lkm" : "integer, unique",
                "nama_lembaga" : "string",
                "vendor" : "string",
                "alamat" : "integer, unique",
                "kontak" : "string",
                "apex_norek" : "string",
                "saldo_akhir" : "integer",
                "plafond" : "integer",
                "status_tab" : "string"
              }
          ]
}
```

## Get LKM Info By lkmCode
Request :
- Method : GET
- Endpoint : `/v1/api/institutions/{kode_lkm}`
- Header : 
    - Accept : application/json

Response :
```json
{
  "code" : "number",
  "data" : {
                "kode_lkm" : "integer, unique",
                "nama_lembaga" : "string",
                "vendor" : "string",
                "alamat" : "integer, unique",
                "kontak" : "string",
                "apex_norek" : "string",
                "saldo_akhir" : "integer",
                "plafond" : "integer",
                "status_tab" : "string"
            }
}
```

## List All SCGroup
Request :
- Method : GET
- Endpoint : `/v1/api/vendors`
- Header : 
    - Accept : application/json
Response :
```json
{
  "code" : "number",
  "data" : [
              {
                "kode_group" : "string",
                "deskripsi_group" : "string"
              },
              {
                "kode_group" : "string",
                "deskripsi_group" : "string"
              }
          ]
}
```

## Reset Password LKM
Request :
- Method : PUT
- Endpoint : `/v1/api/user/reset-password`
- Header : 
    - Content-Type : multipart/form-data
    - Accept : application/json
- Body :
key:value
[kode_lkm]      = [string]

Response :
```json
{
  "code" : "number",
  "data" : {
                "kode_lkm" : "string, unique",
                "password_smec" : "string"
           }
}
```