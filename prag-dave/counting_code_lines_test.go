package code_kata

import "testing"

func TestCountingCodeLines(t *testing.T) {
	for _, c := range []struct {
		want      int
		file_path string
	}{
		{85, "tmp/CountLines.java"},
	} {
		got := countCodeLines(c.file_path)
		if got != c.want {
			t.Errorf("test failed, expected %d, got %d", c.want, got)
		}
	}
}
