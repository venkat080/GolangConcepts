package models

import "sync"

var TotalMessages int

var Mu sync.Mutex
var MessageLog = make(map[string]string)
var LogMutex sync.RWMutex
