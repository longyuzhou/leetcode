package main

import "fmt"

// 实现LRU Cache
//
// Solution: 双向链表
//

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

// 将此节点从链表中脱离
func (this *Node) Detach() {
	prev := this.prev
	next := this.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	this.prev = nil
	this.next = nil
}

// 将此节点附加在指定节点之前
func (this *Node) AppendBefore(node *Node) {
	if node == nil {
		return
	}

	prev := node.prev
	if prev != nil {
		prev.next = this
	}
	node.prev = this
	this.prev = prev
	this.next = node
}

type LRUCache struct {
	capacity int
	data     map[int]*Node
	head     *Node
	tail     *Node
}

func Constructor(capacity int) LRUCache {
	inst := LRUCache{}
	inst.capacity = capacity
	inst.data = make(map[int]*Node)
	return inst
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.data[key]
	if !ok {
		return -1
	}
	if node != this.head {
		if node == this.tail {
			this.tail = node.prev
		}
		node.Detach()
		node.AppendBefore(this.head)
		this.head = node
	}
	return node.value
}

func (this *LRUCache) Put(key int, value int) {
	if this.Get(key) != -1 {
		this.data[key].value = value
	} else {
		size := len(this.data)
		for i := 0; i <= size-this.capacity; i++ {
			node := this.tail
			this.tail = node.prev
			node.Detach()
			delete(this.data, node.key)
		}
		node := &Node{key: key, value: value}
		this.data[key] = node
		node.AppendBefore(this.head)
		this.head = node
		if this.tail == nil {
			this.tail = node
		}
	}
}

func main() {
	capacity := 2
	cache := Constructor(capacity)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1) == 1)  // returns 1
	cache.Put(3, 3)                 // evicts key 2
	fmt.Println(cache.Get(2) == -1) // returns -1 (not found)
	cache.Put(4, 4)                 // evicts key 1
	fmt.Println(cache.Get(1) == -1) // returns -1 (not found)
	fmt.Println(cache.Get(3) == 3)  // returns 3
	fmt.Println(cache.Get(4) == 4)  // returns 4
}
