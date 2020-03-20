package dao

import (
	"crawlers/crawler-core/data"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
)

type Storage struct {
	Db *sqlx.DB
}

func ConnectMysql() *Storage {
	account := os.Getenv("account")
	passwd := os.Getenv("passwd")
	url := os.Getenv("url")
	db, err := sqlx.Open("mysql", account+":"+passwd+"@tcp("+url+")/house?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}
	return &Storage{Db: db}
}

func (storage *Storage) AddErShouFangRecord(data *data.ErShouFangInfo) {
	_, err := storage.Db.Exec("insert into ershoufang(`title`,`address`,`houseInfo`,`totalPrice`,`unitPrice`) values(?,?,?,?,?)", data.Title, data.Address, data.HouseInfo, data.TotalPrice, data.UnitPrice)
	if err != nil {
		log.Fatal(err)
	}
}

func (storage *Storage) AddErShouFangRecords(datas []*data.ErShouFangInfo) {
	for i := 0; i < len(datas); i++ {
		record := datas[i]
		storage.AddErShouFangRecord(record)
	}
}

func (storage *Storage) AddZuFangRecord(data *data.ZuFangInfo) {
	_, err := storage.Db.Exec("insert into zufang(`title`,`address`,`houseInfo`,`price`,`link`) values(?,?,?,?,?)", data.Title, data.Address, data.HouseInfo, data.Price, data.Link)
	if err != nil {
		log.Fatal(err)
	}
}

func (storage *Storage) AddZuFangRecords(datas []*data.ZuFangInfo) {
	for i := 0; i < len(datas); i++ {
		record := datas[i]
		storage.AddZuFangRecord(record)
	}
}
