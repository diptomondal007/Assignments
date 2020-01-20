package main

import "fmt"

type evictionAlgorithm interface {
	evict(c *cache)
}

type fifo struct {

}

func (f *fifo) evict(c *cache){
	fmt.Println("Evicting by fifo strategy")
}

type lru struct {
}

func (l *lru) evict(c *cache) {
	fmt.Println("Evicting by lru strategy")
}

type lfu struct {

}

func (l *lfu) evict(c *cache){
	fmt.Println("Evicting by lfu strategy")
}

type cache struct {
	storage map[string]string
	evictionAlgorithm evictionAlgorithm
	capacity int
	maxCapacity int
}

func initCache(e evictionAlgorithm) *cache{
	storage := make(map[string]string)
	return &cache{
		storage:           storage,
		evictionAlgorithm: e,
		capacity:          0,
		maxCapacity:       2,
	}
}

func (c *cache) setEvictionAlgorithm(e evictionAlgorithm){
	c.evictionAlgorithm = e
}

func (c *cache) add(key, value string){
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *cache) evict() {
	c.evictionAlgorithm.evict(c)
	c.capacity--
}


func main(){
	lfu := &lfu{}
	cache := initCache(lfu)
	cache.add("1","one")
	cache.add("2","two")
	cache.add("2","two")
	lru := &lru{}
	cache.setEvictionAlgorithm(lru)
	cache.add("4","four")
}