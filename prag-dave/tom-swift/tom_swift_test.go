package code_kata

import (
	"reflect"
	"testing"
)

func TestLoadWords(t *testing.T) {
	want := []string{"I", "wish", "I", "may", "I", "wish", "I", "might"}

	str := loadWords("../tmp/tom_swift_list.txt")

	if !reflect.DeepEqual(str, want) {
		t.Errorf("test failed, expect %s, got %s", want, str)
	}
}

func TestMakeDict(t *testing.T) {
	want := make(map[string][]string)
	want["I wish"] = []string{"I", "I"}
	want["wish I"] = []string{"may", "might"}
	want["may I"] = []string{"wish"}
	want["I may"] = []string{"I"}

	str := makeDict("../tmp/tom_swift_list.txt")
	if !reflect.DeepEqual(str, want) {
		t.Errorf("test failed, expect %s, got %s", want, str)
	}
}
