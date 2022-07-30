package randNumber

import (
	"math/rand"
	"time"
)

func GetRandNumber(maxNum int) int {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// set the rand seed
	randNo := random.Intn(maxNum - 1)
	//fmt.Println("Randon Number:", randNo)

	return randNo
}
