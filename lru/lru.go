package lru

import (
	"LRU-Cache-in-golang/list"
	"fmt"
)

type LRU struct {
	hashMap   map[int]*list.Node
	cacheSize int
	capacity  int
	head      *list.Node
	last      *list.Node
}

func NewCache(size int)*LRU{
	n := list.NewNode(-1)
	return &LRU{
		hashMap: make(map[int]*list.Node,size),
		cacheSize: 0,
		capacity: size,
		head: n,
		last: nil,
	}
}

func (l *LRU)AddToCache(key int){
	fmt.Printf("Add %v to Cache",key)
	n := l.getFromCache(key)
	if n == nil {
		l.refreshCache()
		n = list.NewNode(key)
		l.head.AddNodeToHead(n)
		if l.last == nil {
			l.last = n
		}
		l.hashMap[key] = n
		l.cacheSize++
		return
	} else {
		l.updateList(n)
	}
}

func (l *LRU)GetFromCache(key int)int{
	if val := l.getFromCache(key); val != nil {
		l.updateList(val)
		return val.GetValue()
	}
	return -1
}

func (l *LRU)getFromCache(key int)*list.Node{
	val,ok := l.hashMap[key]
	if !ok {
		return nil
	}
	return val
}

func (l *LRU)updateList(n *list.Node){
	if n == l.last {
		l.last = n.RemoveNode()
	} else {
		n.RemoveNode()
	}
	l.head.AddNodeToHead(n)
}

func (l *LRU) refreshCache(){
	if l.cacheSize == l.capacity {
		delete(l.hashMap, l.last.GetValue())
		l.last = l.last.RemoveNode()
		l.cacheSize--
	}
}

func(l *LRU)Print(){
	fmt.Printf("Print LRU :\nsize -> %v\n", l.cacheSize)
	fmt.Print(l.hashMap)
	fmt.Print("\nLinked List -> ")
	l.head.Print()
	fmt.Printf("\n")
}