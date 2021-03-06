package code_kata

import (
	"bufio"
	"os"
	"xojoc.pw/bitset"
)

func joaatHash(key []byte) int {
	var hash int
	var i, key_len int

	key_len = len(key)
	hash = 0

	for i = 0; i < key_len; i++ {
		hash += int(key[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}

	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)
	return hash
}

func djb2Hash(key []byte) int {
	var hash int
	var i, key_len int

	key_len = len(key)
	hash = 5381

	for i = 0; i < key_len; i++ {
		hash = ((hash << 5) + hash) + int(key[i])
	}

	return hash
}

func sdbmHash(key []byte) int {
	var hash int
	var i, key_len int

	key_len = len(key)

	for i = 0; i < key_len; i++ {
		hash = int(key[i]) + (hash << 6) + (hash << 6) - hash
	}

	return hash
}

func loseLoseHash(key []byte) int {
	var hash int
	var i, key_len int

	key_len = len(key)

	for i = 0; i < key_len; i++ {
		hash += int(key[i])
	}

	return hash
}

func baseHashes(data []byte) [4]int {
	return [4]int{
		joaatHash(data),
		djb2Hash(data),
		sdbmHash(data),
		loseLoseHash(data),
	}
}

func add(data []byte, s *bitset.BitSet) {
	h := baseHashes(data)
	for i := int(0); i < 4; i++ {
		s.Set(h[i])
	}
}

func byteBloomFilters(file_path string, target []byte) bool {
	s := &bitset.BitSet{}
	dat, _ := os.Open(file_path)
	scanner := bufio.NewScanner(dat)

	for scanner.Scan() {
		word := scanner.Text()
		add([]byte(word), s)
	}

	h := baseHashes(target)
	for i := int(0); i < 4; i++ {
		if !s.Get(int(h[i])) {
			return false
		}
	}
	return true
}

func stringBloomFilters(file_path string, target string) bool {
	return byteBloomFilters(file_path, []byte(target))
}
