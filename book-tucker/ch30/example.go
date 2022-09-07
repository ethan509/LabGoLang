package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	XMLName xml.Name `xml:"response"`
	Text    string   `xml:",chardata"`
	Header  struct {
		Text       string `xml:",chardata"`
		ResultCode string `xml:"resultCode"`
		ResultMsg  string `xml:"resultMsg"`
	} `xml:"header"`
	Body struct {
		Text  string `xml:",chardata"`
		Items struct {
			Text string `xml:",chardata"`
			Item []struct {
				Text         string `xml:",chardata"`
				Career       string `xml:"career"`       // 진로 및 전망
				EngJmNm      string `xml:"engJmNm"`      // 자격영문명
				Hist         string `xml:"hist"`         // 변천과정
				ImplNm       string `xml:"implNm"`       // 시행기관
				InstiNm      string `xml:"instiNm"`      // 관련부처
				JmNm         string `xml:"jmNm"`         // 자격명
				Job          string `xml:"job"`          // 수행직무
				MdobligFldNm string `xml:"mdobligFldNm"` // 직종
				SeriesNm     string `xml:"seriesNm"`     // 자격등급
				Summary      string `xml:"summary"`      // 개요
				Trend        string `xml:"trend"`        // 출제경향
			} `xml:"item"`
		} `xml:"items"`
	} `xml:"body"`
}

func main() {
	url := "http://testapi.q-net.or.kr/api/service/rest/InquiryQualInfo/"
	urlCmd := "getList?"
	urlServiceKey := "serviceKey=tzOBycybN9XChfAO%2Fbx%2BG0aY3OrfyYq4zUowu2HUJYTiaeEl%2FiISOuNXmFMBxB%2Bj1d6VKXswysBzLaewj1WzQg%3D%3D"
	urlSeriesCd := "&seriesCd=01"
	urlAllIncluded := url + urlCmd + urlServiceKey + urlSeriesCd
	fmt.Println("urlAllIncluded =", urlAllIncluded)

	resp, err := http.Get(urlAllIncluded)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response Response
	err = xml.Unmarshal(data, &response)
	if err != nil {
		panic(err)
	}

	for i, item := range response.Body.Items.Item {
		fmt.Printf("[%d]%s\n", i, item.JmNm)
	}
	//fmt.Println(response)
}
