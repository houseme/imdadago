# Dada Logistics Open Platform SDK For Golang Language

[![Go Report Card](https://goreportcard.com/badge/github.com/houseme/imdadago)](https://goreportcard.com/report/github.com/houseme/imdadago)
[![Go Reference](https://pkg.go.dev/badge/github.com/houseme/imdadago.svg)](https://pkg.go.dev/github.com/houseme/imdadago)
[![ImDada-CI](https://github.com/houseme/imdadago/actions/workflows/go.yml/badge.svg)](https://github.com/houseme/imdadago/actions/workflows/go.yml)
![GitHub](https://img.shields.io/github/license/houseme/imdadago?style=flat-square)
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/houseme/imdadago/main?style=flat-square)

Dada logistics information service open platform SDK for Golang language.
Help you save costs and achieve efficient distribution

## Installation

```bash
go get -u -v github.com/houseme/imdadago@main
```

## Usage

```go
package main

import (
    "context"
    "fmt"
    
    "github.com/houseme/imdadago"
)

func main()  {
    ctx := context.Background()
    d := dadago.New(ctx, dada.WithAppKey("xxxxx"), dada.WithAppSecret("xxxxx"))
    
    fmt.Println("Dada:", d)
}

```

## License
FeiE is primarily distributed under the terms of both the [Apache License (Version 2.0)](LICENSE)