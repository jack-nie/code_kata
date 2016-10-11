package main

import (
  "bufio"
  "os"
  "fmt"
)

type Set struct {
  set map[string]bool
}

func (set *Set) add(str string) bool {
  _, found := set.set[str]
  set.set[str] = true
  return !found
}

func checkWords(file_path string) []string {
  var arr []string
  var result []string
  var length int
  var set Set

  set.set = make(map[string]bool)
  dat, _ := os.Open(file_path)
  scanner := bufio.NewScanner(dat)

  for scanner.Scan() {
    word := scanner.Text()
    length = len(word)
    if length > 6 {
      continue
    } else if length < 6 {
      set.add(word)
    } else if length == 6{
      arr = append(arr, word)
    }
  }

  for _, element := range arr {
    for i := 1; i < 6; i++ {
      first := element[0:i]
      last  := element[i:]
      if !set.add(first) && !set.add(last) {
        result = append(result, element)
        break
      }
    }
  }
  return result
}

func main() {
  s := checkWords("tmp/wordlist.txt")
  fmt.Println(s)
}
