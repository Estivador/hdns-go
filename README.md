# hdns: A Go library for the Hetzner DNS API

[![GitHub Actions status](https://github.com/Estivador/hdns-go/workflows/Continuous%20Integration/badge.svg)](https://github.com/Estivador/hdns-go/actions)
[![GoDoc](https://godoc.org/github.com/Estivador/hdns-go/hdns?status.svg)](https://godoc.org/github.com/Estivador/hdns-go/hdns)

Package `hdns` is a library for the Hetzner DNS API.

The library’s documentation is available at [GoDoc](https://godoc.org/github.com/Estivador/hdns-go/hdns),
the public API documentation is available at [dns.hetzner.com](https://dns.hetzner.com/api-docs/).

## Example

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/Estivador/hdns-go/hdns"
)

func main() {
    client := hdns.NewClient(hdns.WithToken("token"))

    zone, _, err := client.Zone.GetByID(context.Background(), 1)
    if err != nil {
        log.Fatalf("error retrieving zone: %s\n", err)
    }
    if zone != nil {
        fmt.Printf("zone 1 is called %q\n", zone.Name)
    } else {
        fmt.Println("zone 1 not found")
    }
}
```

## License

MIT license
