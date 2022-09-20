package utils

import (
	excelcontroller
)

func (info excelcontroller.aaa) Len() int {
	return len(info)
}

func (info excelcontroller.aaa) Swap(i, j int) {
	info[i], info[j] = info[j], info[i]
}

func (info excelcontroller.aaa) Less(i, j int) bool {
	return info[i].NewPhoneNumber < info[j].NewPhoneNumber
}
