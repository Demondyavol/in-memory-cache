package inMemoryCache

type cache struct {
	memory map[string]interface{}
}

func Main() {
}

func New() *cache {
	return &cache{
		memory: make(map[string]interface{}),
	}
}

func (c *cache) Set(key string, value interface{}) {
	c.memory[key] = value
}

func (c *cache) Get(key string) interface{} {
	val, exist := c.memory[key]
	if !exist {
		return "Not found"
	}
	return val
}

func (c *cache) Delete(key string) {
	delete(c.memory, key)
}
