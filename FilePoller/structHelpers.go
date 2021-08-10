package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Data is needed (lint)
type Data struct {
	ViewID        string `json:"viewId"`
	EventDateTime string `json:"eventDateTime"`
}

// Event is needed (lint)
type Event struct {
	Type      string `json:"type"`
	InnerData Data   `json:"data"`
}

func readfile(file string, types []string) {
	// sleep for 1 sec before trying to access the file
	time.Sleep(1 * time.Second)
	content, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}

	var event Event
	scanner := bufio.NewScanner(content)
	var sc string
	for scanner.Scan() {
		sc = scanner.Text()
		err = json.Unmarshal([]byte(sc), &event)
		if err != nil {
			fmt.Println("error:", err)
			os.Exit(2)
		}

		types = append(types, event.Type)
	}

	var viewed int
	var interacted int
	var clicked int
	for _, x := range types {
		if x == "Viewed" {
			viewed++
		}
		if x == "Interacted" {
			interacted++
		}
		if x == "Click-Through" {
			clicked++
		}
	}

	fmt.Println("\"Viewed\" : " + strconv.Itoa(viewed))
	fmt.Println("\"Interacted\" : " + strconv.Itoa(interacted))
	fmt.Println("\"Click-Through\" : " + strconv.Itoa(clicked))
	time.Sleep(5 * time.Second)
}
