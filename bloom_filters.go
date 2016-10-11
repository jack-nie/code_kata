package code_kata

func joaatHash(key string) int{
  var hash, i, key_len int
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

func djb2Hash(key string) int{
  var hash, i, key_len int
  key_len = len(key)
  hash = 5381

  for i = 0; i < key_len; i++ {
    hash = ((hash << 5) + hash) + int(key[i])
  }

  return hash
}

func sdbmHash(key string) int{
  var hash, i, key_len int

  key_len = len(key)

  for i = 0; i < key_len; i++ {
    hash = int(key[i]) + (hash << 6) + (hash << 6) - hash
  }

  return hash
}

func loseLoseHash(key string) int{
  var hash, i, key_len int

  key_len = len(key)

  for i = 0; i < key_len; i++ {
    hash += int(key[i])
  }

  return hash
}
