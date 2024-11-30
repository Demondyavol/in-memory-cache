package main

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

func (c *cache) Get(key string) (string, interface{}, bool) {
	val, exist := c.memory[key]
	if !exist {
		return "Not found", nil, false
	}
	return key, val, true
}

func (c *cache) Delete(key string) (string, interface{}) {
	if value, exist := c.memory[key]; !exist {
		return "Nothing delete", nil
	} else {
		delete(c.memory, key)
		return key + "was deleted", value
	}
}
