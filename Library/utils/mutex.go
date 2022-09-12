package utils

import "sync"

var RwMutexUser = new(sync.RWMutex)
var RwMutexBook = new(sync.RWMutex)
var RwMutexSession = new(sync.RWMutex)
