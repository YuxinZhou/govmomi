package main

import (
	"context"
	"fmt"
	vim "github.com/vmware/govmomi/vim25/types"
	"github.com/vmware/govmomi/vsan/examples"
	"github.com/vmware/govmomi/vsan/methods"
	"github.com/vmware/govmomi/vsan/types"
	"log"
	"os"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	c, err := examples.NewVSANClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()

	healthSystem := vim.ManagedObjectReference{
		Type:  "VsanVcClusterHealthSystem",
		Value: "vsan-cluster-health-system",
	}

	cluster := vim.ManagedObjectReference{
		Type:  "ClusterComputeResource",
		Value: "domain-c8",
	}

	fetchFromCache := true

	res, err := methods.VsanQueryVcClusterHealthSummary(ctx, c,
		&types.VsanQueryVcClusterHealthSummary{
			This:           healthSystem,
			Cluster:        &cluster,
			Fields:         []string{"OverallHealth", "overallHealthDescription"},
			FetchFromCache: &fetchFromCache,
		})
	if err != nil {
		fmt.Fprintf(os.Stdout, "error: %+v\n", err)
		log.Fatal(err)
	}

	fmt.Println(res.Returnval.OverallHealth)
	fmt.Println(res.Returnval.OverallHealthDescription)

}
