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
	"strings"
)

func ProcessBeiKeZuFang() {
	storage := dao.ConnectMysql()
	for i := 1; i <= 5; i++ {
		url := constant.Zufang_beike + strconv.Itoa(i)
		res, err := http.Get(url)
		if err != nil {
			fmt.Printf("fail %+v", err)
		}
		if res.StatusCode != http.StatusOK {
			log.Fatal(res.StatusCode)
		}
		doc, err := goquery.NewDocumentFromReader(res.Body)
		var datas = make([]*data.ZuFangInfo, 0)
		doc.Find(".wrapper").Find("#content").Find(".content__list .content__list--item").Each(func(i int, selection *goquery.Selection) {
			title := utils.ReplaceAllEmpty(selection.Find(".content__list--item--main").Find("p[class=content__list--item--title\\ twoline]").Text())
			houseInfo := utils.ReplaceAllEmpty(selection.Find(".content__list--item--main .content__list--item--des").Text())
			address := strings.Split(houseInfo, "-")[0]
			link, _ := selection.Find(".content__list--item--main").Find("p[class=content__list--item--title\\ twoline]").Find("a").Attr("href")
			price := utils.ReplaceAllEmpty(selection.Find(".content__list--item--main .content__list--item-price").Text())
			info := &data.ZuFangInfo{Title: title, Address: address, HouseInfo: houseInfo, Price: price, Link: "https://bj.zu.ke.com" + link}
			log.Printf("%+v", info)
			datas = append(datas, info)
		})
		storage.AddZuFangRecords(datas)
	}
	storage.Db.Close()
}
