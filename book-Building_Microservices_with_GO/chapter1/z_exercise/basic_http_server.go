package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type helloWorldRequest struct {
	Name string `json:"name"`
}

type helloWorldResponse struct {
	Message string `json:"message"`    // 출력 필드를 변수명 그대로인 Message 가 아닌 message로 바꿔준다.
	Author  string `json:"-"`          // 이 필드를 출력하지 않는다.
	Date    string `json:",omitempty"` // 값이 비어있으면 출력하지 않는다.
	Id      int    `json:"id,string"`  // 출력을 문자열로 변환하고 이름을 "id"로 바꾼다.
}

// 실행방법 : http://localhost:8080/helloworld
func main() {
	port := 8080

	// HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	// DefaultServeMux에 주어진 패턴에 대한 핸들러 함수를 등록한다. (DefaultServeMux에 함수를 등록)
	http.HandleFunc("/helloworld", func(w http.ResponseWriter, r *http.Request) {

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		var request helloWorldRequest
		err = json.Unmarshal(body, &request)
		if err != nil {
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		response := helloWorldResponse{Message: "HelloWorld" + request.Name}
		encoder := json.NewEncoder(w)
		encoder.Encode(&response)
		/* 아래는 바이트 배열로 마샬링 하는 방식. 위처럼 JSON을 바로 쓸 수 있는 Encoder 객체를 사용하는 것이 더 빠르다.
		// data, err := json.Marshal(response)
		// if err != nil {
		// 	panic("Ooops")
		// }
		// fmt.Fprint(w, string(data))
		*/
	})

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldImageHandler(w http.ResponseWriter, r *http.Request) {
}
