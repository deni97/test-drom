package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"sync"
)

func main() {
	filename := flag.String("file", "", "file name")
	concurrency := flag.Int("concurrency", 1, "во сколькеро браузеров запро́сить?")
	flag.Parse()

	if *filename == "" {
		log.Fatalln("usage: ./dromtest --file=filepath")
	}

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	numChan := numProvider(scanner)

	wg := &sync.WaitGroup{}
	wg.Add(*concurrency)
	for i := 0; i < *concurrency; i++ {
		go processNums(numChan, wg)
	}

	// ожидание воспаляет страсть
	wg.Wait()
}

func numProvider(scanner *bufio.Scanner) <-chan string {
	numChan := make(chan string)

	go func() {
		for scanner.Scan() {
			numChan <- scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		close(numChan)
	}()

	return numChan
}
