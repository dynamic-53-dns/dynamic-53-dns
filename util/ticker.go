package util

import (
	"fmt"
	"sync"
	"time"
)

// 互斥锁
var wg sync.WaitGroup

// 启动定时器
func Star() {
	// 设置定时器时间
	ticker := time.NewTicker(time.Second)
	wg.Add(1)
	go func() {
		var i int
		for do := range ticker.C {
			fmt.Println("sssssss====", i, "---", do)
			if i == 5 {
				wg.Done()
				ticker.Stop()
			}
			i++
		}
	}()
	// 阻塞主线程
	wg.Wait()
}
