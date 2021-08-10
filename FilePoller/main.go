package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fsnotify/fsnotify"
)

func main() {
	var types []string
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(2)
	}
	defer watcher.Close()

	var file string
	done := make(chan bool)
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					file = event.Name
					fmt.Println(file)

				}
				if strings.Contains(file, "events") && strings.Contains(file, ".json") {
					readfile(file, types)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
				os.Exit(2)
			}
		}
	}()

	err = watcher.Add("C:/Users/jtolmay/Desktop/Watchedfile")
	if err != nil {
		log.Fatal(err)
		os.Exit(2)
	}
	<-done
}
