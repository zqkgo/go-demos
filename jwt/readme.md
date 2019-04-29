使用方法

```go
package main

import (
	"github.com/kongoole/go-demos/jwt"
)

func main() {
	go jwt.Jwt()
	select {}
}
```

身份验证获取token

```go
curl http://localhost:8080/jwt/auth
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7Im5hbWUiOiJXaWxsIiwiYWdlIjoyNywiZ2VuZGVyIjoiTWFsZSJ9fQ.DrmcrCNkw6oeaDYMNkGXuMqBmjQmyFpUmjZyVuRsjIw"}%```
```

访问资源

```go
curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyIjp7Im5hbWUiOiJXaWxsIiwiYWdlIjoyNywiZ2VuZGVyIjoiTWFsZSJ9fQ.DrmcrCNkw6oeaDYMNkGXuMqBmjQmyFpUmjZyVuRsjIw" http://localhost:8080/jwt/visit
you are accessing authorized resource
```