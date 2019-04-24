package cache

import (
	"sync"
	"github.com/pkg/errors"
	"time"
	"runtime"
)

// 缓存的value
type item struct {
	val    interface{}
	expire int64
}

// 保存所有的缓存k/v
type Pool map[string]item

// 缓存接口
type ICache interface {
	Set(key string, val interface{}, ttl time.Duration) error
	Get(key string) (interface{}, error)
	Del(key string)
	Flush()
	Collect()              // 执行回收
	Collector() ICollector // 返回回收器
	Pool() Pool
	RWMutex() sync.RWMutex
	KeyNum() int
}

// 回收器接口，清理过期的缓存
type ICollector interface {
	Collect()            // 执行回收
	Cycle(time.Duration) // 返回回收周期
	Stop()               // 停止回收器
	Running() bool       // 是否正在运行
}

type Cache struct {
	pool      Pool
	mu        sync.RWMutex
	collector ICollector
}

func NewCache(cycle time.Duration) ICache {
	cache := &Cache{
		pool: Pool(make(map[string]item)),
	}
	// 启动回收器
	collector := NewCollector(cache, cycle)
	collector.Collect()
	cache.collector = collector
	// cache对象被销毁时停止回收器
	runtime.SetFinalizer(cache, stopCollector)
	return cache
}

func (c *Cache) Set(key string, val interface{}, ttl time.Duration) error {
	if len(key) == 0 {
		return errors.New("Invalid key")
	}

	c.mu.Lock()
	defer c.mu.Unlock()

	item := item{
		val:    val,
		expire: time.Now().Add(ttl).UnixNano(),
	}
	c.pool[key] = item

	return nil
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.pool[key]
	if !ok {
		return nil, errors.New("No such key")
	}
	if val.Expire() < time.Now().UnixNano() {
		delete(c.pool, key)
		return nil, errors.New("No such key")
	}
	return val.Val(), nil
}

func (c *Cache) Del(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, ok := c.pool[key]
	if ok {
		delete(c.pool, key)
	}
}

func (c *Cache) Flush() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.pool = Pool(make(map[string]item))
}

func (c *Cache) Collect() {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, item := range c.pool {
		if item.Expire() < time.Now().UnixNano() {
			delete(c.pool, key)
		}
	}
}

func (c *Cache) RWMutex() sync.RWMutex {
	return c.mu
}

func (c *Cache) Pool() Pool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.pool
}

func (c *Cache) KeyNum() int {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return len(c.pool)
}

func (c *Cache) Collector() ICollector {
	return c.collector
}

func (i *item) Val() interface{} {
	return i.val
}

func (i *item) Expire() int64 {
	return i.expire
}

var defaultCollectCycle = 1 * time.Second

type Collector struct {
	cache   ICache
	exit    chan bool
	cycle   time.Duration
	running bool
	mu      sync.RWMutex
}

func NewCollector(cache ICache, cycle time.Duration) ICollector {
	if cycle.Seconds() <= 0 {
		cycle = defaultCollectCycle
	}
	return &Collector{
		cache: cache,
		cycle: cycle,
		exit:  make(chan bool),
	}
}

func (c *Collector) Collect() {
	c.mu.RLock()
	if c.running {
		c.mu.RUnlock()
		return
	}
	c.mu.RUnlock()
	cycle := c.cycle
	if c.cycle.Seconds() == 0 {
		cycle = defaultCollectCycle
	}
	ticker := time.NewTicker(cycle)
	go func() {
		for {
			select {
			case <-ticker.C:
				c.cache.Collect()
			case <-c.exit:
				c.mu.Lock()
				c.running = false
				c.mu.Unlock()
				return
			}
		}
	}()
	c.mu.Lock()
	c.running = true
	c.mu.Unlock()
}

func (c *Collector) Cycle(cycle time.Duration) {
	c.cycle = cycle
}

func (c *Collector) Stop() {
	c.exit <- true
}

func (c *Collector) Running() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.running
}

func stopCollector(cache ICache) {
	cache.Collector().Stop()
}
