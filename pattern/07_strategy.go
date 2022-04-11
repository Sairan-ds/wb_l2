package pattern

import "fmt"

//_______________________________________________________________________
// Интерфейс стратегии


type evictionAlgo interface {
    evict(c *cache)
}


//_______________________________________________________________________
// Конкретная стратегия

type Fifo struct {
}

func (l *Fifo) evict(c *cache) {
    fmt.Println("Evicting by fifo strtegy")
}

//_______________________________________________________________________
// Конкретная стратегия

type Lru struct {
}

func (l *Lru) evict(c *cache) {
    fmt.Println("Evicting by lru strtegy")
}
//_______________________________________________________________________
// Конкретная стратегия
type Lfu struct {
}

func (l *Lfu) evict(c *cache) {
    fmt.Println("Evicting by lfu strtegy")
}
//_______________________________________________________________________
// Контекст

type cache struct {
    storage      map[string]string
    evictionAlgo evictionAlgo
    capacity     int
    maxCapacity  int
}

func InitCache(e evictionAlgo) *cache {
    storage := make(map[string]string)
    return &cache{
        storage:      storage,
        evictionAlgo: e,
        capacity:     0,
        maxCapacity:  2,
    }
}

func (c *cache) SetEvictionAlgo(e evictionAlgo) {
    c.evictionAlgo = e
}

func (c *cache) Add(key, value string) {
    if c.capacity == c.maxCapacity {
        c.evict()
    }
    c.capacity++
    c.storage[key] = value
}

func (c *cache) Get(key string) {
    delete(c.storage, key)
}

func (c *cache) evict() {
    c.evictionAlgo.evict(c)
    c.capacity--
}


