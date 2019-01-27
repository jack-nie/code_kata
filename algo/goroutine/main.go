package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	chs := make([]chan int, 4)
	files := make([]*os.File, 4)
	var err error
	for i := 0; i < len(files); i++ {
		files[i], err = os.OpenFile(fmt.Sprintf("%d.txt", i), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0655)
		if err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan int)
		go func(i int) {
			for {
				chs[i] <- i + 1
			}
		}(i)
	}

	for i := 0; i < 10; i++ {
		for j := 0; j < len(files); j++ {
			files[j].Write([]byte(strconv.Itoa(<-chs[(i+j)%(len(chs))])))
		}
	}
	for i := 0; i < len(files); i++ {
		defer files[i].Close()
	}
}
