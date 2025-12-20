package goinsights

import (
	"fmt"
	"log"
)

func main() {
	var result int

	cond := 1
	if cond == 1 {
		result, err := FetchResultData1()
		if err != nil {
			log.Fatal("failed to retriev data")
		}
	} else {
		result, err := FetchResultData2()
		if err != nil {
			log.Fatal("failed to retriev data")
		}
	}

	fmt.Println("The result is ->", result)
}

func FetchResultData1() (bool, error) {
	return true, nil
}

func FetchResultData2() (bool, error) {
	return true, nil
}
