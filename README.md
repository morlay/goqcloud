## Tencent APIs of go version

Generated form [Tencent API docs](https://cloud.tencent.com/document/api)


## How to

* run `make` to generate all clients


```go
package main

import (
	"os"
	"fmt"

	"github.com/morlay/goqcloud"
	"github.com/morlay/goqcloud/clients/cvm"
)

func main() {
	client := goqcloud.NewClientWithCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
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