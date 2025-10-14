package roundzero

import "container/list"

type KV2 [2]int
type LRUCacheV1 struct {
	Size    int
	Capcity int
	m       map[int]*list.Element
	l       *list.List
}

func ConstructorV1(capacity int) LRUCacheV1 {
	return LRUCacheV1{
		Size:    0,
		Capcity: capacity,
		m:       make(map[int]*list.Element),
		l:       list.New(),
	}
}

func (this *LRUCacheV1) Get(key int) int {
	elemPtr, ok := this.m[key]
	if !ok {
		return -1
	}

	this.l.MoveToFront(elemPtr)

	kv := elemPtr.Value.(KV2)
	return kv[1]
}

func (this *LRUCacheV1) Put(key int, value int) {
	elemPtr, ok := this.m[key]
	if ok {
		this.l.MoveToFront(elemPtr)
		kv := elemPtr.Value.(KV2)
		kv[1] = value
		return
	}

	for this.Size >= this.Capcity {
		kvPtr := this.l.Back()
		oldKey := kvPtr.Value.(KV2)[0]
		delete(this.m, oldKey)
		this.l.Remove(kvPtr)
		this.Size--
	}

	newKV := KV2{key, value}
	this.l.PushFront(newKV)
	
}
