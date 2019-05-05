package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {
	urls := []string{
		"http://dummy.restapiexample.com/api/v1/employee/22063",
		"http://dummy.restapiexample.com/api/v1/employee/22064",
		"http://dummy.restapiexample.com/api/v1/employee/22067",
	}
	jsonResponses := make(chan string, 1)

	var wg sync.WaitGroup

	wg.Add(len(urls))

	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			res, err := http.Get(url)
			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				//var p = make([]rune, 1)
				var reader = bufio.NewReader(res.Body)
				for {
					r, _, err := reader.ReadRune()
					if err == io.EOF {
						return
					}
					jsonResponses <- string(r)
				}

				if err != nil {
					log.Fatal(err)
				}
			}
		}(url)
	}

	go func() {
		for response := range jsonResponses {
			fmt.Print(response)
		}
	}()

	wg.Wait()
}
