

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		lru.go
//	Author:			Zeke
//	Date:			2022.08.20
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package lru

import (
	"container/list"
)

type Cache struct {
	// maximum bytes that can be used for memory cache
	maxBytes		int64
	// bytes that have been used in the memory
	nbytes			int64
	// maitain a queue to decide which element to be removed from cache
	// the element of ll is the address of entry struct (i.e. *entry)
	ll				*list.List
	// elements cached in the memory
	// the value of cache is the address of elements in ll
	cache			map[string]*list.Element
	// call-back func called when an element is removed from cache
	OnEvicted		func(key string, value Value)
}

type entry struct {
	key				string
	value			Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(string, Value)) *Cache {
	return &Cache{
		maxBytes:	maxBytes,
		ll:			list.New(),
		cache:		make(map[string]*list.Element),
		OnEvicted:	onEvicted,
	}
}

// Get the number of elements of the queue
func (c *Cache) Len() int {
	return c.ll.Len()
}

// Get the kv pair of Cache c and update the queue
func (c *Cache) Get(key string) (value Value, ok bool){
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		return kv.value, true
	}

	return
}

// remove the least recent used element
func (c *Cache) RemoveOldest(){
	ele := c.ll.Back()
	if ele != nil {
		c.ll.Remove(ele)
		kv := ele.Value.(*entry)
		delete(c.cache, kv.key)
		c.nbytes -= int64(len(kv.key)) + int64(kv.value.Len())
		if c.OnEvicted != nil {
			c.OnEvicted(kv.key, kv.value)
		}
	}
}


func (c *Cache) Add(key string, value Value) {
	// if the key exists in the cache, then update it
	if ele, ok := c.cache[key];ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		c.nbytes += int64(value.Len()) - int64(kv.value.Len())
		kv.value = value
	} else {
		ele := c.ll.PushFront(&entry{key, value})
		c.cache[key] = ele
		c.nbytes += int64(len(key)) + int64(value.Len())
	}
	for c.maxBytes != 0 && c.maxBytes < c.nbytes {
		c.RemoveOldest()
	}
}

