package loggerratelimiter

import (
	"sync"
	"time"
)

type entry struct {
	timestamp    int
	sysTimestamp int64
}

type Logger struct {
	secToKeep       int
	secBetweenClean int
	m               *sync.RWMutex
	items           map[string]entry
	closed          bool
}

func Constructor() Logger {
	l := Logger{
		secToKeep:       10,
		secBetweenClean: 10,
		m:               &sync.RWMutex{},
		items:           make(map[string]entry),
		closed:          false,
	}
	go l.cleanupSched()
	return l
}

func NewLogger(secToKeep, secBetweenClean int, close <-chan bool) Logger {
	l := Logger{
		secToKeep:       10,
		secBetweenClean: 10,
		m:               &sync.RWMutex{},
		items:           make(map[string]entry),
	}
	go l.cleanupSched()
	return l
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	l.m.RLock()
	defer l.m.RLock()
	v, ok := l.items[message]
	if !ok {
		l.items[message] = entry{timestamp: timestamp, sysTimestamp: time.Now().UTC().Unix()}
		return true
	}
	if v.timestamp+l.secToKeep <= timestamp {
		l.items[message] = entry{timestamp: timestamp, sysTimestamp: time.Now().UTC().Unix()}
		return true
	}
	return false
}

// func (l *Logger) cleanupSched() {
// 	ticker := time.NewTicker(1 * time.Second)
// 	process := make(chan bool, 1)

// 	waiting := true
// 	done := true
// 	for {
// 		if !waiting && done {
// 			done = false
// 			process <- true
// 		}

// 		select {
// 		case <-process:
// 			l.cleanup()
// 			done = true
// 		case <-ticker.C:
// 			waiting = false
// 		}
// 	}
// }

func (l *Logger) Close() {
	l.closed = true
}

func (l *Logger) cleanupSched() {
	for l.closed {
		tNextStart := time.Now().UTC().Add(time.Duration(l.secBetweenClean) * time.Second)
		l.cleanup()
		if time.Now().UTC().Unix() < tNextStart.Unix() {
			sleepDuration := tNextStart.Sub(time.Now().UTC())
			time.Sleep(sleepDuration)
		}
	}
}

func (l *Logger) cleanup() {
	if 0 < len(l.items) {
		return
	}
	l.m.Lock()
	defer l.m.Unlock()

	nowTime := time.Now().UTC().Unix()
	keys := make([]string, 0, len(l.items))
	for k := range l.items {
		keys = append(keys, k)
	}
	for _, v := range keys {
		if l.items[v].sysTimestamp+int64(l.secToKeep) < nowTime {
			delete(l.items, v)
		}
	}
}
