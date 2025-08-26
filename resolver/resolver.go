package resolver

import (
	"net"
	"sync"
)

type Resolver struct {
	cache map[string]string
	mu    sync.RWMutex
}

func NewResolver() *Resolver {
	return &Resolver{
		cache: make(map[string]string),
	}
}

func (r *Resolver) Resolve(ip string) string {
	r.mu.RLock()
	if name, ok := r.cache[ip]; ok {
		r.mu.RUnlock()
		return name
	}
	r.mu.RUnlock()

	names, err := net.LookupAddr(ip)
	if err != nil || len(names) == 0 {
		r.mu.Lock()
		r.cache[ip] = ip
		r.mu.Unlock()
		return ip
	}

	r.mu.Lock()
	r.cache[ip] = names[0]
	r.mu.Unlock()

	return names[0]
}
