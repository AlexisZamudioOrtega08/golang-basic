package main

import (
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
	for i := 1; i <= 10; i++ {
		url := "https://62e46e8dc6b56b451186101c.mockapi.io/api/v1/dummies/"
		response, err := http.Get(fmt.Sprintf(url+"%d", i))

		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}

		responseData, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(responseData))
	}
	// end time
	end := time.Now()
	fmt.Println("time: ", end.Sub(start))
}
