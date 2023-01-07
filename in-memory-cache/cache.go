package cache

import "errors"

type Cache struct {
	memory map[string]interface{}
}

func New() *Cache {
	return &Cache{memory: make(map[string]interface{})}
}

func (c *Cache) Set(key string, value interface{}) error {
	c.memory[key] = value
	if !c.isExist(key) {
		return errors.New("Нет данных об этом кэщэ ")
	}
	return nil
}

func (c Cache) Get(key string) (interface{}, error) {
	if c.isExist(key) {
		return c.memory[key], nil
	}
	return "", errors.New("Нет данных об этом кэщэ ")
}

func (c *Cache) Delete(key string) error {
	if !c.isExist(key) {
		return errors.New("Нет данных об этом кэщэ ")
	}
	delete(c.memory, key)
	if !c.isExist(key) {
		return errors.New("Кэш не был удален ")
	}
	return nil
}

func (c Cache) isExist(key string) bool {
	_, exist := c.memory[key]
	return exist
}
