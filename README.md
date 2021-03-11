# envstruct
![badge](https://github.com/yihau/envstruct/actions/workflows/go.yml/badge.svg?branch=main)

a simple way to get your environment variable.

## Getting Started


### Installing

```sh
go get -v github.com/yihau/envstruct
```

### Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/yihau/envstruct"
)

type Config struct {
	Host string `env:"HOST"`
	PORT int    `env:"PORT"`
}

func main() {
	var c Config
	err := envstruct.FillIn(&c)
	if err != nil {
		log.Fatalf("parse env config error, err: %v", err)
	}
	fmt.Printf("%+v\n", c)
}

```

```sh
export HOST=yihau.dev PORT=8080 && go run main.go
# {Host:yihau.dev PORT:8080}
```

## Running the tests

```sh
go run cmd/gen_decode_test/main.go

go test -v ./...
```

## License

This project is licensed under the MIT License
