package process

import (
	"api"
	"common"
	"encoding/xml"
	"fmt"
	"mongo"
	"sync"
	"time"
)

func getSeriesCode(code common.Serise, countList []int, wg *sync.WaitGroup) {
	defer wg.Done()

	startApi := time.Now()
	time.Sleep(time.Second)

	// openapi.q-net.or.kr 에서 '국가기술자격 종목 정보' 가져오기
	data, _ := api.GetSeriseCode(int(code))
	endApi := time.Since(startApi)

	var response common.Response
	err := xml.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	var mongoItems []interface{}
	for _, item := range response.Body.Items.Item {
		//fmt.Printf("[%d]%s\t", i, item.JmNm)
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
	var endDb time.Duration

	if countList[code] == len(mongoItems) {
		fmt.Printf("[%02d]%-3s Already inserted(API Search result: %3v, insertedCount: %3d, apiTime:%v, dbTime:%v)\n",
			code, code, len(mongoItems), insertedCount, endApi.Milliseconds(), endDb.Milliseconds())
	} else if len(mongoItems) == 0 {
		fmt.Printf("[%02d]%-3s not Searched(API Search result: %3v, insertedCount: %3d, apiTime:%v, dbTime:%v)\n",
			code, code, len(mongoItems), insertedCount, endApi.Milliseconds(), endDb.Milliseconds())
	} else if len(mongoItems) > 0 {
		startDb := time.Now()
		time.Sleep(time.Second)
		insertedCount = mongo.InsertMany(mongoItems)
		endDb = time.Since(startDb)

		fmt.Printf("[%02d]%-3s API Search result: %3v, insertedCount: %3d, apiTime:%v, dbTime:%v\n",
			code, code, len(mongoItems), insertedCount, endApi.Milliseconds(), endDb.Milliseconds())
	} else {
		fmt.Printf("[%02d]%-3s API Search result: %3v, insertedCount: %3d, apiTime:%v, dbTime:%v\n",
			code, code, len(mongoItems), insertedCount, endApi.Milliseconds(), endDb.Milliseconds())
	}

	//fmt.Printf("[%d] go end of %s\n", code, code)
}

func SaveCertificate() {
	var wg sync.WaitGroup

	// [Step2] 이미 저장된 정보 개수 파악(serise code별로)
	countList := mongo.GetCountCollectionCount()

	start := time.Now()
	time.Sleep(time.Second)

	wg.Add(common.EndSeriseCode - 1)
	for i := common.Serise(common.BeginSeriseCode) + 1; i < common.Serise(common.EndSeriseCode); i++ {
		go getSeriesCode(i, countList, &wg)
	}

	wg.Wait()

	end := time.Since(start)
	fmt.Println(end)
}
