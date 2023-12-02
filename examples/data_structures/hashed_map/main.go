package main

import "fmt"

type HashedMaps struct {
	items [100][]Node
}

type Node struct {
	Key   string
	Value string
}

func newHashedMaps() HashedMaps {
	return HashedMaps{items: [100][]Node{}}
}

func (h HashedMaps) getHash(key string) int {
	totalSum := 0
	for _, v := range key {
		totalSum = totalSum + int(v)
	}
	hashKey := totalSum % 100
	return hashKey
}

func (h *HashedMaps) set(key, val string) {
	hashedKey := h.getHash(key)

	if len(h.items[hashedKey]) == 0 {
		h.items[hashedKey] = []Node{{Key: key, Value: val}}
	}

	for _, i := range h.items[hashedKey] {
		if i.Key == key {
			return
		}
	}
	h.items[hashedKey] = append(h.items[hashedKey], Node{Key: key, Value: val})
}

func (h *HashedMaps) get(key string) string {
	hashedKey := h.getHash(key)

	if len(h.items[hashedKey]) == 0 {
		return ""
	}

	for _, item := range h.items[hashedKey] {
		if item.Key == key {
			return item.Value
		}
	}

	return ""
}

func main() {
	aa := newHashedMaps()
	aa.set("a", "sample value")
	aa.set("ABB", "unknown sample value")
	fmt.Println(aa.get("a"))
	fmt.Println(aa.get("ABB"))
}
