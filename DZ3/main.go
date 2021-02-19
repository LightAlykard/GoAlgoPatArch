package main

import (
	"bytes"
	"flag"
	"fmt"
	"net/http"
	"time"
)

var timeAll float64
var goodResp float64

func httpGet(url string) (int, float64) {
	start := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("httpGet error")
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()
	return resp.StatusCode, elapsed
}

func httpPost(url string, body []byte) (int, float64) {
	start := time.Now()

	r := bytes.NewReader(body)
	resp, err := http.Post(url, "application/json", r)
	if err != nil {
		log.Fatalf("httpPost error")
	}
	defer resp.Body.Close()

	elapsed := time.Since(start).Seconds()
	return resp.StatusCode, elapsed
}

func worker(w int, url string, jobs <-chan int, results chan<- int, method string, body []byte) {
	if method == "GET" {
		for j := range jobs {
			code, t := httpGet(url)
			if code >= 200 && code <= 299 {
				fmt.Println("Время отклика: ", t)
				goodResp++
			}
			timeAll += t

			results <- j * 2
		}
	}
	if method == "POST" {
		for j := range jobs {
			code, t := httpPost(url, body)
			if code >= 200 && code <= 299 {
				fmt.Println("Время отклика: ", t)

				goodResp++
			}
			timeAll += t

			results <- j * 2
		}
	}

}

func main() {

	var url string
	var workerCount int
	var jobCount int	
	var httpMethod string
	var bodyReques string

	flag.StringVar(&url, "url", "https://yandex.ru", "URL сайта")
	flag.IntVar(&workerCount, "t", 5, "number of threads")
	flag.IntVar(&jobCount, "r", 5, "number of requests")	
	flag.StringVar(&httpMethod, "m", "GET", "method GET or POST")
	flag.StringVar(&bodyReques, "b", "", "Body for post request")

	flag.Parse()

	if httpMethod != "GET" && httpMethod != "POST" {
		panic("Вы ввели не верный метод. Разрешенный метод \"GET\" или \"POST\"")
	}

	if workerCount <= 0 {
		panic("Колличество потоков должно быть больше 0")
	}

	if jobCount <= 0 {
		panic("Колличество запросов должно быть больше 0")
	}

	jobs := make(chan int, jobCount)
	results := make(chan int, jobCount)

	for w := 0; w < workerCount; w++ {
		go worker(w, url, jobs, results, httpMethod, []byte(bodyReques))
	}

	for j := 0; j < jobCount; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 0; a < jobCount; a++ {
		<-results
	}

	avg := timeAll / goodResp
	rpc := goodResp / timeAll

	fmt.Println("RPC: ", rpc)
	fmt.Println("AVG time: ", avg)
}
