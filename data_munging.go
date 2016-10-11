package code_kata

import (
  "bufio"
  "log"
  "os"
  "strings"
  "strconv"
  "unicode"
  "math"
)

func check(err error) {
  if err != nil {
    log.Fatal(err)
  }
}

func isNotNumeric(str string) bool {
  _, err := strconv.ParseInt(str, 10, 64)
  return err != nil
}

func parseIntFromStr(str string) string {
  f := func(c rune) bool {
    return !unicode.IsNumber(c) && (c != 46)
  }

  output := strings.FieldsFunc(str, f)
  if len(output) > 0 {
    return output[0]
  } else {
    return " "
  }
}

func calculateGap(max_idx int, min_idx int, arr []string) int64{
  max_t := parseIntFromStr(arr[max_idx])
  min_t := parseIntFromStr(arr[min_idx])
  max, _ := strconv.ParseInt(max_t, 10, 64)
  min, _ := strconv.ParseInt(min_t, 10, 64)
  gap := int64(math.Abs(float64(max - min)))
  return gap
}

func check_scanner_error(scanner *bufio.Scanner) {
  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }
}

func processWeather(file_path string) (int64, int64) {
  m := make(map[string]int64)

  dat, err := os.Open(file_path)
  check(err)

  scanner := bufio.NewScanner(dat)

  m["min"] = 10000
  m["day"] = -1

  for scanner.Scan() {
    arr := strings.Fields(scanner.Text())

    if len(arr) == 0 || isNotNumeric(arr[0]) {
      continue
    }

    gap := calculateGap(1,2, arr)
    day, _ := strconv.ParseInt(arr[0], 10, 64)

    if gap < m["min"] {
      m["min"] = gap
      m["day"] = day
    }
  }

  check_scanner_error(scanner)

  return m["min"], m["day"]
}



func processSoccerLeagueTable(file_path string) (interface{}, interface{}) {
  m := make(map[string]interface{})

  dat, err := os.Open(file_path)
  check(err)

  m["min"] = int64(10000)
  m["league"] = -1

  scanner := bufio.NewScanner(dat)

  for scanner.Scan() {
    arr := strings.Fields(scanner.Text())

    if len(arr) == 0 || !strings.Contains(arr[0], ".") {
      continue
    }

    gap := calculateGap(6, 8, arr)

    if gap < m["min"].(int64) {
      m["min"] = gap
      m["league"] = arr[1]
    }
  }

  check_scanner_error(scanner)

  return m["min"], m["league"]
}
