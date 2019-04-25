**cache**提供进程内kv缓存

* 支持自动和手动过期回收
* 支持Go复杂数据类型：`struct`、`map`、`array`、`func`等
* 支持并发读写

用法
```go
package main

import (
	"time"
	"github.com/kongoole/go-demos/cache"
	"fmt"
)

func main() {
	double := func(n int) int {
		return 2 * n
	}
	c := cache.NewCache(1)
	c.Set("fk", double, 3 * time.Second)
	v,_ := c.Get("fk")
	cv := v.(func(int) int)
	fmt.Println(cv(8)) // 16

	<-time.After( 3 * time.Second)
	c.Collect()
	_,err := c.Get("fk")
	if err != nil {
		fmt.Println(err) // "No such key"
	}
}

```