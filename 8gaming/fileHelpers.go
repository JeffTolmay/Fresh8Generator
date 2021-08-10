package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// func writeToFile(directoryString string, allEvents []Event, interval *int, batchSize *int) {

// 	batchSizeAsInt := *batchSize

// 	for i := 0; i < len(allEvents); i += batchSizeAsInt {
// 		if i != 0 {
// 			intervalAsSec := *interval
// 			time.Sleep(time.Duration(intervalAsSec) * time.Second)
// 		}
// 		j := i + batchSizeAsInt
// 		if j > len(allEvents) {
// 			j = len(allEvents)
// 		}
// 		writeAsBatch(allEvents[i:j], directoryString)
// 	}

// }

func writeAsBatch(allEvents []Event, directoryString string) {

	currentTime := time.Now()
	fileTime := currentTime.String()[0:26]
	fileTime = strings.ReplaceAll(fileTime, " ", "-")
	fileTime = strings.ReplaceAll(fileTime, ":", "-")
	fileTime = strings.ReplaceAll(fileTime, ".", "-")

	file, err := os.Create(directoryString + "/events-" + fileTime + ".json")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	defer file.Close()

	for _, v := range allEvents {

		b, err := json.Marshal(v)
		if err != nil {
			fmt.Println("error:", err)
		}
		_, err = file.WriteString(string(b) + "\n")
		if err != nil {
			fmt.Println("error:", err)
		}
	}

}
