

//===================================================================#
//	Copyright (C) 2022 Zeke. All rights reserved
// 
//	Filename:		hypercache.go
//	Author:			Zeke
//	Date:			2022.08.28
//	E-mail:			hypersus@outlook.com
//	Discription:	test script
//	
//===================================================================#

package hypercache

type Getter interface {
	Get(key string) ([]byte, error)
}

type GettrFunc Get(key string) ([]byte, error) 

func (g GetterFunc) Get(key string) ([]byte, error) {
	return g(key)
}

type Group struct{
	name		string
	getter		Getter
	mainCache	*cache
}

var (
	mu		sync.RWMutex
	groups = make(map[string]*Group),

)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group{
	if getter == nil {
		panic("nil Getter")
	}
	c := &cache{cacheBytes: cacheBytes}
	mu.Lock()
	defer mu.Unlock()
	group := &Group{
		name:		name,
		getter:		getter,
		mainCache:	c,
	}
	groups[name] = group
	return group
}

func GetGroup(key string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// Get value for a key from cache
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmr.Errorf("key is required")
	}

	if v, ok := g.mainCache.get(key); ok {
		log.Println("[HyperCache] hit")
		return v, nil
	}
	
	return g.load(key)
}

func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, error := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView(b: cloneBytes(bytes))
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
