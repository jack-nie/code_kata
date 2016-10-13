package code_kata

import (
	"bufio"
	"fmt"
	"github.com/willf/bitset"
	"os"
)

func joaatHash(key string) uint {
	var hash uint
	var i, key_len int

	key_len = len(key)
	hash = 0

	for i = 0; i < key_len; i++ {
		hash += uint(key[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}

	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return hash
}

func djb2Hash(key string) uint {
	var hash uint
	var i, key_len int

	key_len = len(key)
	hash = 5381

	for i = 0; i < key_len; i++ {
		hash = ((hash << 5) + hash) + uint(key[i])
	}

	return hash
}

func sdbmHash(key string) uint {
	var hash uint
	var i, key_len int

	key_len = len(key)

	for i = 0; i < key_len; i++ {
		hash = uint(key[i]) + (hash << 6) + (hash << 6) - hash
	}

	return hash
}

func loseLoseHash(key string) uint {
	var hash uint
	var i, key_len int

	key_len = len(key)

	for i = 0; i < key_len; i++ {
		hash += uint(key[i])
	}

	return hash
}

func add(value string, s *bitset.BitSet) {
	//s.Set(joaatHash(value))
	//s.Set()
	//s.Set(sdbmHash())
	fmt.Println(joaatHash(value))
	fmt.Println(djb2Hash(value))
	fmt.Println(sdbmHash(value))
	s.Set(loseLoseHash(value))
}

func bloomFilters(file_path string, target string) bool {
	s := new(bitset.BitSet)
	dat, _ := os.Open(file_path)
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		word := scanner.Text()
		add(word, s)
	}
	a := joaatHash(target)
	b := djb2Hash(target)
	c := sdbmHash(target)
	d := loseLoseHash(target)
	return s.Test(a) && s.Test(b) && s.Test(c) && s.Test(d)
}
