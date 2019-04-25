`heartbeat` 作为一个独立`goroutine`提供`http`服务，客户端可定时发起探测请求。

```go
package main

import (
	"github.com/kongoole/go-demos/heartbeat"
	"time"
	"log"
)

func main() {
	go heartbeat.Heartbeat("localhost:8080")
	for {
		<- time.After(1 * time.Second)
		msg, err := heartbeat.Client().Check("http://localhost:8080")
		if err != nil {
			log.Fatal("errors: ", err)
		}
		log.Println(msg.UntilNow)
	}
}

```