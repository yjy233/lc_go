package roundzero

import "container/list"

type KV6 [2]int
type LRUCacheV6 struct {
	L    *list.List
	M    map[int]*list.Element
	Size int
	Cap  int
}

func ConstructorV6(capacity int) LRUCacheV6 {
	return LRUCacheV6{
		Cap:  capacity,
		Size: 0,
		L:    list.New(),
		M:    make(map[int]*list.Element),
	}
}

func (this *LRUCacheV6) Get(key int) int {

	oldElem, ok := this.M[key]
	if !ok {
		return -1
	}

	this.L.MoveToFront(oldElem)
	return oldElem.Value.(KV6)[1]
}

func (this *LRUCacheV6) Put(key int, value int) {

	oldElem, ok := this.M[key]
	if ok {
		oldElem.Value = KV6{key, value}
		this.L.MoveToFront(oldElem)
		return
	}

	if this.Size >= this.Cap {
		oldElem := this.L.Back()
		oldKey := oldElem.Value.(KV6)[0]

		delete(this.M, oldKey)
		this.L.Remove(oldElem)

		this.Size--
	}

	newKV := KV6{key, value}
	this.L.PushFront(newKV)
	this.M[key] = this.L.Front()

	this.Size++
}
