package code_kata

import "testing"

func TestJoaatHash(t *testing.T) {
	for _, c := range []struct {
		want uint
		key  string
	}{
		{8847632551651054667, "hello"},
		{1758146220456233818, "word"},
	} {
		got := joaatHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestDjb2Hash(t *testing.T) {
	for _, c := range []struct {
		want uint
		key  string
	}{
		{210714636441, "hello"},
		{6385842145, "word"},
	} {
		got := djb2Hash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestSdbmHash(t *testing.T) {
	for _, c := range []struct {
		want uint
		key  string
	}{
		{27263685106, "hello"},
		{245562474, "word"},
	} {
		got := sdbmHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestLoseLoseHash(t *testing.T) {
	for _, c := range []struct {
		want uint
		key  string
	}{
		{532, "hello"},
		{444, "word"},
	} {
		got := loseLoseHash(c.key)

		if got != c.want {
			t.Errorf("test faild, expected %d, got %d", c.want, got)
		}
	}
}

func TestBloomFilters(t *testing.T) {
	for _, c := range []struct {
		want      bool
		file_path string
		target    string
	}{
		{true, "tmp/wordlist.txt", "hello"},
		{true, "tmp/wordlist.txt", "world"},
		{true, "tmp/wordlist.txt", "say"},
	} {
		got := bloomFilters(c.file_path, c.target)

		if got != c.want {
			t.Errorf("test failed, expected %s, got %s", c.want, got)
		}
	}
}
