

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		cache.go
//	Author:			Zeke
//	Date:			2022.08.27
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package hypercache

import (
	"sync"
	"zekehypersus.top/hypercache/lru"
)

type cache struct{
	mu			sync.Mutex
	lru			*lru.Cache
	cacheBytes	int64
}

func (c *cache) add(key string, value ByteView) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		c.lru = lru.New(c.cacheBytes, nil)
	}
	c.lru.Add(key, value)
}

func (c *cache) get(key string) (value Byteview, ok bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.lru == nil {
		return
	}

	if v, ok := c.lru.Get(key); ok{
		return v.(ByteView), ok
	}

	return 
}
