

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		lru_test.go
//	Author:			Zeke
//	Date:			2022.08.21
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package lru

import (
	"testing"
	"reflect"
)

type String string

func (s String) Len() int {
	return len(s)
}


func TestGet(t *testing.T) {
	lru := New(int64(0), nil)
	lru.Add("key1", String("1234"))
	if v, ok := lru.Get("key1"); !ok || string(v.(String)) != "1234" {
		t.Fatalf("cache hit key1=1234 failed!")
	}
	if _, ok := lru.Get("key2"); ok {
		t.Fatalf("cache miss key2 failed")
	}

}

func TestRemoveoldest(t *testing.T) {
	k1, k2, k3 := "key1", "key2", "key3"
	v1, v2, v3 := "value1", "value2", "value3"
	capacity := len(k1+k2+v1+v2)
	lru := New(int64(capacity), nil)
	lru.Add(k1, String(v1))
	lru.Add(k2, String(v2))
	lru.Add(k3, String(v3))

	if _, ok := lru.Get("key1"); ok || lru.Len() != 2 {
		t.Fatalf("Removeoldest key1 failed")
	}
}

func TestOnEvicted(t *testing.T) {
	keys := make([]string, 0)
	callback := func(key string, value Value) {
		keys = append(keys, key)
	}
	lru := New(int64(16), callback)
	lru.Add("key1", String("val1"))
	lru.Add("key2", String("val2"))
	lru.Add("key3", String("val3"))
	lru.Add("key4", String("val4"))

	expect := []string{"key1", "key2"}
	if !reflect.DeepEqual(expect, keys) {
		t.Fatalf("Call OnEvicted failed, expect keys to equals to %s", expect)
	}
}


