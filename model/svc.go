package model

import "sync"

type Service interface {
}

type serviceManager struct {
	store *Storager

	mu  sync.RWMutex
	avS map[string]interface{}
}
