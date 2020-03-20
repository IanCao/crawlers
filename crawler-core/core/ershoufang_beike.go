package core

import (
	"crawlers/crawler-core/constant"
	"crawlers/crawler-core/dao"
	"crawlers/crawler-core/data"
	"crawlers/crawler-core/utils"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

func ProcessBeiKeErShouFang() {
	storage := dao.ConnectMysql()
	for i := 1; i <= 3; i++ {
		url := constant.Ershoufang_beike + strconv.Itoa(i)
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("fail %+v", err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatal(res.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		var datas = make([]*data.ErShouFangInfo, 0)
		doc.Find(".sellListPage .content .leftContent").Find("div[data-component=list]").Find(".sellListContent .clear").Each(func(i int, selection *goquery.Selection) {
			clear := selection.Find("div[class=info\\ clear]")
			if clear.Nodes == nil {
				return
			}
			title := utils.ReplaceAllEmpty(clear.Find(".title").Find("a[class=VIEWDATA\\ CLICKDATA\\ maidian-detail]").Text())
			link, _ := clear.Find(".title").Find("a[class=VIEWDATA\\ CLICKDATA\\ maidian-detail]").Attr("href")
			addressWithPrice := clear.Find(".address")
			address := utils.ReplaceAllEmpty(addressWithPrice.Find("div[class=flood]").Find("a").Text())
			houseInfo := utils.ReplaceAllEmpty(addressWithPrice.Find("div[class=houseInfo]").Text())
			totalPrice := utils.ReplaceAllEmpty(addressWithPrice.Find("div[class=priceInfo]").Find("div[class=totalPrice]").Text())
			unitPrice := utils.ReplaceAllEmpty(addressWithPrice.Find("div[class=priceInfo]").Find("div[class=unitPrice]").Text())
			info := &data.ErShouFangInfo{Title: title, Address: address, HouseInfo: houseInfo, TotalPrice: totalPrice, UnitPrice: unitPrice, Link: link}
			log.Printf("%+v", info)
			datas = append(datas, info)
		})
		storage.AddErShouFangRecords(datas)
	}
	storage.Db.Close()
}
