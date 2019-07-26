package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

type Counter struct {
	occurrence map[string]int
	mutex      sync.Mutex
}

func (counter *Counter) print() {
	for word, count := range counter.occurrence {
		fmt.Println(word, ":", count)
	}
}

func (counter *Counter) updateCounter(channel chan string) {
	for {
		word := <-channel

		counter.mutex.Lock()
		counter.occurrence[word]++
		counter.mutex.Unlock()
	}
}

func scanFile(filePath string, channel chan string, group *sync.WaitGroup) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		channel <- scanner.Text()
	}

	group.Done()
}

func scanFiles(files []string) (counter Counter) {
	channel := make(chan string)

	counter = Counter{occurrence: make(map[string]int)}
	go counter.updateCounter(channel)

	var group sync.WaitGroup
	for _, file := range files {
		group.Add(1)
		go scanFile(file, channel, &group)
	}
	group.Wait()
	return
}

func scanDirectory(directoryPath string) []string {
	var files []string
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return files
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("go run main.go [DIRECTORY_PATH]")
	}

	directory := os.Args[1]

	files := scanDirectory(directory)

	counter := scanFiles(files)
	counter.print()
}
