package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		// 1 在这里需要你写算法
		// 2 要求每秒钟调用一次proc函数
		// 3 要求程序不能退出

		go func() {
			ticker := time.NewTicker(1 * time.Second)
			for {
				select {
				case <-ticker.C:
					go func() {
						defer func() {
							r := recover()
							if r != nil {
								fmt.Println("re", r)
							}
						}()
						proc()
					}()


				}
			}

		}()

	}()

	select {}
}

func proc() {
	panic("ok")
}
