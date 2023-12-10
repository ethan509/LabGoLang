package main

import (
	"api"
	"common"
	"encoding/xml"
	"fmt"
	"mongo"
	"sync"
	"time"
)

func getSeriesCode(code common.Serise, wg *sync.WaitGroup) {
	defer wg.Done()
	//fmt.Printf("[%d] go start of %s\n", code, code)

	data, _ := api.GetSeriseCode(int(code))

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

	var insertedCount int
	if len(mongoItems) > 0 {
		insertedCount = mongo.InsertMany(mongoItems)
	}

	fmt.Printf("[%02d]%-4s API Search result: %3v, insertedCount: %3d\n", code, code, len(mongoItems), insertedCount)

	//fmt.Printf("[%d] go end of %s\n", code, code)
}

func main() {
	mongo.DbClient = mongo.Connect()

	var wg sync.WaitGroup

	start := time.Now()
	time.Sleep(time.Second)

	wg.Add(common.EndSeriseCode - 1)
	for i := common.Serise(common.BeginSeriseCode) + 1; i < common.Serise(common.EndSeriseCode); i++ {
		go getSeriesCode(i, &wg)
	}

	wg.Wait()

	end := time.Since(start)
	fmt.Println(end)
}
