package main

import (
	"container/list"
)

type LRUCache struct {
	Map      map[int]*list.Element
	Size     int
	Capacity int
	List     *list.List
}

type Node struct {
	key int
	val int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Map:      make(map[int]*list.Element),
		Capacity: capacity,
		List:     list.New(),
	}
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.Map[key]
	if !ok {
		return -1
	}
	c.List.MoveToFront(node)
	return node.Value.(Node).val

}

func (c *LRUCache) Put(key int, value int) {
	node, ok := c.Map[key]
	if !ok {
		c.List.PushFront(Node{val: value, key: key})
		c.Map[key] = c.List.Front()
		c.Size++
		if c.Size > c.Capacity {
			tail := c.List.Back()
			c.List.Remove(tail)
			delete(c.Map, tail.Value.(Node).key)
			c.Size--
		}
		return
	}
	node.Value = Node{val: value, key: key}
	c.List.MoveToFront(node)
}
