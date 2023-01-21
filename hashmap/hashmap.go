package hashmap

import "sync"

// HashMap represents a hash map
type HashMap struct {
	items     map[string]interface{}
	lock      sync.RWMutex
	size      int
	threshold int
}

// NewHashMap creates and returns a new hash map
func NewHashMap(initialSize int, threshold float64) *HashMap {
	return &HashMap{
		items:     make(map[string]interface{}),
		size:      0,
		threshold: int(float64(initialSize) * threshold),
	}
}

// Set sets the value for the given key
func (h *HashMap) Set(key string, value interface{}) {
	h.lock.Lock()
	defer h.lock.Unlock()
	h.items[key] = value
	h.size++
	if h.size >= h.threshold {
		h.Resize()
	}
}

// Get gets the value for the given key
func (h *HashMap) Get(key string) (interface{}, bool) {
	h.lock.RLock()
	defer h.lock.RUnlock()
	value, ok := h.items[key]
	return value, ok
}

// Delete deletes the value for the given key
func (h *HashMap) Delete(key string) {
	h.lock.Lock()
	defer h.lock.Unlock()
	delete(h.items, key)
	h.size--
}

// Len returns the number of items in the hash map
func (h *HashMap) Len() int {
	h.lock.RLock()
	defer h.lock.RUnlock()
	return h.size
}

// LoadFactor returns the current load factor of the hashmap
func (h *HashMap) LoadFactor() float64 {
	return float64(h.size) / float64(len(h.items))
}

// Resize increases the size of the hashmap and rehashes all the items
func (h *HashMap) Resize() {
	h.lock.Lock()
	defer h.lock.Unlock()
	newItems := make(map[string]interface{}, len(h.items)*2)
	for key, value := range h.items {
		newItems[key] = value
	}
	h.items = newItems
	h.threshold = int(float64(len(h.items)) * h.LoadFactor())
}
