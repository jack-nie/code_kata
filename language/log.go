package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/facebookgo/atomicfile"
)

func main() {
	flag.Parse()
	f, err := os.OpenFile(os.Args[1], os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	mu := new(sync.Mutex)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		mu.Lock()
		defer mu.Unlock()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Print(err)
			return
		}
		_, err = f.Write(body)
		if err != nil {
			log.Print(err)
			return
		}

		_, err = f.WriteString("\n")
		if err != nil {
			log.Print(err)
			return
		}
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	go func() {
		<-c
		f.Close()
		os.Exit(0)
	}()
	err = Write()
	if err != nil {
		log.Print(err)
	}
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Listen and serve: %s", err)
	}
}

var (
	errNotConfigured = errors.New("pidfile not configured")
	pidfile          = flag.String("pidfile", "", "If specified, write pid to file.")
)

// IsNotConfigured returns true if the error indicates the pidfile location has
// not been configured.
func IsNotConfigured(err error) bool {
	if err == errNotConfigured {
		return true
	}
	return false
}

// GetPidfilePath returns the configured pidfile path.
func GetPidfilePath() string {
	return *pidfile
}

// SetPidfilePath sets the pidfile path.
func SetPidfilePath(p string) {
	*pidfile = p
}

// Write the pidfile based on the flag. It is an error if the pidfile hasn't
// been configured.
func Write() error {
	if *pidfile == "" {
		return errNotConfigured
	}

	if err := os.MkdirAll(filepath.Dir(*pidfile), os.FileMode(0755)); err != nil {
		return err
	}

	file, err := atomicfile.New(*pidfile, os.FileMode(0644))
	if err != nil {
		return fmt.Errorf("error opening pidfile %s: %s", *pidfile, err)
	}
	defer file.Close() // in case we fail before the explicit close

	_, err = fmt.Fprintf(file, "%d", os.Getpid())
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}

	return nil
}

// Read the pid from the configured file. It is an error if the pidfile hasn't
// been configured.
func Read() (int, error) {
	if *pidfile == "" {
		return 0, errNotConfigured
	}

	d, err := ioutil.ReadFile(*pidfile)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(d)))
	if err != nil {
		return 0, fmt.Errorf("error parsing pid from %s: %s", *pidfile, err)
	}

	return pid, nil
}
