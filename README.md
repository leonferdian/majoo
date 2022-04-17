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
