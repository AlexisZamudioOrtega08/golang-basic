package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	// init time
	start := time.Now()
	jobs := make(chan int, 10)
	results := make(chan string, 10)
	for j := 1; j <= 10; j++ {
		jobs <- j
		go worker(jobs, results)
	}
	close(jobs)
	for a := 1; a <= 10; a++ {
		select {
		case r := <-results:
			fmt.Println(r)
		//add a case when is time out
		case <-time.After(time.Second * 1):
			fmt.Println("timeout")
			break
		}
	}
	// end time
	end := time.Now()
	fmt.Println("time: ", end.Sub(start))
}

func worker(jobs <-chan int, results chan<- string) {
	id := <-jobs
	url := "https://62e46e8dc6b56b451186101c.mockapi.io/api/v1/dummies/"
	response, err := http.Get(fmt.Sprintf(url+"%d", id))

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	var jsonResponse map[string]interface{}
	json.Unmarshal([]byte(responseData), &jsonResponse)
	results <- string("id " + jsonResponse["id"].(string) + " completed")
}
