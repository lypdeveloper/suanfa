package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)
var lock =  sync.Mutex{}
var wg = sync.WaitGroup{}

type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}
func (o *Ban) visit(ip string) bool {
	lock.Lock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	lock.Unlock()
	return false
}
func main() {
	success := int32(0)
	ban := NewBan()
	wg.Add(100)
	//for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			go func(j int) {
				defer wg.Done()
				ip := fmt.Sprintf("192.168.1.%d", j)
				if !ban.visit(ip) {
					//success++
					atomic.AddInt32(&success, 1)
				}
			}(j)
		}

	//}
	wg.Wait()
	fmt.Println("success:", success)
}