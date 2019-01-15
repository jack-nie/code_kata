package main

import (
	"context"
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type favContextKey string

func main() {
	wg := &sync.WaitGroup{}
	values := []string{"http://www.baidu.com", "http://www.sina.com"}
	ctx, cancel := context.WithCancel(context.Background())
	for _, url := range values {
		wg.Add(1)
		subCtx := context.WithValue(ctx, favContextKey("url"), url)
		go reqUrl(subCtx, wg)
	}

	go func() {
		time.Sleep(3 * time.Second)
		cancel()
	}()
	wg.Wait()
	fmt.Println("exit main goroutine")
}

func reqUrl(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	url, _ := ctx.Value(favContextKey("url")).(string)
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("stopping getting url: %s\n", url)
			return
		default:
			r, err := http.Get(url)
			if r.StatusCode == http.StatusOK && err == nil {
				body, _ := ioutil.ReadAll(r.Body)
				subCtx := context.WithValue(ctx, favContextKey("resp"), fmt.Sprintf("%s : %x", url, md5.Sum(body)))
				wg.Add(1)
				go showResp(subCtx, wg)
			}
			r.Body.Close()
			time.Sleep(1 * time.Second)
		}
	}
}

func showResp(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopping showing resp")
			return
		default:
			fmt.Printf("Printing: %s\n", ctx.Value(favContextKey("resp")).(string))
			time.Sleep(time.Second * 1)
		}
	}
}
