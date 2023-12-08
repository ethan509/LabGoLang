package main

import (
	"common"
	"encoding/xml"
	"fmt"
	"io"
	"mongo"
	"net/http"
	"strconv"
)

func GetInfo(seriesCd int) ([]byte, error) {
	url := "http://openapi.q-net.or.kr/api/service/rest/InquiryQualInfo/"
	urlCmd := "getList?"
	urlServiceKey := "serviceKey=tzOBycybN9XChfAO%2Fbx%2BG0aY3OrfyYq4zUowu2HUJYTiaeEl%2FiISOuNXmFMBxB%2Bj1d6VKXswysBzLaewj1WzQg%3D%3D"
	// urlSeriesCd : (계열코드) 01:기술사, 02:기능장, 03:기사, 04:기능사
	urlSeriesCd := "&seriesCd=0" + strconv.Itoa(seriesCd)
	urlAllIncluded := url + urlCmd + urlServiceKey + urlSeriesCd

	fmt.Println("urlAllIncluded =", urlAllIncluded)

	resp, err := http.Get(urlAllIncluded)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		//panic(err)

	}

	return data, err
}

func main() {
	mongo.DbClient = mongo.MongoConn()

	for i := 1; i <= 4; i++ {
		data, _ := GetInfo(i)

		var response common.Response
		err := xml.Unmarshal(data, &response)
		if err != nil {
			panic(err)
		}

		var mongoItems []interface{}
		for i, item := range response.Body.Items.Item {
			fmt.Printf("[%d]%s\t", i, item.JmNm)
			var mongoItem common.MongoItem

			mongoItem.Career = item.Career
			mongoItem.EngJmNm = item.Career
			mongoItem.Hist = item.Hist
			mongoItem.ImplNm = item.ImplNm
			mongoItem.InstiNm = item.InstiNm
			mongoItem.JmNm = item.JmNm
			mongoItem.Job = item.Job
			mongoItem.MdobligFldNm = item.MdobligFldNm
			mongoItem.SeriesNm = item.SeriesNm
			mongoItem.Summary = item.Summary
			mongoItem.Trend = item.Trend

			mongoItems = append(mongoItems, mongoItem)
			//mongo.MongoInsertOne(mongoItem)
		}

		mongo.MongoInsertMany(mongoItems)
	}
}
