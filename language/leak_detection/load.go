package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"sync"
)

const leaky_server = "http://localhost:8080"

func getReq(endPoint string) ([]byte, error) {
	res, err := http.Get(leaky_server + endPoint)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func getRoutineCount() (int, error) {
	body, err := getReq("/_count")
	if err != nil {
		return -1, err
	}
	count, err := strconv.Atoi(string(body))
	if err != nil {
		return -1, err
	}
	return count, nil
}

func getStackTrace() (string, error) {
	body, err := getReq("/_stack")
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func main() {
	count, err := getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d goroutines before the test begin.\n", count)

	trace, err := getStackTrace()
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("stack trace after the load test: \n %s", trace)
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := getReq("/sum")
			if err != nil {
				log.Fatal(err)
			}
		}()
	}
	wg.Wait()

	count, err = getRoutineCount()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d goroutines after the test finished.\n", count)

	trace, err = getStackTrace()
	if err != nil {
		log.Fatal()
	}
	fmt.Printf("stack trace after the load test: \n %s", trace)
}
