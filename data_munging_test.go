package code_kata

import "testing"

func TestProcessWeather(t *testing.T) {
  for _, c := range []struct{
    want_one int64
    want_two int64
    file_path string
  } {
    {2, 14, "tmp/weather.dat"},
  } {
     got_one, got_two := processWeather(c.file_path)
     if got_one != c.want_one || got_two != c.want_two {
       t.Errorf("test faild, expected %d, %d, got %d, %d", c.want_one, c.want_two, got_one, got_two)
     }
  }
}

func TestProcessSoccerLeagueTable(t *testing.T) {
  for _, c := range []struct{
    want_one int64
    want_two string
    file_path string
  } {
    {1, "Aston_Villa", "tmp/football.dat"},
  } {
    got_one, got_two := processSoccerLeagueTable(c.file_path)
    if got_one != c.want_one || got_two != c.want_two {
      t.Errorf("test faild, expected %d, %s, got %d, %s", c.want_one, c.want_two, got_one, got_two)
    }
  }
}
