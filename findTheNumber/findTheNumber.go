package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// set the rand seed
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	// set the rand seed
	randNo := random.Intn(100) + 1
	//fmt.Println("Randon Number:", randNo)

	inputMinNo := 0
	inputMaxNo := 100
	for cnt := 10; cnt > 0; cnt-- {
		fmt.Print("[", cnt, "] enter the number: ")
		stdin := bufio.NewReader(os.Stdin)
		inputStr, err := stdin.ReadString('\n')
		inputStr = strings.TrimSuffix(inputStr, "\r\n")
		if err != nil {
			log.Fatal(err)
		}

		inputNo, err := strconv.Atoi(inputStr)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(input)

		if inputNo > randNo {
			fmt.Println("[too High]")
			inputMaxNo = inputNo
		} else if inputNo < randNo {
			fmt.Println("[too Low]")
			inputMinNo = inputNo
		} else {
			fmt.Println("That's right !!!")
			break
		}
		fmt.Println(inputMinNo, "< right answer <", inputMaxNo)
	}
}
