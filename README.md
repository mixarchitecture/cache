# this repository has been moved into [cillop](https://github.com/cilloparch/cillop).

# Mix Arch Cache Client

This is a client for the cache server. It is used to store and retrieve data from the cache server.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/mixarchitecture/cache.svg)](https://pkg.go.dev/github.com/mixarchitecture/cache)

## Installation

You can install the package using the following command:

```shell
go get github.com/mixarchitecture/cache
```

## Usage

To use the cache package, import it in your Go code:

```go
import "github.com/mixarchitecture/cache"
```

## Full Example

```go
package main

import (
 "context"
 "fmt"

 "github.com/mixarchitecture/cache"
 "github.com/mixarchitecture/i18np"
 "github.com/mixarchitecture/mredis"
)

type Entity struct {
 ID string `json:"id"`
}

func main() {
 redis := mredis.New(&mredis.Config{
  Host: "localhost",
  Port: "6379",
  DB:   0,
 })
 c := cache.New[*Entity](redis)

 key := "my-cache-key"
 cacheHandler := func() (*Entity, *i18np.Error) {
  return &Entity{ID: "my-id"}, nil
 }
 res, err := c.Creator(createEntity).Handler(cacheHandler).Get(context.Background(), key)
 if err != nil {
  fmt.Println(err)
  // handle error
 }
 fmt.Println(res.ID)
}

func createEntity() *Entity {
 return &Entity{}
}
```

## Error Keys

this package uses i18np to show error messages. [Click here](https://github.com/mixarchitecture/i18np) to learn how to use i18np.

Here is the list of error messages to be forwarded to i18np:

| Key | Description |
| --- | --- |
| cache_an_error_on_exist | An error occurred while checking the existence of the cache |
| cache_an_error_on_get | An error occurred while getting the cache |
| cache_an_error_on_set | An error occurred while setting the cache |
| cache_not_runnable | The cache is not runnable |

## Recommended Service

The service cache uses in the background is flexible and you can add your own service if you want. However, we recommend [mredis](https://github.com/mixarchitecture/redis) using redis.

## Custom Service

If you want to write your own service, make sure your service implements the following interface!

```go
// Service is an interface that defines the methods of a cache service
// It's implemented by the cache service
type Service interface {
 // Get returns a value and an error.
 Get(ctx context.Context, k string) (string, error)

 // Set sets a value and returns an error.
 SetEx(ctx context.Context, k string, v interface{}, d time.Duration) error

 // Set sets a value and returns an error.
 Exist(ctx context.Context, k string) (bool, error)
}
```

## Contributing

Contributions are welcome! If you find a bug or want to add a new feature, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
