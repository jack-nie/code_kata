package code_kata

import "testing"

func TestJoaatHash(t *testing.T) {
	for _, c := range []struct {
		want int
		key  []byte
	}{
		{8847632551651054667, []byte("hello")},
		{409105252719273154, []byte("world!")},
	} {
		got := joaatHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestDjb2Hash(t *testing.T) {
	for _, c := range []struct {
		want int
		key  []byte
	}{
		{210714636441, []byte("hello")},
		{6954182107950, []byte("world!")},
	} {
		got := djb2Hash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestSdbmHash(t *testing.T) {
	for _, c := range []struct {
		want int
		key  []byte
	}{
		{27263685106, []byte("hello")},
		{3960677284911, []byte("world!")},
	} {
		got := sdbmHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestLoseLoseHash(t *testing.T) {
	for _, c := range []struct {
		want int
		key  []byte
	}{
		{532, []byte("hello")},
		{585, []byte("world!")},
	} {
		got := loseLoseHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestStringBloomFilters(t *testing.T) {
	for _, c := range []struct {
		want      bool
		file_path string
		target    string
	}{
		{true, "tmp/wordlist.txt", "hello"},
		{true, "tmp/wordlist.txt", "world"},
		{true, "tmp/wordlist.txt", "say"},
	} {
		got := stringBloomFilters(c.file_path, c.target)

		if got != c.want {
			t.Errorf("test failed, expected %s, got %s", c.want, got)
		}
	}
}
