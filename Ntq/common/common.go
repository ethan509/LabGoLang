package common

import "encoding/xml"

type Serise int

const (
	BeginSeriseCode       int = iota
	ProfessionalEngineers     // 기술사
	MasterCraftsment          // 기능장
	Engineers                 // 기사
	Technicians               // 기능사
	EndSeriseCode
)

// 01:기술사, 02:기능장, 03:기사, 04:기능사
func (code Serise) String() string {
	return [...]string{
		"BeginSeriseCode",
		"기술사",
		"기능장",
		"기사",
		"기능사",
		"EndSeriseCode"}[code]
}

// https://www.data.go.kr/data/15041600/openapi.do?recommendDataYn=Y
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

type MongoItem struct {
	Career       string `bson:"career"`       // 진로 및 전망
	EngJmNm      string `bson:"engJmNm"`      // 자격영문명
	Hist         string `bson:"hist"`         // 변천과정
	ImplNm       string `bson:"implNm"`       // 시행기관
	InstiNm      string `bson:"instiNm"`      // 관련부처
	JmNm         string `bson:"jmNm"`         // 자격명
	Job          string `bson:"job"`          // 수행직무
	MdobligFldNm string `bson:"mdobligFldNm"` // 직종
	SeriesNm     string `bson:"seriesNm"`     // 자격등급
	Summary      string `bson:"summary"`      // 개요
	Trend        string `bson:"trend"`        // 출제경향
}
