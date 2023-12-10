package api

import (
	"io"
	"net/http"
	"strconv"
)

func getSeriseCodeUrl(seriesCd int) string {
	const OPEN_API_URL string = "http://openapi.q-net.or.kr/api/service/rest/InquiryQualInfo/"
	const URL_SERVICE_KEY string = "tzOBycybN9XChfAO%2Fbx%2BG0aY3OrfyYq4zUowu2HUJYTiaeEl%2FiISOuNXmFMBxB%2Bj1d6VKXswysBzLaewj1WzQg%3D%3D"

	const SUB_CMD_GETLIST string = "getList"
	const VN_SERVICEKEY string = "serviceKey" // variable name serviceKey
	const VN_GETLIST string = "getList"       // variable name getList
	const VN_SERISECD string = "seriesCd"     // variable name seriesCd

	urlSeriesCd := "0" + strconv.Itoa(seriesCd)
	urlAllIncluded := OPEN_API_URL +
		SUB_CMD_GETLIST + "?" +
		VN_SERVICEKEY + "=" + URL_SERVICE_KEY + "&" +
		VN_SERISECD + "=" + urlSeriesCd

	//fmt.Println("urlAllIncluded =", urlAllIncluded)

	return urlAllIncluded
}

func GetSeriseCode(seriesCd int) ([]byte, error) {

	resp, err := http.Get(getSeriseCodeUrl(seriesCd))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		//panic(err)

	}

	return data, err
}