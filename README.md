# Dada logistics information service open platform SDK for Golang language

[![Go Report Card](https://goreportcard.com/badge/github.com/houseme/imdada-go)](https://goreportcard.com/report/github.com/houseme/imdada-go)
[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/imdada-go.svg)](https://pkg.go.dev/github.com/houseme/imdada-go)
[![ImDada-CI](https://github.com/houseme/imdada-go/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/imdada-go/actions/workflows/go.yml)
![GitHub](https://img.shields.io/github/license/houseme/imdada-go?style=flat-square)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/imdada-go/main?style=flat-square)

Dada logistics information service open platform SDK for Golang language.

## Installation

```bash
go get -u -v github.com/houseme/imdada-go@main
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/houseme/imdada-go"
)

func main()  {
    ctx := context.Background()
    d := dada.New(ctx, dada.WithAppKey("xxxxx"), dada.WithAppSecret("xxxxx"))
    
    fmt.Println("Dada:", d)
}

```

## License
FeiE is primarily distributed under the terms of both the [Apache License (Version 2.0)](LICENSE)