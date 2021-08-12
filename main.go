package main

import (
	"container/list"
	"fmt"
	"log"
)

type Data struct {
	Key   string
	Value string
}

type InMemoryCache struct {
	limit           int
	evictionManager EvictionManager
}

type EvictionManager interface {
	Push(key string) int
	Pop() string
	Clear() int
}

var (
	MapData = make(map[string]string)
	Queue   = list.New()
)

func (i *InMemoryCache) Add(key, value string) int {
	// get key
	valueGet := i.Get(key)

	// if valueGet is empty string, then do next step.
	if valueGet == "" {
		// check MapData < limit or not. If true, then insert new data. Otherwise print error
		if len(MapData) < i.limit {
			MapData[key] = value
			data := Data{
				Key:   key,
				Value: MapData[key],
			}

			Queue.PushFront(data)
			return 0
		} else {
			log.Fatal("key_limit_exceeded")
		}
	}

	return 1
}

func (i *InMemoryCache) Get(key string) string {
	return MapData[key]
}

func (i *InMemoryCache) Clear() int {
	return Clear()
}

func (i *InMemoryCache) Keys() []string {
	var keys []string

	for key := range MapData {
		keys = append(keys, key)
	}

	return keys
}

func Push(key string) int {
	// check if key exist, then return 1. Otherwise, 0
	_, isExist := MapData[key]
	if isExist {
		return 1
	}

	return 0
}

func Pop() string {
	// get and remove last list
	last := Queue.Back()
	Queue.Remove(last)

	return last.Value.(Data).Key
}

func Clear() int {
	var totalKey int

	for key := range MapData {
		// remove queue
		last := Queue.Back()
		Queue.Remove(last)

		// delete map
		delete(MapData, key)

		// count deleted key
		totalKey += 1
	}
	return totalKey
}

func main() {
	var NoneEvictionManager EvictionManager

	cache := InMemoryCache{
		limit:           3,
		evictionManager: NoneEvictionManager,
	}

	fmt.Println(cache.Add("key1", "value1"))
	// fmt.Println(cache.Add("key2", "value2"))   // return 0
	// fmt.Println(cache.Add("key3", "value3"))   // return 0
	// fmt.Println(cache.Add("key2", "value2.1")) // return 1
	// fmt.Println(cache.Get("key3"))             // return value3
	// fmt.Println(cache.Get("key1"))             // return value1
	// fmt.Println(cache.Get("key3"))             // return value3
	// fmt.Println(cache.Keys())                  // return ['key1', 'key2', 'key3']
	// fmt.Println(cache.Add("key4", "ABC"))      // return / throw Error('key_limit_exceeded')
	// fmt.Println(cache.Keys())                  // return ['key1', 'key2', 'key3']
	// fmt.Println(cache.Clear()) // return 3
	// fmt.Println(cache.Keys())  // return []

}
