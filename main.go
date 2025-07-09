package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	path := flag.String("file", "url.txt", "path to file with url")
	flag.Parse()

	var wg sync.WaitGroup
	var (
		muBad sync.Mutex
		muGood sync.Mutex
	)

	file, err := os.ReadFile(*path)
	
	if err != nil {
		fmt.Println("Reading file err -", err)
		return
	}

	badURL := map[string]error{}
	goodURL := map[string]string{}
	urls := strings.Split(string(file), "\n")

	for _, v := range urls {
		wg.Add(1)
		go func(v string) {
			cleanedURL := strings.TrimRight(v, "\r\n") // url may have invisible signs, when u just copy them from brauser
			res, err := http.Get(cleanedURL)
			if err != nil {
				muBad.Lock()
				badURL[cleanedURL] = err
				muBad.Unlock()
			} else {
				muGood.Lock()
				goodURL[cleanedURL] = res.Status
				muGood.Unlock()
			}
			wg.Done()
		}(v)
	}

	wg.Wait()
	resFile, err := os.Create("result.txt")
	
	if err != nil {
		fmt.Println("Creating resultt file err -", err)
		return
	}
	defer resFile.Close()

	//Write err ping to res file
	resFile.WriteString("Error while ping next urls: \n\n")

	for k, v := range badURL {
		wg.Add(1)
		go func(k string, v error) {
			defer wg.Done() 
			_, writeErr := fmt.Fprintf(resFile, "URL: %s, err: %v\n", k, v)
			if writeErr != nil {
				fmt.Printf("Failed to write to file: %v\n", writeErr)
			}
		}(k, v)
	}
	wg.Wait()

	resFile.WriteString("\n\nSuccessfully pinged: \n\n")
	//Write good ping to res file
	for k, v := range goodURL {
		wg.Add(1)
		go func(k string, v string) {
			defer wg.Done() 
			_, writeErr := fmt.Fprintf(resFile, "URL: %s, status: %v\n", k, v)
			if writeErr != nil {
				fmt.Printf("Failed to write to file: %v\n", writeErr)
			}
		}(k, v)
	}
	wg.Wait()

	fmt.Println("Result file was created")
	

}
