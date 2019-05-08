/*
 * 146. LRU Cache
 * Design and implement a data structure for Least Recently Used (LRU) cache. It should support the following operations: get and put.
 * get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
 * put(key, value) - Set or insert the value if the key is not already present. When the cache reached its capacity, it should invalidate the least recently used item before inserting a new item.
 *
 * Follow up:
 * Could you do both operations in O(1) time complexity?
 *
 * Example:
 * LRUCache cache = new LRUCache( 2 * capacity );
 *
 * cache.put(1, 1);
 * cache.put(2, 2);
 * cache.get(1);
 * cache.put(3, 3);
 * cache.get(2);
 * cache.put(4, 4);
 * cache.get(1);
 * cache.get(3);
 * cache.get(4);
 */
package solutions

import (
	"fmt"
)

type DoubleLinkListNode struct {
	Key  int
	Val  int
	Pre  *DoubleLinkListNode
	Next *DoubleLinkListNode
}

func (this *LRUCache) addNode(node *DoubleLinkListNode) {
	node.Pre = this.Head
	node.Next = this.Head.Next

	this.Head.Next.Pre = node
	this.Head.Next = node
}

func (this *LRUCache) removeNode(node *DoubleLinkListNode) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (this *LRUCache) moveToHead(node *DoubleLinkListNode) {
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCache) discardTail() {
	if this.Head.Next == this.Tail {
		return
	}
	node := this.Tail.Pre
	delete(this.KeyMap, node.Key)
	this.removeNode(node)
}

type LRUCache struct {
	DlList  *DoubleLinkListNode
	Head    *DoubleLinkListNode
	Tail    *DoubleLinkListNode
	KeyMap  map[int]*DoubleLinkListNode
	Capcity int
	Size    int
}

func Constructor(capacity int) LRUCache {
	var ret LRUCache
	ret.Head = &DoubleLinkListNode{}
	ret.Tail = &DoubleLinkListNode{}
	ret.Head.Next = ret.Tail
	ret.Tail.Pre = ret.Head
	ret.Capcity = capacity
	ret.KeyMap = make(map[int]*DoubleLinkListNode)

	return ret
}

func (this *LRUCache) Get(key int) int {
	if val, ok := this.KeyMap[key]; ok {
		this.moveToHead(val)
		return val.Val
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if val, ok := this.KeyMap[key]; ok {
		val.Val = value
		this.moveToHead(val)
		return
	}

	node := &DoubleLinkListNode{}
	node.Key = key
	node.Val = value
	this.KeyMap[key] = node

	this.addNode(node)

	this.Size++
	if this.Size > this.Capcity {
		this.discardTail()
	}

}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	fmt.Println("LRU Cache")
}
