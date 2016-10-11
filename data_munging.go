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

func processWeather() (int64, int64) {

  var gap int64
  var max int64
  var min int64
  var day int64
  m := make(map[string]int64)

  dat, err := os.Open("tmp/weather.dat")
  check(err)

  scanner := bufio.NewScanner(dat)

  m["min"] = 10000
  m["day"] = -1

  for scanner.Scan() {
    arr := strings.Fields(scanner.Text())

    if len(arr) == 0 || isNotNumeric(arr[0]) {
      continue
    }

    max_t := parseIntFromStr(arr[1])
    min_t := parseIntFromStr(arr[2])
    max, _ = strconv.ParseInt(max_t, 10, 64)
    min, _ = strconv.ParseInt(min_t, 10, 64)

    gap = max - min
    day, _ = strconv.ParseInt(arr[0], 10, 64)

    if gap < m["min"] {
      m["min"] = gap
      m["day"] = day
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return m["min"], m["day"]
}

func processSoccerLeagueTable() (interface{}, interface{}) {
  m := make(map[string]interface{})

  dat, err := os.Open("tmp/football.dat")
  check(err)

  m["min"] = int64(10000)
  m["league"] = -1

  scanner := bufio.NewScanner(dat)

  for scanner.Scan() {
    arr := strings.Fields(scanner.Text())

    if len(arr) == 0 || !strings.Contains(arr[0], ".") {
      continue
    }

    for_t := parseIntFromStr(arr[6])
    against_t := parseIntFromStr(arr[8])
    for_score, _ := strconv.ParseInt(for_t, 10, 64)
    against_score, _ := strconv.ParseInt(against_t, 10, 64)

    gap := int64(math.Abs(float64(for_score - against_score)))

    if gap < m["min"].(int64) {
      m["min"] = gap
      m["league"] = arr[1]
    }
  }

  if err := scanner.Err(); err != nil {
    log.Fatal(err)
  }

  return m["min"], m["league"]
}
