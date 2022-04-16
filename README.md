# majoos
1. 
a. login system pada file main.go
b. Implementasi jwt oauth
c. representasi report merchant dan omzet pada function get_mercant_omzet
d. membuat API untuk implementasi report merchant, outlet, dan omzet pada function get_mercant_outlet_omzet
e. menampilkan point c berdasarkan user yang login pada function get_mercant_omzet_by_user
f. menampilkan point d berdasarkan user yang login pada function get_mercant_outlet_omzet_by_user
g. belum, seharusnya transaksi cukup punya relasi dengan outlet
h. 
Query soal c
SELECT a.merchant_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id

menggabungkan antara tabel merchant dengan tabel transaksi dengan relasi id merchant

Query soal d
SELECT a.merchant_name, b.outlet_name, case when c.created_at BETWEEN '2022-11-01' and '2022-11-30' and c.bill_total = 0 then 0 else c.bill_total end as omzet FROM `merchants` a left join outlets b on a.id = b.merchant_id left join transactions c on a.id = c.merchant_id and b.id = c.outlet_id

menggabungkan 3 tabel (merchant, outlet, transaksi) untuk representasi report merchant, outlet, dan omzet

2. Koreksi untuk code pada soal nomer 2

package main

//menambahkan beberapa modul
import (
  "fmt"
  "log"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
)

model pada soal

type (
    Area struct {
    ID int64 `gorm:"column:id;primaryKey;"`
    AreaValue int64 `gorm:"column:area_value"`
    AreaType string `gorm:"column:type"`
    }
)

Implementasi model yang sehrusnya
type Model struct (
    ID int64 `gorm:"column:id;primaryKey;"`
    AreaValue int64 `gorm:"column:area_value"`
    AreaType string `gorm:"column:type"`
)

func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area) (err error) {
        inst := _r.DB.Model(ar)
        Var area int
        area = 0
        switch type {
            case 'persegi panjang':
            var area := param1 * param2 //harusnya tidak pakai var, karena =: sudah mewakili deklarasi
            ar.AreaValue = area
            ar.AreaType = 'persegi panjang'
            err = _r.DB.create(&ar).Error
            if err != nil {
                return err
            }
            case 'persegi':
            var area = param1 * param2
            ar.AreaValue = area
            ar.AreaType = 'persegi'
            err = _r.DB.create(&ar).Error
            if err != nil {
                return err
            }
            case segitiga: //kata segitiga harusnya berupa string yang mana harus disertai tanda petik
            area = 0.5 * (param1 * param2)
            ar.AreaValue = area
            ar.AreaType = 'segitiga'
            err = _r.DB.create(&ar).Error
            if err != nil {
                return err
            }
            default:
            ar.AreaValue = 0
            ar.AreaType = 'undefined data'
            err = _r.DB.create(&ar).Error
            if err != nil {
                return err
            }
        }
    }

func main() {
    err = _u.repository.InsertArea(10, 10, 'persegi')
    if err != nil {
        log.Error().Msg(err.Error())
        err = errors.New(en.ERROR_DATABASE)
        return err
    }
}