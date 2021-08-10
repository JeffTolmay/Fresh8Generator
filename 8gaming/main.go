package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
)

func main() {
	generateEvents()
}

func generateEvents() {
	groupNo := flag.Int("number-of-groups", 0, "Number of groups")
	batchSize := flag.Int("batch-size", 0, "Batch size")
	interval := flag.Int("interval", 0, "interval")
	directory := flag.String("output-directory", "", "Directory where file will be saved")

	flag.Parse()
	directoryString := *directory

	// Checking if directory is accessible and valid.
	_, err := os.Stat(directoryString)
	if err != nil {
		fmt.Println("Error with specified field " + err.Error())
		os.Exit(2)
	}

	if *groupNo <= 0 {
		fmt.Println("groupNo must be bigger than 0")
	}
	if *batchSize <= 0 {
		fmt.Println("batchSize must be bigger than 0")
	}
	if *interval <= 0 {
		fmt.Println("interval must be bigger than 0")
	}

	handleInput(groupNo, batchSize, interval, directoryString)
}

func handleInput(groupNo *int, batchSize *int, interval *int, directoryString string) {
	batch := *batchSize
	group := *groupNo

	// Handling batch via group size so each file contains a batch with all types of events
	for {
		if batch > group && group <= 0 {
			fmt.Println("Done")
			break
		} else {
			group = group - batch

			viewedOnly := getPercentages(float64(*batchSize), 85)
			interactedOnly := getPercentages(float64(*batchSize), 5)
			clickOnly := getPercentages(float64(*batchSize), 5)
			clickAndInteracted := getPercentages(float64(*batchSize), 5)

			/*
				Rounding the percentage values as we cannot have .5 lines in file, if a group is a small number.
			*/
			viewedOnly = math.Round(viewedOnly)
			interactedOnly = math.Round(interactedOnly)
			clickOnly = math.Round(clickOnly)
			clickAndInteracted = math.Round(clickAndInteracted)

			currentTime := time.Now()
			// Datatime is different to filedate as requested (Format).
			dataTime := currentTime.Format("2006-01-02T15:04:05Z")

			createViewedData(dataTime, viewedOnly, interactedOnly, clickOnly, clickAndInteracted, directoryString, interval)
			// sleep as requested by input.
			intervalAsSec := *interval
			time.Sleep(time.Duration(intervalAsSec) * time.Second)
		}

	}
}
