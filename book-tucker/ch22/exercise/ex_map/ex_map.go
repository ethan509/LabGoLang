package ex_map

import "fmt"

func ExMap() {
	fmt.Println("ExMap")

	m := make(map[string]string)

	m["이화랑"] = "서울시 광진구"
	m["송하나"] = "서울시 강남구"
	m["백두산"] = "서울시 서대문구"
	m["최번개"] = "서울시 마포구"

	m["최번개"] = "제주도 제주시"

	fmt.Printf("송하나의 주소는 %s입니다.", m["송하나"])
	fmt.Printf("백두산의 주소는 %s입니다.", m["백두산"])

	// v = 값, ok = 존재여부
	v, ok := m["송하나"]

	fmt.Println()
	fmt.Printf("%s(%t)", v, ok)

	// v = 값, ok = 존재여부
	v, ok = m["송새벽"]
	fmt.Println()
	fmt.Printf("%s(%t)", v, ok)
}
