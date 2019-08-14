Using vSAN go binding: 
```go
package main 
import (
    "log"
    "os"
	"context"
	"fmt"
	"github.com/vmware/govmomi"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan/examples"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
    "github.com/vmware/govmomi/vim25"
    "github.com/vmware/govmomi/vim25/soap"
)


const (
	vsanPath      = "/vsanHealth"
	vsanNamespace = "/vsan"
)

func main() {
	urlFlag := "https://<user>:<password>@<vcenter-ip>/sdk"
    u, err := soap.ParseURL(urlFlag)

    // Connect and log in to ESX or vCenter
    ctx := context.Background()
    govmomiClient, err := govmomi.NewClient(ctx, u, true)
    
    // Create vSAN client
    vsanClient := govmomiClient.Client.Client.NewServiceClient(vsanPath, vsanNamespace)
    
    healthSystem := vim.ManagedObjectReference{
        Type:  "VsanVcClusterHealthSystem",
        Value: "vsan-cluster-health-system",
    }
    
    cluster := vim.ManagedObjectReference{
        Type:  "ClusterComputeResource",
        Value: "domain-c8",
    }

    fetchFromCache := true

    res, err := methods.VsanQueryVcClusterHealthSummary(ctx, vsanClient,
        &types.VsanQueryVcClusterHealthSummary{
            This:           healthSystem,
            Cluster:        &cluster,
            Fields:         []string{"OverallHealth", "overallHealthDescription"},
            FetchFromCache: &fetchFromCache,
        })
    
    if err != nil {
        fmt.Println(fmt.Errorf("error querying vSAN health: %v", err))
        log.Fatal(err)
    }

    fmt.Println(res.Returnval.OverallHealth)
    fmt.Println(res.Returnval.OverallHealthDescription)
}
```

Run vSAN examples:
```
go run  ~/govmomi/vsan/examples/health/health.go --url https://<user>:<password>@<vcenter-ip>/sdk
go run  ~/govmomi/vsan/examples/cns/cns.go --url https://<user>:<password>@<vcenter-ip>/sdk
go run  ~/govmomi/vsan/examples/space/space.go --url https://<user>:<password>@<vcenter-ip>/sdk
```

Notes:
1. you might need to escape special characters in the username or password with a backslash (\)
2. you might need to change the moid of the cluster in the code:

```go
cluster := vim.ManagedObjectReference{
    Type:  "ClusterComputeResource",
    Value: "domain-c8", // vSAN MOID
}
```

