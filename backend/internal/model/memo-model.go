package model

import "time"

type Memory struct {
	MemoryId   string
	MemoryDate time.Time
	MemoryPics []byte
	MemoryNote string
}
