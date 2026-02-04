//go:build !solution

package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"sync"
	"time"
)

func main() {
	args := os.Args[1:]
	now := time.Now()

	var wg sync.WaitGroup

	for _, url := range args {
		wg.Add(1)

		go func(url string) {
			defer wg.Done()

			resp, err := http.Get(url)
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()

			bytes, _ := io.Copy(io.Discard, resp.Body)
			log.Printf("%6.3fs %7d %s\n", time.Since(now).Seconds(), bytes, url)
		}(url)
	}
	wg.Wait()
	log.Printf("Total time: %6.3fs\n", time.Since(now).Seconds())
}
