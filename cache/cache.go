/*
@Time : 26/10/2020
@Author : GC
@Desc : 
*/

package cache

import (
	"fmt"
	"sync"
)

type Cache interface {
	Set(string, []byte) error
	Get(string) ([]byte, error)
	Del(string) error
	GetStat() Stat
}

// implement Cache interface
type memCache struct {
	c map[string][]byte // TODO 基准测试, 和 sync.Map 对比性能
	sync.RWMutex
	Stat
}

func (c *memCache) Set(k string, v []byte) error {
	c.Lock()
	defer c.Unlock()
	tmp, exist := c.c[k]
	// TODO 基准测试, 对比不统计valueSize 时的性能
	if exist {
		c.del(k, tmp)
	}
	c.c[k] = v
	c.add(k, v)
	return nil
}

func (c *memCache) Get(k string) ([]byte, error) {
	// TODO 没有 lru 特性, 缓存功能基本是不可用的
	c.RLock()
	defer c.RUnlock()
	// TODO 处理数据不存在的情况
	return c.c[k], nil
}

func (c *memCache) Del(k string) error {
	c.Lock()
	defer c.Unlock()
	tmp, exist := c.c[k]
	if exist {
		delete(c.c, k)
		c.del(k, tmp)
	}
	return nil
}

func (c *memCache) GetStat() Stat {
	return c.Stat
}

func newMemCache() *memCache {
	return &memCache{
		make(map[string][]byte),
		sync.RWMutex{},
		Stat{},
	}
}

func New() (c Cache) {
	// TODO 可扩展多种cache
	c = newMemCache()
	fmt.Println("memcache init")
	return c
}
