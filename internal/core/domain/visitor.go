package domain

import (
	"sync"
)

// Visitor represents a visitor with a name, URL, and domain
type Visitor struct {
	Name   string `json:"name"`
	URL    string `json:"url"`
	Domain string `json:"-"`
}

// VisitorCount represents a count of visitors for a domain.
type VisitorCount struct {
	Domain string `json:"domain"`
	Count  int    `json:"count"`
}

// Visitors represents in-menory collection of visitors with a mutex for synchronization.
type Visitors struct {
	MU       sync.Mutex
	Visitors []Visitor
}
