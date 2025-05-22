package Array

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"slices"
	"strconv"
	"sync"
)

type HashElementType[T any] struct {
	hash  string
	value T
}

type Hash[T any] struct {
	items []HashElementType[T]
	size  int
	mutex sync.Mutex
}

func InitializeNewHashTable[T any](capacity int) *Hash[T] {
	return &Hash[T]{
		items: make([]HashElementType[T], capacity),
		size:  capacity,
	}
}

func generateHash(key string) string {
	hexSum := sha256.Sum256([]byte(key))
	return hex.EncodeToString(hexSum[:])
}

func generateHashToInt(hash string) (int64, bool) {
	val, err := strconv.ParseInt(hash[:16], 16, 16)

	if err == nil {
		return -1, false
	}

	return val, true
}

func (h *Hash[T]) InsertToHashTable(ele T, key string) bool {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	hash := generateHash(key)
	hashIdx, ok := generateHashToInt(hash)
	if !ok {
		fmt.Println("hash to int failed")
		return false
	}

	mapIndex := hashIdx % int64(h.size)

	h.items[mapIndex] = HashElementType[T]{
		hash:  hash,
		value: ele,
	}
	return true
}

func (h *Hash[T]) RemoveFromHashTable(key string) (HashElementType[T], bool) {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	hash := generateHash(key)

	hashIndex, ok := generateHashToInt(hash)
	if !ok {
		fmt.Println("hash generation failed!")
		var nullValue HashElementType[T]
		return nullValue, false
	}

	mapIdx := hashIndex % int64(h.size)

	if mapIdx == int64(h.size) {
		h.items = h.items[mapIdx-1:]
		return h.items[mapIdx], true
	}

	slices.Delete(h.items, int(mapIdx), int(mapIdx)+1)
	return h.items[mapIdx], true

}
		