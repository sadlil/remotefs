package main

import (
	"time"
)

type MemoryFile struct {
	IsDir       bool
	Content     []byte
	LastUpdated time.Time
	Created     time.Time
}
