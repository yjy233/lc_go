package roundzero

import "container/list"

type KV [2]int

type LRUCache struct {
	Size int
	Cap  int

	K2Elem   map[int]*list.Element
	ElemList *list.List
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Size: 0,
		Cap:  capacity,

		K2Elem:   make(map[int]*list.Element),
		ElemList: list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	elem, ok := this.K2Elem[key]
	if !ok || elem == nil {
		return -1
	}

	this.ElemList.MoveToFront(elem)
	return elem.Value.(KV)[1]
}

func (this *LRUCache) Put(key int, value int) {
	// 如果原来就在
	if oldElem, ok := this.K2Elem[key]; ok && oldElem != nil {
		oldElem.Value = KV{key, value}

		this.ElemList.MoveToFront(oldElem)
		return
	}

	if this.Size >= this.Cap {
		elem := this.ElemList.Back()
		this.ElemList.Remove(elem)

		oldKey := elem.Value.(KV)[0]
		delete(this.K2Elem, oldKey)

		this.Size--
	}

	kv := KV{key, value}
	this.ElemList.PushFront(kv)
	this.K2Elem[key] = this.ElemList.Front()

	this.Size++
}
