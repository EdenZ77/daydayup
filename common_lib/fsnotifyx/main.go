package main

import (
	"log"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.Println("main....")
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Create) {
					createOrWriteEvent(event.Name)
					log.Printf("listened create file: %s\n", event.Name)
				}
				if event.Has(fsnotify.Write) {
					createOrWriteEvent(event.Name)
					log.Printf("listened write file: %s\n", event.Name)
				}
				if event.Has(fsnotify.Remove) {
					removeEvent(event.Name)
					log.Printf("listened remove file: %s\n", event.Name)
				}
				if event.Has(fsnotify.Rename) {
					log.Printf("listened rename file: %s\n", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	// Add a path.
	err = watcher.Add("D:\\Mount\\sk-hkgateway-autotools")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	<-make(chan struct{})
}

func createOrWriteEvent(eventName string) {
	//ParseFilesMutex.Lock()
	//defer ParseFilesMutex.Unlock()

	fileName := getFileNameFromPath(eventName)
	log.Println(fileName)
	//if parseFile, ok := ParseFiles[fileName]; ok {
	//	parseFile.found = true
	//	if content, err := os.ReadFile(eventName); err != nil {
	//		log.Errorf("failed to read file: %s, err:%+v", eventName, err)
	//	} else {
	//		parseFile.content = content
	//	}
	//} else {
	//	log.Infof("file: %s not found in ParseFiles map", eventName)
	//}
}

func removeEvent(eventName string) {
	//ParseFilesMutex.Lock()
	//defer ParseFilesMutex.Unlock()

	fileName := getFileNameFromPath(eventName)
	log.Println(fileName)
	//if parseFile, ok := ParseFiles[fileName]; ok {
	//	parseFile.found = false
	//	parseFile.content = nil
	//} else {
	//	log.Infof("file: %s not found in ParseFiles map", eventName)
	//}
}

// 获得文件名
func getFileNameFromPath(pathStr string) string {
	_, file := filepath.Split(pathStr)
	return file
}
