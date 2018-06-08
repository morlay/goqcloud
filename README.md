## Tencent Cloud SDK for go

[![GoDoc Widget](https://godoc.org/github.com/morlay/goqcloud?status.svg)](https://godoc.org/github.com/morlay/goqcloud)
[![Go Report Card](https://goreportcard.com/badge/github.com/morlay/goqcloud)](https://goreportcard.com/report/github.com/morlay/goqcloud)

Generated form [Tencent Cloud API docs](https://cloud.tencent.com/document/api)


## How to

* run `make` to generate all clients


```go
package main

import (
    "os"
    "fmt"
    "time"

    "github.com/morlay/goqcloud"
    "github.com/morlay/goqcloud/clients/cvm"
)

func main() {
    client := goqcloud.NewClientWithCredential(
        os.Getenv("TENCENTCLOUD_SECRET_ID"),
        os.Getenv("TENCENTCLOUD_SECRET_KEY"),

        // optional config timeout
        goqcloud.ClientOptionWithTimeout(1 * time.Hour),
        // log each request        
        goqcloud.ClientOptionWithTransports(
        	goqcloud.NewLogTransport(),
        	// or define your own transport
        ),
    )
    
    req := cvm.DescribeInstancesRequest{
        Region: "cn-beijing",
    }
    
    resp, err := req.Invoke(client)
    
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("%v", resp.InstanceSet)
}
```

## Rules

* required field will not be pointer