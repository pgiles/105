package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
)

var urls = []string{
	"https://google.com",
	"https://bing.com",
	"https://sony.com",
}

func main() {
	fmt.Println("Running")
	log.SetHandler(text.Default)
	http.HandleFunc("/", notifyMe)
	log.Fatalf("%+v\n", http.ListenAndServe(":8080", nil))
}

func notifyMe(writer http.ResponseWriter, request *http.Request) {
	defer log.WithField("path", request.URL).Trace("tracking").Stop(new(error))
	log.Info("notifying")

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			defer log.WithField("path", url).Trace("track url").Stop(new(error))
			resp, err := http.Head(url)
			if err != nil {
				fmt.Println(err)
			}
			_, _ = fmt.Fprintf(writer, "%s %+v\n", url, resp.Status)
			wg.Done()
		}(url)
	}
	wg.Wait()
	log.Info("done")
}
