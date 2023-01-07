# Cache
Package for processing and creating a cache

# V1 preliminary implementation

```go
  Set(key string, value any) error
  Get(key string) (any, error)
  Delete(key string) error
```
```shell
go get -u -v  https://github.com/OrxanKerimov/Cache
```
___

# example 1
[play](https://goplay.space/#2xH2YJBUk_B "goplay.space")
```go
package main

import (
	"fmt"
	"github.com/OrxanKerimov/Cache/cache"
)

func main() {
	cache := cache.New()

	cache.Set("userId", 42)
	userId, _ := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, _ = cache.Get("userId")

	fmt.Println(userId)
}
```
```text
42
<nil>
```

___
