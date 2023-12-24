package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

func processNums(numChan <-chan string, wg *sync.WaitGroup) {
	for n := range numChan {
		processNum(n)
	}
	wg.Done()
}

type numData struct {
	num         string
	info        string
	preview     []byte
	previewLink string
}

func processNum(num string) {
	data := getAutohistoryJSON(num)

	nd := parseData(data, num)

	if len(nd.previewLink) > 0 {
		nd.preview = []byte(download(nd.previewLink))
	}

	writeOutput(nd)
}

func parseData(dataJson, number string) numData {
	nd := numData{
		num: number,
	}

	var dataMap map[string]any
	if !json.Valid([]byte(dataJson)) {
		log.Fatalln("invalid JSON string:", dataJson)
	}
	_ = json.Unmarshal([]byte(dataJson), &dataMap)

	if cd, ok := dataMap["carData"]; ok {
		carData := cd.(map[string]interface{})

		cdBytes, _ := json.Marshal(carData)
		nd.info = string(cdBytes)

		if image, ok := carData["image"]; ok {
			link, ok := image.(string)
			if ok && strings.HasSuffix(link, ".jpg") {
				nd.previewLink = link
			}
		}
	}

	return nd
}

func writeOutput(nd numData) {
	dirName := filepath.Join("output", nd.num)

	if err := os.MkdirAll(dirName, os.ModePerm); err != nil { // возможен ли тут гон многопоточный?
		log.Fatalln(err)
	}

	if err := os.WriteFile(filepath.Join(dirName, "info.json"), []byte(nd.info), os.ModePerm); err != nil {
		log.Fatalln(err)
	}
	if len(nd.preview) != 0 {
		if err := os.WriteFile(filepath.Join(dirName, "preview.jpg"), nd.preview, os.ModePerm); err != nil {
			log.Fatalln(err)
		}
	}
}

func download(link string) string {
	req, _ := http.NewRequest(http.MethodGet, link, nil)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	buf := new(strings.Builder)
	if _, err = io.Copy(buf, resp.Body); err != nil {
		log.Fatalln(err)
	}

	return buf.String()
}
