

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		test_hypercache.go
//	Author:			Zeke
//	Date:			2022.09.05
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package hypercache

import (
	"testing"
	"log"
	"fmt"
)

var db = map[string]string {
	"testing1":"test1",
	"testing2":"test2",
	"testing3":"test3",
}

func TestGet(t *testing.T) {
	loadCounts := make(map[string]int, len(db))
	cache := NewGroup("testing", 1<<10, GetterFunc(
		func (key string) ([]byte, error) {
			// every time Getter is called, loadCount++
			log.Printf("[HyperCache] key: %s Load from device\n", key)
			// load from cache
			if v, ok := db[key]; ok {
				if _, ok := loadCounts[key]; !ok {
					loadCounts[key] = 0
				} else {
					loadCounts[key]++
				}
				return []byte(v), nil
			}
			return nil, fmt.Errorf("key %s not existing!\n", key)
		}))
	for k, v := range db {
		if view, err := cache.Get(k); err != nil || view.String() != v {
			t.Fatal("Failed to load key from device")
		}
		if _, err := cache.Get(k); err != nil || loadCounts[k] > 1 {
			t.Fatalf("Key %s missed from cache", k)
		}
	}
}
