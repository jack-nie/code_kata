package code_kata

import (
  "os"
  "bufio"
)

const SIZE =  2 << 25

type bits uint64
type BitSet []bits

func joaatHash(key string) uint{
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

func djb2Hash(key string) uint{
  var hash uint
  var i, key_len int

  key_len = len(key)
  hash = 5381

  for i = 0; i < key_len; i++ {
    hash = ((hash << 5) + hash) + uint(key[i])
  }

  return hash
}

func sdbmHash(key string) uint{
  var hash uint
  var i, key_len int

  key_len = len(key)

  for i = 0; i < key_len; i++ {
    hash = uint(key[i]) + (hash << 6) + (hash << 6) - hash
  }

  return hash
}

func loseLoseHash(key string) uint{
  var hash uint
  var i, key_len int

  key_len = len(key)

  for i = 0; i < key_len; i++ {
    hash += uint(key[i])
  }

  return hash
}

func (s *BitSet) Set(i uint) {
  if len(*s) < int(i/SIZE+1) {
    r := make([]bits, i/SIZE+1)
    copy(r, *s)
    *s = r
  }
  (*s)[i/SIZE] |= 1 << (i % SIZE)
}

func (s *BitSet) Clear(i uint) {
  if len(*s) >= int(i/SIZE+1) {
    (*s)[i/SIZE] &^= 1 << (i % SIZE)
  }
}

func (s *BitSet) IsSet(i uint) bool {
  return (*s)[i/SIZE]&(1<<(i%SIZE)) != 0
}

func (s *BitSet) add(value string) {
  s.Set(joaatHash(value))
  s.Set(djb2Hash(value))
  s.Set(sdbmHash(value))
  s.Set(loseLoseHash(value))
}

func (s *BitSet) contains(value string) bool {
  return s.IsSet(joaatHash(value)) &&
    s.IsSet(djb2Hash(value)) &&
    s.IsSet(sdbmHash(value)) &&
    s.IsSet(loseLoseHash(value))

}

func bloomFilters(file_path string, target string) bool {
  s := new(BitSet)
  dat, _ := os.Open(file_path)
  scanner := bufio.NewScanner(dat)

  for scanner.Scan() {
    word := scanner.Text()
    s.add(word)
  }

  return s.contains(target)
}
